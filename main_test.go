package main

import (
	"bytes"
	"fmt"
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

func TestGetTransactions(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/transactions.pdf")
	if err != nil {
		t.Fatal(err)
	}
	s, err := extractPDFText(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("s", s)
	t.Fail()
}
