# Double Hashing

Up: [Hash Collision Resolution](hash_collision_resolution)
Brother(s):
TARGET DECK

Double hashing involves using another [Hashing Algorithms|hashing algorithm](hashing_algorithms|hashing_algorithm) to find a hopefully unique number of spaces to check after the one we couldn't go to in order to place our value in.

It looks crazy op ngl:

$$ h(x) + (i * h_2(x)) $$
The first hashing algorithm depends on the value itself, but the 2nd one depends on the key.

People normally use $h_2(x) = R - (x\%R)$ where $R$ is a prime smaller than the table size.





































#### Why:
#### How:









