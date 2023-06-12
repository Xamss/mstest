package usecase

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	entities "xamss/microservices/test/services/contact/internal/domain"
	"xamss/microservices/test/services/contact/internal/repository"
)

type ContactModel struct {
	pool *pgxpool.Pool
	repo repository.Repo
}

type Contact interface {
	Insert(contact *entities.Contact) error
	Get(id int64) (*entities.Contact, error)
	Update(contact *entities.Contact) error
	Delete(id int64) error
}

func (c ContactModel) Insert(contact *entities.Contact) error {

	query := `
		INSERT INTO contacts (firstname, surname, patronymic, phone)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	args := []interface{}{contact.Value().Firstname, contact.Value().Surname, contact.Value().Patronymic, contact.Phone}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.pool.QueryRow(ctx, query, args...).Scan(&contact.ID)
}

func (c ContactModel) Get(id int64) (*entities.Contact, error) {
	query := `
	SELECT id, firstname, surname, patronymic, phone
	FROM contacts
	WHERE id = $1`
	var contact entities.Contact
	var fullname entities.Fullname
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := c.pool.QueryRow(ctx, query, id).Scan(
		&contact.ID,
		&fullname.Firstname,
		&fullname.Surname,
		&fullname.Patronymic,
		&contact.Phone,
	)

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, pgx.ErrNoRows
		default:
			return nil, err
		}
	}
	return &contact, nil
}

func (c ContactModel) Update(contact *entities.Contact) error {
	query := `
	UPDATE movies
	SET firstname = $1, surname = $2, patronymic = $3, phone = $4
	WHERE id = $5`

	var fullname entities.Fullname
	args := []interface{}{
		fullname.Firstname,
		fullname.Surname,
		fullname.Patronymic,
		contact.Phone,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := c.pool.QueryRow(ctx, query, args...).Scan()
	if err != nil {
		return err
	}

	return nil
}

func (c ContactModel) Delete(id int64) error {
	query := `
	DELETE FROM contacts
	WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := c.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
