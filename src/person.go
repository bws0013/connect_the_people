package main

import (
  "fmt"
  "log"
  "strings"
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

var (
  people_map = make(map[string]Person)
)

func new_person(name_in string) *Person {
  p := new(Person)
  p.Name = name_in
  jsonObj := gabs.New()
  jsonObj.Set(name_in, "person", "name")
  // jsonObj.Set("c", "person", "traits")
  // jsonObj.ArrayAppend("", "person", "tags")
  jsonObj.Array("person", "tags")
  p.Json = jsonObj
  return p
}

// See about comparing input to what already exists
func (p Person) add_trait(trait_path, trait_name, trait_text string) {
  current_json := p.Json

  // Get the full dotted path to use for all operations
  path := create_path(trait_path, trait_name)

  // 2 situations, the trait already exists or it doesnt
  // If it exists array it (there is a check for the array also)
  // If it doesnt exist add it

  // There are 2 additional situations if the trait exists
  // The array already exists (add to it)
  // The array doesn't already exist (make it and add to it)

  // It it exists, else it does not
  if current_json.ExistsP(path) {

    // Get an error if there is no array
    _, err := current_json.Path(path).Children()

    // If there is an error we need to make the array
    if err != nil {
      // Get the original value (that isnt an array)
      current_val := current_json.Path(path).Data()
      // Create an array
      current_json.ArrayP(path)
      // Add the original value to the array
      current_json.ArrayAppendP(current_val, path)
    }

    // Add the value we had originally wanted to add
    current_json.ArrayAppendP(trait_text, path)
  } else {
    current_json.SetP(trait_text, path)
  }
}

// If the name will be unique it will change
func (p Person) change_name(name string) error {
  if acceptable_person_name(name) {
    p.Name = name
    return nil
  } else {
    return errors.New("This name has been previously assinged, try another.")
  }
}

// This will have some content at some point to ensure uniquity among names
func acceptable_person_name(name string) bool {
  if true != false {
    return true
  } else {
    return false
  }
}

// See about comparing input to what already exists
func (p Person) add_tag(tag_name string) error {
  current_json := p.Json
  if p.acceptable_tag_name(tag_name) {
    current_json.ArrayAppend(tag_name, "person", "tags")
    return nil
  } else {
    return errors.New("This tag has been previously assinged, try another.")
  }

}

// This will have some content at some point to ensure uniquity among tags
func (p Person) acceptable_tag_name(name string) bool {

  tag_set := p.get_person_tags()

  // If the tag is currently associated with the person do not add it again
  // If the statement is true then the tag is associated with the person
  if tag_set[name] == false {
    return true
  } else {
    return false
  }
}

func get_all_people_with_tag(tag_name string) {
  names_with_tag := make([]string, 0)

  for _, p := range people_map {
    local_tags := p.get_person_tags()
    if local_tags[tag_name] == true {
      names_with_tag = append(names_with_tag, p.get_name())
    }
  }

  fmt.Println(names_with_tag)
}

func (p Person) get_person_tags() map[string]bool {
  // _, err := current_json.Path(path).Children()
  current_json := p.Json
  tags := make(map[string]bool)
  children, err := current_json.Path("person.tags").Children()
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

func (p Person) get_name() string {
  return p.Json.Path("person.name").Data().(string)
}

// This may be un-needed and we may just store directly to map
// More investigation is required
func (p Person) add_to_people_map() {
  if _, exists := people_map[p.get_name()]; !exists {
    people_map[p.get_name()] = p
  }
}

func (p Person) t_delete_trait(trait_name string) {
  err := p.Json.DeleteP("person.traits.relative")
  check_err(err)
}

func main() {

  traits_path := "person.traits"

  p1 := new_person("ben")
  p1.add_tag("friend")
  p1.add_tag("geog 1000")

  p2 := new_person("steve")
  p2.add_trait(traits_path, "relative.brother", "uno")
  // p2.add_trait(traits_path, "relative.sister", "dos")
  p2.add_trait(traits_path, "relative.brother", "tres")
  p2.add_trait(traits_path, "relative.brother", "quad")
  p2.add_trait(traits_path, "location.current", "md")

  p1.add_tag("dog person")
  p1.add_tag("cat person")
  p2.add_tag("cat person")
  //p2.add_trait(traits_path, "relative", "tre")

  // fmt.Println(p2.Json.ExistsP("person.traits.relative"))

  // p2.t_delete_trait("bob")

  // fmt.Println(p1.Name, "->", p2.Name)
  // fmt.Println(p1.Json.String())
  // fmt.Println(p2.Json.String())
  //
  //
  //
  // p2.get_person_trails()
  p1.add_to_people_map()
  p2.add_to_people_map()

  get_all_people_with_tag("cat person")

}

// Support methods

// Import people json objects from a file
func import_people_from_file() {

}

// Export people json objects to a file
func export_people_to_file() {

}

// Create a path to a particular location within the json
func create_path(path, addition string) string {
  if !strings.HasSuffix(path, ".") {
    path = path + "."
  }
  if strings.HasPrefix(addition, ".") {
    addition = addition[1:]
  }
  return (path + addition)
}

// Break if there is an error passed in
func check_err(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
