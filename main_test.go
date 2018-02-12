package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/unidoc/unidoc/common"
)

func init() {
	common.Log = common.ConsoleLogger{LogLevel: common.LogLevelDebug}
}

func TestGetViewState(t *testing.T) {
	val, err := findViewState(bytes.NewReader(loginPage))
	if err != nil {
		t.Fatal(err)
	}
	if val != "5428792554773752026:-479288318579711101" {
		t.Errorf("wrong val: %v", val)
	}
}

func TestGetCards(t *testing.T) {
	r := bytes.NewReader(dashboard)
	cards, err := GetCards(r)
	if err != nil {
		t.Fatal(err)
	}
	if len(cards) != 3 {
		t.Errorf("len(cards) should be 3, got %d", len(cards))
	}
	if cards[0].Nickname != "Personal" {
		t.Errorf("bad card nickname: %q", cards[0].Nickname)
	}
}

func TestParsePDF(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/transactions.pdf")
	if err != nil {
		t.Fatal(err)
	}
	s, err := extractPDFText(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if len(s) != 2 {
		t.Errorf("expected 2 pages, got %d", len(s))
	}
}

func TestGetTransactions(t *testing.T) {
	records, err := getCSV(samplePages)
	if err != nil {
		t.Fatal(err)
	}
	if len(records) != 32 {
		t.Errorf("bad record length: want %d, got %d", 32, len(records))
	}
	if records[0][1] != "Single-tag fare payment" {
		t.Errorf("first record: want Single-tag, got %q", records[0][1])
	}
	if records[31][2] != "Civic Center (BART)" {
		t.Errorf("last record: want Civic Center, got %q", records[31][2])
	}
}

var tabsTests = []struct {
	prev, cur float64
	tabs      int
}{
	{28, 133, 1},
	{28, 359, 2},
	{28, 300, 2},
	{28, 400, 2},
	{655, 722, 2},
	{550, 722, 3},
}

func TestHowManyTabs(t *testing.T) {
	for _, tt := range tabsTests {
		got := howManyTabs(tt.prev, tt.cur)
		if got != tt.tabs {
			t.Errorf("howManyTabs(%v, %v): got %d, want %d", tt.prev, tt.cur, got, tt.tabs)
		}
	}
}
