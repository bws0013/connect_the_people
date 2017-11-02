package main

import (
  "os"
  "fmt"
  "log"
  "errors"
  "strings"
  "io/ioutil"
  "github.com/Jeffail/gabs"
)

// Name OR Maybe some kind of id and just have name be its own thang
// Traits
// Tags/Groups
type Person struct {
  Name string // We may want this to be some kind of id
  Json *gabs.Container
}

// The global people_map is just used for testing purposes at the moment
var (
  people_map = make(map[string]Person)
)

// Create a new person given a name, the name will probably be some kind of id later
func new_person(name_in string) *Person {
  p := new(Person)
  p.Name = name_in
  new_json := gabs.New()
  new_json.Set(name_in, "person", "name")
  // new_json.Set("c", "person", "traits")
  // new_json.ArrayAppend("", "person", "tags")
  new_json.Array("person", "tags")
  p.Json = new_json
  return p
}

// Create a new person from a byte array
func new_person_from_data(person_data []byte) *Person {

  p := new(Person)

  new_json, err := gabs.ParseJSON(person_data)
  check_err(err)

  p.Name = new_json.Path("person.name").String()
  p.Json = new_json

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

// Given a tag return the name of all of those people on the param map who have it
func get_all_people_with_tag(tag_name string, local_people_map map[string]Person) []string {
  names_with_tag := make([]string, 0)

  for _, p := range local_people_map {
    local_tags := p.get_person_tags()
    if local_tags[tag_name] == true {
      names_with_tag = append(names_with_tag, p.get_name())
    }
  }

  return names_with_tag
}

// Returns a map (set) of all of the tags of a particular person
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

// TODO the get_person_traits method
// This should get the traits of a person, it currently doesnt
func (p Person) get_person_traits() map[string]*gabs.Container {
  current_json := p.Json

  json_map := current_json.Path("person.traits")

  children, err := json_map.ChildrenMap()
  check_err(err)
  return children
}

// Get the name of a person
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

// Simplifying delete
func (p Person) delete_trait(trait_name string) {
  current_json := p.Json

  path := create_path("person.traits", trait_name)

  // If we can see the object, ie it has a mapped value
  if current_json.ExistsP(path) {
    err := current_json.DeleteP(path)
    if err == nil {
      return
    }
  } else {
    // If we cant see the object, ie it might not have a mapped value
    p.delete_single_trait_object(trait_name)
    return
  }

  path_elements := strings.Split(path, ".")
  path = strings.Join(path_elements[:len(path_elements) - 1], ".")
  element_to_delete := path_elements[len(path_elements) - 1]

  my_data := current_json.Path(path)
  element_depth := obtain_array_count(my_data.String())
  current_val := current_json.Path(path)
  for i := 0; i < element_depth - 1; i++ {
    current_val = current_val.Index(0)
  }

  kinder, err := current_val.Children()
  check_err(err)

  for _, kind := range kinder {
    if kind.ExistsP(element_to_delete) {
      err = kind.DeleteP(element_to_delete)
      check_err(err)
    }
  }
}

/*
  Delete a trait from an array where the trait is the smallest objects
  Example: Deleting "1" from ["1", {"2":"yo"}, "3"]
*/
func (p Person) delete_single_trait_object(trait_name string) {
  current_json := p.Json

  path := create_path("person.traits", trait_name)

  path_elements := strings.Split(path, ".")
  path = strings.Join(path_elements[:len(path_elements) - 1], ".")
  element_to_delete := path_elements[len(path_elements) - 1]

  // fmt.Println("delete ->", element_to_delete)

  if current_json.ExistsP(path) {

    // Get an error if there is no array
    children, err := current_json.Path(path).Children()

    // Return if there is no array
    if err != nil {
      return
    }

    current_json.ArrayP(path)
    for _, child := range children {
      arr_string := child.Data().(string)
      if arr_string != element_to_delete {
        current_json.ArrayAppendP(arr_string, path)
      }
    }
  }
}

// Delete a tag from a user
func (p Person) delete_tag(tag_name string) {
  current_json := p.Json
  tag_path := "person.tags"

  children, err := current_json.Path(tag_path).Children()
  check_err(err)
  current_json.ArrayP(tag_path)
  for _, child := range children {
    tag_string := child.Data().(string)
    if tag_string != tag_name {
      current_json.ArrayAppendP(tag_string, tag_path)
    }
  }
}

// Uncomment for main
// func main() {
//
//   // traits_path := "person.traits"
//   //
//   // p1 := new_person("ben")
//   // p1.add_tag("friend")
//   // p1.add_tag("geog 1000")
//   //
//   // p2 := new_person("steve")
//   // p2.add_trait(traits_path, "relative.brother", "uno")
//   // // p2.add_trait(traits_path, "relative.sister", "dos")
//   // p2.add_trait(traits_path, "relative.brother", "tres")
//   // p2.add_trait(traits_path, "relative.brother", "quad")
//   // p2.add_trait(traits_path, "location.current", "md")
//   //
//   // p1.add_tag("dog person")
//   // p1.add_tag("cat person")
//   // p2.add_tag("cat person")
//   // //p2.add_trait(traits_path, "relative", "tre")
//   //
//   // // fmt.Println(p2.Json.ExistsP("person.traits.relative"))
//   //
//   // // p2.t_delete_trait("bob")
//   //
//   // // fmt.Println(p1.Name, "->", p2.Name)
//   // // fmt.Println(p1.Json.String())
//   // // fmt.Println(p2.Json.String())
//   // //
//   // //
//   // //
//   // // p2.get_person_trails()
//   // p1.add_to_people_map()
//   // p2.add_to_people_map()
//
//   fmt.Println(get_all_people_with_tag("cat person", people_map))
//   // get_all_people_with_tag("dog person")
//   // get_all_people_with_tag("rhino person")
//
//   // export_people_to_file()
//   import_people_from_file()
//
//   for _, p := range people_map {
//     fmt.Print(p.get_name(), " ")
//     fmt.Println(p.get_person_tags())
//     p.delete_tag("cat person")
//     fmt.Print(p.get_name(), " ")
//     fmt.Println(p.get_person_tags())
//     fmt.Println("==========================")
//   }
//   fmt.Println(get_all_people_with_tag("cat person", people_map))
// }

// Support methods

// Obtain how many layers deep an array is
func obtain_array_count(text string) int {
  count := 0
  for _, elem := range text {
    if elem == '[' {
      count++
    } else {
      break
    }
  }
  return count
}

// Import people json objects from a file
func import_people_from_file() {
  storage_path := "./../storage/People_Storage/"
  file_type := ".json"

  files, err := ioutil.ReadDir(storage_path)
  check_err(err)
  for _, f := range files {
    if strings.HasSuffix(f.Name(), file_type) {
      person_data, err := ioutil.ReadFile(storage_path + f.Name())
      check_err(err)
      p := new_person_from_data(person_data)
      people_map[p.get_name()] = *p
    }
  }
}

// Export people json objects to a file
func export_people_to_file() {
  storage_path := "./../storage/People_Storage/"
  file_type := ".json"

  for _, p := range people_map {
    jsonOutput := p.Json.String()
    file, err := os.Create(storage_path + p.get_name() + file_type)
    check_err(err)
    defer file.Close()
    fmt.Fprintf(file, jsonOutput)
  }
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
