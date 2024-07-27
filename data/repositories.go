package data

import (
	"errors"
	"math/rand"
)

type userRepository struct {
	db Database
}

type UserRepository interface {
	Checkuser(username string, password string) error
}

func NewUserRepository(db Database) *userRepository {
	return &userRepository{db}
}

func (repo *userRepository) Checkuser(username string, password string) error {
	var user *User
	repo.db.GetInstance().Model(&User{}).Where("username = ? AND password = ?", username, password).First(&user)
	if user.ID == 0 {
		return errors.New("NO USER EXISTS")
	}
	return nil
}

type quoteRepository struct {
	db Database
}

type QuoteRepository interface {
	GetQuotes() *Quote
}

func NewQuoteRepository(db Database) *quoteRepository {
	return &quoteRepository{db}
}

func (repo *quoteRepository) GetQuotes() *Quote {
	id := rand.Int63()
	var count int64
	var quote *Quote
	repo.db.GetInstance().Model(&Quote{}).Count(&count)
	repo.db.GetInstance().Find(&quote, id%count)
	return quote
}
