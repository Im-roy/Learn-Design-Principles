package main

import (
	"fmt"
	"sync"
)

// we are keeping the struct name in lower case making sure
// its object can't be created outside package.
type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPoplulation(city string) int {
	return db.capitals[city]
}

// this sync.Once make sure that function inside it runs just once
var once sync.Once
var dbClient *singletonDatabase

func ReadFile(filename string) (map[string]int, error) {
	return map[string]int{
		"patna":    20000,
		"Delhi":    40000,
		"Banglore": 9000,
	}, nil
}

func GetSingletonDatabaseObj() *singletonDatabase {
	once.Do(func() {
		caps, _ := ReadFile("../filename")
		db := singletonDatabase{capitals: caps}
		dbClient = &db
	})
	return dbClient
}

func GetTotalPopulation(cities []string) int {
	total := 0
	for _, city := range cities {
		total += GetSingletonDatabaseObj().GetPoplulation(city)
	}
	return total
}

func main() {

	fmt.Println("************Sigleton design pattern************")

	db := GetSingletonDatabaseObj()
	fmt.Println(db.GetPoplulation("Banglore"))

	ok := GetTotalPopulation()
}
