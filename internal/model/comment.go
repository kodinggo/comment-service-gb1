package model

type Comment struct {
	Id        string `json:"id"`
	StoryId   string `json:"story_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CommentUsecase interface {
	FindComments(storyID string, from string) ([]Comment, error)
}

type CommentRepository interface {
	FindComments(storyID string) ([]Comment, error)
	FindCommentFromDocument(storyID string) ([]Comment, error)
}
