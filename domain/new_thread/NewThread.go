package new_thread

import (
	"github.com/unabara-org/go-keijiban/domain/comment"
	"github.com/unabara-org/go-keijiban/domain/thread"
)

type NewThread struct {
	Id      thread.Id
	Title   string
	Comment comment.Comment
}

func NewNewThread(id thread.Id, title string, comment comment.Comment) NewThread {
	return NewThread{
		Id:      id,
		Title:   title,
		Comment: comment,
	}
}
