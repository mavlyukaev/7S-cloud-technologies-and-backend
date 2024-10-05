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
   // Если контекст отменен, возвращаем ошибку
   return nil, ctx.Err()
  default:
   // Симулируем задержку, чтобы привести к тайм-ауту
   time.Sleep(100 * time.Millisecond)
   // Генерируем случайные данные
   data = append(data, byte(rand.Intn(256)))
  }
 }
 return data, nil
}

func getAvailableMemory() int {
 var m runtime.MemStats
 runtime.ReadMemStats(&m)
 // Возвращаем доступную память (в байтах)
 return int(m.Sys)
}

func main() {
 // Проверяем доступную память в байтах
 availableMemory := getAvailableMemory()
 // Переводим доступную память в МБ
 availableMemoryMB := availableMemory / (1024 * 1024)

 // Выводим доступную память
 fmt.Printf("Доступная память в этой среде: %d MB\n", availableMemoryMB)

 // Ограничим размер генерируемых данных доступной памятью
 // Устанавливаем максимальный размер данных, который можно генерировать
 maxSize := availableMemoryMB * 1024 * 1024

 // Устанавливаем таймаут в 10 секунд
 timeout := 10 * time.Second
 ctx, cancel := context.WithTimeout(context.Background(), timeout)
 defer cancel()

 // Ожидаемый размер данных - 100 МБ (но не больше доступной памяти)
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