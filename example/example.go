package main

import (
  "fmt"
  sqwiggle "github.com/lukeroberts1990/sqwiggle-go"
)

func main(){
  r := new(sqwiggle.Request)
  fmt.Println("%T hello", r)
}
