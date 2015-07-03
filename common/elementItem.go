package common

type ElementItem struct {
	UrlStr     string
	FaildCount int
}

func NewElementItem(urlStr string) ElementItem {
	return ElementItem{
		UrlStr:     urlStr,
		FaildCount: 0,
	}
}
