# Lab II Key commands

Up: [Computer Systems Labs](computer_systems_labs)



- `ps` -  writes the status of active processes, their Process IDs and other info (as a snapshot, not continuous)
- `top` - stands for table of processes, gives a real-time view of running processes
- `kill` - sends a signal to a job
- `bg` - Move jobs to the background
- `fg` - Move jobs to the foreground
- `taskset` - used to set or retrieve the CPU affinity of a running process given its Process ID, or to launch a new command with a given CPU affinity
- `dd` - Copies and converts a file

# Additional notes:
- `./test.py &` - Runs `test.py` in the background
- `taskset -c 1 ./test.py` - Runs `test.py` on CPU core 1
- `dd if=dev/urandom of=dev/null` - Sends a constant stream of random data `urandom` to a "black hole" `null`