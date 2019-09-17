```
// These logs are for enqueuing 
[task] 4 will be completed after 188 ms or timeout after 773 ms or failed after 505ms
[task] 1 will be completed after 781 ms or timeout after 891 ms or failed after 412ms
[task] 0 will be completed after 495 ms or timeout after 831 ms or failed after 831ms
[task] 2 will be completed after 12 ms or timeout after 921 ms or failed after 519ms
[task] 3 will be completed after 428 ms or timeout after 604 ms or failed after 64ms
[task] 2 completed
[task] 3 failed
[task] 4 completed
[task] 1 failed
[task] 0 completed


// A purge cycle starts 
[manager] starting purge
[manager] task found 3 failed
[manager] requeue task 3
[manager] task found 4 completed
[manager] task found 0 completed
[manager] task found 1 failed


// Failed tasks are being enqueued
[manager] requeue task 1
[manager] task found 2 completed
[task] 3 will be completed after 358 ms or timeout after 515 ms or failed after 762ms
[task] 1 will be completed after 965 ms or timeout after 261 ms or failed after 538ms
[manager] starting purge
[manager] task found 1 running
[manager] task found 3 running
[task] 1 timeout occured
[task] 3 completed
[manager] starting purge
[manager] task found 1 timeout
[manager] requeue task 1
[manager] task found 3 completed
[task] 1 will be completed after 198 ms or timeout after 302 ms or failed after 401ms
[manager] starting purge
[manager] task found 1 running
[task] 1 completed
[manager] starting purge
[manager] task found 1 completed
[manager] starting purge
[manager] no more items remaining
[manager] starting purge
[manager] no more items remaining
```