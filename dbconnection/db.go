package dbconnection

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection تابعی است که به پایگاه داده متصل شده و کالکشن "users" را بازمی‌گرداند.
func Collection() (*mongo.Collection, error) {
	// تنظیمات اتصال به پایگاه داده
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// ایجاد کلاینت و اتصال به سرور
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// بررسی اتصال
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	// دسترسی به کالکشن "users" در دیتابیس "nostr"
	collection := client.Database("nostr").Collection("users")
	return collection, nil
}
