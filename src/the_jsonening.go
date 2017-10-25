package main

import (
  "fmt"
  "strings"
)

/*
  lol, when you dont read the instructions
*/
func (p Person) change_json(trait_name string) {
  current_json := p.Json

  path := create_path("person.traits", trait_name)

  if current_json.ExistsP(path) {
    fmt.Println("we here")
  } else {
    fmt.Println("Nothing to delete")
    return
  }

  err := current_json.DeleteP(path)
  if err == nil {
    fmt.Println("It has been done.")
  }

  path_elements := strings.Split(path, ".")

  path = strings.Join(path_elements[:len(path_elements) - 1], ".")

  path = create_path("person.traits", trait_name)

  my_data := current_json.Path(path).Index(0)
  children, err := my_data.Children()

  current_json.ArrayP(path)

  check_err(err)
  for key, child := range children {
  	fmt.Println(key, "->", child.Index(0))
    current_json.ArrayAppendP(child, path)
  }


  /*
  elements, ok := my_data.([]interface{})
  if !ok {
    fmt.Println("yeah, something is definitely broken")
    return
  }

  sub_elements := elements[0].(interface{})
  sub_sub_elements := sub_elements.([]interface{})
  sub_sub_sub_elements := sub_sub_elements[0].(map[string]interface{})
  cat_elements := sub_sub_sub_elements["kittens"]
  kitten_elements := cat_elements.([]interface{})
  sub_kitten_elements := kitten_elements[0].(map[string]interface{})
  name := sub_kitten_elements["male"]
  fmt.Println(name)
  */

}
