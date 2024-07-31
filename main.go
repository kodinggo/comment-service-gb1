package main

import (
	"log"
	"net"

	grpcSvc "github.com/kodinggo/comment-service-gb1/internal/delivery/grpc"
	"github.com/kodinggo/comment-service-gb1/internal/repository"
	"github.com/kodinggo/comment-service-gb1/internal/usecase"
	pb "github.com/kodinggo/comment-service-gb1/pb/comment"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	commentRepo := repository.NewCommentRepository()
	commentUsecase := usecase.NewCommentUsecase(commentRepo)

	commentService := grpcSvc.NewCommentService(commentUsecase)

	pb.RegisterCommentServiceServer(s, commentService)

	listen, err := net.Listen("tcp", ":3100")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server at :3100")

	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
