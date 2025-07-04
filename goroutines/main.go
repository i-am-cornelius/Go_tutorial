package main

import (
	"fmt"
	"sync"
	"time"
)

var database = []string{"id 1", "id 2", "id 3", "id 4", "id 5"}
var wg = sync.WaitGroup{}
var results = []string{}
var mutex = sync.Mutex{}

func database_call(id int32) {
	var delay float32 = 2000 // 2 seconds (2000 milliseconds)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The response from the database is ", database[id])
	mutex.Lock()
	results = append(results, database[id])
	mutex.Unlock()
	wg.Done()
}

func main() {
	t0 := time.Now()
	for i := range database {
		wg.Add(1)
		go database_call(int32(i))
	}
	wg.Wait()
	fmt.Println("Execution time is ", time.Since(t0))
	fmt.Println("Results: ", results)

	var nums = [...]int16{78, 86, 71, 87, 94, 73, 82}
	for i := range nums {
		if nums[i]%2 != 0 {
			fmt.Println(nums[i], "is an odd number!")
		} else {
			fmt.Println(nums[i], "is an even number!")
		}
	}
}
