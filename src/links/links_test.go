package links

import "testing"

func TestExtract(t *testing.T) {
    tests := []struct{
        url string
        failure bool
    }{
        {"http://gopl.io", false},
        {"wrongURL",true},
    }
    for _, test := range tests {
        result, err := Extract(test.url)
        if test.failure && err == nil {
            t.Errorf("Expected to failed but successeded: %s", test.url)
        } else if !test.failure && err != nil {
            t.Errorf("Expected to successed but failed: %s", test.url)
        } else if !test.failure && len(result) == 0 {
            t.Errorf("No links discovered: %s", test.url)
        }
    }
}
