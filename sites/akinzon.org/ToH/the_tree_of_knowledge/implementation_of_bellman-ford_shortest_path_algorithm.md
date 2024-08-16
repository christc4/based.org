# Implementation of Bellman-Ford Shortest Path Algorithm

Up: [Bellman-Ford Shortest Path Algorithm](bellman-ford_shortest_path_algorithm)
Brother(s):
TARGET DECK

```Java
public ArrayList<Edge> bellmanFordShortestPath(T value, T destination) {  
  Node<T> node = findNode(value);  
  Node<T> destNode = findNode(destination);  
  ArrayList<Edge> path = new ArrayList<>();  
  if (node != null && destNode != null) {  
    node.setKey(0);  
    ArrayList<Edge> sortedEdges = getSortedEdges();  
    for (int i = 0; i < vertices.size() - 1; i++) {  
      for (Edge edge: sortedEdges) {  
        relax(edge);  
      }  
    }  
    while (destNode != node) {  
      path.add(new Edge(destNode, destNode.getPrevious(), 0));  
      destNode = destNode.getPrevious();  
    }  
  }  
  return path;  
}
```



































#### Why:
#### How:









