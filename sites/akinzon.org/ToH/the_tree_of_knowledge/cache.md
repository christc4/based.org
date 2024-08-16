# Cache (Money Money Cache Cache!!!)

Up: [Internal Memory](internal_memory)
Brother(s): [RAM](ram)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Computer Systems

Caches are small blocks of fast [SRAM](sram) located on the [The Processor|CPU](the_processor|cpu) chip 

In truth, memory requests go there, and not to [DRAM](dram)

For example:
![Pasted image 20231123200307.png](pasted_image_20231123200307.png)


## How do it work?

1. CPU tries to read a memory location
2. address goes to cache
3. If it's in cache, the cache sends it to the CPU (called a cache  hit) 
4. If not, the required block is read from RAM to cache (called a cache miss)
5. Then from cache it is sent to the CPU


FYI, cache lines are chunks of information stored in cache, and are much larger than typical words.

More cache = less cache miss, but only before a certain threshold.

Cache acts as in intermediary for [RAM](ram) and the [The Processor|CPU](the_processor|cpu)

Furthermore, it makes sense to split instructions and data into their own cache so that they can be fetched in parallel.

The benefit of Level-2 cache is that it reduces the penalty of reading from/writing to RAM (51 ns), keeping the latency down.

The benefit of Level-3 cache is that it acts as a shield that all cores communicate with before going to RAM first (and hopefully they find what they're looking for in the cache!)

The Level-3 cache is used so multiple cores share a large cache.

**Cache Coherency** - The mechanism by which CPUs are sure they are using the most up-to-date value in cache.

![Pasted image 20231123201939.png](pasted_image_20231123201939.png)

Cache isn't as big as RAM - need [Cache Mapping](cache_mapping) to make it work

What is cache and how is it used #flashcard 
Caches are small blocks of SRAM on the CPU chip and the CPU will request to read from it before going to RAM if it cannot find what it is looking for in cache.
<!--ID: 1700775011485-->


What are the benefits of Levels-1,2,3 cache #flashcard 
Level-1 cache - splits instructions and data so they can be read in parallel
Level-2 cache - maintains the low latency of reading cache
Level-3 cache - allows multiple cores to share cache
<!--ID: 1700775011492-->






















#### Why:
#### How:









