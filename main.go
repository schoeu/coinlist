package main

import (
	"fmt"
	"./config"
	"./utils"
	//"database/sql"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type coinsInfo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Website_slug string `json:"website_slug"`
}

type coins struct {
	Data []coinsInfo `json:"data"`
}

func main() {
	// dbStr := config.DBString
	//db := utils.OpenDb("mysql", dbStr)
	url := config.URL
	getCoinsList(url)
}

// get request
func getCoinsList(url string) coins {
	res, err := http.Get(url)
	utils.ErrHandle(err)

	body, err := ioutil.ReadAll(res.Body)
	val := coins{}
	json.Unmarshal(body, &val)
	res.Body.Close()
	utils.ErrHandle(err)
	fmt.Println(val)

	return val
}

func storeData() {

}