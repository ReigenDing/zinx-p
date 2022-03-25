package ziface

type IRequest interface {
	GetData() []byte
	GetConn() IConnection
}
