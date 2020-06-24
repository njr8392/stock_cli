package main

import ("os"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

type Stock struct {
	c float32 `json: "c"`	// current price
	h float32 `json: "h"`	// high price of the day
	l float32 `json: "l"`	// low price of the day
	o float32 `json: "o"`	// open price of the day
	pc float32 `json: pc`   // previous close 
	t float64 `json: t`     // unix timestamp 
}


func getStockPrice (body []byte) (*Stock, error){
	var api = new(Stock)
	err := json.Unmarshal(body, &api)
	if err != nil {
		log.Fatal(err)
	}
	return api, err
}

func main () {
	// api key for finhub = token=bq94sovrh5rc96c0mkmg
	fmt.Printf("args are %v\n", os.Args[1:])

	res, err := http.Get("https://finnhub.io/api/v1/quote?symbol=AAPL&token=bq94sovrh5rc96c0mkmg")
	if err != nil {
		log.Fatal(err)
	}
	price, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Printf("%s\n", price)
	
	if err != nil {
		log.Fatal(err)
	}
	
	call, err := getStockPrice(price)

	
	fmt.Printf("%+v\n", call)
	
}

