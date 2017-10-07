package main

import (
  "fmt"
  "log"
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
  // jsonObj.Set("c", "person", "traits")
  jsonObj.ArrayAppend("", "person", "tags")
  jsonObj.Array("person", "tags")
  p.Json = jsonObj
  return p
}

// See about comparing input to what already exists
func (p Person) add_trait(trait_name, trait_text string) {
  current_json := p.Json

  current_json.Set(trait_text, "person", "traits", trait_name)
  // current_json.SetP("si", "person.traits.e")
  // current_json.Set(trait_name, "person", "traits")
  // current_json.Set(trait_text, "person", "traits")

  fmt.Println("=======================")
  fmt.Println(current_json.String())
  fmt.Println("=======================")
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

func get_all_people_with_tag(tag_name string) {

}

func (p Person) get_person_tags() map[string]bool {
  tags := make(map[string]bool)
  children, err := p.Json.S("person", "tags").Children()
  check_err(err)
  for _, child := range children {
  	tags[child.Data().(string)] = true
  }
  return tags
}

func (p Person) get_person_trails() {
  children, err := p.Json.S("tags").ChildrenMap()
  check_err(err)
  for key, child := range children {
  	fmt.Printf("key: %v, value: %v\n", key, child.Data().(string))
  }
}

func (p Person) t_delete_trait(trait_name string) {
  err := p.Json.DeleteP("person.traits.relative")
  check_err(err)
}


func main() {

  p1 := new_person("ben")
  p1.add_tag("friend")
  p1.add_tag("geog 1000")

  p2 := new_person("steve")
  p2.add_trait("relative", "bob")
  p2.add_trait("relative2", "dod")


  fmt.Println(p2.Json.ExistsP("person.traits.relative"))

  p2.t_delete_trait("bob")

  // fmt.Println(p1.Name, "->", p2.Name)
  // fmt.Println(p1.Json.String())
  fmt.Println(p2.Json.String())
  //
  //
  //
  // p2.get_person_trails()


}

func check_err(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
