package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	// _ "github.com/go-sql-driver/mysql"
	handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type busStop struct {
	busStopCode int    `json: "busStopCode"`
	desc        string `json: "desc"`
}

func getBusStop(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	busStopCodeGet := params["busStopCode"]
	busStopCode, err := strconv.Atoi(busStopCodeGet)
	if err != nil {
		println(err.Error())
	}
	var data busStop
	data.busStopCode = busStopCode
	if r.Method == "GET" {
		data.desc = "Hello"
		json.NewEncoder(w).Encode(data.desc)
		w.WriteHeader(http.StatusAccepted)
	} else if r.Method == "PUT" {
		//Unmarshal JSON
		reqBody, err1 := ioutil.ReadAll(r.Body)
		if err1 != nil {
			println(err1.Error())
		}
		defer r.Body.Close()
		err = json.Unmarshal(reqBody, &data)
		w.WriteHeader(http.StatusAccepted)
		return
	}

}

func main() {
	//API
	router := mux.NewRouter()
	//Web front-end CORS
	headers := handlers.AllowedHeaders([]string{"X-REQUESTED-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/api/v1/busStops/{busStopCode}", getBusStop).Methods("GET", "PUT")

	//Establish port
	//checkMicroservices()
	fmt.Println("Listening at port 6030")
	log.Fatal(http.ListenAndServe(":6030", handlers.CORS(headers, methods, origins)(router)))
}
