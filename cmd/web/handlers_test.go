package main

import (
		"io/ioutil"
		"log"
		"net/http" 	
		"net/http/httptest"
		"testing"
)

func TestPing(t *testing.T) {
		app := &application {
				errorLog: log.New(ioutil.Discard, "", 0),
				infoLog: log.New(ioutil.Discard, "", 0),
		}
		ts := httptest.NewTLSServer(app.routes())
		defer ts.Close()
		rs, err := ts.Client().Get(ts.URL + "/ping")
		if err != nil {
				t.Fatal(err)
		}
		if rs.StatusCode != http.StatusOK {
				t.Errorf("Expected %d. Got %d", http.StatusOK, rs.StatusCode)
		}
		defer rs.Body.Close()
		body, err := ioutil.ReadAll(rs.Body)
		if err != nil {
				t.Fatal(err)
		}
		if string(body) != "OK" {
				t.Errorf("Expected %q", "OK")
		}
}
