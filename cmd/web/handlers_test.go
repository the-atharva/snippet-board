package main

import (
		"io/ioutil"
		"net/http" 	
		"net/http/httptest"
		"testing"
)

func TestPint(t *testing.T) {
		rr := httptest.NewRecorder()
		r, err := http.NewRequest("Get", "/", nil)
		if err != nil {
				t.Fatal(err)
		}
		ping(rr, r)
		rs := rr.Result()
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
