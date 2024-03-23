# KC_QUEUE

Non production grade queue to use for projects


# Example (not tested, sorry!)

```go
package main

type MessageService struct {
    queue *Queue
}

func MessageServiceCreate() *MessageService {
    result := &MessageService{}
    result.messageQueue = New()
    return result
}

// Get the latest message in the queue
func (ms* MessageService) Get() Message {
    item := ms.messageQueue.Deque()
    msg, _ := item.(Message)
    return msg
}

// Get the latest message in the queue
func (ms* MessageService) Set(m Message) {
    ms.messageQueue.Enqueue(m)
}
```
