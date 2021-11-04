package main

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"

	xerrors "github.com/pkg/errors"
)

func PrintDog(wg *sync.WaitGroup, fish, dog chan bool) {
	defer wg.Done()
	defer close(dog)

	for i:=0; i<4; i++ {
		<-fish
		fmt.Println("dog...")
		dog<-true
	}
	
}

func PrintCat(wg *sync.WaitGroup, dog, cat chan bool) {
	defer wg.Done()
	defer close(cat)

	for i:=0; i<4; i++ {
		<-dog
		fmt.Println("cat...")
		cat<-true
	}
}

func PrintFish(wg *sync.WaitGroup, cat, fish chan bool) {
	defer wg.Done()
	defer close(fish)

	for i:=0; i<4; i++ {
		<-cat
		fmt.Println("fish...")
		fish<-true
	}
}

func QueryRows() ([]int, error) {
	return nil, xerrors.Wrap(sql.ErrNoRows, "no rows")
}

func QueryCount() int {
	ret, err := QueryRows()
	if errors.Is(err, sql.ErrNoRows) {
		return 0
	} 
	return len(ret)
}

func main1() {
	wg := &sync.WaitGroup{}

    dogChan := make(chan bool, 1)
    catChan := make(chan bool, 1)
    fishChan := make(chan bool, 1)
	fishChan <- true

	go PrintDog(wg, fishChan, dogChan)
	go PrintCat(wg, dogChan, catChan)
	go PrintFish(wg, catChan, fishChan)

	wg.Add(3)
	wg.Wait()
	fmt.Println("done")
}
