package main

import (
	"fmt"
	"log"

	//"os"
	"context"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/joho/godotenv"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Register struct {
	Name         string `json:name`
	Location     string `json:location`
	Age          int    `json:age`
	Vaccine_type string `json:vaccine_type`
	N_dose       int    `json:n_dose`
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			next.ServeHTTP(w, req)
		})
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodOptions)

	router.Use(middlewareCors)
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

func publisherHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error con Read Body")
		log.Fatal(err)
	}

	var nuevo Register
	err = json.Unmarshal(body, &nuevo)
	if err != nil {
		fmt.Println("Error con Unmarshal")
		log.Fatal(err)
	}

	if nuevo.Age != 0 && nuevo.Name != "" {
		//opt, err := redis.ParseURL(os.Getenv("REDIS_ADDRESS"))
		opt, err := redis.ParseURL("redis://default:redisgrupo11@34.125.118.239:6379")
		//opt, err := redis.ParseURL("redis://172.17.0.3:6379")
		if err != nil {
			fmt.Println("Error con URL de redis en handler")
			log.Fatal(err)
		}
		rdb := redis.NewClient(opt)

		// Contador de edades
		_, err = rdb.Incr(ctx, keyEdad(nuevo.Age)).Result()
		if err != nil {
			fmt.Println("Error con Incr")
			log.Fatal(err)
		}

		// Agregar Nombre a la lista
		_, err = rdb.LPush(ctx, "lNombres", nuevo.Name).Result()
		if err != nil {
			fmt.Println("Error con LPush")
			log.Fatal(err)
		}
		rdb.LTrim(ctx, "lNombres", 0, 4)

		// Publicar Mensaje
		rdb.Publish(ctx, "Registros", string(body))
		fmt.Println("Registro Publicado")
	} else {
		fmt.Println("Algunos datos incompletos")
	}
}

func main() {
	/*
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	*/

	router := mux.NewRouter().StrictSlash(true)
	enableCORS(router)

	router.HandleFunc("/entrada", publisherHandler).Methods("POST")

	fmt.Println("Servidor pub en puerto 3050")
	if err := http.ListenAndServe(":3050", router); err != nil {
		log.Fatal(err)
		return
	}
}
