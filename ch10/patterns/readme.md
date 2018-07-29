# 并行模式 Go

设计原则
- Design your program as a collection of independent processes
- Design these processes to eventually run in parallel
- Design your code so that the outcome is always the same

目标
- no race conditions
- no deadlocks
- more goroutine
- readable

go CSP工具
- go 
    - gmp
- chan
    - buffer
    - sender
    - receiver
    - close channels
        - Close sends "closed" message
        - try send more panic
- select
    >more than one chan
    - no order
    - non-blocking
- sync
    - mutexes not good
    - atomic not good
    - pool
    - waitgroup