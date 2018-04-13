package clipper

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestParsePDF(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/2017-12-clipper-year.pdf")
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
	num, records, err := getCSV(samplePages)
	if err != nil {
		t.Fatal(err)
	}
	if len(records) != 32 {
		t.Errorf("bad record length: want %d, got %d", 32, len(records))
	}
	if records[0][1] != "Transaction Type" {
		t.Errorf("header record: want Transaction Type, got %q", records[0][1])
	}
	if records[1][1] != "Single-tag fare payment" {
		t.Errorf("first record: want Single-tag, got %q", records[1][1])
	}
	if records[31][2] != "Civic Center (BART)" {
		t.Errorf("last record: want Civic Center, got %q", records[31][2])
	}
	if num != 1202728442 {
		t.Errorf("bad account number: got %d, want %d", num, 1202728442)
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
