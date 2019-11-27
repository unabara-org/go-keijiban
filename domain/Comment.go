package domain

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

type Comment struct {
	Id       string
	Nickname string
	Body     string
}

func NewComment(id string, nickname string, body string) Comment {

	return Comment{
		Id:       id,
		Nickname: nickname,
		Body:     body,
	}
}

func NewCommentId() string {
	// ここ見たけど、まあそんなもんかって理解でちゃんと分かってない
	// http://tocsato.hatenablog.com/entry/2016/09/06/080743
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	return id.String()
}
