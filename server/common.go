package main

import "errors"

//  STRUCTS  /////////////////////////////////////////////////////////////////

type Item struct {
	Id   int
	Path string
}

type Status struct {
	Status string
	List   []Item
	Id     int
}

const (
	StateStopped = iota
	StatePlaying
)

//  VARS  ////////////////////////////////////////////////////////////////////

var items []Item
var nextId int
var currentTrack int
var currentState int

//  HELPER FUNCTIONS  ////////////////////////////////////////////////////////

func addItem(path string) (id int) {
	id = nextId
	items = append(items, Item{Id: nextId, Path: path})
	nextId = nextId + 1
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

func initialize() {
	items = []Item{}
	nextId = 1
	currentTrack = 0
	currentState = StateStopped
}
