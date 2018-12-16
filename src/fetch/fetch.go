package fetch

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

type Response struct {
    Content *string
    Success bool
    Error *string
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
        errMsg := fmt.Sprintf("%v\n", err)
        return failure(errMsg)
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        errMsg := fmt.Sprintf("Status code: %d\n", resp.StatusCode)
        return failure(errMsg)
    }
    bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return failure("Cannot read response bytes\n")
    }
    return success(string(bytes))
}

