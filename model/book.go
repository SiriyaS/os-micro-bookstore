package model

import (
	"os-micro-bookstore/database"
	"os-micro-bookstore/form"
)

type BookModel struct{}

func (bm BookModel) ReadAll() ([]form.Book, error) {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return nil, err
	}
	defer database.CloseConnection(conn)

	var books []form.Book
	if err = conn.
		Table("books b").
		Select("b.ISBN, b.name, a.name AS author, b.unit_price, b.publish_year, p.name AS publisher, edition, c.name AS category").
		Joins("INNER JOIN authors a ON b.author = a.id").
		Joins("INNER JOIN publishers p ON b.publisher = p.id").
		Joins("INNER JOIN category c ON b.category = c.id ").
		Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (bm BookModel) ReadByID(isbn string) (form.Book, error) {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return form.Book{}, err
	}
	defer database.CloseConnection(conn)

	var book form.Book
	if err = conn.
		Table("books b").
		Select("b.ISBN, b.name, a.name AS author, b.unit_price, b.publish_year, p.name AS publisher, edition, c.name AS category").
		Joins("INNER JOIN authors a ON b.author = a.id").
		Joins("INNER JOIN publishers p ON b.publisher = p.id").
		Joins("INNER JOIN category c ON b.category = c.id ").
		Where("b.isbn = ?", isbn).
		Find(&book).Error; err != nil {
		return form.Book{}, err
	}

	return book, nil
}

func (bm BookModel) ReadByCategoryName(category string) ([]form.Book, error) {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return nil, err
	}
	defer database.CloseConnection(conn)

	var books []form.Book
	subQuery := conn.Table("category c").Select("id").Where("c.name = ?", category)
	if err = conn.
		Table("books b").
		Select("b.ISBN, b.name, a.name AS author, b.unit_price, b.publish_year, p.name AS publisher, edition, c.name AS category").
		Joins("INNER JOIN authors a ON b.author = a.id").
		Joins("INNER JOIN publishers p ON b.publisher = p.id").
		Joins("INNER JOIN category c ON b.category = c.id ").
		Where("b.category = (?)", subQuery).
		Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (bm BookModel) ReadByAuthorName(author string) ([]form.Book, error) {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return nil, err
	}
	defer database.CloseConnection(conn)

	var books []form.Book
	subQuery := conn.Table("authors a").Select("id").Where("a.name LIKE ?", "%"+author+"%")
	if err = conn.
		Table("books b").
		Select("b.ISBN, b.name, a.name AS author, b.unit_price, b.publish_year, p.name AS publisher, edition, c.name AS category").
		Joins("INNER JOIN authors a ON b.author = a.id").
		Joins("INNER JOIN publishers p ON b.publisher = p.id").
		Joins("INNER JOIN category c ON b.category = c.id ").
		Where("b.author = (?)", subQuery).
		Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (bm BookModel) ReadByPublisherName(publisher string) ([]form.Book, error) {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return nil, err
	}
	defer database.CloseConnection(conn)

	var books []form.Book
	subQuery := conn.Table("publishers p").Select("id").Where("p.name LIKE ?", "%"+publisher+"%")
	if err = conn.
		Table("books b").
		Select("b.ISBN, b.name, a.name AS author, b.unit_price, b.publish_year, p.name AS publisher, edition, c.name AS category").
		Joins("INNER JOIN authors a ON b.author = a.id").
		Joins("INNER JOIN publishers p ON b.publisher = p.id").
		Joins("INNER JOIN category c ON b.category = c.id ").
		Where("b.publisher = (?)", subQuery).
		Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (bm BookModel) Add(book form.BookRequest) (err error) {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return err
	}
	defer database.CloseConnection(conn)

	if err = conn.
		Table("books").
		Create(&book).Error; err != nil {
		return
	}

	return
}

func (bm BookModel) UpdateByID(isbn string, book form.BookRequest) (err error) {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return err
	}
	defer database.CloseConnection(conn)

	if book.Name != "" {
		if err = conn.
			Table("books").
			Where("isbn = ?", isbn).
			Update("name", book.Name).Error; err != nil {
			return
		}
	}
	if book.Author != 0 {
		if err = conn.
			Table("books").
			Where("isbn = ?", isbn).
			Update("author", book.Author).Error; err != nil {
			return
		}
	}
	if book.UnitPrice != 0 {
		if err = conn.
			Table("books").
			Where("isbn = ?", isbn).
			Update("unit_price", book.UnitPrice).Error; err != nil {
			return
		}
	}
	if book.PublishYear != 0 {
		if err = conn.
			Table("books").
			Where("isbn = ?", isbn).
			Update("publish_year", book.PublishYear).Error; err != nil {
			return
		}
	}
	if book.Publisher != 0 {
		if err = conn.
			Table("books").
			Where("isbn = ?", isbn).
			Update("publisher", book.Publisher).Error; err != nil {
			return
		}
	}
	if book.Edition != 0 {
		if err = conn.
			Table("books").
			Where("isbn = ?", isbn).
			Update("edition", book.Edition).Error; err != nil {
			return
		}
	}
	if book.Category != 0 {
		if err = conn.
			Table("books").
			Where("isbn = ?", isbn).
			Update("category", book.Category).Error; err != nil {
			return
		}
	}

	return nil
}

func (bm BookModel) DeleteByID(isbn string) (err error) {
	// connect to database
	conn, err := database.NewConnection()
	if err != nil {
		return err
	}
	defer database.CloseConnection(conn)

	if err = conn.
		Table("books").
		Delete(&form.BookRequest{}, isbn).Error; err != nil {
		return
	}

	return
}
