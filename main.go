package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/michaelvial/hmapgen"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":9001", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()

	service := queryString.Get("service")
	apiKey := queryString.Get("key")
	rawPrecision := queryString.Get("precision")
	precision, _ := strconv.ParseFloat(rawPrecision, 64)
	rawArea := queryString.Get("area")
	area := strings.Split(rawArea, ",")

	options := hmapgen.Options{
		Service:   service,
		Key:       apiKey,
		Precision: precision,
		File:      "images/" + rawArea + "-" + rawPrecision + ".png",
	}

	res, err := hmapgen.GenerateHeightMap(area, options)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

	json.NewEncoder(w).Encode(res)
}
