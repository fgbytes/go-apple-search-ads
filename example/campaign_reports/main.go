package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fgbytes/go-apple-search-ads/searchads"
)

func main() {
	orgID := int64(378940)
	pemdat, _ := ioutil.ReadFile("../cert.pem")
	keydat, _ := ioutil.ReadFile("../cert.key")
	client, err := searchads.NewClient(nil, pemdat, keydat, &orgID)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	filter := searchads.ReportFilter{
		StartTime:   "2019-09-06",
		EndTime:     "2019-09-06",
		Granularity: searchads.DAILY,
		Selector: searchads.Selector{
			OrderBy: []searchads.OrderBySelector{
				searchads.OrderBySelector{
					Field:     searchads.OrderByImpressions,
					SortOrder: searchads.DESCENDING,
				},
			},
			Conditions: []searchads.Condition{},
			Pagination: searchads.PaginationSelector{
				Offset: 0,
				Limit:  1000,
			},
		},
		GroupBy: []searchads.GroupBy{
			searchads.GroupByCountryOrRegion,
		},
		TimeZone:                   searchads.ORTZ,
		ReturnRecordsWithNoMetrics: true,
		ReturnRowTotals:            true,
		ReturnGrandTotals:          true,
	}
	report, rs, err := client.Report.Campaigns(context.Background(), &filter)
	if err != nil {
		log.Fatalf("Campaign Reports error: %s", err)
		panic(err)
	}
	fmt.Println("----------------")

	log.Printf("%20.20s\tStatus\tImpr\tTap\tInst\tTTR\t\tCPA\tCPT\tCost\tCR", "campaign name")

	for _, row := range report.ReportingDataResponse.Row {
		if row.Total.Impressions != 0 {
			log.Println(row)
		}

	}
	fmt.Println("----------------")
	fmt.Println(report.ReportingDataResponse.GrandTotals.Total.LocalSpend.Amount, " ", report.ReportingDataResponse.GrandTotals.Total.Installs, " ", report.ReportingDataResponse.GrandTotals.Total.AvgCPA.Amount)
	fmt.Println("----------------")
	fmt.Println(rs.Pagination.ItemsPerPage)
	fmt.Println(rs.Pagination.StartIndex)
	fmt.Println(rs.Pagination.TotalResults)

}
