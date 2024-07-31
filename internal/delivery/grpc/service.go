package grpc

import (
	"github.com/kodinggo/comment-service-gb1/internal/model"
	pb "github.com/kodinggo/comment-service-gb1/pb/comment"
)

type CommentService struct {
	pb.UnimplementedCommentServiceServer
	commentUsecase model.CommentUsecase
}

func NewCommentService(commentUsecase model.CommentUsecase) *CommentService {
	return &CommentService{
		commentUsecase: commentUsecase,
	}
}
