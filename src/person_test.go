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

func Test_remove_item_present(t *testing.T) {
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

func Test_remove_item_not_present(t *testing.T) {

}
