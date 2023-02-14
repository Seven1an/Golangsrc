package main

import (
    "flag"
    "fmt"
    "net"
    "time"
)

func main() {
    // 解析命令行参数
    var listenPort string
    flag.StringVar(&listenPort, "port", "8888", "port to listen on")
    flag.Parse()

    // 监听指定端口
    listener, err := net.Listen("tcp", ":"+listenPort)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        return
    }
    defer listener.Close()

    fmt.Println("Listening on port " + listenPort)

    for {
        // 等待客户端连接
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting:", err.Error())
            continue
        }

        // 获取客户端的IP地址
        remoteAddr := conn.RemoteAddr().String()

        // 打印客户端的IP地址和当前时间
        fmt.Printf("Received connection from %s at %s\n", remoteAddr, time.Now().Format(time.RFC3339))

        // 关闭连接
        conn.Close()
    }
}
