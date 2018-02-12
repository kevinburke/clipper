package clipper

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"golang.org/x/net/html"
)

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
						num, err := strconv.ParseInt(z.Token().Data, 10, 64)
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

func getCards(r io.Reader) ([]Card, error) {
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
