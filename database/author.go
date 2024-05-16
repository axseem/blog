package database

import (
	"golang.org/x/crypto/bcrypt"
)

func (s *Storage) ValidAuthor(username string, password string) error {
	var validPasswordHash []byte

	const query = "SELECT password FROM author WHERE username = ?"
	err := s.db.QueryRow(query, username).Scan(&validPasswordHash)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword(validPasswordHash, []byte(password)); err != nil {
		return err
	}

	return nil
}

func (s *Storage) CreateAuthor(username string, password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		return err
	}

	const query = "INSERT INTO author (username, password) VALUES (?, ?)"
	_, err = s.db.Exec(query, username, passwordHash)
	return err
}
