package stockchecker

import (
	"github.com/gin-gonic/gin"
	
    "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// MSFT data=[110.56, 111.25, 115.78], average=112.50
type stock struct {
	StockSymbol string 		`json:Symbol`
	Close  		[]float64 	`json:data`
	Average 	float64 	`json:average`
}


type entry struct {
	Open 			 string `json:"1. open"`
	High 			 string `json:"2. high"`
	Low  			 string `json:"3. low"`
	Close 			 string `json:"4. close"`
	Volume 			 string `json:"5. volume"`
}

type alphavantagedata struct {
	MetaData    map[string]interface{} `json:"Meta Data"`
	TimeSeries 	map[string]entry `json:"Time Series (Daily)"`
}

func getData(symbol string, apikey string) alphavantagedata {
	url := fmt.Sprintf("https://www.alphavantage.co/query?apikey=%s&function=TIME_SERIES_DAILY&symbol=%s", apikey, symbol)
	// log.Println(url)

	// create the request, add the header asking for json, get the content
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	
	// unmarshall to the structs to use later
	var result alphavantagedata
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	// log.Println(result.TimeSeries["2021-09-16"].Close)

	return result
}

func getAverage(values []float64) float64 {
	// get the average (how is this not built in?!)
	var sum float64
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}
	return sum/float64(len(values))
}


func getLastNDays(data alphavantagedata, ndays int) []float64 {
	var close []float64
	// you can't just iterate over all the keys as its a map, they don't always come out in date order...
	// so we will do it by day, starting today and going back a day. However we want a break if this
	// doesn't find three valid dates.
	d := time.Now()
	max_attempts := len(data.TimeSeries)
	attempt := 0
	for {
		if len(close) >= ndays || attempt > max_attempts {
			break;
		}
		// get the date in the right format
		date := d.Format("2006-01-02")
		// add the value to the list
		close_value, _ := strconv.ParseFloat(data.TimeSeries[date].Close, 64)
		if close_value != 0 {
			close = append(close, close_value)
		}
		// go back a day
		d = d.AddDate(0, 0, -1)
		attempt++
	}
	return close
}

func getDisplayData(data alphavantagedata, ndays int, symbol string) stock {
	close := getLastNDays(data, ndays)

	// create the struct and return
	return stock{
		StockSymbol: symbol,
		Close: close,
		Average: getAverage(close),
	}
}


func GetStock(c *gin.Context) {
// func getStock() {
	//normally here we would handle what happens if the environment
	// variables are not set - but quick and dirty example
	symbol := os.Getenv("SYMBOL")
	ndays, err := strconv.Atoi(os.Getenv("NDAYS"))
	if err != nil {
		log.Fatal(err)
	}
	apikey := os.Getenv("APIKEY")
	
	data := getData(symbol, apikey)

	// run the get request
	c.IndentedJSON(http.StatusOK, getDisplayData(data, ndays, symbol))
}