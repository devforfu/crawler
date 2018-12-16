package fetch

import "testing"

func TestFetch(t *testing.T) {
    tests := []struct {
        url string
        success bool
    }{
        {"http://golang.org", true},
        {"https://google.com", true},
        {"nonExistingPage", false},
    }
    for _, test := range tests {
        response := Fetch(test.url)
        if response.Success != test.success {
            t.Errorf("Test case failed: %v", test)
        }
    }
}