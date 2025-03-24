package repository

import "database/sql"

type User struct {
	ID       int
	Username string
	Email    string
}

type UserRepository interface {
	GetByID(id int) (*User, error)
	Create(user *User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByID(id int) (*User, error) {
	user := &User{}
	query := "SELECT id, username, email FROM users WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Create(user *User) error {
	query := "INSERT INTO users (username, email) VALUES (?, ?)"
	_, err := r.db.Exec(query, user.Username, user.Email)
	return err
}
