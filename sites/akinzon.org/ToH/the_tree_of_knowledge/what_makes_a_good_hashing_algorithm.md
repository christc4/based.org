# What Makes a Good Hashing Algorithm

Up: [Hashing Algorithms](hashing_algorithms)
Brother(s):
TARGET DECK

Good Hashing Algorithms:
	- Compute the address fast
	- Distribute values evenly across the hash table to minimise [Hash Collisions|collisions](hash_collisions|collisions)
	- Ensure that equal objects have equal key
	- Use a prime number for the table size to aid in avoiding clusters (if not, multiples of its factors will hash to the same place, but prime numbers have as little factors as possible)

**General form of hash function:**

`h(k) = k % M`

Where $k$ is the value and $M$ is the size of the table.

































#### Why:
#### How:









