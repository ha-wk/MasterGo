📌 Understanding Worker Pools
A worker pool is a concurrency pattern that helps manage and control the number of goroutines running simultaneously. Instead of spawning thousands of goroutines (which can overwhelm the system), we create a fixed number of "workers" that process tasks from a queue.

🛠️ How It Works:
A job queue (channel) holds incoming tasks.
Multiple workers (goroutines) fetch and process tasks from the queue.
Once a worker finishes, it takes the next available task.
✅ Why Use Worker Pools?

Prevents excessive goroutines that can slow down performance.
Efficient resource usage by reusing goroutines.
Helps maintain system stability under heavy load.







📌 Fan-out & Fan-in Patterns
These patterns help manage data flow efficiently in concurrent systems.

🔷 Fan-out
Multiple goroutines read from the same channel (distributing work).
Used to increase parallelism.
✅ Example: Multiple goroutines process requests from a queue.

🔷 Fan-in
Multiple channels merge into a single channel.
Used to combine results.
✅ Example: Multiple worker goroutines send results to a single output channel.



🛠️ Hands-on Practice Questions
1️⃣ Worker Pool Implementation

Write a worker pool where 4 workers process 10 numbers and return their squares.
2️⃣ Fan-out Pattern

Implement 4 goroutines that fetch data from a common channel and process it.
3️⃣ Fan-in Pattern

Merge the outputs of three separate channels into a single channel.
4️⃣ Worker Pool with Timeout

Modify the worker pool example to stop processing jobs after 3 seconds using the context package.
