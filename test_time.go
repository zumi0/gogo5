package main

import (
  "fmt"
  "time"
)

func main() {
  date := time.Now()
  fmt.Println(date)
  fix_date := fmt.Sprintf("%d年%d月%d日の情報",date.Year(),int(date.Month()),int(date.Day()))
  fmt.Println(fix_date)
}
