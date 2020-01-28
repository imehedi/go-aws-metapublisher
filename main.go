package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

func metaPublisher(resourceKey string) {
    if resourceKey == "list" {
        resourceKey = ""
    }

    res, err := http.Get("http://169.254.169.254/latest/meta-data/" + resourceKey)
    if err != nil {
        fmt.Printf("an error has occurred: %t\n", err)
        return
    }

    b, e := ioutil.ReadAll(res.Body)
    content := string(b)
    if e != nil {
        fmt.Printf("an error has occurred: %t\n", err)
        return
    }

    if resourceKey == "list" {
        fmt.Printf("%s\n\n", content)
    } else {
        fmt.Printf("%s: %s\n\n", resourceKey, content)
    }
}

func showHelp() {
    fmt.Println("Please supply resource name or resource group name as an argument to the program. " +
        "A full list can be found by passing " +
        "list' as an argument (i.e. ./app list) or by visiting the following link: " +
        "https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instancedata-data-retrieval.html")
}

func main() {
    if len(os.Args) == 1 {
        showHelp()
        os.Exit(0)
    }

    metaPublisher(strings.ToLower(os.Args[1]))
}
