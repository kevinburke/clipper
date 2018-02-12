package server

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"

	"github.com/kevinburke/clipper-api/clipper"
	"github.com/kevinburke/handlers"
	"github.com/kevinburke/nacl"
	"github.com/kevinburke/rest"
)

// getSingleValue returns the single value, or writes a 400 Bad Request to the
// client. Returns (value, wroteError).
func getSingleValue(w http.ResponseWriter, r *http.Request, key string, emptyOk bool) (string, bool) {
	vals, ok := r.MultipartForm.Value[key]
	if !ok || len(vals) == 0 {
		if emptyOk {
			return "", false
		}
		rest.BadRequest(w, r, &rest.Error{
			Title:      fmt.Sprintf("Please provide a '%s' parameter", key),
			ID:         "missing_parameter",
			Instance:   r.URL.Path,
			StatusCode: 400,
		})
		return "", true
	}
	if len(vals) > 1 {
		rest.BadRequest(w, r, &rest.Error{
			Title:      fmt.Sprintf("Too many '%s' parameters provided", key),
			ID:         "too_many_parameters",
			Instance:   r.URL.Path,
			StatusCode: 400,
		})
		return "", true
	}
	return vals[0], false
}

var key nacl.Key

func csvUpload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 * 1024 * 1024); err != nil {
		rest.BadRequest(w, r, &rest.Error{
			Title: err.Error(),
			ID:    "bad_upload",
		})
		return
	}
	//csv, wroteError := getSingleValue(w, r, "csv", true)
	//if wroteError {
	//return
	//}
	fileDatas := make(map[string][]byte)

	csvHeaders := r.MultipartForm.File["csv"]
	for i := range csvHeaders {
		header := csvHeaders[i]
		file, err := header.Open()
		if err != nil {
			// todo: hide the error here?
			FlashError(w, err.Error(), key)
			http.Redirect(w, r, r.URL.RequestURI(), http.StatusFound)
			return
		}
		buf := new(bytes.Buffer)
		if _, err := io.Copy(buf, file); err != nil {
			// todo: hide the error here?
			FlashError(w, err.Error(), key)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		fileDatas[header.Filename] = buf.Bytes()
	}
	if len(fileDatas) != 1 {
		FlashError(w, "Please provide exactly one file", key)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	var txnData clipper.TransactionData
	for filename := range fileDatas {
		// should be exactly one key
		contents := fileDatas[filename]
		var err error
		txnData, err = clipper.ParsePDF(bytes.NewReader(contents))
		if err != nil {
			FlashError(w, err.Error(), key)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		break
	}
	buf := new(bytes.Buffer)
	csvWriter := csv.NewWriter(buf)
	if err := csvWriter.WriteAll(txnData.Transactions); err != nil {
		FlashError(w, err.Error(), key)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="clipper-transactions-%d.csv"`, txnData.AccountNumber))
	w.Header().Set("Content-Type", "encoding/csv; charset=utf-8")
	w.WriteHeader(200)
	if _, err := w.Write(buf.Bytes()); err != nil {
		handlers.Logger.Warn("error writing response", "err", err)
	}
}
