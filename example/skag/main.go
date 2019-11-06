package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/fgbytes/go-apple-search-ads/searchads"
)

func main() {
	//use your specific orgID
	orgID := int64(1111111)
	csvFile, err := os.Open(fmt.Sprintf("%s", "keywords.csv"))
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	if err != nil {
		log.Fatal(err)
	}

	pemdat, _ := ioutil.ReadFile("crt/cert.pem")
	keydat, _ := ioutil.ReadFile("crt/cert.key")
	client, err := searchads.NewClient(nil, pemdat, keydat, &orgID)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	//
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal("failed to read line with email ", error)
		}
		campaignID, err := strconv.ParseInt(line[2], 10, 64)
		now := time.Now().UTC()
		startTime := fmt.Sprintf("%4d-%02d-%02dT%02d:%02d:%02d.000",
			now.Year(), now.Month(), now.Day(),
			now.Hour(), now.Minute(), now.Second())
		data := searchads.AdGroup{
			CampaignID:             campaignID,
			Name:                   line[0],
			StartTime:              startTime,
			AutomatedKeywordsOptIn: false,
			// CpaGoal: &searchads.Amount{
			// 	Amount:   "5",
			// 	Currency: "USD",
			// },
			DefaultCpcBid: &searchads.Amount{
				Amount:   line[1],
				Currency: "USD",
			},
		}

		createdAdGroup, _, err := client.AdGroup.Create(context.Background(), campaignID, &data)
		if err != nil {
			log.Fatalf("adgroup Create error: %s", err)
			panic(err)
		}
		res, _ := json.Marshal(&createdAdGroup)
		fmt.Println(string(res))

		log.Println(createdAdGroup.ID)
		// ___
		status, _ := searchads.ParseKeywordStatus("ACTIVE")
		matchType, _ := searchads.ParseMatchType("EXACT")
		input := []searchads.TargetingKeyword{
			searchads.TargetingKeyword{
				AdGroupID: createdAdGroup.ID,
				Text:      line[0],
				Status:    status,
				BidAmount: searchads.Amount{
					Amount:   line[1],
					Currency: "USD",
				},
				MatchType: matchType,
			},
		}
		createdKeyword, rs, err := client.AdGroupTargetingKeyword.CreateBulk(context.Background(), campaignID, createdAdGroup.ID, input)
		if err != nil {
			log.Fatalf("TargetingKeyword error: %s", err)
			panic(err)
		}
		res, _ = json.Marshal(&createdKeyword)
		fmt.Println(string(res))
		fmt.Println("----------------")
		fmt.Println(rs.Pagination.ItemsPerPage)
		fmt.Println(rs.Pagination.StartIndex)
		fmt.Println(rs.Pagination.TotalResults)
	}
}
