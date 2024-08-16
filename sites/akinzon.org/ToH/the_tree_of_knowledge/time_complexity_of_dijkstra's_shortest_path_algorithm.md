# Time Complexity of Dijkstra's Shortest Path Algorithm

Up: [Dijkstra's Shortest Path Algorithm](dijkstra's_shortest_path_algorithm)
Brother(s):
TARGET DECK

Typically, the specific running time of Dijkstra's algorithm is $O((|V|+|E|)log(|V|))$.

$O(|V|log(|V|)$ comes from executing `removeMin()` on the [Priority Queues|priority queue](priority_queues|priority_queue) implemented as a [Heaps|heap](heaps|heap), which runs in $O(log(|V|))$ time,  $|V|$ times.

$O(|E|log(|V|)$ comes from the [Edge Relaxation|relax()](edge_relaxation|relax()) function which runs in $O(log(|V|)$ time due to having to update the priority queue with the new weights, $|E|$ times.

Most people will just say it is $O(m log (n))$.

































#### Why:
#### How:









