package repository

import (
	"database/sql"
	"errors"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/entity"
)

type UserRepository interface {
	FindAll() ([]entity.User, error)
	FindSimilarUserByName(name string) ([]entity.User, error)
	FindUserById(id int64) (*entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindAll() ([]entity.User, error) {
	users := []entity.User{}

	q := `SELECT u.id, u.user_name, u.email, u.phone, u.created_at, u.updated_at, u.deleted_at from users u`

	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := entity.User{}

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindSimilarUserByName(name string) ([]entity.User, error) {
	users := []entity.User{}

	q := `SELECT u.id, u.user_name, u.email, u.phone, u.created_at, u.updated_at, u.deleted_at from users u
	where u.user_name ILIKE '%' ||$1|| '%'`

	rows, err := r.db.Query(q, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := entity.User{}

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil

}

func (r *userRepository) FindUserById(id int64) (*entity.User, error) {
	var user entity.User

	q := `SELECT u.id, u.user_name, u.email,u.phone,u.created_at, u.updated_at
	 from users u where u.id = $1`

	row := r.db.QueryRow(q, id)
	if row == nil {
		return nil, errors.New("error query")
	}
	row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)

	return &user, nil

}
