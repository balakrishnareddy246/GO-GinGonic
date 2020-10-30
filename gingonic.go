package main

import (
    "bitbucket-repository-management-service/internal/gRPC/impl"
    "bitbucket-repository-management-service/internal/gRPC/service"
    "fmt"
    "log"
    "net"

    "google.golang.org/gRPC"
)

func main() {
    netListener := getNetListener(7000)
    gRPCServer := gRPC.NewServer()

    repositoryServiceImpl := impl.NewRepositoryServiceGrpcImpl()
    service.RegisterRepositoryServiceServer(gRPCServer, repositoryServiceImpl)

    // start the server
    if err := gRPCServer.Serve(netListener); err != nil {
        log.Fatalf("failed to serve: %s", err)
    }

}

func getNetListener(port uint) net.Listener {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
        panic(fmt.Sprintf("failed to listen: %v", err))
    }

    return lis
}