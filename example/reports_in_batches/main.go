package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/fgbytes/go-apple-search-ads/searchads"
)

func main() {
	start := time.Now()
	pemdat, _ := ioutil.ReadFile("../cert.pem")
	keydat, _ := ioutil.ReadFile("../cert.key")
	client, err := searchads.NewClient(nil, pemdat, keydat, nil)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	opt := searchads.ListOptions{Limit: 1000, Offset: 0}
	list, _, err := client.Campaign.List(context.Background(), &opt)
	if err != nil {
		log.Fatalf("Campaign List error: %s", err)
		panic(err)
	}
	campaignIDs := []int64{}
	for _, c := range list {
		campaignIDs = append(campaignIDs, c.ID)
	}
	searchTerms, err := AsyncSearchTerms(campaignIDs)
	fmt.Println(searchTerms)
	res, _ := json.Marshal(&searchTerms)
	fmt.Println(string(res))

	duration := time.Since(start)
	fmt.Println("duration")
	fmt.Println(duration)
}

func buildFilter() searchads.ReportFilter {
	return searchads.ReportFilter{
		StartTime:   "2019-04-01",
		EndTime:     "2019-06-29",
		Granularity: searchads.DAILY,
		TimeZone:    searchads.UTC,
		Selector: searchads.Selector{
			OrderBy: []searchads.OrderBySelector{
				searchads.OrderBySelector{
					Field:     searchads.OrderByImpressions,
					SortOrder: searchads.ASCENDING,
				},
			},
			Pagination: searchads.PaginationSelector{
				Offset: 0,
				Limit:  1000,
			},
		},
		GroupBy:                    []searchads.GroupBy{},
		ReturnRecordsWithNoMetrics: false,
		ReturnRowTotals:            false,
		ReturnGrandTotals:          false,
	}
}

func getReport(id int64, ch chan<- *searchads.SearchTermsReport, wg *sync.WaitGroup) {
	fmt.Println("Reports for ", id)
	defer wg.Done()
	filter := buildFilter()
	pemdat, _ := ioutil.ReadFile("../cert.pem")
	keydat, _ := ioutil.ReadFile("../cert.key")
	client, _ := searchads.NewClient(nil, pemdat, keydat, nil)
	report, _, err := client.Report.SearchTerms(context.Background(), id, &filter)
	if err != nil {
		log.Println("err handle it")
	}
	ch <- report
}

func AsyncSearchTerms(campaignIDs []int64) ([]*searchads.SearchTermsReport, error) {
	ch := make(chan *searchads.SearchTermsReport)
	var responses []*searchads.SearchTermsReport
	var campaignID int64
	var wg sync.WaitGroup
	for _, campaignID = range campaignIDs {
		wg.Add(1)
		go getReport(campaignID, ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for res := range ch {
		responses = append(responses, res)
	}

	return responses, nil
}
