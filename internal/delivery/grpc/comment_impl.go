package grpc

import (
	"context"
	"errors"

	pb "github.com/kodinggo/pb/comment"
)

// FindComments is a function to find comments
func (s *CommentService) FindComments(ctx context.Context, in *pb.CommentRequest) (*pb.CommentList, error) {
	if in.StoryId != "" {
		return nil, errors.New("story_id is required")
	}

	// hit usecase

	return &pb.CommentList{
		Comments: []*pb.Comment{
			{
				Id:      "1",
				StoryId: "1",
				Content: "This story is amazing",
			},
			{
				Id:      "2",
				StoryId: "1",
				Content: "Not my cup of tea",
			},
		},
	}, nil
}
