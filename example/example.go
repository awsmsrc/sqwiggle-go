package main

import (
  "fmt"
  sqwiggle "github.com/lukeroberts1990/sqwiggle-go"
)

func main(){
  cli := sqwiggle.Client { "Ps4P4pghZFpqPMgDXvzD" }
  s := cli.Users()
  fmt.Println("%v", s)
}
