package main

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kevinburke/rest"
	pdfcontent "github.com/unidoc/unidoc/pdf/contentstream"
	"github.com/unidoc/unidoc/pdf/core"
	pdf "github.com/unidoc/unidoc/pdf/model"
	"golang.org/x/net/html"
	"golang.org/x/net/publicsuffix"
	"golang.org/x/text/encoding/charmap"
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

var positions = []float64{
	28,
	133.71,
	359.24,
	479.05,
	528.38,
	655.88,
	685.78,
	722.22,
}

func findPositionIdx(pos float64) int {
	if pos < 0 || pos > 1100 {
		panic(fmt.Sprintf("invalid pos %v", pos))
	}
	if pos <= positions[0] {
		return 0
	}
	if pos >= positions[len(positions)-1] {
		return len(positions) - 1
	}
	for i := 0; i < len(positions)-1; i++ {
		halfway := positions[i] + (positions[i+1]-positions[i])/2
		if pos < halfway {
			return i
		}
	}
	return len(positions) - 1
}

func howManyTabs(prevPos, curPos float64) int {
	if prevPos >= curPos {
		panic("not expected")
	}
	idx := findPositionIdx(prevPos)
	idx2 := findPositionIdx(curPos)
	return idx2 - idx
}

func extractText(parser *pdfcontent.ContentStreamParser) (string, error) {
	operations, err := parser.Parse()
	if err != nil {
		return "", err
	}
	xPos, yPos := float64(-1), float64(-1)
	inText := false
	txt := ""
	// columnStarts:
	//  28.00 date
	// 133.71 transaction type
	// 359.24 location
	// 479.05 route
	// 528.38 product
	// 655.88 debit
	// 685.78 credit
	// 722.22 balance
	for _, op := range *operations {
		if op.Operand == "BT" {
			inText = true
		} else if op.Operand == "ET" {
			inText = false
		}
		if op.Operand == "Tm" {
			// Text matrix. See here for an explanation of how this relates to
			// drawn software:
			// https://stackoverflow.com/a/17202701/329700
			if len(op.Params) != 6 {
				continue
			}
			// 0-3 are scale/shear for x and y. Typical values are 1 0 0 1.
			// 4 is X offset from the left side.
			// 5 is Y offset from the bottom (origin in doc bottom left corner).
			xfloat, ok := op.Params[4].(*core.PdfObjectFloat)
			if !ok {
				xint, ok := op.Params[4].(*core.PdfObjectInteger)
				if !ok {
					continue
				}
				xfloat = core.MakeFloat(float64(*xint))
			}
			yfloat, ok := op.Params[5].(*core.PdfObjectFloat)
			if !ok {
				yint, ok := op.Params[5].(*core.PdfObjectInteger)
				if !ok {
					continue
				}
				yfloat = core.MakeFloat(float64(*yint))
			}
			if yPos == -1 {
				yPos = float64(*yfloat)
			} else if yPos > float64(*yfloat) {
				txt += "\n"
				xPos = float64(*xfloat)
				yPos = float64(*yfloat)
				continue
			}
			if xPos == -1 {
				xPos = float64(*xfloat)
			} else if xPos < float64(*xfloat) {
				numTabs := howManyTabs(xPos, float64(*xfloat))
				txt += strings.Repeat("\t", numTabs)
				xPos = float64(*xfloat)
			}
		}

		if op.Operand == "Td" || op.Operand == "TD" || op.Operand == "T*" {
			// Move to next line...
			txt += "\n"
		}
		if inText && op.Operand == "TJ" {
			if len(op.Params) < 1 {
				continue
			}
			paramList, ok := op.Params[0].(*core.PdfObjectArray)
			if !ok {
				return "", fmt.Errorf("Invalid parameter type, no array (%T)", op.Params[0])
			}
			for _, obj := range *paramList {
				switch v := obj.(type) {
				case *core.PdfObjectString:
					txt += string(*v)
				case *core.PdfObjectFloat:
					if *v < -100 {
						txt += " "
					}
				case *core.PdfObjectInteger:
					if *v < -100 {
						txt += " "
					}
				}
			}
		} else if inText && op.Operand == "Tj" {
			if len(op.Params) < 1 {
				continue
			}
			param, ok := op.Params[0].(*core.PdfObjectString)
			if !ok {
				return "", fmt.Errorf("Invalid parameter type, not string (%T)", op.Params[0])
			}
			txt += string(*param)
		}
	}

	return txt, nil
}

func extractPDFText(r io.ReadSeeker) ([]string, error) {
	pdfReader, err := pdf.NewPdfReader(r)
	if err != nil {
		return nil, err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return nil, err
	}
	pages := make([]string, numPages)
	decoder := charmap.Windows1252.NewDecoder()
	for i := 1; i <= numPages; i++ {
		page, err := pdfReader.GetPage(i)
		if err != nil {
			return nil, err
		}
		contentStreams, err := page.GetContentStreams()
		if err != nil {
			return nil, err
		}
		pageContentStr := ""

		// If the value is an array, the effect shall be as if all of the
		// streams in the array were concatenated, in order, to form a
		// single stream.
		for _, cstream := range contentStreams {
			pageContentStr += cstream + "\n"
		}

		cstreamParser := pdfcontent.NewContentStreamParser(pageContentStr)
		txt, err := extractText(cstreamParser)
		if err != nil {
			return nil, err
		}
		s, err := decoder.String(txt)
		if err != nil {
			fmt.Printf("Error decoding stream: %q\n", txt)
			return nil, err
		}
		pages[i-1] = strings.TrimSpace(s)
	}
	return pages, nil
}

func parseLine(text string) ([]string, error) {
	parts := strings.Split(text, "\t")
	if len(parts) < 8 && (strings.Contains(parts[0], "If there is a discrepancy in the listing of the card balance") || strings.Contains(parts[1], "Page ")) {
		return nil, io.EOF
	}
	if len(parts) != 8 {
		return nil, fmt.Errorf("invalid line: %q", text)
	}
	return parts, nil
}

func getCSV(pages []string) ([][]string, error) {
	records := make([][]string, 0)
	for i := range pages {
		if i == 0 {
			bs := bufio.NewScanner(strings.NewReader(pages[i]))
			line := 0
			for bs.Scan() {
				text := bs.Text()
				if line == 0 {
					if text != "TRANSACTION HISTORY FOR" {
						rest.Logger.Warn("Unexpected line text", "line", line, "text", text)
					}
					line++
					continue
				}
				if line == 1 {
					if !strings.HasPrefix(text, "CARD ") {
						rest.Logger.Warn("Unexpected line text", "line", line, "text", text)
					}
					line++
					continue
				}
				if line == 2 {
					if !strings.HasPrefix(text, "TRANSACTION TYPE\tLOCATION\tROUTE") {
						rest.Logger.Warn("Unexpected line text", "line", line, "text", text)
					}
					line++
					continue
				}
				parts, err := parseLine(text)
				if err == io.EOF {
					break
				}
				if err != nil {
					return nil, err
				}
				records = append(records, parts)
				line++
			}
		} else {
			bs := bufio.NewScanner(strings.NewReader(pages[i]))
			line := 0
			for bs.Scan() {
				text := bs.Text()
				if line == 0 {
					if !strings.HasPrefix(text, "TRANSACTION TYPE\tLOCATION\tROUTE") {
						rest.Logger.Warn("Unexpected line text", "line", line, "text", text)
					}
					line++
					continue
				}
				parts, err := parseLine(text)
				if err == io.EOF {
					break
				}
				if err != nil {
					return nil, err
				}
				records = append(records, parts)
				line++
			}
		}
	}
	return records, nil
}
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
	_, discardErr := io.Copy(ioutil.Discard, resp.Body)
	checkError(discardErr, "reading entire body")
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
	resp2, err := client.Do(req)
	checkError(err, "making login request")
	if resp2.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Bad status: want 200 got %d\n", resp2.StatusCode)
		io.Copy(os.Stderr, resp2.Body)
		os.Exit(2)
	}
	dashboardData, err := ioutil.ReadAll(resp2.Body)
	checkError(err, "reading response body")
	closeErr = resp2.Body.Close()
	checkError(closeErr, "closing body")
	cards, err := GetCards(bytes.NewReader(dashboardData))
	checkError(err, "finding cards")
	viewState2, err := findViewState(bytes.NewReader(dashboardData))
	checkError(err, "finding viewState value")
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	for i := range cards {
		time.Sleep(2 * time.Second)
		history := fmt.Sprintf("mainForm:j_idt95:%d:seeHistorySixty", i)
		data := url.Values{}
		data.Set("javax.faces.ViewState", viewState2)
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
		checkError(err, "creating request")
		req = req.WithContext(ctx)
		req.Header.Set("User-Agent", userAgent)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Referer", "https://www.clippercard.com/ClipperCard/dashboard.jsf")
		resp, err := client.Do(req)
		checkError(err, "making data request")
		if resp.StatusCode != 200 {
			fmt.Fprintf(os.Stderr, "Bad status: want 200 got %d\n", resp.StatusCode)
			io.Copy(os.Stderr, resp.Body)
			os.Exit(2)
		}
		ctype := resp.Header.Get("Content-Type")
		typ, _, err := mime.ParseMediaType(ctype)
		checkError(err, "reading content type")
		if typ != "application/pdf" {
			fmt.Fprintf(os.Stderr, "req %d: Bad response content-type: want pdf got %s\n", i, ctype)
			resp.Header.Write(os.Stderr)
			io.Copy(os.Stderr, resp.Body)
			os.Exit(2)
		}
		pdfBody, err := ioutil.ReadAll(resp.Body)
		checkError(err, "reading response body")
		pages, err := extractPDFText(bytes.NewReader(pdfBody))
		checkError(err, "parsing PDF number "+strconv.Itoa(i)+": "+string(pdfBody))
		closeErr := resp.Body.Close()
		checkError(closeErr, "closing body")
		csv, err := getCSV(pages)
		checkError(err, "parsing CSV")
		fmt.Printf("%#v\n", csv)
	}
}
