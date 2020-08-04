package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Stock struct {
	C  float32 `json: "c"` // current price
	H  float32 `json: "h"` // high price of the day
	L  float32 `json: "l"` // low price of the day
	O  float32 `json: "o"` // open price of the day
	Pc float32 `json: pc`  // previous close
	T  float64 `json: t`   // unix timestamp
}

type Earnings struct {
	Date            string
	EpsActual       float32
	EpsEstimate     float32
	Hour            string
	Quarter         int8
	RevenueActual   float64
	RevenueEstimate float64
	Symbol          string
	Year            int32
}

type EarningsSearch struct {
	EarningsCalendar []*Earnings
}

// get Earnings calender
func getEarnings(date string) (*EarningsSearch, error) {
	var result EarningsSearch
	url := "https://finnhub.io/api/v1/calendar/earnings?from=" + date + "&to=" + date + "&token=" + os.Getenv("APIKEY")
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	er := json.Unmarshal(data, &result)
	if er != nil {
		log.Fatal(err)
	}

	return &result, nil
}

// unmarshal json response and store in Stock variable
func getStockPrice(body []byte) (*Stock, error) {
	var api *Stock
	err := json.Unmarshal(body, &api)
	if err != nil {
		log.Fatal(err)
	}
	return api, err
}

func printEarnings(earn *EarningsSearch) {
	for _, earnings := range earn.EarningsCalendar {
		fmt.Printf("Date:\t%s\n", earnings.Date)
		fmt.Printf("EPS Acutal:\t%.2f\n", earnings.EpsActual)
		fmt.Printf("EPS Estimate:\t%.2f\n", earnings.EpsEstimate)
		fmt.Printf("Earnings will be released:\t%s\n", earnings.Hour)
		fmt.Printf("Revenue Acutal:\t%.2f\n", earnings.RevenueActual)
		fmt.Printf("Revenue Estimate:\t%.2f\n", earnings.RevenueEstimate)
		fmt.Printf("Stockticker is :\t%s\n", earnings.Symbol)
		fmt.Println("**************************************\n")
	}
}

func main() {
	// takes first argument and returns stock price

	if args := os.Args[1:]; len(args) < 2 {
		tick := os.Args[1]
		res, err := http.Get("https://finnhub.io/api/v1/quote?symbol=" + tick + "&token=" + os.Getenv("APIKEY"))
		if err != nil {
			log.Fatal(err)
		}
		price, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		if err != nil {
			log.Fatal(err)
		}

		call, err := getStockPrice(price)

		fmt.Printf("Current Price of %s is %.2f\n", tick, call.C)
	} else {
		date := os.Args[2]
		earnings, err := getEarnings(date)
		if err != nil {
			fmt.Println(err)
		}

		printEarnings(earnings)

	}

}
