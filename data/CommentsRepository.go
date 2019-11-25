package data

import (
	"../domain"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type RequestComment struct {
	Nickname string
	Body     string
}

func Create(requestComment RequestComment) (*domain.Comment, error) {
	db, err := sql.Open("mysql", "root:@/go_keijiban")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO comments SET nickname='?', body='?'", requestComment.Nickname, requestComment.Body)

	if err != nil {
		return nil, err
	}

	insertId, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &domain.Comment{
		Id:       uint(insertId),
		Nickname: requestComment.Nickname,
		Body:     requestComment.Body,
	}, nil
}
