package repository

import (
	"github.com/trombettamoacyr/garage-api/entity"

	"cloud.google.com/go/firestore"
	"context"
	"github.com/google/uuid"
	"log"
)

type CarRepository interface {
	Save(car *entity.Car) (*entity.Car, error)
	FindAll() ([]entity.Car, error)
}

type repo struct{}

func NewCarRepository() CarRepository {
	return &repo{}
}

const (
	projectId      = "garage-api"
	collectionName = "car"
)

func (*repo) Save(car *entity.Car) (*entity.Car, error) {
	ctx, client, err := createFirestoreClient()
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"id":      car.Id,
		"model":   car.Model,
		"brand":   car.Brand,
		"hp":      car.Hp,
		"license": car.License,
	})

	if err != nil {
		log.Fatalf("Failed adding new Car: %v", err)
		return nil, err
	}
	return car, nil
}

func (*repo) FindAll() ([]entity.Car, error) {
	ctx, client, err := createFirestoreClient()
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()

	var cars []entity.Car
	iterator := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Fatalf("Failed to iterate the list of cars: %v", err)
			return nil, err
		}
		car := entity.Car{
			Id:      doc.Data()["id"].(uuid.UUID),
			Model:   doc.Data()["model"].(string),
			Brand:   doc.Data()["brand"].(string),
			Hp:      doc.Data()["hp"].(int),
			License: doc.Data()["license"].(string),
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func createFirestoreClient() (context.Context, *firestore.Client, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	return ctx, client, err
}
