package mysql

import (
	"database/sql"

	"abahjoseph.com/books/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) GetAll() ([]*models.User, error) {
	sql := `SELECT id, first_name, last_name, email FROM users`

	rows, err := m.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []*models.User{}
	for rows.Next() {
		b := &models.User{}
		err := rows.Scan(&b.ID, &b.FirstName, &b.LastName, &b.Email)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}
