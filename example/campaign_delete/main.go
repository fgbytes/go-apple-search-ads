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
	orgID := int64(1405310)
		pemdat, _ := ioutil.ReadFile("crt/cert.pem")
	keydat, _ := ioutil.ReadFile("crt/cert.key")
	client, err := searchads.NewClient(nil, pemdat, keydat, &orgID)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}

	resp, err := client.Campaign.Delete(context.Background(), 259428355)
	if err != nil {
		log.Fatalf("Campaign Delete error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(resp)
	fmt.Println(string(res))
}
