package ziface

type IRouter interface {
	PreHandle(*IRequest) error
	Handle(*IRequest) error
	PostHandle(*IRequest) error
}
