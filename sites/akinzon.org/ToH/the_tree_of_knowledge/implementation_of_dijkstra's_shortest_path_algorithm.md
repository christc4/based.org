# Implementation of Dijkstra's Shortest Path Algorithm

Up: [Dijkstra's Shortest Path Algorithm](dijkstra's_shortest_path_algorithm)
Brother(s):
TARGET DECK

```Java
public ArrayList<Edge> dijkstraShortestPath(T value, T destination) {  
  // declare the unvisited list  
  MinHeap<Node<T>> minHeap = new MinHeap<>();  
  Node<T> destNode = findNode(destination);  
  ArrayList<Edge> path = new ArrayList<>();  
  // add all edges to the unvisited list with key max_value and previous of null  
  for (Node<T> node: vertices) {  
    minHeap.add(node, node.getKey());  
  }  
  // set start node to key 0  
  Node<T> source = findNode(value);  
  source.setKey(0);  
  // while unvisited list is not empty  
  while (minHeap.size() > 0) {  
    //  current node = removeMin()  
    Node<T> currentNode = minHeap.removeMin();  
    //  for each neighbour of current node  
    for (Node<T> adjacentNode: adj(currentNode)){  
      //    relax the neighbour  
      relax(currentNode, adjacentNode, minHeap);  
    }  
    currentNode.setColour("visited");  
  }  
  while (destNode != source) {  
    path.add(new Edge(destNode, destNode.getPrevious(), 0));  
    destNode = destNode.getPrevious();  
  }  
  return path;  
}
```



































#### Why:
#### How:









