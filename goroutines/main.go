package main

import (
	"fmt"
	"sync"
	"time"
)

var database = []string{"id 1", "id 2", "id 3", "id 4", "id 5"}
var wg = sync.WaitGroup{}
var results []string
var mutex = sync.RWMutex{}

func databaseCall(id int32) {
	var delay float32 = 2000 // 2 seconds (2000 milliseconds)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The response from the database is ", database[id])
	mutex.Lock()
	results = append(results, database[id])
	mutex.Unlock()
	save(database[id])
	log()
	wg.Done()
}

func save(result string) {
	mutex.Lock()
	results = append(results, result)
	mutex.Unlock()
}

func log() {
	mutex.RLock()
	fmt.Println("The current results is", results)
	mutex.RUnlock()
}

func main() {
	t0 := time.Now()
	for i := range database {
		wg.Add(1)
		go databaseCall(int32(i))
	}
	wg.Wait()
	fmt.Println("Execution time is ", time.Since(t0))
	fmt.Println("Results: ", results)

}
