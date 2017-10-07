package main

import (
  "fmt"
  "errors"
  "github.com/Jeffail/gabs"
)

// Name
// Traits
// Tags/Groups
type Person struct {
  Name string
  Json *gabs.Container
}

func new_person(name_in string) *Person {
  p := new(Person)
  p.Name = name_in
  jsonObj := gabs.New()
  jsonObj.Set(name_in, "person", "name")
  jsonObj.Set("", "person", "traits")
  jsonObj.ArrayAppend("", "person", "tags")
  jsonObj.Array("person", "tags")
  p.Json = jsonObj
  return p
}

// See about comparing input to what already exists
func (p Person) add_trait(trait_name, trait_text string) {
  current_json := p.Json
  current_json.Set(trait_text, "person", "traits", trait_name)
}

// See about comparing input to what already exists
func (p Person) add_tag(tag_name string) {
  current_json := p.Json
  current_json.ArrayAppend(tag_name, "person", "tags")
}

// If the name will be unique it will change
func (p Person) change_name(name string) error {
  if acceptable_name(name) {
    p.Name = name
    return nil
  } else {
    return errors.New("This name has been previously assinged, try another.")
  }
}

// This will have some content at some point to ensure uniquity among names
func acceptable_name(name string) bool {
  if true != false {
    return true
  } else {
    return false
  }
}

func (p Person) get_person_tags() {

}

func (p Person) get_person_trails() {

}

func main() {

  p1 := new_person("ben")

  p2 := new_person("steve")

  fmt.Println(p1.Name, "->", p2.Name)
  fmt.Println(p1.Json.String())
  fmt.Println(p2.Json.String())

  p1.add_tag("friend")
  p1.add_tag("geog 1000")
  fmt.Println(p1.Name, "->", p2.Name)
  fmt.Println(p1.Json.String())
  fmt.Println(p2.Json.String())
}
