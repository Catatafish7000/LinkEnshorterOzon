package main

import (
	"LinkEnshorter/internal/controller"
	"LinkEnshorter/internal/generator"
	"LinkEnshorter/internal/service"
	"LinkEnshorter/internal/strategy"
	"LinkEnshorter/pb"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runRest(ctx context.Context) {

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterUrlShorterHandlerFromEndpoint(ctx, mux, "localhost:12201", opts)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}

func runGrpc(ctx context.Context, server pb.UrlShorterServer) {
	port := 12201
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUrlShorterServer(s, server)
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		panic(err)
	}

}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	repo, err := strategy.RepoStrategy(os.Args)
	if err != nil {
		log.Fatalf("failed to init repo. err:%v", err)
	}
	gen := generator.NewGenerator()
	usecase := service.NewService(repo, gen)
	server := controller.NewServer(usecase)
	go runRest(ctx)
	runGrpc(ctx, server)
}
