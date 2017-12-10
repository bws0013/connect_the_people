package main

import (
    "fmt"
    "log"
    // "strings"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/Jeffail/gabs"
)

var people []Person

// Copying from here: https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo

// our main function
func main() {


  pm := get_example_person_map()
  for _, p := range pm {
    people = append(people, p)
  }

	router := mux.NewRouter()

  router.HandleFunc("/people", GetPeople).Methods("GET")
  router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
  router.HandleFunc("/people/{ild}", DeletePerson).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8000", router))
}

func get_example_person_map() map[string]Person {
  sample_ben := []byte(`{"person":{"name":"ben","tags":["friend","geog 1000","dog person","cat person"]}}`)
  sample_steve := []byte(`{"person":{"name":"steve","tags":["cat person"],"traits":{"location":{"current":"md"},"relative":{"brother":["uno","tres","quad"]}}}}`)
	sample_dave := []byte(`{"person":{"name":"dave","tags":["cat person"],"traits":{"location":{"current":"md"},"relative":{"brother":[{"uno":{"age":"18"}},{"tres":"nil"},{"quad":"nil"}]}}}}`)
	sample_stan := []byte(`{"person":{"name":"stan","tags":["cat person"],"traits":{"location":{"current":"md"},"relative":{"brother":[{"uno":{"age":"18","pet":[{"cat":"cat-do"},{"dog":"dog-do"},{"rat": "rat-do"}]}},{"tres":"dc"},{"quad":"yo"}]}}}}`)
	sample_edwin := []byte(`{"person":{"name":"edwin","traits":{"relative":{"brother":[{"uno":{"pet":[{"cat":"cat-do","kittens":[{"male":"one"},{"female":"two"},"third"]},{"dog":"dog-do"},{"rat":"rat-do"}]}}]}}}}`)

  p_ben := new_person_from_data(sample_ben)
  p_steve := new_person_from_data(sample_steve)
	p_dave := new_person_from_data(sample_dave)
	p_stan := new_person_from_data(sample_stan)
	p_edwin := new_person_from_data(sample_edwin)

  test_people_map := make(map[string]Person)
  test_people_map[p_ben.get_name()] = *p_ben
  test_people_map[p_steve.get_name()] = *p_steve
	test_people_map[p_dave.get_name()] = *p_dave
	test_people_map[p_stan.get_name()] = *p_stan
	test_people_map[p_edwin.get_name()] = *p_edwin

	return test_people_map
}


// ***** RESTful stuff below *****


func GetPeople(w http.ResponseWriter, r *http.Request) {
    jsonObj := gabs.New()
    jsonObj.Array("Names")
    for _, v := range people {
      jsonObj.ArrayAppend(v.get_name(), "Names")
    }

    // jsonParsedObj, err := gabs.ParseJSON([]byte(people_string[0]))
    // check_err(err)
    //
    // for i := 1; i < len(people_string); i++ {
    //   current_parsed_obj, err := gabs.ParseJSON([]byte(people_string[i]))
    //   check_err(err)
    //   err = jsonParsedObj.Merge(current_parsed_obj)
    //   check_err(err)
    // }
    //
    json_string := jsonObj.String()
    // fmt.Println(json_string)


    // fmt.Println(string(bytes))
    fmt.Println("I have been chosen")
    // json.NewEncoder(w).Encode(ff)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(json_string))
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for _, item := range people {
    if item.get_name() == params["id"] {
      // jsonParsedObj, err := gabs.ParseJSON())
      // check_err(err)
      w.Header().Set("Content-Type", "application/json")
      w.Write([]byte(item.Json.String()))
      return
    }
  }
  json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  person := new_person(params["id"])
  people = append(people, *person)
  json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for index, item := range people {
  	if item.get_name() == params["id"] {
    	people = append(people[:index], people[index+1:]...)
      break
  	}
  	json.NewEncoder(w).Encode(people)
	}

}
