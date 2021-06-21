# [WIP] 채널

This line

```go
c <- 1
```
blocks because the channel is unbuffered. As there's no other goroutine to receive the value, the situation can't resolve, this is a deadlock.

You can make it not blocking by changing the channel creation to

```go
c := make(chan int, 1) 
```
so that there's room for one item in the channel before it blocks.

But that's not what concurrency is about. Normally, you wouldn't use a channel without other goroutines to handle what you put inside. You could define a receiving goroutine like this :

```go
func main() {
    c := make(chan int)    
    go func() {
        fmt.Println("received:", <-c)
    }()
    c <- 1   
}
```
