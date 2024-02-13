package utils

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/entity"
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
	if book.Cover.Valid{
		converted.Cover = book.Cover.String
	}
	return converted
}
