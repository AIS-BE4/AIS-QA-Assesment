package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "ais-qa-assesment/gen"
	"ais-qa-assesment/lib/database_transaction"
	"ais-qa-assesment/stores/repository"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

const (
	defaultPort = "60002"
)

type Server struct {
	db                 *gorm.DB
	repository         repository.RepoMethod
	transactionManager database_transaction.Client
}

func NewServer(ctx context.Context, db *gorm.DB, repository repository.RepoMethod, transactionManager database_transaction.Client) (*Server, error) {
	return &Server{
		db:                 db,
		repository:         repository,
		transactionManager: transactionManager,
	}, nil
}

func (s *Server) Run() {
	go func() {
		mux := runtime.NewServeMux()

		pb.RegisterStoresAPIHandlerServer(context.Background(), mux, s)

		log.Printf("Starting Stores server on port %s", "8003")
		log.Fatalln(http.ListenAndServe("0.0.0.0:8003", mux))
	}()

	port := "8000"
	if port == "" {
		port = defaultPort
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Print("net.Listen failed")
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterStoresAPIServer(grpcServer, s) // use authogenerated code to register the server
	reflection.Register(grpcServer)

	log.Printf("Starting Stores server on port %s", port)
	go func() {
		grpcServer.Serve(listener)
	}()
}
