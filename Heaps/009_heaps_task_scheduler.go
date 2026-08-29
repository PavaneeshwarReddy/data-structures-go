package heaps

/*
Task Scheduler
- Let's consider 2 things ,
- MaxHeap -> Which stores frequenciesLeft ( total no of tasks per character)
- Queue -> [frequenceLeft, atWhichTimeWeCanProcess ]
- Let's take this as example ['A', 'A'] and waiting period is 2
- time:0 -> maxHeap: [2], queue: Nil
- time:1 -> maxHeap: Nil, queue: [1 (2 - 1), 3 (1 + 2)]
- time:2 -> maxHeap: Nil, queue:[1,3] - we cannot pick this task as time is still 2
- time:3 -> maxHeap: Nil, queue: Nil
*/

func TasksScheduler(tasks []byte, n int) int {
	maxHeap := MaxHeap{}
	counts := [26]int{}

	for _, val := range tasks {
		counts[val-'A']++
	}

	for _, val := range counts {
		if val != 0 {
			maxHeap.Insert(val)
		}
	}

	queue := [][2]int{}
	time := 0

	for maxHeap.Length() > 0 || len(queue) > 0 {
		time += 1
		if maxHeap.Length() > 0 {
			taskFreq := maxHeap.Pop() - 1
			if taskFreq > 0 {
				queue = append(queue, [2]int{taskFreq, time + n})
			}
		}

		if len(queue) > 0 && queue[0][1] == time {
			maxHeap.Insert(queue[0][0])
			queue = queue[1:]
		}
	}

	return time

}
