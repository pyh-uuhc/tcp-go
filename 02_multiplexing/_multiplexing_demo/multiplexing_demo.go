// multiplexing demo using goroutine and channel

package main

import (
	"fmt"
	"time"
)

// worker 함수는 작업을 처리하는 고루틴
// jobs 채널에서 작업을 받아 처리한 후 결과를 results 채널에 보낸다.
// id는 각 고루틴의 식별자
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 세 개의 worker 고루틴을 시작하여 동시에 작업을 처리
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 5개의 작업을 jobs 채널에 보내고, 모든 작업이 완료될 때까지 기다렸다가 결과를 출력
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Printf("Result: %d\n", <-results)
	}
}
