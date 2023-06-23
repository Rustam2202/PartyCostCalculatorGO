package repository

import (
	"database/sql"
	"errors"
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

func (r *PersonRepository) Create(per *models.Person) (int64, error) {
	var lastInsertedId int64
	err := r.db.DB.QueryRow(`INSERT INTO persons (name) VALUES($1) RETURNING Id`, per.Name).
		Scan(&lastInsertedId)
	if err != nil {
		logger.Logger.Error("Failed to Execute Insert to 'persons' table: ", zap.Error(err))
		return 0, err
	}
	return lastInsertedId, nil
}

func (r *PersonRepository) Get(per *models.Person) (models.Person, error) {
	var result models.Person
	var row *sql.Row
	if per.Id != 0 {
		row = r.db.DB.QueryRow(`SELECT * FROM persons WHERE id = $1`, per.Id)
	} else if per.Name != "" {
		row = r.db.DB.QueryRow(`SELECT * FROM persons WHERE name = $1`, per.Name)
	} else {
		return models.Person{}, errors.New("empty input Person model")
	}
	err := row.Scan(&result.Id, &result.Name)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from persons: ", zap.Error(err))
		return models.Person{}, err
	}
	return result, nil
}

func (r *PersonRepository) Update(perOld, perNew *models.Person) error {
	_, err := r.db.DB.Exec(
		`UPDATE persons SET name=$1 WHERE id=$2`,
		perNew.Name, perOld.Id)
	if err != nil {
		logger.Logger.Error("Failed to Execute Update operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *PersonRepository) Delete(per *models.Person) error {
	_, err := r.db.DB.Exec(`DELETE FROM persons WHERE name=$1`, per.Name)
	if err != nil {
		logger.Logger.Error("Failed to Execute Delete operation: ", zap.Error(err))
		return err
	}
	return nil
}

