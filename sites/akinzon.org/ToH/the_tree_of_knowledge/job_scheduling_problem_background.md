# Job Scheduling Problem Background

Up: [The Job Scheduling Algorithm](the_job_scheduling_algorithm)
Brother(s):
TARGET DECK

Say we had $n$ jobs, that all had a weight $w_i$ and a length $l_i$, $1 \le i \le n$.

Also say that the jobs share the same resource (like the [The Processor|CPU](the_processor|cpu)) and so they must run sequentially.

By running the jobs in the order 1, 2, 3, a job $i$ has the completion time $c_i = \sum_{k=1}^i = l_k$ (this is the completion time for a particular job $c_i$ and the total completion times of all jobs that came before it.)

The problem of job scheduling is ordering the jobs such that we can minimise $f = \sum_{k=1}^n w_k \cdot c_k$, the weighted/average completion time for all jobs.

> [!TIP] Akin, while it seemed like it would take the same time for all jobs to complete no matter the order, this isn't true. It makes sense, doesn't it? Why should shorter jobs wait for longer jobs to finish?

This problem has 2 special cases:
	**All jobs have equal weighting/priority**: Shortest job goes first
	**All jobs have equal length**: Highest priority job goes first

The general case is derived from both special cases, and forms the Job Scheduling Algorithm.



































#### Why:
#### How:









