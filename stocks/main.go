package main

import (
	"flag"
	"fmt"
	"github.com/SimpleApplicationsOrg/stock/alphavantage"
)

// export ALPHA_VANTAGE_URL="https://www.alphavantage.co"
// export ALPHA_VANTAGE_KEY_NAME="apikey"
// export ALPHA_VANTAGE_KEY_VALUE="xxx"
// https://www.alphavantage.co/query?apikey=xxx&function=TIME_SERIES_DAILY&symbol=VMW - example endpoint - PREMIUM
// Random tickers to test with, some are no longer traded:
// "VMW", "CLGX", "GOOG", "AMZN", "NFLX", "AAPL", "IBM", "GE", "CAT", "AA", "META", "YHOO", "HPE"}

func main() {
	symbolFlag := flag.String("symbol", "VMW", "Specify any stock symbol")
	// TIME_SERIES_DAILY is a premium feature, therefore excluded for use here
	functionFlag := flag.String("function", "TIME_SERIES_INTRADAY", "The following values are supported: TIME_SERIES_INTRADAY, " + 
								"TIME_SERIES_WEEKLY, TIME_SERIES_MONTHLY, TIME_SERIES_DAILY_ADJUSTED")
	intervalFlag := flag.String("interval", "1min", "The following values are supported: 1min, 5min, 15min, 30min, 60min")

	flag.Parse()

	avClient, err := alphavantage.NewAVClient()
	if err != nil {
		fmt.Printf("error getting client: %s", err.Error())
		return
	}

	if *functionFlag == "TIME_SERIES_INTRADAY" && *intervalFlag == "1min" || *intervalFlag == "5min" || *intervalFlag == "15min" || *intervalFlag == "30min" || *intervalFlag == "60min" {
		response, err := avClient.TimeSeriesIntraday(*symbolFlag, *intervalFlag)
		if err != nil {
			fmt.Printf("error calling api: %s", err.Error())
			return
		}

		metaData := *response.MetaData
		fmt.Println(metaData.Information(), metaData.OutputSize())
		fmt.Println(metaData.LastRefreshed(), metaData.TimeZone())
		fmt.Println(metaData.Interval())
	
		timeSeries := *response.TimeSeries
		for _, timeStamp := range timeSeries.TimeStamps() {
			value := (timeSeries)[timeStamp]
			// TimeSeriesIntraDay
			fmt.Println(timeStamp, value.Open(), value.High(), value.Low(), value.Close(), value.Volume())
		}
	}

	// TIME_SERIES_WEEKLY and TIME_SERIES_MONTHLY functions 
	if *functionFlag == "TIME_SERIES_WEEKLY" || *functionFlag == "TIME_SERIES_MONTHLY" {
		response, err := avClient.TimeSeries(*functionFlag, *symbolFlag)
		if err != nil {
			fmt.Printf("error calling api: %s", err.Error())
			return
		}

		metaData := *response.MetaData
		fmt.Println(metaData.Information(), metaData.OutputSize())
		fmt.Println(metaData.LastRefreshed(), metaData.TimeZone())
		fmt.Println(metaData.Interval())
	
		timeSeries := *response.TimeSeries
		for _, timeStamp := range timeSeries.TimeStamps() {
			value := (timeSeries)[timeStamp]
			fmt.Println(timeStamp, value.Open(), value.High(), value.Low(), value.Close(), value.Volume())
		}
	}	

	// TIME_SERIES_DAILY_ADJUSTED function
	if *functionFlag == "TIME_SERIES_DAILY_ADJUSTED" {
		response, err := avClient.TimeSeries(*functionFlag, *symbolFlag)
		if err != nil {
			fmt.Printf("error calling api: %s", err.Error())
			return
		}

		metaData := *response.MetaData
		fmt.Println(metaData.Information(), metaData.OutputSize())
		fmt.Println(metaData.LastRefreshed(), metaData.TimeZone())
		fmt.Println(metaData.Interval())
	
		timeSeries := *response.TimeSeries
		for _, timeStamp := range timeSeries.TimeStamps() {
			value := (timeSeries)[timeStamp]
			fmt.Println(timeStamp, value.Open(), value.High(), value.Low(), value.Close(), value.Volume())
		}
	}	
}