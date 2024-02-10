package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type responseMap struct {
	Next      string      `json:"next"` 
	Prev 			string			`json:"previous"`
	Locations []location  `json:"results"`
}

type location struct {
	Name string					  `json:"name"`
}

var responseObject responseMap = responseMap{}
var firstCall bool = true

func innerGetMaps(url string) {
	client := &http.Client{}

	res, err := client.Get(url)
	if err != nil 					{ fmt.Println(err) }
	
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 { fmt.Printf("Response failed with status code: %d", res.StatusCode) }
	if err != nil 					{ fmt.Println(err) }

	json.Unmarshal(body, &responseObject)

	for _, location := range responseObject.Locations {
		fmt.Println(location.Name)
	}
}

func callbackMap() {
	if firstCall {
		innerGetMaps("https://pokeapi.co/api/v2/location/?limit=20&offset=0")
		firstCall = false
	} else {
		if responseObject.Next == "" {
			fmt.Println("there are no further pages to current page")
			return
		}
		innerGetMaps(responseObject.Next)
	}
}

func callbackMapB() {
	if firstCall {
		fmt.Println("map has not been run and there is no prior page")
		return
	} 

	if responseObject.Prev == "" {
		fmt.Println("there are no prior pages to current page")
		return
	}

	innerGetMaps(responseObject.Prev)
}