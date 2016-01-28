package error


type ResponseError struct {
	Exist bool
}

func (self ResponseError) Error() string {
	return "client force response error..."
}