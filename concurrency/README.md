## Goroutine
Goroutines can be thought of as lightweight threads. They share memory, so writing to the same global variables needs special care.

Because of their weight, you can run even hundreds of thousands of them.

Their relatively small size comes from the fact that they're defined and handled by go runtime. They're created with initial size of 2KB and can become larger when needed.  
On the other hand, OS threads have a fixed size, often as large as 2MB.

## Channel
A channel is a thread-safe queue of values of a given type.

It's used to communicate between goroutines.


