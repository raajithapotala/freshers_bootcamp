package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rate int

func findrating(ind_rate int, wg *sync.WaitGroup) {

	defer wg.Done()
	//fmt.Println(ind_rate)
	response_time := rand.Intn(5)
	time.Sleep(time.Duration(response_time))
	rate += ind_rate

}
func main() {
	var wg sync.WaitGroup
	nos := 200
	rate = 0
	for i := 1; i <= nos; i++ {
		wg.Add(1)
		ind_rate := rand.Intn(10)
		go findrating(ind_rate, &wg)
	}
	wg.Wait()
	avgrate := rate / nos
	fmt.Println("The average rating of the teacher is ", avgrate)
}
