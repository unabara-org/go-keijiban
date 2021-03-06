package data

import (
	"database/sql"
	"fmt"
	"os"
)

func NewDatabase() (*sql.DB, error) {
	// 下のようになる
	// "root:password@tcp(localhost:3306)/go_keijiban"
	source := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	return sql.Open("mysql", source)
}
