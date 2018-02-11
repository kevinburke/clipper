package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kevinburke/rest"
	"golang.org/x/net/html"
	"golang.org/x/net/publicsuffix"
	"golang.org/x/sync/errgroup"
)

const host = "https://www.clippercard.com"
const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:58.0) Gecko/20100101 Firefox/58.0"

type Card struct {
	Nickname            string
	SerialNumber        int
	Status              string
	Reason              string
	Type                string
	CashValueCents      int
	AutoloadAmountCents int
}

func findViewState(r io.Reader) (string, error) {
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return "", errors.New("ViewState not found")
		case html.SelfClosingTagToken:
			tok := z.Token()
			if tok.Data != "input" {
				continue
			}
			foundViewState := false
			for i := range tok.Attr {
				if tok.Attr[i].Key == "name" && tok.Attr[i].Val == "javax.faces.ViewState" {
					foundViewState = true

				}
			}
			if !foundViewState {
				continue
			}
			for i := range tok.Attr {
				if tok.Attr[i].Key == "value" {
					return tok.Attr[i].Val, nil
				}
			}
		}
	}
}

func checkError(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s: %v\n", msg, err)
		os.Exit(2)
	}
}

func setNickSerialNumber(z *html.Tokenizer, card *Card) error {
	depth := 1
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return fmt.Errorf("reached document end, nothing found: %v", z.Token())
		case html.StartTagToken:
			depth++
			tok := z.Token()
			if tok.Data != "div" {
				continue
			}
			for i := range tok.Attr {
				if tok.Attr[i].Key == "class" && tok.Attr[i].Val == "infoDiv" {
					tt = z.Next()
					for tt == html.TextToken {
						tt = z.Next()
					}
					if tt != html.StartTagToken {
						return fmt.Errorf("expected start tag token, got %#v", z.Token().String())
					}
					tok = z.Token()
					depth++
					if tok.Data != "div" || len(tok.Attr) != 1 || tok.Attr[0].Key != "class" || tok.Attr[0].Val != "fieldName" {
						return fmt.Errorf("expected start tag token, got %#v", tok.String())
					}
					tt = z.Next()
					if tt != html.TextToken {
						return errors.New("expected text token")
					}
					name := z.Token().Data
					switch name {
					case "Serial Number:":
						tt = z.Next()
						if tt != html.EndTagToken {
							return fmt.Errorf("expected end tag token, got %#v", z.Token().String())
						}
						depth--
						tt = z.Next()
						for tt == html.TextToken {
							tt = z.Next()
						}
						if tt != html.StartTagToken {
							return fmt.Errorf("expected start tag token, got %#v", z.Token().String())
						}
						depth++
						tt = z.Next()
						if tt != html.TextToken {
							return errors.New("expected text token")
						}
						num, err := strconv.Atoi(z.Token().Data)
						if err != nil {
							return err
						}
						card.SerialNumber = num
						continue

					case "Card Nickname:":
						tt = z.Next() // <div class="fieldData field90">
						if tt != html.EndTagToken {
							return fmt.Errorf("expected end tag token, got %#v", z.Token().String())
						}
						depth--
						tt = z.Next()
						for tt == html.TextToken {
							tt = z.Next()
						}
						if tt != html.StartTagToken {
							return fmt.Errorf("expected start tag token, got %#v", z.Token().String())
						}
						tok = z.Token()
						depth++
						if tok.Data != "div" || len(tok.Attr) != 1 || tok.Attr[0].Key != "class" || tok.Attr[0].Val != "fieldData field90" {
							return errors.New("expected fieldData field90 token")
						}
						tt = z.Next() // <span class="displayName">
						for tt == html.TextToken {
							tt = z.Next()
						}
						if tt != html.StartTagToken {
							return fmt.Errorf("expected start tag token, got %#v", z.Token().String())
						}

						tok = z.Token()
						depth++
						if tok.Data != "span" || len(tok.Attr) != 1 || tok.Attr[0].Key != "class" || tok.Attr[0].Val != "displayName" {
							return errors.New("expected span tag token")
						}
						tt = z.Next() // the actual name
						if tt == html.EndTagToken {
							// no nickname
							depth--
							continue
						}
						if tt != html.TextToken {
							return fmt.Errorf("expected text token, got %#v\n", z.Token().String())
						}
						tok = z.Token()
						card.Nickname = tok.Data
					}
				}
			}
		case html.EndTagToken:
			depth--
			if depth <= 0 {
				return nil
			}
		}
	}
}

func setCardInfo(z *html.Tokenizer, card *Card) error {
	depth := 1
	hitSpacer := false
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return fmt.Errorf("reached document end, nothing found: %v", z.Token())
		case html.StartTagToken:
			tok := z.Token()
			depth++
			if hitSpacer || tok.Data != "div" {
				continue
			}
			for i := range tok.Attr {
				if tok.Attr[i].Key == "class" && tok.Attr[i].Val == "spacer" {
					hitSpacer = true
					continue
				}
				if tok.Attr[i].Key == "class" && tok.Attr[i].Val == "infoDiv" {
					tt = z.Next()
					for tt == html.TextToken {
						tt = z.Next()
					}
					if tt != html.StartTagToken {
						return fmt.Errorf("expected start tag token, got %#v", z.Token().String())
					}
					tok = z.Token()
					depth++
					if tok.Data != "div" || len(tok.Attr) != 1 || tok.Attr[0].Key != "class" || tok.Attr[0].Val != "fieldName" {
						return fmt.Errorf("expected start tag token, got %#v", tok.String())
					}
					tt = z.Next()
					if tt != html.TextToken {
						return errors.New("expected text token")
					}
					name := z.Token().Data
					tt = z.Next()
					if tt != html.EndTagToken {
						return fmt.Errorf("expected end tag token, got %#v", z.Token().String())
					}
					depth--
					tt = z.Next()
					for tt == html.TextToken {
						tt = z.Next()
					}
					if tt != html.StartTagToken {
						return fmt.Errorf("expected start tag token, got %#v", z.Token().String())
					}
					depth++
					tt = z.Next()
					if tt != html.TextToken {
						return errors.New("expected text token")
					}
					data := z.Token().Data
					switch name {
					case "Type:":
						card.Type = data
					case "Status:":
						card.Status = data
					case "Reason:":
						card.Reason = data
					default:
						fmt.Println("unknown name", name)
					}
					continue
				}
			}
		case html.EndTagToken:
			depth--
			if depth <= 0 {
				return nil
			}
		}
	}
}

func GetCards(r io.Reader) ([]Card, error) {
	z := html.NewTokenizer(r)
	cards := make([]Card, 0)
	card := new(Card)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return cards, nil
		case html.StartTagToken:
			tok := z.Token()
			if tok.Data != "div" {
				continue
			}
			for i := range tok.Attr {
				if tok.Attr[i].Key == "class" && tok.Attr[i].Val == "darkGreyCardHeader" {
					if err := setNickSerialNumber(z, card); err != nil {
						return nil, err
					}
				}
				if tok.Attr[i].Key == "class" && tok.Attr[i].Val == "cardInfo" {
					if err := setCardInfo(z, card); err != nil {
						return nil, err
					}
					cards = append(cards, *card)
					card = new(Card)
				}
			}
		}
	}
}

var email = "kev@inburke.com"
var password = "g82N99GCJ37Mdqo"

func main() {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	checkError(err, "creating cookie jar")

	client := &http.Client{
		Jar:       jar,
		Transport: rest.DefaultTransport,
	}

	req, err := http.NewRequest("GET", host+"/ClipperCard/loginFrame.jsf", nil)
	checkError(err, "creating req")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "*/*")
	resp, err := client.Do(req)
	checkError(err, "getting Clipper site")
	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Bad status: want 200 got %d\n", resp.StatusCode)
		io.Copy(os.Stderr, resp.Body)
		os.Exit(2)
	}
	viewState, err := findViewState(resp.Body)
	checkError(err, "finding viewState value")
	closeErr := resp.Body.Close()
	checkError(closeErr, "closing body")

	data := url.Values{}
	data.Set("j_idt13", "j_idt13")
	data.Set("j_idt13:username", email)
	data.Set("j_idt13:password", password)
	data.Set("javax.faces.behavior.event", "action")
	data.Set("javax.faces.partial.ajax", "true")
	data.Set("javax.faces.partial.execute", "j_idt13:submitLogin j_idt13:username j_idt13:password")
	data.Set("javax.faces.partial.render", "j_idt13:err")
	data.Set("javax.faces.source", "j_idt13:submitLogin")
	data.Set("javax.faces.ViewState", viewState)

	req, err = http.NewRequest("POST", host+"/ClipperCard/loginFrame.jsf", strings.NewReader(data.Encode()))
	checkError(err, "making POST request")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Referer", "https://www.clippercard.com/ClipperCard/loginFrame.jsf")
	req.Header.Set("Faces-Request", "partial-ajax")
	resp, err = client.Do(req)
	checkError(err, "making login request")
	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Bad status: want 200 got %d\n", resp.StatusCode)
		io.Copy(os.Stderr, resp.Body)
		os.Exit(2)
	}
	dashboardData, err := ioutil.ReadAll(resp.Body)
	checkError(err, "reading response body")
	closeErr = resp.Body.Close()
	checkError(closeErr, "closing body")
	cards, err := GetCards(bytes.NewReader(dashboardData))
	checkError(err, "finding cards")
	fmt.Printf("cards: %#v\n", cards)
	viewState, err = findViewState(bytes.NewReader(dashboardData))
	checkError(err, "finding viewState value")
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	group, errctx := errgroup.WithContext(ctx)
	for i := range cards {
		i := i
		history := fmt.Sprintf("mainForm:j_idt95:%d:seeHistorySixty", i)
		group.Go(func() error {
			data := url.Values{}
			data.Set("javax.faces.ViewState", viewState)
			data.Set("mainForm", "mainForm")
			data.Set("mainForm:password", "")
			data.Set("mainForm:j_idt65", strconv.Itoa(i))
			for j := range cards {
				data.Set(fmt.Sprintf("mainForm:j_idt95:%d:cardName", j), cards[j].Nickname)
			}
			data.Set("mainForm:newEcashAmtVal", "0.0")
			data.Set("mainForm:newParkingPurseAmtVal", "0.0")
			data.Set(history, history)
			data.Set("mainForm:username", email)
			req, err := http.NewRequest("POST", host+"/ClipperCard/dashboard.jsf", strings.NewReader(data.Encode()))
			if err != nil {
				return err
			}
			req = req.WithContext(errctx)
			req.Header.Set("User-Agent", userAgent)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
			req.Header.Set("Accept", "*/*")
			req.Header.Set("Referer", "https://www.clippercard.com/ClipperCard/dashboard.jsf")
			resp, err = client.Do(req)
			checkError(err, "making data request")
			if resp.StatusCode != 200 {
				fmt.Fprintf(os.Stderr, "Bad status: want 200 got %d\n", resp.StatusCode)
				io.Copy(os.Stderr, resp.Body)
				os.Exit(2)
			}
			f, err := os.Create("data-" + strconv.Itoa(i))
			checkError(err, "opening data file")
			if _, err := io.Copy(f, resp.Body); err != nil {
				return err
			}
			closeErr := resp.Body.Close()
			checkError(closeErr, "closing body")
			fmt.Println("wrote", f.Name())
			return nil
		})
	}
	groupErr := group.Wait()
	checkError(groupErr, "making parallel requests")
}
