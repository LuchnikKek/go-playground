package processors

import (
	"context"
	"errors"
	"web-server/internals/app/db"
	"web-server/internals/app/models"
)

type CarsProcessor struct {
	storage *db.CarsStorage
}

func NewCarsProcessor(storage *db.CarsStorage) *CarsProcessor {
	processor := new(CarsProcessor)
	processor.storage = storage
	return processor
}

func (processor *CarsProcessor) CreateCar(ctx context.Context, car models.Car) error {
	if car.Colour == "" {
		return errors.New("colour can not be empty")
	}
	if car.Brand == "" {
		return errors.New("brand can not be empty")
	}
	if car.Owner.Id < 0 {
		return errors.New("user must be filled")
	}

	return processor.storage.CreateCar(ctx, car)
}

func (processor *CarsProcessor) FindCar(ctx context.Context, id int64) (models.Car, error) {
	car := processor.storage.GetCarById(ctx, id)

	if car.Id != id {
		return car, errors.New("car not found")
	}

	return car, nil
}

func (processor *CarsProcessor) ListCars(ctx context.Context, userId int64, brandFilter string, colourFilter string, licenseFilter string) ([]models.Car, error) {
	return processor.storage.GetCarsList(ctx, userId, brandFilter, colourFilter, licenseFilter), nil
}
