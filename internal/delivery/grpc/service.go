package grpc

import (
	pb "github.com/kodinggo/comment-service-gb1/pb/comment"
)

type CommentService struct {
	pb.UnimplementedCommentServiceServer
}

func NewCommentService() *CommentService {
	return &CommentService{}
}
