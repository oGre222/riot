package tikv

import (
	"context"
	"github.com/hashicorp/golang-lru"
	pb "github.com/oGre222/tea/tikv/lru_grpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type NoticeType uint
const (
	None NoticeType = iota
	RemoveKey
	BatchRemoveKey
)

type Notice struct {
	Type   NoticeType
	Data   interface{}
}

type Lru struct {
	lruCache *lru.TwoQueueCache
	receiveCh []chan <- Notice
	noticeCh  chan Notice
}

func NewLru (lruCache *lru.TwoQueueCache, address string, peers []string) chan <- Notice {
	l := &Lru{
		lruCache: lruCache,
		noticeCh: make(chan Notice),
	}
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLruGRpcServer(s, l)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	go l.startListen(peers)
	return l.noticeCh
}

func (s *Lru) startListen(peers []string) {
	for _, v := range peers  {
		s.receiveCh = append(s.receiveCh, NewLruClient(v))
	}
	for  {
		select {
		case n := <- s.noticeCh:
			if n.Data == nil {
				continue
			}
			for _, c := range s.receiveCh  {
				c <- n
			}
		}
	}
}

func (s *Lru) Remove(ctx context.Context, in *pb.Key) (*pb.Response, error) {
	log.Printf("receive key:%s", in.Data)
	s.lruCache.Remove(in.Data)
	return &pb.Response{Success:true}, nil
}

func (s *Lru) RemoveBatch(ctx context.Context, in *pb.Keys) (*pb.Response, error) {
	log.Println("receive key:", in.Data)
	for _, v := range in.Data  {
		s.lruCache.Remove(v)
	}
	return &pb.Response{Success:true}, nil
}

type LruClient struct {
	client pb.LruGRpcClient
	receiveC chan Notice
}

func NewLruClient (address string)  chan <- Notice {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	lc := LruClient{
		client: pb.NewLruGRpcClient(conn),
		receiveC: make(chan Notice),
	}
	go lc.run()
	return lc.receiveC
}

func (lc *LruClient) run() {
	for  {
		select {
		case n := <- lc.receiveC:
			lc.transNotice(n)
		}
	}
}

func (lc *LruClient) transNotice(n Notice) {
	var r *pb.Response
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	switch n.Type {
	case RemoveKey:
		d, ok := n.Data.(string)
		if !ok {
			log.Println("type error", n)
		}
		r, err = lc.client.Remove(ctx, &pb.Key{Data:d})
	case BatchRemoveKey:
		d, ok := n.Data.([]string)
		if !ok {
			log.Println("type error", n)
		}
		r, err = lc.client.RemoveBatch(ctx, &pb.Keys{Data:d})
	}
	if err != nil {
		log.Printf("could not greet: %v", err)
	} else {
		log.Println("Greeting:", r.Success)
	}
}
