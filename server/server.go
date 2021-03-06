// Package server serves a PDF-to-CSV downloader for Clipper card data.
package server

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"html/template"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/kevinburke/clipper/assets"
	"github.com/kevinburke/handlers"
	"github.com/kevinburke/nacl"
	"github.com/kevinburke/rest"
)

// DefaultPort is the listening port if no other port is specified.
var DefaultPort = 7065

// The server's Version.
const Version = "0.6"

var homepageTpl *template.Template
var digests map[string][sha256.Size]byte

// hashurl returns a hash of the resource with the given key
func hashurl(key string) template.URL {
	d, ok := digests[strings.TrimPrefix(key, "/")]
	if !ok {
		return ""
	}
	// we don't actually need the whole hash.
	return template.URL("s=" + b64(d[:12]))
}

func b64(digest []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(digest), "=")
}

func init() {
	var err error
	digests, err = assets.Digests()
	if err != nil {
		panic(err)
	}
	homepageHTML := assets.MustAssetString("templates/index.html")
	homepageTpl = template.Must(
		template.New("homepage").Option("missingkey=error").Funcs(template.FuncMap{
			"hashurl": hashurl,
		}).Parse(homepageHTML),
	)

	// Add more templates here.
}

// A HTTP server for static files. All assets are packaged up in the assets
// directory with the go-bindata binary. Run "make assets" to rerun the
// go-bindata binary.
type static struct {
	modTime time.Time
}

var expires = time.Date(2050, time.January, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC1123)

func (s *static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		r.URL.Path = "/static/favicon.ico"
	}
	bits, err := assets.Asset(strings.TrimPrefix(r.URL.Path, "/"))
	if err != nil {
		rest.NotFound(w, r)
		return
	}
	// with the hashurl implementation below, we can set a super-long content
	// expiry and ensure content is never stale.
	if query := r.URL.Query(); query.Get("s") != "" {
		w.Header().Set("Expires", expires)
	}
	http.ServeContent(w, r, r.URL.Path, s.modTime, bytes.NewReader(bits))
}

// Render a template, or a server error.
func render(w http.ResponseWriter, r *http.Request, tpl *template.Template, name string, data interface{}) {
	buf := new(bytes.Buffer)
	if err := tpl.ExecuteTemplate(buf, name, data); err != nil {
		rest.ServerError(w, r, err)
		return
	}
	w.Write(buf.Bytes())
}

type flashMessage struct {
	Error   string
	Success string
}

// NewServeMux returns a HTTP handler that covers all routes known to the
// server.
func NewServeMux(key nacl.Key) http.Handler {
	staticServer := &static{
		modTime: time.Now().UTC(),
	}

	r := new(handlers.Regexp)
	r.Handle(regexp.MustCompile(`(^/static|^/favicon.ico$)`), []string{"GET"}, handlers.GZip(staticServer))
	r.HandleFunc(regexp.MustCompile(`^/$`), []string{"GET"}, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		render(w, r, homepageTpl, "homepage", flashMessage{GetFlashError(w, r, key), GetFlashSuccess(w, r, key)})
	})
	r.HandleFunc(regexp.MustCompile(`^/csv$`), []string{"POST"}, csvUpload(key))
	return r
}
