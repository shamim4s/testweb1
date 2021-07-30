package main

import (
    "log"
    "os"
    // "io/ioutil"
    // "fmt"
)

func main() {
    // If the file doesn't exist, create it, or append to the file
    f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    if _, err := f.Write([]byte("appended some data\n")); err != nil {
        log.Fatal(err)
    }

    // text, err := os.ReadFile("access.log")
    // if err != nil {
    //     log.Fatal(err)
    // }


    // log.Printf ( "test %s", f)

  

    // os.Stdout.Write("tetst %s", text)

    if err := f.Close(); err != nil {
        log.Fatal(err)
    }

    // read,err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)

    
    //  data, err := ioutil.ReadFile("access.log")
    // if err != nil {
    //     log.Panicf("failed reading data from file: %s", err)
    // }
    // fmt.Printf("\nLength: %d bytes", len(data))
    
    text, err := os.ReadFile("access.log")
    if err != nil {
        log.Fatal(err)
    }

    log.Printf ( "\ntest 1111111111111111111111111111%s", text)


}