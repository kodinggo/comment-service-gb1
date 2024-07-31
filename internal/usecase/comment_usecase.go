package usecase

import "github.com/kodinggo/comment-service-gb1/internal/model"

type commentUsecase struct {
	commentRepo model.CommentRepository
}

func NewCommentUsecase(cr model.CommentRepository) model.CommentUsecase {
	return &commentUsecase{
		commentRepo: cr,
	}
}

func (cu *commentUsecase) FindComments(storyId string, from string) ([]model.Comment, error) {
	if from == "elastic" {
		return cu.commentRepo.FindCommentFromDocument(storyId)
	}

	return cu.commentRepo.FindComments(storyId)
}
