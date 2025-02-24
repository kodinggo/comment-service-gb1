package repository

import "github.com/kodinggo/comment-service-gb1/internal/model"

type commentRepository struct {
}

func NewCommentRepository() model.CommentRepository {
	return &commentRepository{}
}

func (cr *commentRepository) FindComments(storyID string) ([]model.Comment, error) {
	var comments []model.Comment

	// select * from comments where story_id = 3
	if storyID != "3" {
		return comments, nil
	}

	comments = getDataFromDB()
	return comments, nil
}

func (cr *commentRepository) FindCommentFromDocument(storyID string) ([]model.Comment, error) {
	var comments []model.Comment

	// get from no sql, elastic search
	if storyID != "3" {
		return comments, nil
	}

	comments = getDataFromDB()
	return comments, nil
}

func getDataFromDB() []model.Comment {
	return []model.Comment{
		{
			Id:        "1",
			StoryId:   "1",
			Content:   "I like this story",
			CreatedAt: "2021-01-01T00:00:00Z",
			UpdatedAt: "2021-01-01T00:00:00Z",
		},
		{
			Id:        "2",
			StoryId:   "1",
			Content:   "Not my cup of tea",
			CreatedAt: "2021-01-01T00:00:00Z",
			UpdatedAt: "2021-01-01T00:00:00Z",
		},
	}
}
