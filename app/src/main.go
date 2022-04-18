package main

import (
	aservice "github.com/ekinbulut-yemeksepeti/samples/app/arraylist"
	hservice "github.com/ekinbulut-yemeksepeti/samples/app/hashtable"
)

type App struct {
	Name      string
	Version   string
	ArrayList *aservice.ArrayList
	HashTable *hservice.HashTable
}

func NewApp() *App {
	return &App{
		Name:      "App",
		Version:   "1.0.0",
		ArrayList: aservice.NewArrayList(),
		HashTable: hservice.NewHashTable(10),
	}
}

func main() {

	a := NewApp()
	a.ArrayList.Add(1)
	a.ArrayList.Print()
}
