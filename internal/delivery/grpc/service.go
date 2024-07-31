package grpc

import (
	pb "github.com/kodinggo/pb/comment"
)

type CommentService struct {
	pb.UnimplementedCommentServiceServer
}

func NewCommentService() *CommentService {
	return &CommentService{}
}
