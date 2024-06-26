package utils

import (
	"github.com/frederikdaniel7/go-gin-library-api/dto"
	"github.com/frederikdaniel7/go-gin-library-api/entity"
)

func ConvertBookToJson(book entity.Book) dto.Book {
	converted := dto.Book{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Quantity:    book.Quantity,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}

	if book.DeletedAt.Valid {
		converted.DeletedAt = &book.DeletedAt.Time
	}
	if book.Cover.Valid {
		converted.Cover = &book.Cover.String
	}
	return converted
}

func ConvertBookDetailToJson(book entity.BookDetail) dto.BookDetail {
	converted := dto.BookDetail{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Quantity:    book.Quantity,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}

	if book.DeletedAt.Valid {
		converted.DeletedAt = &book.DeletedAt.Time
	}
	if book.Cover.Valid {
		converted.Cover = &book.Cover.String
	}
	if book.Author != nil {
		converted.Author = &dto.Author{
			ID:   book.Author.ID,
			Name: book.Author.Name,
		}
	}

	return converted
}

func ConvertUserToJson(user entity.User) dto.User {
	converted := dto.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	if user.DeletedAt.Valid {
		converted.DeletedAt = &user.DeletedAt.Time
	}
	return converted
}

func ConvertBorrowRecordToJson(record entity.BorrowRecord) dto.BorrowRecord {
	converted := dto.BorrowRecord{
		ID:            record.ID,
		UserID:        record.UserID,
		BookID:        record.BookID,
		Status:        record.Status,
		BorrowingDate: record.BorrowingDate,
		CreatedAt:     record.CreatedAt,
		UpdatedAt:     record.UpdatedAt,
	}

	if record.DeletedAt.Valid {
		converted.DeletedAt = &record.DeletedAt.Time
	}
	if record.ReturningDate.Valid {
		converted.ReturningDate = &record.ReturningDate.Time
	}

	return converted
}
