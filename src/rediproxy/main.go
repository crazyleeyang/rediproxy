package main

import (
    "gen-go/redis/String"
    "fmt"
    "rediproxy/cmd"
    "git.apache.org/thrift.git/lib/go/thrift"
    "github.com/Unknwon/goconfig"
    "os"
)

var(
    serverPort      string
    serverIp        string
)

func main() {
    err := loadConf()
    if err != nil{
        fmt.Println("read conf failed", err.Error())
        os.Exit(1)
    }

    transportFactory        := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    protocolFactory         := thrift.NewTBinaryProtocolFactoryDefault()
    serverTransport, err    := thrift.NewTServerSocket(serverIp + ":" + serverPort)
    if err != nil {
        fmt.Println("Error!", err)
        os.Exit(1)
    }

    handler     := &cmd.Cmd{}
    processor   := String.NewStringRpcProcessor(handler)
    server      := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
    fmt.Println("server start at", serverIp, serverPort)
    server.Serve()
    fmt.Println("server over", serverIp, serverPort)
}

func loadConf()(err error){
    cfg, err := goconfig.LoadConfigFile("conf/rediproxy.ini")
    if err != nil {
        return
    }

    serverPort, err = cfg.GetValue("rpc", "serverPort")
    if err != nil {
        return
    }

    serverIp, err = cfg.GetValue("rpc", "serverIp")
    if err != nil {
        return
    }

    return
}

