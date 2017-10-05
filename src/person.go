package main

import (
  "fmt"
  "encoding/json"
)

// Build basic person object json object
// Give the capability to custimize the object

func main() {
  template_extreme()
  template_rev()
}

func template_extreme() {
  data := `
    {
  	"id": "0001",
  	"type": "donut",
  	"name": "Cake",
  	"ppu": 0.55,
  	"batters":
  		{
  			"batter":
  				[
  					{ "id": "1001", "type": "Regular" },
  					{ "id": "1002", "type": "Chocolate" },
  					{ "id": "1003", "type": "Blueberry" },
  					{ "id": "1004", "type": "Devil's Food" }
  				]
  		},
  	"topping":
  		[
  			{ "id": "5001", "type": "None" },
  			{ "id": "5002", "type": "Glazed" },
  			{ "id": "5005", "type": "Sugar" },
  			{ "id": "5007", "type": "Powdered Sugar" },
  			{ "id": "5006", "type": "Chocolate with Sprinkles" },
  			{ "id": "5003", "type": "Chocolate" },
  			{ "id": "5004", "type": "Maple" }
  		]
  }
  `
  byt := []byte(data)

  var dat map[string]interface{}
  if err := json.Unmarshal(byt, &dat); err != nil {
    panic(err)
  }
  for k, v := range dat {
    fmt.Println("Key: ", k)
    fmt.Println("Val: ", v)
    fmt.Println("====================")
  }

  dat2 := dat["topping"].([]interface{})
  // var dat2 map[string]interface{}
  // byt = []byte(dat["topping"])
  // if err = json.Unmarshal(byt, &dat2); err != nil {
  //   panic(err)
  // }
  //
  for k, v := range dat2 {
    fmt.Println("Key: ", k)
    fmt.Println("Val: ", v)
    fmt.Println("====================")
  }

}

func template_rev() {
  mapD := map[string]int{"apple": 5, "lettuce": 7}
  mapB, _ := json.Marshal(mapD)
  fmt.Println(string(mapB))
}


// Use this as a template
func template() {

  byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
  var dat map[string]interface{}

  if err := json.Unmarshal(byt, &dat); err != nil {
    panic(err)
  }
  fmt.Println(dat)

  num := dat["num"].(float64)
  fmt.Println(num)

  strs := dat["strs"].([]interface{})
  str1 := strs[0].(string)
  fmt.Println(str1)

}
