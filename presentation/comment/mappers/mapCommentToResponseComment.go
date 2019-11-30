package mappers

import (
	"github.com/unabara-org/go-keijiban/domain/comment"
	presentationComment "github.com/unabara-org/go-keijiban/presentation/comment"
)

func MapCommentToResponseComment(comment comment.Comment) presentationComment.ResponseComment {
	return presentationComment.ResponseComment{
		Id:       comment.Id.String(),
		Nickname: comment.Nickname,
		Body:     comment.Body,
	}
}
