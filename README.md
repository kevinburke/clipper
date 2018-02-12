# Clipper API

Use this tool to download data about your Cards, as well as transaction history
for each card.

## Usage

```go
client := clipper.NewClient("email", "password")
// You can only access this page twice per day, per Clipper.
transactions := client.Transactions(context.TODO())
for card := range transactions {
	fmt.Println("nickname:", card.Nickname)
	fmt.Printf("txns: %#v\n", transactions[card].Transactions
}
```

## PDF-to-CSV

You can run a server that converts PDF's to CSV files; it's the one that runs at
[clipper-csv.appspot.com](https://clipper-csv.appspot.com).

```
make serve
```

## Install

Use "go get" to install the server.

```
go get github.com/kevinburke/clipper/...
```
