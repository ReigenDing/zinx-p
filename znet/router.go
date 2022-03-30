package znet

import (
	"fmt"

	"zinx.dev/v0/ziface"
)

type BaseRouter struct {
}

// template pattern
func (br *BaseRouter) PreHandle(request ziface.IRequest) error {
	fmt.Printf("prehandle\n")
	return nil
}

func (br *BaseRouter) Handle(request ziface.IRequest) error {
	fmt.Printf("handle\n")
	return nil
}

func (br *BaseRouter) PostHandle(request ziface.IRequest) error {
	fmt.Printf("posthandle\n")
	return nil

}
