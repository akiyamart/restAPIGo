package queries

import (
    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
    "github.com/koddr/tutorial-go-fiber-rest-api/app/models"
)

// BookQueries struct for queries from Book model.
type BookQueries struct { 
	*sqlx.DB
}

// GetBooks method for getting all books.
func (q *BookQueries) GetBooks() ([]models.Book, error) { 
	books := []models.Book{}

	query := `SELECT * FROM books`

	err := q.Get(&books, query)
	if err != nil { 
		return books, err
	}

	return books, nil 
}

// GetBook method for getting one book by given ID.
func (q *BookQueries) GetBook(id uuid.UUID) (models.Book, error) { 
	book := models.Book{}

	query := `SELECT * FROM books WHERE id = $1`
	err := q.Get(&book, query, id)
	if err != nil { 
		return book, err
	}

	return book, nil
}

// CreateBook method for creating book by given Book object.
func (q *BookQueries) CreateBook(b *models.Book) error { 
	query := `INSERT INTO books VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := q.Exec(
		query, b.ID, b.CreatedAt,
		b.UpdatedAt, b.UserID, b.Title,
		b.Author, b.BookStatus, b.BookAttrs,
	)

	if err != nil { 
		return err 
	}

	return nil 
}

// UpdateBook method for updating book by given Book object.
func (q *BookQueries) UpdateBook(id uuid.UUID, b *models.Book) error { 
    query := `UPDATE books SET updated_at = $2, title = $3, author = $4, book_status = $5, book_attrs = $6 WHERE id = $1`

	_, err := q.Exec(
		query, id, b.UpdatedAt,
		b.Title, b.Author, b.BookStatus, b.BookAttrs,
	)
	if err != nil { 
		return err
	}

	return nil 
}

// DeleteBook method for delete book by given ID.
func (b *BookQueries) DeleteBook(id uuid.UUID) error { 
	query := `DELETE FROM books WHERE id = $1`

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}

	return nil 
}

