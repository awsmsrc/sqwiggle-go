package main

import (
  "fmt"
  sqwiggle "github.com/lukeroberts1990/sqwiggle-go"
)

func main(){
  cli := sqwiggle.Client { "FV8dgj7eXcEwjsHP9QgN" }
  s := cli.Get("messages")
  fmt.Println("%v", s)
  s = cli.Get("users")
  fmt.Println("%v", s)
  s = cli.Get("rooms")
  fmt.Println("%v", s)
}
