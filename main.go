package main

import (
    "log"
    "net/http"
    "net/url"
    "strings"
    "io/ioutil"
    "encoding/json"
    "os"
    "fmt"
)

type Bible struct {
    Scriptures []Scripture `json:"scriptures"`
}

type Scripture struct {
    Title string           `json:"title"`
    Desc string            `json:"desc"`
}

func main() {
    accessToken := "<access_key>"

    raw, err := ioutil.ReadFile("./bible.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var b Bible
    json.Unmarshal(raw, &b)

    counter := 0
    for _, sc := range b.Scriptures {
        fmt.Println(sc.Title)
        fmt.Println(sc.Desc)

        msg := sc.Title + "\n" + sc.Desc
        URL := "https://notify-api.line.me/api/notify"

        u, err := url.ParseRequestURI(URL)
        if err != nil {
            log.Fatal(err)
        }

        c := &http.Client{}

        form := url.Values{}
        form.Add("message", msg)

        body := strings.NewReader(form.Encode())

        req, err := http.NewRequest("POST", u.String(), body)
        if err != nil {
            log.Fatal(err)
        }

        req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
        req.Header.Set("Authorization", "Bearer "+accessToken)

        _, err = c.Do(req)
        if err != nil {
            log.Fatal(err)
        }

        if counter >= 5 {
            break
        }
        counter++
    }
}
