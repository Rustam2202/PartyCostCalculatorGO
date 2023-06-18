package repository

import (
	"party-calc/internal/database"
	"party-calc/internal/database/models"
	"party-calc/internal/logger"

	"go.uber.org/zap"
)

type PersonRepository struct {
	db *database.DataBase
}

func NewPersonRepository(db *database.DataBase) *PersonRepository {
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
	return lastInsertedId, nil
}

func (r *PersonRepository) Get(p *models.Person) (models.Person, error) {
	var per models.Person
	err := r.db.DB.QueryRow(`SELECT * FROM persons WHERE name = $1`, p.Name).
		Scan(&per.Id, &per.Name)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from persons: ", zap.Error(err))
		return models.Person{}, err
	}
	return per, nil
}

func (r *PersonRepository) Update(p *models.Person) error {
	_, err := r.db.DB.Exec(
		`UPDATE persons SET name = $1 WHERE id = $2`,
		p.Name, p.Id)
	if err != nil {
		logger.Logger.Error("Failed to Execute Update operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *PersonRepository) Delete(p *models.Person) error {
	_, err := r.db.DB.Exec(`DELETE FROM persons WHERE name = $1`, p.Name)
	if err != nil {
		logger.Logger.Error("Failed to Execute Delete operation: ", zap.Error(err))
		return err
	}
	return nil
}
