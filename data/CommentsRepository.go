package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	domainComment "github.com/unabara-org/go-keijiban/domain/comment"
)

type CommentsRepository struct {
	db *sql.DB
}

type queryRecord struct {
	Id string
	Nickname string
	Body string
	PostedAt string
}

func NewCommentsRepository(db *sql.DB) CommentsRepository {
	return CommentsRepository{db: db}
}

func (r CommentsRepository) Create(comment domainComment.Comment) error {
	_, err := r.db.Exec("INSERT INTO comments SET id=?, nickname=?, body=?", comment.Id.String(), comment.Nickname, comment.Body)

	if err != nil {
		return err
	}

	return nil
}

func (r CommentsRepository) Read() ([]domainComment.Comment, error) {
	var comments []domainComment.Comment
	rows, err := r.db.Query("SELECT id, nickname, body FROM comments")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var record queryRecord

		if err := rows.Scan(&record.Id, &record.Nickname, &record.Body); err != nil {
			return nil, err
		}

		comment, err := mapQueryRecordToComment(record)

		if err != nil {
			return nil, err
		}

		comments = append(comments, *comment)
	}

	return comments, nil
}

func (r CommentsRepository) Find(id domainComment.Id) (*domainComment.Comment, error) {
	rows, err := r.db.Query("SELECT id, nickname, body FROM comments WHERE id=?", id.String())

	if err != nil {
		return nil, err
	}

	rows.Next()
	var record queryRecord

	if err := rows.Scan(&record.Id, &record.Nickname, &record.Body); err != nil {
		return nil, err
	}

	comment, err := mapQueryRecordToComment(record)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func mapQueryRecordToComment(record queryRecord) (*domainComment.Comment, error) {
	id, err := domainComment.NewIdFromString(record.Id)

	if err != nil {
		return nil, err
	}

	return &domainComment.Comment{
		Id: *id,
		Nickname: record.Nickname,
		Body: record.Body,
	}, nil
}

func (r CommentsRepository) Update(comment domainComment.Comment) error {
	result, err := r.db.Exec("UPDATE comments SET nickname=?, body=? WHERE id=?", comment.Nickname, comment.Body, comment.Id.String())

	if err != nil {
		return err
	}

	var rows int64
	if rows, err = result.RowsAffected(); err != nil {
		if rows == 0 {
			return nil
		}
	}

	return err
}
