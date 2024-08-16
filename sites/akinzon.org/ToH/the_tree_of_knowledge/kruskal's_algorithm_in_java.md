# Kruskal's Algorithm in Java

Up: [High-Level Implementation of Kruskal's Algorithm](high-level_implementation_of_kruskal's_algorithm)
Brother(s):
TARGET DECK

```java
public ArrayList<Edge> minimumSpanningTreeKruskal() {  
  ArrayList<Edge> mst = new ArrayList<>();  
  UnionFindMakeSet<Node<T>> uf = new UnionFindMakeSet<>(vertices.size());  
  for (Node<T> vertex: vertices) {  
    uf.makeSet(vertex);  
  }  
  MinHeap<Edge> minHeap = new MinHeap<>();  
  ArrayList<Edge> nodeEdges = getEdges();  
  // heap sort on edges  
  for (Edge edge: nodeEdges) {  
    minHeap.add(edge, edge.getWeight());  
  }  
  ArrayList<Edge> sortedEdges = new ArrayList<>();  
  for (int i = 0; i < nodeEdges.size(); i++) {  
    sortedEdges.add(minHeap.removeMin());  
  }  
  for(Edge edge: sortedEdges) {  
    // if u and v are in the same set but the edge (u,v) is not in the MST, adding it will create a cycle.  
    if (uf.getSet(edge.getU()) != uf.getSet(edge.getV())) {  
      mst.add(edge);  
      uf.setUnion(edge.getU(), edge.getV());  
    }  
  }  
  return mst;
```



































#### Why:
#### How:









