package singoriensis

type UrlHeap struct{}

func NewUrlHeap() *UrlHeap {
	return &UrlHeap{}
}

func (self *UrlHeap) Contain(urlStr string) bool {
	return true
}
