package main

import (
	"encoding/json"
	"fmt"
	"sync"
)
type count struct {
	m map[string] int
	mu sync.Mutex
}
func (countMap *count)countLetters(s string,wg *sync.WaitGroup) {
	defer wg.Done()
	for _,i:=range s{

		countMap.mu.Lock()
		countMap.m[string(i)] = countMap.m[string(i)]+1
		countMap.mu.Unlock()
	}
}
func main() {
	data := [5]string{"quick", "brown", "fox", "lazy", "dog"}
	var countMap = count{make(map[string]int),sync.Mutex{}}

	var wg sync.WaitGroup
	//fmt.Println("Check")
	for  _,v := range data {
		wg.Add(1)
		countMap.countLetters(v,&wg)
	}

	wg.Wait()
	empData, err := json.Marshal(countMap.m)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonStr := string(empData)
	fmt.Println("The JSON data is:")
	fmt.Println(jsonStr)
	//fmt.Println(countmap.m)
}