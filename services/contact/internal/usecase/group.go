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

type GroupModel struct {
	repo repository.Repo
	pool *pgxpool.Pool
}

type Group interface {
	Insert(groups *entities.Group) error
	Get(id int64) (*entities.Group, error)
	AddContact(contact *entities.Contact) error
}

func (g GroupModel) Insert(groups *entities.Group) error {

	query := `
		INSERT INTO groups (name)
		VALUES ($1)
		RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return g.pool.QueryRow(ctx, query, groups.Name).Scan(&groups.ID)
}

func (g GroupModel) Get(id int64) (*entities.Group, error) {
	query := `
	SELECT id, name, contact
	FROM groups inner join contacts on contacts.id = groups.contactID  
	WHERE id = $1`

	var group entities.Group

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := g.pool.QueryRow(ctx, query, id).Scan(
		&group.ID,
		&group.Contact,
		&group.Name,
	)

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, pgx.ErrNoRows
		default:
			return nil, err
		}
	}
	return &group, nil
}

func (g GroupModel) AddContact(contact *entities.Contact) error {
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

	err := g.pool.QueryRow(ctx, query, args...).Scan()
	if err != nil {
		return err
	}

	return nil
}
