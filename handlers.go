package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}

func HanldeHome(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "api endpoint")
}

func PostRequest (w http.ResponseWriter, r * http.Request){
	decoder := json.NewDecoder(r.Body)
	var metaData MetaData
	err := decoder.Decode(&metaData)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	
	fmt.Fprintf(w, "PayLoad %v", metaData)
}


func UserPostRequest (w http.ResponseWriter, r * http.Request){
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(response)
	//fmt.Fprintf(w, "PayLoad %v", user)
}