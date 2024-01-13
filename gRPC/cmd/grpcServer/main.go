package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/leonardogregoriocs/gRPC/internal/database"
	"github.com/leonardogregoriocs/gRPC/internal/pb"
	"github.com/leonardogregoriocs/gRPC/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}

	fmt.Println("Rodando na port 50051")
}
