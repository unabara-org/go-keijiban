package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/unabara-org/go-keijiban/domain"
)

type CommentsRepository struct {
	db *sql.DB
}

func NewCommentsRepository(db *sql.DB) CommentsRepository {
	return CommentsRepository{db: db}
}

func (commentsRepository CommentsRepository) Create(comment domain.Comment) error {
	_, err := commentsRepository.db.Exec("INSERT INTO comments SET id=?, nickname=?, body=?", comment.Id, comment.Nickname, comment.Body)

	if err != nil {
		return err
	}

	return nil
}

func (commentsRepository CommentsRepository) Read() ([]domain.Comment, error) {
	comments := []domain.Comment{}
	rows, err := commentsRepository.db.Query("SELECT id, nickname, body FROM comments")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		comment := domain.Comment{}

		if err := rows.Scan(&comment.Id, &comment.Nickname, &comment.Body); err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
