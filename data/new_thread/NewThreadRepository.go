package new_thread

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	domainNewThread "github.com/unabara-org/go-keijiban/domain/new_thread"
)

type NewThreadRepository struct {
	db *sql.DB
}

func NewNewThreadRepository(db *sql.DB) NewThreadRepository {
	return NewThreadRepository{db: db}
}

func (r NewThreadRepository) Create(newThread domainNewThread.NewThread) error {
	tx, err := r.db.Begin()

	if err != nil {
		return err
	}

	if _, err = tx.Exec("INSERT INTO threads(id, title) VALUES (?, ?)", newThread.Id.String(), newThread.Title); err != nil {
		_ = tx.Rollback()
		return err
	}

	if _, err = tx.Exec(
		"INSERT INTO comments(id, thread_id, nickname, body) VALUES (?, ?, ?, ?)",
		newThread.Comment.Id.String(),
		newThread.Id.String(),
		newThread.Comment.Nickname,
		newThread.Comment.Body,
	); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}
