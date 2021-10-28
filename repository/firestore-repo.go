package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
	"log"

	"github.com/trombettamoacyr/garage-api/entity"
)

type repo struct{}

func NewFirestoreRepository() CarRepository {
	return &repo{}
}

const (
	projectId      = "garage-api-e1e76"
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
		"id":              car.Id.String(),
		"model":           car.Model,
		"brand":           car.Brand,
		"hp":              car.Hp,
		"license":         car.License,
		"insurance_price": car.InsurancePrice,
		"owner_id":        car.OwnerId,
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
	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of cars: %v", err)
			return nil, err
		}

		car := entity.Car{
			Id:             uuid.MustParse(doc.Data()["id"].(string)),
			Model:          doc.Data()["model"].(string),
			Brand:          doc.Data()["brand"].(string),
			Hp:             int(doc.Data()["hp"].(int64)),
			License:        doc.Data()["license"].(string),
			InsurancePrice: doc.Data()["insurance_price"].(string),
			OwnerId:        doc.Data()["owner_id"].(string),
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func (*repo) FindById(id uuid.UUID) (*entity.Car, error) {
	ctx, client, err := createFirestoreClient()
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()

	iter := client.Collection(collectionName).Where("id", "==", id.String()).Documents(ctx)

	doc, err := iter.Next()
	if err != nil {
		err := errors.New("Failed to find the car.")
		return nil, err
	}

	car := entity.Car{
		Id:             uuid.MustParse(doc.Data()["id"].(string)),
		Model:          doc.Data()["model"].(string),
		Brand:          doc.Data()["brand"].(string),
		Hp:             int(doc.Data()["hp"].(int64)),
		License:        doc.Data()["license"].(string),
		InsurancePrice: doc.Data()["insurance_price"].(string),
		OwnerId:        doc.Data()["owner_id"].(string),
	}
	return &car, nil
}

func createFirestoreClient() (context.Context, *firestore.Client, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	return ctx, client, err
}
