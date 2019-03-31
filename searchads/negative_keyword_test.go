package searchads

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestCampaignNegativeKeywordServive_List(t *testing.T) {
	client, mux, _, teardown := setup()
	t.Log("Setup Done")
	defer teardown()

	mux.HandleFunc("/campaigns/1234/negativekeywords", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(loadFixture("campaign_negative_keywords.json"))
	})
	opt := ListOptions{}
	got, _, err := client.CampaignNegativeKeyword.List(context.Background(), 1234, &opt)
	if err != nil {
		t.Errorf("CampaignNegativeKeyword.List returned error: %v", err)
	}

	want := []*NegativeKeyword{}
	responseToInterface(loadFixture("campaign_negative_keywords.json"), &want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CampaignNegativeKeyword.List = %+v, want %+v", got, want)
	}
}
func TestCampaignNegativeKeywordServive_CreateBulk(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	nk := NegativeKeyword{
		CampaignID: 1234,
		Text:       "i do negative keywords",
		MatchType:  EXACT,
		Status:     KEYWORD_ACTIVE,
	}

	input := []*NegativeKeyword{&nk}

	wantAcceptHeaders := []string{"application/json"}
	mux.HandleFunc("/campaigns/1234/negativekeywords/bulk", func(w http.ResponseWriter, r *http.Request) {
		v := []*NegativeKeyword{}
		json.NewDecoder(r.Body).Decode(&v)
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write(loadFixture("campaign_negative_keyword_bulk.json"))
	})

	got, _, err := client.CampaignNegativeKeyword.CreateBulk(context.Background(), 1234, input)

	if err != nil {
		t.Errorf("CampaignNegativeKeyword.CreateBulk returned error: %v", err)
	}
	want := []*NegativeKeyword{
		&NegativeKeyword{
			ID:               1,
			CampaignID:       1234,
			Text:             "i do negative keywords",
			MatchType:        EXACT,
			Status:           KEYWORD_ACTIVE,
			ModificationTime: "2019-02-22T15:25:46.851",
			Deleted:          false,
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CampaignNegativeKeyword.CreateBulk returned %+v, want %+v", got, want)
	}
}

func TestCampaignNegativeKeywordServive_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	wantAcceptHeaders := []string{"application/json"}
	mux.HandleFunc("/campaigns/1234/negativekeywords/1234", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		w.WriteHeader(http.StatusOK)
		w.Write(loadFixture("campaign_negative_keyword_delete.json"))
	})

	resp, err := client.CampaignNegativeKeyword.Delete(context.Background(), 1234, 1234)
	if err != nil {
		t.Errorf("CampaignNegativeKeyword.Delete returned error: %v", err)
	}
	want := http.StatusOK
	got := resp.StatusCode
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CampaignNegativeKeyword.Delete returned %+v, want %+v", got, want)
	}
}
