package dbconnection

import (
	"Panda/repositories"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(uri string, timeout time.Duration) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	return client, nil
}

// متد اتصال به Repository ها
func ConnectToRepo(repoName string) (interface{}, error) {
	client, err := ConnectMongo("mongodb://localhost:27017", 10*time.Second)
	if err != nil {
		log.Println("خطا در اتصال به MongoDB:", err)
		return nil, err
	}

	dic := map[string]interface{}{
		"domains":    repositories.NewDomainRepository(client, "domains", 10*time.Second),
		"users":      repositories.NewUserRepository(client, "users", 10*time.Second),
		"identifier": repositories.NewUserRepository(client, "identifier", 10*time.Second),
	}

	repo, exists := dic[repoName]
	if !exists {
		return nil, errors.New("repository یافت نشد")
	}
	return repo, nil
}
