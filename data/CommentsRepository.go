package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/unabara-org/go-keijiban/domain"
)

type CommentsRepository struct {
}

func NewCommentsRepository() CommentsRepository {
	return CommentsRepository{}
}

func (_ CommentsRepository) Create(comment domain.Comment) error {
	db, err := NewDatabase()
	if err != nil {
		panic(err.Error())
	}

	// https://github.com/go-sql-driver/mysql/issues/910 を参考にした
	defer func() {
		if err := db.Close(); err != nil {
			// error handle
		}
	}()

	_, err = db.Exec("INSERT INTO comments SET id=?, nickname=?, body=?", comment.Id, comment.Nickname, comment.Body)

	if err != nil {
		return err
	}

	return nil
}
