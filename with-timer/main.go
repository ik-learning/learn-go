package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

func DoWork(ctx context.Context) error {
	time.Sleep(5 * time.Second)
	fmt.Println("Work completed")
	return nil
}

func DoWorkConcurrent() chan struct{} {
	ch := make(chan struct{})
	go func() {
		DoWork(context.Background())
		close(ch)
	}()
	return ch
}

func doSomething(ctx context.Context) error {
	ctx, cancel := context.WithTimeoutCause(context.Background(), 2*time.Second, errors.New("longer then 2 seconds"))
	defer cancel() // releases resources if slowOperation completes before timeout elapses
	return DoWork(ctx)

	// select {
	// case <-ctx.Done():
	// 	return ctx.Err()
	// default:
	// 	// Simulate a long-running task
	// 	time.Sleep(5 * time.Second)
	// 	fmt.Println("Task completed")
	// }

	// return nil
}

func doSomethingSelect(ctx context.Context, timeout time.Duration) error {
	select {
	case <-DoWorkConcurrent():

		return nil
	case <-time.After(timeout):
		return fmt.Errorf("timed out waiting for")
	}
}

// https://golang.cafe/blog/golang-context-with-timeout-example.html
// func slowOperationWithTimeout(ctx context.Context) (Result, error) {
// 	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
// 	defer cancel()  // releases resources if slowOperation completes before timeout elapses
// 	return slowOperation(ctx)
// }

func main1() {
	fmt.Println("haha")
	// doSomething(context.TODO())
	err := doSomethingSelect(context.TODO(), time.Duration(2*time.Second))

	// ctx, cancel := context.WithTimeoutCause(context.Background(), 2*time.Second, errors.New("longer then 2 seconds"))
	// defer cancel()

	// err := doSomething(ctx)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("task completed successfuly")
	}
}

func doSomethingV2(ch chan string) {
	fmt.Println("doSomething Sleeping...")
	time.Sleep(time.Second * 5)
	fmt.Println("doSomething Wake up...")
	ch <- "Did Something"
}

// https://medium.com/geekculture/timeout-context-in-go-e88af0abd08d
// https://dev.to/hgsgtk/timeout-using-context-package-in-go-1b3c

func main() {
	ch := make(chan string, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go doSomethingV2(ch)

	select {
	case <-ctx.Done():
		fmt.Printf("Context cancelled: %v\n", ctx.Err())
	case result := <-ch:
		fmt.Printf("Received: %s\n", result)
	}

	ctx = context.Background()
	ctx, cancel = context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := execute(ctx); err != nil {
		log.Fatalf("error: %#v\n", err)
	}
	log.Println("Success to process in time")
}

func execute(ctx context.Context) error {
	proc1 := make(chan struct{}, 1)
	proc2 := make(chan struct{}, 1)

	go func() {
		// Would be done before timeout
		time.Sleep(1 * time.Second)
		proc1 <- struct{}{}
	}()

	go func() {
		// Would not be executed because timeout comes first
		time.Sleep(3 * time.Second)
		proc2 <- struct{}{}
	}()

	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-proc1:
			fmt.Println("process 1 done")
		case <-proc2:
			fmt.Println("process 2 done")

		}
	}

	return nil
}
