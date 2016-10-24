package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Index basic endpoint
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	dat, err := ioutil.ReadFile("./public/439322/info.json")
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(string(dat)); err != nil {
		panic(err)
	}
}

func MatchInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	dat, err := ioutil.ReadFile("./public/439322/info.json")
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(string(dat)); err != nil {
		panic(err)
	}
}

func MatchStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	dat, err := ioutil.ReadFile("./public/439322/stats.json")
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(string(dat)); err != nil {
		panic(err)
	}
}
