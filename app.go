package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

type Req struct {
	Method      string   `json:"method"`
	Url         string   `json:"url"`
	Body        string   `json:"body"`
	Headers     []Header `json:"headers"`
	ContentType string   `json:"contentType"`
}

type Resp struct {
	Body        string         `json:"body"`
	Cookies     []*http.Cookie `json:"cookies"`
	Log         []string       `json:"log"`
	Status      string         `json:"status"`
	ContentType string         `json:"contentType"`
	Headers     []Header       `json:"headers"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var FileName = "kanapi-pref.json"

var pref *Pref

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// read Preferences from pref.go
}

func (a *App) GetPref() (*Pref, error) {
	var err error
	pref, err = readPreference(FileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return pref, nil
}

func (a *App) SetPref(pref Pref) (*Pref, error) {
	return nil, updatePreference(FileName, pref)
}

// CallApi is the main function that is called by the runtime
func (a *App) CallApi(req *Req) (*Resp, error) {
	var res Resp

	log.Println(req.Method + " call initiated for: " + req.Url)

	//  replace with https://github.com/monaco-io/request or equivalent
	doreq, err := http.NewRequest(req.Method, req.Url, bytes.NewReader([]byte(req.Body)))
	if err != nil {
		log.Fatal("Failed to create request. ", err, '\n')
		return nil, err
	}

	doreq.Header.Set("User-Agent", "kanapi")
	doreq.Header.Set("Accept", "*/*")
	doreq.Header.Set("Connection", "keep-alive")

	doreq.Header.Set("Content-Type", req.ContentType)
	for _, header := range req.Headers {
		if header.Key != "" {
			doreq.Header.Set(header.Key, header.Value)
		}
	}

	dores, err := http.DefaultClient.Do(doreq)
	if err != nil {
		log.Fatal("Failed to make http request: ", err, '\n')
		return nil, err
	}
	defer dores.Body.Close()

	log.Println(req.Method, " call completed for: ", req.Url)
	log.Println("> Status: ", dores.Status)

	res.Status = dores.Status
	res.ContentType = dores.Header.Get("Content-Type")
	res.Headers = parseHeaders(&dores.Header)
	buf, err := io.ReadAll(dores.Body)
	if err != nil {
		log.Fatal("Failed to process response body. ", err, '\n')
		return nil, err
	}
	res.Body = string(buf)

	dump, err := httputil.DumpRequestOut(doreq, true)
	if err != nil {
		log.Fatal(err)
	}
	res.Log = append(res.Log, string(dump))

	return &res, nil

}

func parseHeaders(headers *http.Header) []Header {
	var headarr []Header
	for k, v := range *headers {
		h := Header{}
		h.Key = k
		h.Value = strings.Join(v, "")
		headarr = append(headarr, h)
	}
	return headarr

}
