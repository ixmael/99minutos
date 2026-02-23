package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func main() {
	healthURL := "http://localhost/health"

	// order := map[string]interface{}{
	// 	"product_id": "09871d35-1a66-46df-b3b2-1f0aa23763ca",
	// 	"quantity":   6,
	// }
	// orderJSON, err := json.Marshal(order)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	var wg sync.WaitGroup
	concurrent := 500

	wg.Add(concurrent)

	for i := range concurrent {
		go func(i int) {
			defer wg.Done()

			response, err := http.Get(healthURL)
			if err != nil {
				fmt.Printf("Failed to load health %d: %s\n", i, err)
				return
			}
			defer response.Body.Close()

			body, err := io.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			var okResult struct {
				Status      string `json:"status"`
				Version     string `json:"version"`
				Environment string `json:"environment"`
				Host        string `json:"host"`
			}

			err = json.Unmarshal(body, &okResult)
			if err != nil {
				fmt.Println("error on read body", err)
			}

			fmt.Print("\n\n\n\n")
			fmt.Println("status: ", response.StatusCode)
			fmt.Printf("response: %+v\n", okResult)
			fmt.Print("\n\n\n\n")

			/*
				response, err := http.Post(healthURL, "application/json", bytes.NewBuffer(orderJSON))
				if err != nil {
					fmt.Printf("Failed to create order %d: %s\n", i, err)
					return
				}
				defer response.Body.Close()

				fmt.Println("status: ", response.StatusCode)
			*/
		}(i)
	}

	wg.Wait()
}
