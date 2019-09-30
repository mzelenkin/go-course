package main

import (
	"errors"
	"math/rand"
	"time"
)

const (
	MaxErrors               = 2    // Максимальное кол-во ошибок. При превышении этого значения выполнение будет приостановлено
	MaxThreads              = 2    // Максимальное количество параллельно выполняющихся заданий
	MinTimeExecution        = 500  // Минимальное время выполнение фиктивного задания
	MaxTimeExecution        = 1000 // Максимальное время выполнение фиктивного задания
	SuccessFunctionsPercent = 80   // Кол-во фиктивных функций с успешным результатом выполнения (остальные с ошибками)
)

func main() {
	tasks := genTasks(15)
	RunTasks(tasks, MaxThreads, MaxErrors)
}

// genTasks генерирует список задач длинной n
func genTasks(n int) []Task {
	var tasks []Task

	for i := 0; i < n; i++ {
		task := func() error {
			// Случайное время сна
			timeExec := MinTimeExecution*time.Millisecond +
				time.Duration(rand.Intn(MaxTimeExecution-MinTimeExecution))*time.Microsecond

			time.Sleep(timeExec)

			// Указаное кол-во функций возвращает ошибку
			if rand.Intn(100) > SuccessFunctionsPercent {
				println("Error")
				return errors.New("error in function")
			}

			// Остальные success
			println("Success")
			return nil
		}

		tasks = append(tasks, task)
	}

	return tasks
}
