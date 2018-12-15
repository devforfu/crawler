package fetch

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

type Response struct {
    content *string
    success bool
    error *string
}

func success(content string) Response {
    return Response{&content,true,nil}
}

func failure(err string) Response {
    return Response{nil, false,&err}
}

func Fetch(url string) Response {
    resp, err := http.Get(url)
    if err != nil {
        errMsg := fmt.Sprintf("%v", err)
        return failure(errMsg)
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        errMsg := fmt.Sprintf("Status code: %d", resp.StatusCode)
        return failure(errMsg)
    }
    bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return failure("Cannot read response bytes")
    }
    return success(string(bytes))
}

