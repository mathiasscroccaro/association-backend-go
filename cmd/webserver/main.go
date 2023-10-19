package main

import (
	"captcha_example/internal/api"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	server := api.CreateServer()

	go func() {
		defer wg.Done()
		fmt.Printf("Server is running on %s...\n", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Error: %s\n", err)
		}
	}()

	wg.Wait()
}
