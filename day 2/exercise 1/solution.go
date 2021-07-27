package main

import (
	"fmt"
	"sync"
)

func main() {
	data := &[...]string{"quick", "brown", "fox", "lazy", "dog"}
	mymap := make(map[string]int)
	var wg sync.WaitGroup
	// wg := sync.WaitGroup
	for _, v := range data {
		wg.Add(1)
		go func(s string) {
			fmt.Println("Entering ", s)
			defer wg.Done()
			for _, w := range s {
				mymap[string(w)] = mymap[string(w)] + 1
			}
			fmt.Println("Leaving ", s)
		}(v)
	}
	wg.Wait()
	for s := "a"; s <= "z"; s = string(s[0] + 1) {
		fmt.Println(s+" ", mymap[s])
	}
}
