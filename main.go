package main

import (
	"./config"
	"./utils"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type coinsInfo struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Website_slug string `json:"website_slug"`
}

type coins struct {
	Data []coinsInfo `json:"data"`
}

func main() {
	dbStr := config.DBString
	db := utils.OpenDb("mysql", dbStr)
	url := config.URL
	listInfo := getCoinsList(url)

	var sqlArr []string
	for _, v := range listInfo.Data {
		var infos []string
		infos = append(infos, strconv.Itoa(v.Id), "'"+strings.Replace(v.Name, "'", "''", -1)+"'", "'"+v.Symbol+"'", "'"+v.Website_slug+"'", "'"+utils.GetCurrentDate()+"'")
		sqlArr = append(sqlArr, "("+strings.Join(infos, ",")+")")
	}
	storeData(db, sqlArr)
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
	return val
}

func storeData(db *sql.DB, sqlData []string) {
	_, err := db.Exec("delete from bt_listings")
	utils.ErrHandle(err)

	sqlStr := "INSERT INTO bt_listings (pid, name, symbol, website_slug, update_time) VALUES " + strings.Join(sqlData, ",")
	_, err = db.Exec(sqlStr)
	utils.ErrHandle(err)
}
