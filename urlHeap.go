package singoriensis

import (
	"container/list"
	"crypto/md5"
	"encoding/hex"
	"singoriensis/common"
)

type UrlHeap struct {
	hash []*list.List
	size int
}

func NewUrlHeap(heapSize int) *UrlHeap {
	return &UrlHeap{
		hash: make([]*list.List, heapSize+1),
		size: heapSize,
	}
}

func (self *UrlHeap) Contain(elem common.ElementItem) bool {
	var i int64
	flag := false
	crypto := md5.New()
	crypto.Write([]byte(elem.UrlStr))
	hashNum := common.NewDjb2Hash(hex.EncodeToString(crypto.Sum(nil)))

	if hashNum < 0 {
		i = int64(self.size)
	} else {
		i = common.NewDjb2Hash(hex.EncodeToString(crypto.Sum(nil))) % int64(self.size)
	}

	hashList := self.hash[int(i)]
	if hashList != nil {
		nextItem := hashList.Front()
		for ; nextItem != nil; nextItem = nextItem.Next() {
			value := nextItem.Value.(common.ElementItem)
			if value.UrlStr == elem.UrlStr {
				flag = true
			}
		}
	} else {
		hashList = list.New()
		hashList.PushBack(elem)
		self.hash[i] = hashList
	}

	return flag
}
