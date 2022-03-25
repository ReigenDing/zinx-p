package znet

import "zinx.dev/v0/ziface"

type BaseRouter struct {
}

// template pattern
func (br *BaseRouter) PreHandle(request *ziface.IRequest) error {
	return nil
}

func (br *BaseRouter) Handle(request *ziface.IRequest) error {
	return nil
}

func (br *BaseRouter) PostHandle(request *ziface.IRequest) error {
	return nil

}
