package singoriensis

import (
	"container/list"
	"crypto/md5"
	"encoding/hex"
	"math/big"
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
	flag := false
	crypto := md5.New()
	crypto.Write([]byte(elem.UrlStr))
	hashNum := common.NewDjb2Hash(hex.EncodeToString(crypto.Sum(nil)))
	hashNum = hashNum.Mod(hashNum, big.NewInt(int64(self.size)))
	i := hashNum.Int64()

	hashList := self.hash[int(i)]
	if hashList != nil {
		nextItem := hashList.Front()
		for ; nextItem != nil; nextItem = nextItem.Next() {
			value := nextItem.Value.(common.ElementItem)
			if value.UrlStr == elem.UrlStr {
				flag = true
				break
			}
		}

		if !flag {
			hashList.PushBack(elem)
		}

	} else {
		hashList = list.New()
		hashList.PushBack(elem)
		self.hash[i] = hashList
	}

	return flag
}
