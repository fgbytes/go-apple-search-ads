package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fgbytes/go-apple-search-ads/searchads"
)

func main() {
	campaignID := int64(262727835)
	adGroupID := int64(262765077)
	pemdat, _ := ioutil.ReadFile("crt/cert.pem")
	keydat, _ := ioutil.ReadFile("crt/cert.key")
	client, err := searchads.NewClient(nil, pemdat, keydat, nil)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	opt := searchads.ListOptions{Limit: 1000, Offset: 0}
	list, rs, err := client.AdGroupTargetingKeyword.List(context.Background(), campaignID, adGroupID, &opt)
	if err != nil {
		log.Fatalf("AdGroupTargetingKeyword List error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&list)
	fmt.Println(string(res))
	fmt.Println("----------------")
	fmt.Println(len(list))
	fmt.Println(rs.Pagination.ItemsPerPage)
	fmt.Println(rs.Pagination.StartIndex)
	fmt.Println(rs.Pagination.TotalResults)
}
