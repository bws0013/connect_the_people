package main

import (
  "fmt"
  "strings"
  "github.com/Jeffail/gabs"
)

func (p Person) true_delete(trait_name string) {
  current_json := p.Json

  path := create_path("person.traits", trait_name)
  if current_json.ExistsP(path) {
    // fmt.Println("we here")
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
  element_to_delete := path_elements[len(path_elements) - 1]
  // fmt.Println("To Delete:", element_to_delete)

  my_data := current_json.Path(path)

  element_depth := obtain_array_count(my_data.String())

  children, err := my_data.Children()
  check_err(err)
  for i := 0; i < element_depth - 1; i++ {
    for _, child := range children {
      children, err = child.Children()
      check_err(err)
    }
  }

  jsonObj := gabs.New()
  jsonObj.Array("foo", "array")

  for _, child := range children {
    potential_target := child.String()

    target_split := strings.Split(potential_target, ":")
    if clean_text(target_split[0]) == element_to_delete {

    } else {
      jsonObj.ArrayAppend(child.Data(), "foo", "array")
    }
  }

  new_arr := jsonObj.Path("foo.array")
  // fmt.Println(new_arr)

  current_val := current_json.Path(path)
  check_err(err)
  for i := 0; i < element_depth - 1; i++ {
    current_val = current_val.Index(0)
  }

  fmt.Println("->", current_val)
  kinder, err := current_val.Children()

  for _, kind := range kinder {
    check_err(err)
    fmt.Println(kind)
  }

  _, err = current_val.SetIndex(new_arr, 0)
  check_err(err)

  // fmt.Println(my_data)
}

// This may have some bad consequences down the line, warrents further investigation
func (p Person) clean_json() Person {
  current_json_text := p.Json.String()
  if current_json_text == "{}" { return p }
  replacer := strings.NewReplacer(
    "{},", "",
    "{}", "",
    "[],", "",
    "[]", "")
  temp := replacer.Replace(current_json_text)
  new_json, err := gabs.ParseJSON([]byte(temp))
  check_err(err)
  p.Json = new_json
  return p
}

func clean_text(text string) string {
  temp := strings.Replace(text, "{", "", -1)
  if temp[0] == '"' {
    temp = temp[1:]
  }
  if temp[len(temp) - 1] == '"' {
    temp = temp[:len(temp) - 1]
  }
  return temp
}

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
