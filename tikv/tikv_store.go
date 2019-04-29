// Copyright 2016 ego authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Tikv is key-value store use tikv+pd, before this use lru cache data.
// for search engine is huge keyword => []list table, but store to tikv is kv struct
// with id prefix。

package tikv

import (
	"errors"
	"github.com/hashicorp/golang-lru"
	"github.com/pingcap/tidb/config"
	ti "github.com/pingcap/tidb/store/tikv"
	"strings"
)

const DefaultLruSize = 100000

// Bolt bolt store struct
type Tikv struct {
	cli      *ti.RawKVClient
	lruCache *lru.TwoQueueCache
	noticeCh chan <- Notice
}

type KvData struct {
	Key []byte
	Val []byte
}

type KvScanData struct {
	Keys [][]byte
	Values [][]byte
}

// OpenBolt open Bolt store
func OpenTikv(addr string, listen string, peers string,  lruSize ...int) (*Tikv, error) {
	cli, err := ti.NewRawKVClient(strings.Split(addr, ","), config.Security{})
	if err != nil {
		cli.Close()
		return nil, err
	}
	size := DefaultLruSize
	if len(lruSize) > 0 && lruSize[0] > 0 {
		size = lruSize[0]
	}
	l, err := lru.New2Q(size)
	if err != nil {
		cli.Close()
		return nil, err
	}
	return &Tikv{
		cli: cli,
		lruCache: l,
		noticeCh: NewLru(l, listen, strings.Split(peers, ",")),
	}, nil
}

// Set executes a function within the context of a read-write managed
// transaction. If no error is returned from the function then the transaction
// is committed. If an error is returned then the entire transaction is rolled back.
// Any error that is returned from the function or returned from the commit is returned
// from the Update() method.
func (s *Tikv) Set(k []byte, v []byte) error {
	//s.lruCache.Add(k, v)
	//s.lruCache.Remove(string(k))
	if len(k) == 0 || len(v) == 0 {
		return errors.New("empty")
	}
	s.noticeCh <- Notice{Type:RemoveKey, Data: string(k)}
	return s.cli.Put(k, v)
}

// Get executes a function within the context of a managed read-only transaction.
// Any error that is returned from the function is returned from the View() method.
func (s *Tikv) Get(k []byte) (b []byte, err error) {
	if v, ok := s.lruCache.Get(string(k)); ok {
		return v.([]byte), nil
	}

	b, err = s.cli.Get(k)
	if err == nil {
		s.lruCache.Add(string(k), b)
	}

	return b, err
}

// Delete deletes a key. Exposing this so that user does not
// have to specify the Entry directly.
func (s *Tikv) Delete(k []byte) error {
	//s.lruCache.Remove(string(k))
	s.noticeCh <- Notice{Type:RemoveKey, Data: string(k)}
	return s.cli.Delete(k)
}

// Has returns true if the DB does contains the given key.
func (s *Tikv) Has(k []byte) (bool, error) {
	d, err := s.Get(k)

	if err != nil || len(d) == 0 {
		return false, err
	}
	return true, nil
}

func (s *Tikv) BatchPut(data map[string][]byte, reTokens ...string) {
	if len(data) > 0 {
		var keys, values [][]byte
		for k, v := range data {
			keys = append(keys, []byte(k))
			values = append(values, v)
		}
		err := s.cli.BatchPut(keys, values)
		if err != nil {
			return
		}
	}
	if len(reTokens) > 0 {
		s.noticeCh <- Notice{Type:BatchRemoveKey, Data: reTokens}
	}
	/*for _, k := range reTokens  {
		s.lruCache.Remove(k)
	}*/
}

func (s *Tikv) BatchDelete(keys [][]byte) {
	err := s.cli.BatchDelete(keys)
	if err == nil {
		var removeKeys []string
		for _, k := range keys {
			//s.lruCache.Remove(string(k))
			removeKeys = append(removeKeys, string(k))
		}
		if len(removeKeys) > 0 {
			s.noticeCh <- Notice{Type:BatchRemoveKey, Data: removeKeys}
		}
	}

}

//左匹配，开区间
func (s *Tikv) PreLike(key []byte) (keys [][]byte, values [][]byte, err error) {
	v, found := s.lruCache.Get(string(key))
	if found {
		if res, ok := v.(KvScanData); ok {
			return res.Keys, res.Values, nil
		}
		return nil, nil, errors.New("type error")
	}

	startKey := key
	for  {
		k, v, e := s.cli.Scan(startKey, append(key, 255), ti.MaxRawKVScanLimit)
		if e != nil {
			return nil, nil, e
		}
		keys = append(keys, k...)
		values = append(values, v...)
		if len(keys) < ti.MaxRawKVScanLimit {
			break
		}
		startKey = k[len(k) - 1]
	}
	s.lruCache.Add(string(key), KvScanData{Keys: keys, Values:values})
	return
}

// Close releases all database resources. All transactions
// must be closed before closing the database.
func (s *Tikv) Close() error {
	return s.cli.Close()
}
