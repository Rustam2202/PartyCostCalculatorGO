package database

import (
	"party-calc/internal/database/models"
	"party-calc/internal/logger"

	"go.uber.org/zap"
)

type PersonRepository struct {
	db *DataBase
}

func NewPersonRepository(db *DataBase) *PersonRepository {
	return &PersonRepository{db: db}
}

func (r *PersonRepository) Create(p *models.Person) (int64, error) {
	var lastInsertedId int64
	err := r.db.DB.QueryRow(`INSERT INTO persons (name) VALUES($1) RETURNING Id`, p.Name).
		Scan(&lastInsertedId)
	if err != nil {
		logger.Logger.Error("Failed to Execute Insert to 'persons' table: ", zap.Error(err))
		return 0, err
	}
	//p.Id = lastInsertedId
	return lastInsertedId, nil
}
