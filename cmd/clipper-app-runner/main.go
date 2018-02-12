package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/kevinburke/clipper"
)

func checkError(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s: %v\n", msg, err)
		os.Exit(2)
	}
}

var email = flag.String("email", "", "Login email")
var password = flag.String("password", "", "Password")

func writeTransactions(result map[clipper.Card]clipper.TransactionData) {
	for card := range result {
		f, err := os.Open(fmt.Sprintf("clipper-transactions-%d.csv", card.SerialNumber))
		checkError(err, "opening file")
		w := csv.NewWriter(f)
		writeErr := w.WriteAll(result[card].Transactions)
		checkError(writeErr, "writing CSV to "+f.Name())
		closeErr := f.Close()
		checkError(closeErr, "closing file")
		fmt.Println("wrote CSV", f.Name())
	}
}

func main() {
	flag.Parse()
	if *email == "" || *password == "" {
		fmt.Fprintf(os.Stderr, "Please provide an email and a password\n")
		os.Exit(2)
	}
	client, err := clipper.NewClient(*email, *password)
	checkError(err, "creating client")
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	//result, err := client.Transactions(ctx)
	//checkError(err, "getting transactions")
	cards, err := client.Cards(ctx)
	checkError(err, "getting cards")
	fmt.Printf("cards: %#v\n", cards)
}
