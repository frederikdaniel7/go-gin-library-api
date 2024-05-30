package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/frederikdaniel7/go-gin-library-api/constant"
	"github.com/frederikdaniel7/go-gin-library-api/dto"
	"github.com/frederikdaniel7/go-gin-library-api/entity"
	"github.com/frederikdaniel7/go-gin-library-api/exception"
	"github.com/frederikdaniel7/go-gin-library-api/utils"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]entity.User, error)
	FindSimilarUserByName(ctx context.Context, name string) ([]entity.User, error)
	FindUserById(ctx context.Context, id int64) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, body dto.CreateUserBody) (*entity.User, error)
	FindUserPassword(ctx context.Context, body dto.LoginBody) (string, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, body dto.CreateUserBody) (*entity.User, error) {
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
	row := r.db.QueryRowContext(ctx, sb.String(), body.Name, body.Email, hashedPassword, body.Phone)
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if row == nil {
		return nil, exception.NewErrorType(http.StatusInternalServerError, constant.ResponseMsgErrorInternal)
	}
	if err != nil {
		return nil, exception.NewErrorType(http.StatusInternalServerError, constant.ResponseMsgErrorInternal)
	}
	return &user, nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	users := []entity.User{}

	q := `SELECT u.id, u.user_name, u.email, u.phone, u.created_at, u.updated_at, u.deleted_at from users u`

	rows, err := r.db.QueryContext(ctx, q)
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

func (r *userRepository) FindSimilarUserByName(ctx context.Context, name string) ([]entity.User, error) {
	users := []entity.User{}

	q := `SELECT u.id, u.user_name, u.email, u.phone, u.created_at, u.updated_at, u.deleted_at from users u
	where u.user_name ILIKE '%' ||$1|| '%'`

	rows, err := r.db.QueryContext(ctx, q, name)
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

func (r *userRepository) FindUserById(ctx context.Context, id int64) (*entity.User, error) {
	var user entity.User

	q := `SELECT u.id, u.user_name, u.email,u.phone,u.created_at, u.updated_at
	 from users u where u.id = $1 `

	row := r.db.QueryRowContext(ctx, q, id)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, errors.New("error query")
	}
	return &user, nil

}

func (r *userRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	q := `SELECT u.id, u.user_name, u.email,u.phone,u.created_at, u.updated_at
	 from users u where u.email = $1 `

	row := r.db.QueryRowContext(ctx, q, email)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, nil
		}
		return nil, exception.NewErrorType(http.StatusInternalServerError, constant.ResponseMsgErrorInternal)
	}
	return &user, nil
}

func (r *userRepository) FindUserPassword(ctx context.Context, body dto.LoginBody) (string, error) {

	q := `SELECT user_password from users where email = $1 `

	row := r.db.QueryRowContext(ctx, q, body.Email)
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
