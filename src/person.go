package main

import (
  "fmt"
  "github.com/Jeffail/gabs"
)

type Person struct {
  Name string
  Json *gabs.Container
}

func new_person(name_in string) *Person {
  p := new(Person)
  p.Name = name_in
  p.Json = gabs.New()
  return p
}

func main() {

  p1 := new_person("ben")

  p2 := new_person("steve")


  fmt.Println(p1.Name, "->", p2.Name)
}
