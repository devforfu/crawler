package main

import "fmt"

func main() {
    result, err := noPanicDivision(1, 0)
    fmt.Printf("%v, %v", result, err)
}

func noPanicDivision(a, b int) (result int, err error) {
    type zeroDivision struct {}

    defer func() {
        if p := recover(); p != nil {
            switch p {
            case nil:
                // do nothing
                result = 0
            case zeroDivision{}:
                err = fmt.Errorf("attempt to divide by zero: a=%d/b=%d", a, b)
            default:
                panic(p)
            }
        }
    }()

    if b == 0 {
        panic(zeroDivision{})
    } else {
        result = a / b
    }

    return
}
