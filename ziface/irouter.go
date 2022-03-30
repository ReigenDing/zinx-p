package ziface

type IRouter interface {
	PreHandle(request IRequest) error
	Handle(request IRequest) error
	PostHandle(request IRequest) error
}
