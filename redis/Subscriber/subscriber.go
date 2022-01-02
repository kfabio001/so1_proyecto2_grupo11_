package main

import (
	"context"
	"fmt"
	"log"

	//"os"
	"encoding/json"

	//"github.com/joho/godotenv"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Register struct {
	Name         string `json:name`
	Location     string `json:location`
	Age          int    `json:age`
	Vaccine_type string `json:vaccine_type`
	N_dose       int    `json:n_dose`
}

var ctx = context.Background()

func main() {
	// Carga de archivo .env
	/*
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	*/

	// Conexion Redis
	//opt, err := redis.ParseURL(os.Getenv("REDIS_ADDRESS"))
	opt, err := redis.ParseURL("redis://default:redisgrupo11@34.125.118.239:6379")
	//opt, err := redis.ParseURL("redis://172.17.0.3:6379")
	if err != nil {
		fmt.Println("Error con URL de redis")
		log.Fatal(err)
	}
	rdb := redis.NewClient(opt)
	fmt.Println("Conectado a Redis")

	sub := rdb.Subscribe(ctx, "Registros")
	fmt.Println("Suscrito al canal de Redis")

	//Conexion Mongo
	//cOptions := options.Client().ApplyURI(os.Getenv("MONGO_ADDRESS"))
	cOptions := options.Client().ApplyURI("mongodb://sopes1:managerl@34.125.118.239:27017")
	//cOptions := options.Client().ApplyURI("mongodb://172.17.0.2:27017")
	mongoClient, err := mongo.Connect(context.TODO(), cOptions)
	if err != nil {
		fmt.Println("Error creando cliente de Mongo")
		log.Println(err)
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error conectando al servidor")
		log.Fatal(err)
	}
	fmt.Println("Conectado a MongoDB")
	myMongoDB := mongoClient.Database("sopes1-data")
	collection := myMongoDB.Collection("registros")

	for {
		// Subscriber de Redis
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			fmt.Println("Error reciviendo datos")
			log.Fatal(err)
		} else {
			var nuevo Register
			err = json.Unmarshal([]byte(msg.Payload), &nuevo)
			if err != nil {
				fmt.Println("Error Convirtiendo datos")
				log.Fatal(err)
			}
			fmt.Print("Mensaje recibido del canal '" + msg.Channel + "': ")
			fmt.Println(nuevo.Name + " - " + nuevo.Location)

			//Insertar MongoDB
			_, err := collection.InsertOne(context.TODO(), nuevo)
			if err != nil {
				fmt.Println("Error Insertando datos en MongoDB")
				log.Fatal(err)
			}
			fmt.Println("\tRegistro insertado en MongoDB")
		}
	}
}
