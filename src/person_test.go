package main

import (
	"testing"
	"fmt"
)

// Most of the simple tests are just going to need a person or two, this generates those people
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

// Just test that getting the name returns the correct info
func Test_basic_name(t *testing.T) {
  pm := get_example_person_map()

  p := pm["ben"]

  if p.get_name() != "ben" {
    t.Errorf("Wrong value of result: %v", p.get_name())
  }
}

/*
	Testing deleting a tag from a person.
	In this test both users have the tag.
*/
func Test_remove_tag_present(t *testing.T) {
  pm := get_example_person_map()

  tag_to_delete := "cat person"

  p_ben := pm["ben"]
  p_steve := pm["steve"]

  p_ben.delete_tag(tag_to_delete)
  tag_map_ben := p_ben.get_person_tags()
  if tag_map_ben[tag_to_delete] == true {
    t.Errorf("Tag not deleted: %v", tag_to_delete)
  }

  p_steve.delete_tag(tag_to_delete)
  tag_map_steve := p_steve.get_person_tags()
  if tag_map_steve[tag_to_delete] == true {
    t.Errorf("Tag not deleted: %v", tag_to_delete)
  }
}

/*
	Testing deleting a tag from a person.
	In this test neither user has the tag.
	Given the way we are removing tags there should not be any difference
	between removing items that the person has and those that they do not.
*/
func Test_remove_tag_not_present(t *testing.T) {
	pm := get_example_person_map()

  tag_to_delete := "cat person1"

  p_ben := pm["ben"]
  p_steve := pm["steve"]

  p_ben.delete_tag(tag_to_delete)
  tag_map_ben := p_ben.get_person_tags()
  if tag_map_ben[tag_to_delete] == true {
    t.Errorf("Tag not deleted: %v", tag_to_delete)
  }

  p_steve.delete_tag(tag_to_delete)
  tag_map_steve := p_steve.get_person_tags()
  if tag_map_steve[tag_to_delete] == true {
    t.Errorf("Tag not deleted: %v", tag_to_delete)
  }
}

// Testing the adding of traits to a person
func Test_add_some_traits_simple(t *testing.T) {
	pm := get_example_person_map()

	traits_path := "person.traits"

	p_ben := pm["ben"]

	p_ben.add_trait(traits_path, "relative.brother", "uno")
	p_ben.add_trait(traits_path, "relative.brother", "tres")
	p_ben.add_trait(traits_path, "relative.brother", "quad")

	expected_result := `["uno","tres","quad"]`
	result := p_ben.Json.Path("person.traits.relative.brother").String()

	if result != expected_result {
		t.Errorf("Expecting: %v But got: %v", expected_result, result)
	}

	p_ben.add_trait(traits_path, "location.current", "md")
	expected_result = `"md"`
	result = p_ben.Json.Path("person.traits.location.current").String()
	if result != expected_result {
		t.Errorf("Expecting: %v But got: %v", expected_result, result)
	}
}

// Seeing if we can delete a single trait of a person
func Test_delete_single_trait(t *testing.T) {

	pm := get_example_person_map()
	p_steve := pm["steve"]

	current_json := p_steve.Json

	trait_to_delete := "location.current"

	p_steve.delete_trait(trait_to_delete)

	if current_json.ExistsP("person.traits" + trait_to_delete) {
		t.Errorf("We are still seeing the element we want to delete")
	}
}

// Test of my function for getting a persons traits
func Test_get_person_traits(t *testing.T) {

	pm := get_example_person_map()
	simple_person := pm["steve"]

	trait_map := simple_person.get_person_traits()

	_, ok := trait_map["location"]
	if ok != true {
		t.Errorf("Getting traits is not working")
	}

}

// Use this to determine how we are going to end up being able to delete on an array
// They could be whole other objects
func Test_delete_array_trait(t *testing.T) {

	pm := get_example_person_map()
	p_steve := pm["steve"]
	current_json := p_steve.Json

	trait_to_delete := "relative.brother"

	p_steve.delete_trait(trait_to_delete)

	if current_json.ExistsP("person.traits" + trait_to_delete) {
		t.Errorf("We are still seeing the element we want to delete")
	}
}


// This refers to just getting info from elements that within an array
// This test isnt one of my program, but rather the functionality of a library
func Test_deep_search(t *testing.T) {

	pm := get_example_person_map()
	p_dave := pm["dave"]

	r1 := p_dave.Json.ExistsP("person.traits.relative.brother")
	r2 := p_dave.Json.ExistsP("person.traits.relative.brother.uno")
	r3 := p_dave.Json.ExistsP("person.traits.relative.brother.quad")

	if !r1 || !r2 || !r3 {
		t.Errorf("The library must have changed")
	}

}

func done() {
	fmt.Println("done")
}
