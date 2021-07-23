package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rate int

func findRating(wg *sync.WaitGroup) {

	defer wg.Done()
	//fmt.Println(ind_rate)
	indRate := rand.Intn(10)
	responseTime := rand.Intn(5)
	time.Sleep(time.Duration(responseTime))
	rate += indRate

}
func main() {
	var wg sync.WaitGroup
	nos := 200
	rate = 0
	for i := 1; i <= nos; i++ {
		wg.Add(1)

		go findRating( &wg)
	}
	wg.Wait()
	avgrate := rate / nos
	fmt.Println("The average rating of the teacher is ", avgrate)
}
