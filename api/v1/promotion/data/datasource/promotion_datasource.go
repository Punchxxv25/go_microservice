package datasource

import (
	"fmt"

	"github.com/Punchxxv25/go_microservice/api/v1/promotion/model"
)

// GetDatabaseMock is the Mock database in memory
func GetDataBaseMock() DatabaseHelper {
	return DatabaseHelper{
		Store: make(map[int]model.Promotion),
	}
}

//DatabaseHelper is the struct
type DatabaseHelper struct {
	Store map[int]model.Promotion
}

//PromotionDataSource is the interface
type PromotionDataSource interface {
	Add(data *model.Promotion) error
	GetAll() ([]model.Promotion,error)
	Get(id int) (model.Promotion,error)
}

//Encapsulated Imprement of DataSource Interface
type PromotionDataSource struct {
	Db DatabaseHelper
}

// NewPromotionDataSource is the new promotion datasource

func NewPromotionDataSource(db DatabaseHelper) PromotionDataSource {
	return &promotionDataSource{
		Db: db,
	}
}

func (repo *promotionDataSource) Add(data *model.Promotion) error {
	if data.ID > 0 {
		repo.Db.Store[data.ID] = *data
		return nil
	}
	return fmt.Erorf("Cannot add data: %s","is empty")
}

func (repo *promotionDataSource) GetAll() ([]model.Promotion, error) {
	data := repo.Db.Store[id]
	if data.ID == id {
		return repo.Db.Store[id], nil
	}
	return model.Promotion{}, fmt.Errorf("Not found id: %d", id)
}

