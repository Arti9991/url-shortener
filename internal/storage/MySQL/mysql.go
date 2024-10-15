package mysql

import (
	"database/sql"
	"fmt"
	"url-shortener/internal/storage"

	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	db *sql.DB
}

func NewStorage() (*Storage, error) {

	const op = "storage.mysql.New"

	const connStr = "user:password@tcp(localhost:3306)/url"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	query := `
    CREATE TABLE IF NOT EXISTS urls(
        id INT AUTO_INCREMENT PRIMARY KEY,
        alias VARCHAR(100) UNIQUE,
		url TEXT NOT NULL
    );`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) SaveURL(urlToSave string, alias string) error {
	const op = "storage.mysql.SaveURL"

	query := `INSERT IGNORE INTO urls(alias, url) VALUES (?, ?);`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(alias, urlToSave)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	const op = "storage.mysql.SaveURL"
	var str string
	query := `SELECT url FROM urls 
	WHERE alias = ?`

	row := s.db.QueryRow(query, alias)
	err := row.Scan(&str)
	if err == sql.ErrNoRows {
		return "", storage.ErrURLNotFound
	} else if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return str, nil
}

func (s *Storage) DeleteURL(alias string) error {
	const op = "storage.mysql.SaveURL"

	query := `DELETE FROM urls WHERE alias = ?`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(alias)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
