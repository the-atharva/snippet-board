package main 

import (
		"io/ioutil"
		"net/http"
		"net/http/httptest"
		"testing"
)

func TestSecureHeaders(t *testing.T) {
		rr := httptest.NewRecorder()
		r, err := http.NewRequest("Get", "/", nil)
		if err != nil {
				t.Fatal(err)
		}
		next := http.HandlerFunc(ping)
		secureHeaders(next).ServeHTTP(rr, r)
		rs := rr.Result()
		frameOptions := rs.Header.Get("X-Frame-Options")
		if frameOptions != "deny" {
				t.Errorf("Expected: %q. Got %q", "deny", frameOptions)
		}
		xssProtection := rs.Header.Get("X-XSS-Protection")
		if xssProtection != "1; mode=block" {
				t.Errorf("EXpected %q. Got %q", "1; mode=block", xssProtection)
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
				t.Errorf("Expected %qq", "OK")
		}
}
