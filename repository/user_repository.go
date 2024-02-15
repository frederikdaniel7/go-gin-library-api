package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/exception"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/utils"
)

type UserRepository interface {
	FindAll() ([]entity.User, error)
	FindSimilarUserByName(name string) ([]entity.User, error)
	FindUserById(id int64) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	CreateUser(body dto.CreateUserBody) (*entity.User, error)
	FindUserPassword(body dto.LoginBody) (string, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(body dto.CreateUserBody) (*entity.User, error) {
	user := entity.User{}

	hashedPassword, err := utils.HashPassword(body.Password, 12)
	if err != nil {
		return nil, err
	}

	var sb strings.Builder
	sb.WriteString("INSERT INTO users (user_name, email, user_password, phone)")
	sb.WriteString("VALUES (")
	for i := 1; i <= constant.LenCreateUserBody; i++ {
		sb.WriteString("$" + fmt.Sprintf("%d", i))
		if i != constant.LenCreateUserBody {
			sb.WriteString(",")
		}
	}
	sb.WriteString(")returning id, user_name, email, phone, created_at, updated_at")
	row := r.db.QueryRow(sb.String(), body.Name, body.Email, hashedPassword, body.Phone)
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if row == nil {
		return nil, exception.NewErrorType(http.StatusInternalServerError, constant.ResponseMsgErrorInternal)
	}
	if err != nil {
		return nil, exception.NewErrorType(http.StatusInternalServerError, constant.ResponseMsgErrorInternal)
	}
	return &user, nil
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
	 from users u where u.id = $1 `

	row := r.db.QueryRow(q, id)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, errors.New("error query")
	}
	return &user, nil

}

func (r *userRepository) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	q := `SELECT u.id, u.user_name, u.email,u.phone,u.created_at, u.updated_at
	 from users u where u.email = $1 `

	row := r.db.QueryRow(q, email)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, nil
		}
		return nil, exception.NewErrorType(http.StatusInternalServerError, constant.ResponseMsgErrorInternal)
	}
	return &user, nil
}

func (r *userRepository) FindUserPassword(body dto.LoginBody) (string, error) {

	q := `SELECT user_password from users where email = $1 `

	row := r.db.QueryRow(q, body.Email)
	var password string
	err := row.Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return password, nil
		}
		return "", exception.NewErrorType(http.StatusUnauthorized, constant.ResponseMsgErrorCredentials)
	}
	return password, nil

}
