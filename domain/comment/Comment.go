package comment

type Comment struct {
	Id       Id
	Nickname string
	Body     string
}

func NewComment(id Id, nickname string, body string) Comment {

	return Comment{
		Id:       id,
		Nickname: nickname,
		Body:     body,
	}
}

func (c *Comment) UpdateBody(body string) {
	c.Body = body
}
