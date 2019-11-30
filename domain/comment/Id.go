package comment

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

type Id struct {
	id string
}

func (c Id) String() string {
	return c.id
}

func NewId() Id {
	// ここ見たけど、まあそんなもんかって理解でちゃんと分かってない
	// http://tocsato.hatenablog.com/entry/2016/09/06/080743
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	return Id{id: id.String()}
}

func NewIdFromString(id string) (*Id, error) {
	if _, err := ulid.Parse(id); err != nil {
		return nil, err
	}

	return &Id{id: id}, nil
}
