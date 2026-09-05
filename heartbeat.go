package main

import (

	"fmt"
	"log"
	"net/http"
)


func heartbitCheck(w http.ResponseWriter , r *http.Request) {

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w,"Ok , Staus code : %d",http.StatusOK)

}

func main() {

	http.HandleFunc("/heartbit",heartbitCheck)

	port:=":8082"

	//fmt.Println("Server running on ",port)
	log.Println("Server running on http://localhost:8082/heartbit")



	if err:= http.ListenAndServe(port,nil) ; err!=nil {

		log.Fatal(err)
	}
}