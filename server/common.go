package main

import (
	"errors"
)

//  STRUCTS  /////////////////////////////////////////////////////////////////

type Item struct {
	Id int
	Path string
}

type Status struct {
	Status string
	List []Item
	Id int
}

//  VARS  ////////////////////////////////////////////////////////////////////

var items []Item
var next_id int

//  HELPER FUNCTIONS  ////////////////////////////////////////////////////////

func addItem(path string) (id int) {
	id = next_id
	items = append(items, Item{Id: next_id, Path: path})
	next_id = next_id + 1
	return
}

func itemIndex(id int) (i int, err error) {
	for i = 0; i < len(items); i++ {
		if items[i].Id == id {
			return
		}
	}
	err = errors.New("no item with that id found")
	return
}

func removeItem(id int) (err error) {
	i, err := itemIndex(id)
	if err != nil {
		return
	}
	items = append(items[:i], items[i+1:]...)
	return
}
