package grpc

import (
	"context"
	"errors"

	pb "github.com/kodinggo/comment-service-gb1/pb/comment"
)

// FindComments is a function to find comments
func (s *CommentService) FindComments(ctx context.Context, in *pb.CommentRequest) (*pb.CommentList, error) {
	if in.StoryId == "" {
		return nil, errors.New("story_id is required")
	}

	comments, err := s.commentUsecase.FindComments(in.StoryId, "postgres")
	if err != nil {
		return nil, err
	}

	var commentPbs []*pb.Comment

	for _, comment := range comments {
		commentPb := &pb.Comment{
			Id:        comment.Id,
			StoryId:   comment.StoryId,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}

		commentPbs = append(commentPbs, commentPb)
	}

	return &pb.CommentList{
		Comments: commentPbs,
	}, nil
}
