package http_session

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPSessionCookie(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie := r.Header.Get("Cookie")
		if cookie != "" {
			c, err := r.Cookie("testCookie")
			if err != nil {
				t.Fatal(err)
			}
			if c.Value != "testCookieValue" {
				t.Errorf("Cookie value is invalid: %v", cookie)
			}
		} else {
			newCookie := http.Cookie{Name: "testCookie", Value: "testCookieValue"}
			http.SetCookie(w, &newCookie)
		}
	}))
	defer ts.Close()

	session := New()
	response, err := session.Request("get", ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if !checkCookie(response.Raw.Cookies(), "testCookie", "testCookieValue") {
		t.Errorf("Response cookie is invalid: %v", response.Raw.Header.Get("testCookie"))
	}
	if len(session.Cookies) != 1 || session.Cookies[0].Name != "testCookie" || session.Cookies[0].Value != "testCookieValue" {
		t.Errorf("Session cookie is invalid: %v", session.Cookies)
	}

	session2 := New()
	_, err = session2.Request("get", ts.URL, http.Cookie{Name: "testCookie", Value: "testCookieValue"})
	if err != nil {
		t.Fatal(err)
	}

	session3 := New()
	_, err = session3.Request("get", ts.URL, &http.Cookie{Name: "testCookie", Value: "testCookieValue"})
	if err != nil {
		t.Fatal(err)
	}

	_, err = session.Request("get", ts.URL)
	if err != nil {
		t.Fatal(err)
	}
}

func checkCookie(cookies []*http.Cookie, cookieName, cookieValue string) bool {
	for _, v := range cookies {
		if v.Name == cookieName && v.Value == cookieValue {
			return true
		}
	}
	return false
}
