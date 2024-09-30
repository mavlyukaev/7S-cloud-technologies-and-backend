package main

import (
 "fmt"
 "math"
 "sync"
)

type Task func()

type TaskPool struct {
 tasks []Task
 mu    sync.Mutex
}

type ITaskPool interface {
 NextTask() Task
 Push(Task)
}

type Executor interface {
 ExecNext()
}

func (tp *TaskPool) Push(t Task) {
 tp.mu.Lock()
 tp.tasks = append(tp.tasks, t)
 tp.mu.Unlock()
}

func (tp *TaskPool) NextTask() Task {
 tp.mu.Lock()
 defer tp.mu.Unlock()

 if len(tp.tasks) == 0 {
  return nil
 }

 task := tp.tasks[0]
 tp.tasks = tp.tasks[1:]
 return task
}

func (tp *TaskPool) ExecNext() {
 task := tp.NextTask()
 if task != nil {
  task()
 }
}

var primes []int
var primeCount int
var mu sync.Mutex

func isPrime(n int) bool {
 if n < 2 {
  return false
 }
 for _, p := range primes {
  if p > int(math.Sqrt(float64(n))) {
   break
  }
  if n%p == 0 {
   return false
  }
 }
 return true
}

func createPrimeTask(n int) Task {
 return func() {
  if isPrime(n) {
   mu.Lock()
   primes = append(primes, n)
   primeCount++
   mu.Unlock()
  }
 }
}

func GenerateNums(n int) []int {
 nums := make([]int, n)
 for i := 0; i < n; i++ {
  nums[i] = i + 1
 }
 return nums
}

func main() {
 nums := GenerateNums(10)
 var taskPool TaskPool

 for _, n := range nums {
  taskPool.Push(createPrimeTask(n))
 }

 for len(taskPool.tasks) > 0 {
  taskPool.ExecNext()
 }

 fmt.Printf("Общее количество простых чисел: %d\n", primeCount)
 fmt.Println("Простые числа:", primes)
}
