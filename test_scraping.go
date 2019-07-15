package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
)

func main() {
  // get page source
  source_url := "https://www.digital.archives.go.jp/DAS/meta/era"
  doc, err := goquery.NewDocument(source_url)
  if err != nil {
    fmt.Println(err)
  }
  class := ".mng-table-3 mng-table-col mng-table-head-black-1 mng-table-stripe"
  text := doc.Find(class).Text()
  fmt.Println(text)
}
