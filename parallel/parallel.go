package main

import (
	"sync"
	"sync/atomic"
)

type Task func() error

// RunTasks выполняет задания tasks по n заданий параллельно.
// При достижении кол-ва ошибок maxErrors выполнение функций приостанавливается.
// Идея реализации - worker pool (thread pool)
// Мы создаем очередь задач, запускаем n worker'ов. Свободный worker достает
// задание из очереди и выполняет. Если произошла ошибка, увеличивает счетчик ошибок.
// Если кол-во ошибок превышает maxErrors, воркер завершает работу, предварительно очистив очередь.
func RunTasks(tasks []Task, n int, maxErrors int) {
	// Канал jobs используется как очередь задач
	// Буфер канала выставляем равным n, хотя это и не принципиально,
	// т.к. worker'ы все равно будут забирать задания по мере их выполнения
	jobs := make(chan Task, n)

	// Счетчик ошибок
	// Используем sync/atomic т.к. код получается более читаемый чем с mutex
	var errorsCount int32

	wg := sync.WaitGroup{}

	// Запускаем n worker'ов
	for i := 0; i < n; i++ {
		go func() {
			for job := range jobs { // Получение задания из очереди задач

				// Проверка на общее количество ошибок задач
				// По условию д/з, если это число > аргумента maxErrors, то задание не выполняется
				// В этом случае, т.к. мы его уже забрали, но не выполняем и забираем следующее.
				// Фактически в этом случае мы очищаем очередь заданий
				currentErrorsCounter := int(atomic.LoadInt32(&errorsCount))
				if currentErrorsCounter >= maxErrors {
					continue
				}

				// Выполняем задачу (функцию)
				if job() != nil {
					atomic.AddInt32(&errorsCount, 1)
				}
			}

			wg.Done()
		}()
		wg.Add(1)
	}

	// Добавление задач в очередь задач
	for _, task := range tasks {
		jobs <- task

		// Если дистигнут порог ошибок, не добавляем больше заданий
		// Чисто оптимизация
		currentErrorsCounter := int(atomic.LoadInt32(&errorsCount))
		if currentErrorsCounter >= maxErrors {
			break
		}
	}
	close(jobs) // Закрытие очереди

	wg.Wait()
}
