package ziface

type IServer interface {
    // 启动服务起
    Start()

    // 停止服务器
    Stop()

    //运行服务器

    Server()
}