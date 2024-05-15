package database

import (
	"crypto/subtle"
)

func (s *Storage) ValidAuthor(username string, password []byte) bool {
	var validPassword []byte

	const query = "SELECT password FROM author WHERE username = ?"
	err := s.db.QueryRow(query, username).Scan(&validPassword)
	if err != nil {
		return false
	}

	if subtle.ConstantTimeCompare(validPassword, password) == 0 {
		return false
	}

	return true
}
