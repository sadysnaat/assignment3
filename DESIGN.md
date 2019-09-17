### Task
A Task is an abstract representation of a running task in real world load

Task has three states from running to end to

```
                              +-------------+
                              |             |
       +--------------------->+  Completed  |
+-----------+                 |             |
|           |                 +-------------+
|  Running  |
|           |                 +----------+
+-----------+                 |          |
       +--------------------->+  Failed  |
       |                      |          |
       |                      +----------+
       |
       |                      +-----------+
       |                      |           |
       +--------------------->+  Timeout  |
                              |           |
                              +-----------+


``` 

#### Manager
Manger enqueues and dedupe the tasks sync.Map provides dedupe with the help of keys

Manager runs purge 
* at an interval of 1 sec 
* finds the current snapshot of Map and iterates over each key value pair if found
    * if task is running noop
    * if task is completed delete from manager
    * if task is failed or timeout
        * delete task
        * enqueue task
* else mark work as complete and return