package main

import (
	"testing"
)

func setup_simple_test() map[string]Person {

  sample_ben := []byte(`{"person":{"name":"ben","tags":["friend","geog 1000","dog person","cat person"]}}`)
  sample_steve := []byte(`{"person":{"name":"steve","tags":["cat person"],"traits":{"location":{"current":"md"},"relative":{"brother":["uno","tres","quad"]}}}}`)

  p_ben := new_person_from_data(sample_ben)
  p_steve := new_person_from_data(sample_steve)

  test_people_map := make(map[string]Person)
  test_people_map[p_ben.get_name()] = *p_ben
  test_people_map[p_steve.get_name()] = *p_steve
  return test_people_map
}

func Test_basic_name(t *testing.T) {
  pm := setup_simple_test()

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
  pm := setup_simple_test()

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
	pm := setup_simple_test()

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

func Test_add_some_traits_simple(t *testing.T) {
	pm := setup_simple_test()

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
