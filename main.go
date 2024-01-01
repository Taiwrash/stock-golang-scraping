package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type Stock struct {
	company, change, price string
}

func main() {
	tickers := []string{
		"MSFT",
		"IBM",
		"GE",
		"UNP",
		"COST",
		"MCD",
		"V",
		"WMT",
		"DIS",
		"MMM",
		"INTC",
		"AXP",
		"AAPL",
		"BA",
		"CSCO",
		"GS",
		"JPM",
		"CRM",
		"VZ",
		"NQ",
		"SPSK",
		"ISWD.SW",
		"ISUS",
		"SPUS",
	}
	// fmt.Println(ticker)

	c := colly.NewCollector()

	stocks := []Stock{}

	// Find and visit all links
	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {
		stock := Stock{}
		stock.company = e.ChildText("h1")
		fmt.Println("Company:", stock.company)
		stock.price = e.ChildText("fin-streamer[data-field='regularMarketPrice']")
		fmt.Println("Price:", stock.price)
		stock.change = e.ChildText("fin-streamer[data-field='regularMarketChangePercent']")
		fmt.Println("Change:", stock.change)

		stocks = append(stocks, stock)
	})
	c.Wait()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error: %v", err)
	})

	for _, ticker := range tickers {
		URLToBeLocated := fmt.Sprintf("https://finance.yahoo.com/quote/%s/", ticker)
		c.Visit(URLToBeLocated)
	}

	fmt.Println(stocks)

	f, err := os.Create("data.csv")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	defer f.Close()

	w := csv.NewWriter(f)
	headers := []string{"company", "price", "change"}
	w.Write(headers)
	for _, records := range stocks {
		record := []string{
			records.company,
			records.price,
			records.change,
		}
		w.Write(record)
	}
	defer w.Flush()

	// / Server setup to make it an API

	// r := gin.Default()
	// r.GET("/stocks", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": stocks,
	// 	})
	// })
	// r.Run()

}
