package main

import "fmt"

func enqueue(queue []int, element int) []int {
queue = append(queue, element) // Simply append to enqueue.
fmt.Println("Enqueued:", element)
return queue
}

func dequeue(queue []int) (int, []int) {
element := queue[0] // The first element is the one to be dequeued.
if len(queue) == 1 {
var tmp = []int{}
return element, tmp

}

return element, queue[1:] // Slice off the element once it is dequeued.
}

func main() {
var queue = make([]int, 0)

queue = enqueue(queue, 10)

fmt.Println("After pushing 10 ", queue)
queue = enqueue(queue, 20)

fmt.Println("After pushing 20 ", queue)
queue = enqueue(queue, 30)

fmt.Println("After pushing 30 ", queue)

ele, queue := dequeue(queue)
fmt.Println("Queue After removing", ele, " :", queue)

queue = enqueue(queue, 40)
fmt.Println("After pushing 40 ", queue)
}