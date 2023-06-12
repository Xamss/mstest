package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	domain "xamss/microservices/test/services/contact/internal/domain"
)

type Repo interface {
	CreateContact(user *domain.Contact) (int, error)
	ReadContact(id int) (*domain.Contact, error)
	UpdateContact(id int) error
	DeleteContact(id int) error
	CreateGroup(group *domain.Group) (int, error)
	ReadGroup(id int) (*domain.Group, error)
	AddContact(id int) error
}

type repo struct {
	DB *pgxpool.Pool
}

func New(db *pgxpool.Pool) *repo {
	return &repo{DB: db}
}

func (r *repo) CreateContact(user *domain.Contact) (int, error) {
	return 0, nil
}

func (r *repo) ReadContact(id int) (*domain.Contact, error) {
	return nil, nil
}

func (r *repo) UpdateContact(id int) error {
	return nil
}

func (r *repo) DeleteContact(id int) error {
	return nil
}

func (r *repo) CreateGroup(group *domain.Group) (int, error) {
	return 0, nil
}

func (r *repo) ReadGroup(id int) (*domain.Group, error) {
	return nil, nil
}

func (r *repo) AddContact(id int) error {
	return nil
}
