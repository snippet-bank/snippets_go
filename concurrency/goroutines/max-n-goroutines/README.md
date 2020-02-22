## Limiting access to a common resource in a multitasking system.

In computer science, a semaphore is an integer variable used to control access to resource shared by multiple processes.

Typically, a semaphore can never drop below 0 and has two operations:

- acquiring - decreasing the semaphore's value by 1; when trying to decrease the value below 0, it blocks until such an action is possible (i.e. the semaphore is positive)
- releasing - increasing the value

## Setting upper limit for number of goroutines 

It's just another example of the previously described problem.

Goroutines are the _shared resource_.

The semaphore is implemented through a buffered channel:
- the size of the buffer is the maximum value of the semaphore
- "acquiring the semaphore" is sending a message to the channel; if there's no space left in the buffer, the goroutine blocks until there is
- "releasing the semaphore" is reading a message from a channel, thus releasing one space in the buffer (increasing the "value" of the semaphore by 1) 

