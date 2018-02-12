package clipper

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/kevinburke/rest"
	pdfcontent "github.com/unidoc/unidoc/pdf/contentstream"
	"github.com/unidoc/unidoc/pdf/core"
	pdf "github.com/unidoc/unidoc/pdf/model"
	"golang.org/x/text/encoding/charmap"
)

// Found by trial and error from PDF.
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
	if len(parts) == 8 && parts[1] == "" && parts[2] == "" && parts[3] == "" && strings.Contains(parts[7], "Page ") {
		// page number on first PDF
		return nil, io.EOF
	}
	if len(parts) < 8 && strings.Contains(parts[0], "If there is a discrepancy in the listing of the card balance") {
		return nil, io.EOF
	}
	if len(parts) != 8 {
		return nil, fmt.Errorf("invalid line: %q", text)
	}
	return parts, nil
}

func getCSV(pages []string) (int, [][]string, error) {
	records := make([][]string, 0)
	num := -1
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
						line++
						continue
					}
					parts := strings.Split(text, " ")
					if len(parts) != 2 {
						rest.Logger.Warn("Unexpected line text", "line", line, "text", text)
						line++
						continue
					}
					var err error
					num, err = strconv.Atoi(parts[1])
					if err != nil {
						rest.Logger.Warn("error reading account number", "line", line, "text", text)
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
					return num, nil, err
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
					return num, nil, err
				}
				records = append(records, parts)
				line++
			}
		}
	}
	return num, records, nil
}

type TransactionData struct {
	AccountNumber int
	Transactions  [][]string
}

// ParsePDF parses r (a stream of PDF encoded data) and returns a list of
// transaction records suitable for encoding in a CSV file.
//
// Each row in the output will have 8 columns. Note, the transaction data in the
// PDF is not well validated; as long as it has 8 columns (or close to it), the
// file will be returned as is.
func ParsePDF(r io.ReadSeeker) (TransactionData, error) {
	pages, err := extractPDFText(r)
	if err != nil {
		return TransactionData{}, err
	}
	accountNumber, records, err := getCSV(pages)
	if err != nil {
		return TransactionData{}, err
	}
	return TransactionData{
		AccountNumber: accountNumber,
		Transactions:  records,
	}, nil
}
