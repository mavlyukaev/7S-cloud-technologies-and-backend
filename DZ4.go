package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"
)

func generateData(ctx context.Context, size int) ([]byte, error) {
	data := make([]byte, 0, size)
	for len(data) < size {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			time.Sleep(100 * time.Millisecond)
			data = append(data, byte(rand.Intn(256)))
		}
	}
	return data, nil
}

func getAvailableMemory() int {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return int(m.Sys)
}

func main() {
	availableMemory := getAvailableMemory()
	availableMemoryMB := availableMemory / (1024 * 1024)
	fmt.Printf("Доступная память в этой среде: %d MB\n", availableMemoryMB)
	maxSize := availableMemoryMB * 1024 * 1024

	timeout := 1 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	size := 100 * 1024 * 1024
	if size > maxSize {
		size = maxSize
		fmt.Printf("Размер данных ограничен доступной памятью: %d MB\n", size/(1024*1024))
	}

	data, err := generateData(ctx, size)
	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println("Операция отменена: время вышло.")
		} else {
			log.Fatalf("Ошибка при генерации данных: %v", err)
		}
	} else {
		fmt.Printf("Данные успешно сгенерированы. Размер данных: %d MB\n", len(data)/(1024*1024))
	}
}
