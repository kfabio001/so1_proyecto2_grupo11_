package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/go-redis/redis/v8"
	//"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type casoJSON struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	Age         int    `json:"age"`
	VaccineType string `json:"vaccine_type"`
	Dosis       int    `json:"n_dose"`
}

//var ctx = context.Background()

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

func keyEdad(age int) string {

	if age >= 0 && age <= 11 {
		return "ninos"
	} else if age >= 12 && age <= 18 {
		return "adolescentes"
	} else if age >= 19 && age <= 26 {
		return "jovenes"
	} else if age >= 27 && age <= 59 {
		return "adultos"
	} else if age >= 60 {
		return "vejez"
	} else {
		return "vejez"
	}
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	//log.Printf("Received: %v", in.GetName())

	//************************************************************MONGO
	//Toma el json y lo deserealiza
	data := in.GetName()
	info := casoJSON{}
	fmt.Println(info)
	json.Unmarshal([]byte(data), &info)

	//Mongo
	/*err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}*/
	//fmt.Println(os.Getenv("MONGO_ADDRESS"))
	clienteMongo := options.Client().ApplyURI(os.Getenv("MONGO_ADDRESS"))
	cliente, err := mongo.Connect(context.TODO(), clienteMongo)
	if err != nil {
		log.Println(err)
	}

	//insertar los datos en mongo
	collection := cliente.Database("sopes1-data").Collection("registros")
	//collection := cliente.Database("sopes1").Collection("Vacunados")
	insertResult, err := collection.InsertOne(context.TODO(), info)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(insertResult)

	//************************************************************REDIS

	//Conexion a Redis
	/*rdb := redis.NewClient(&redis.Options{
		Addr: "34.125.118.239:6379",
		DB:   0, // default DB
	})*/

	if info.Age != 0 && info.Name != "" {
		opt, err := redis.ParseURL(os.Getenv("REDIS_ADDRESS"))
		if err != nil {
			fmt.Println("Error con URL de redis en handler")
			log.Println(err)
		}
		rdb := redis.NewClient(opt)

		// Contador de edades
		_, err = rdb.Incr(ctx, keyEdad(info.Age)).Result()
		if err != nil {
			fmt.Println("Error con Incr")
			log.Println(err)
		}

		// Agregar Nombre a la lista
		_, err = rdb.LPush(ctx, "lNombres", info.Name).Result()
		if err != nil {
			fmt.Println("Error con LPush")
			log.Println(err)
		}
		rdb.LTrim(ctx, "lNombres", 0, 4)

		// Publicar Mensaje
		//rdb.Publish(ctx, "Registros", string(body))
		fmt.Println("Registro Publicado")
	} else {
		fmt.Println("Algunos datos incompletos")
	}

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Printf("Falló al escuchar: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})

	fmt.Println(">> SERVER: El servidor está escuchando...")

	if err := s.Serve(lis); err != nil {
		log.Printf("Falló el servidor: %v", err)
	}
}
