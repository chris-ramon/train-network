package main

import (
	"net/http"
	"encoding/json"
	"path"
)

type Train struct {
	Code string
}

type Station struct {
	Name string
	Latitude string
	Longitude string
}

func Stations() (models []*Station) {
	return models
}

func NewTrains() (models []*Train) {
	return models
}

func rootHandler(response http.ResponseWriter, request *http.Request) {
	fp := path.Join("static", "main.html")
	http.ServeFile(response, request, fp)
}

func trainsHandler(response http.ResponseWriter, request *http.Request) {
	trainOne := &Train{Code: "ICE 847"}
	trainTwo := &Train{Code: "ICE 756"}
	arr := NewTrains()
	arr = append(arr, trainOne)
	arr = append(arr, trainTwo)
	trainJson, err := json.Marshal(arr)
	if err != nil {

	}
	normalizeResponse(response)
	response.Write(trainJson)
}

func normalizeResponse(response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")
}

func stationsHandler(response http.ResponseWriter, request *http.Request) {
	normalizeResponse(response)
	stationsArr := Stations()
	stationsArr = append(stationsArr, (&Station{Name: "Lima", Latitude: "321", Longitude: "123"}))
	stationsArr = append(stationsArr, (&Station{Name: "Arequipa", Latitude: "421", Longitude: "723"}))
	stationsArrJson, err := json.Marshal(stationsArr)
	if err != nil {

	}
	response.Write(stationsArrJson)
}

func main() {
	http.HandleFunc("/", rootHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/trains", trainsHandler)
	http.HandleFunc("/stations", stationsHandler)

	http.ListenAndServe(":8080", nil)
}
