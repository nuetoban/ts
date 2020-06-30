package main

import (
    "os"
    "fmt"
    "time"
    "strconv"
)

func main() {
    if len(os.Args) > 2 {
        fmt.Println("Usage: ts [TIMESTAMP]")
    }

    if len(os.Args) == 1 {
        fmt.Println(time.Now().Unix())
        return
    }

    i, err := strconv.ParseInt(os.Args[1], 10, 64)
    if err != nil {
        fmt.Println(err)
    }

    unixTimeUTC := time.Unix(i, 0)
    unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC3339)

    fmt.Println(unitTimeInRFC3339)
}
