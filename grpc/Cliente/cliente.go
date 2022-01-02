package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/gorilla/handlers"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

type caso struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	Age         int    `json:"age"`
	VaccineType string `json:"vaccine_type"`
	Dosis       int    `json:"n_dose"`
}

const (
	address     = "servidorgrcp:50051"
	defaultName = "world"
)

func CasoNuevo(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		panic(err)

	}

	var nuevo caso

	err = json.Unmarshal(body, &nuevo)

	if err != nil {

		panic(err)

	}

	//Variable que almacenará el nuevo caso en formato json
	var jsonstr = string(`{"name":"` + nuevo.Name + `","location":"` + nuevo.Location + `","age":` + strconv.Itoa(nuevo.Age) + `,"vaccine_type":"` + nuevo.VaccineType + `","n_dose":` + strconv.Itoa(nuevo.Dosis) + `}`)

	//Metodo gRPC

	/*
		Crea una conexión con el servidor
		grpc.WithInsecure() os permite realizar una conexión sin tener que suar SSL
	*/

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())

	if err != nil {
		log.Printf("No se conectó: %v", err)
	}

	//Realiza la desconexión al final de la ejecución
	defer conn.Close()

	//Crea un cliente con el cual podemos escuchar
	//Se envía como parametro el Dial de gRPC
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	ra, err := c.SayHello(ctx, &pb.HelloRequest{Name: jsonstr})

	if err != nil {

		log.Printf("could not greet: %v", err)

	}
	log.Printf("Greeting: %s", ra.GetMessage())
}

func Inicio(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Conexión exitosa...")
	log.Println("Si inicio el server")

}

//Función principal
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/casoNuevo", CasoNuevo).Methods("POST")
	router.HandleFunc("/inicio", Inicio).Methods("GET")
	log.Println("Si inicio el server")
	log.Fatal(http.ListenAndServe(":3050", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
