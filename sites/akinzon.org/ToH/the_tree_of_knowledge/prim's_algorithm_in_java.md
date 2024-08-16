# Prim's Algorithm in Java

Up:
Related Note(s): [Kruskal's Algorithm in Java](kruskal's_algorithm_in_java)
TARGET DECK

```Java
public ArrayList<Edge> minimumSpanningTreePrim() {  
  ArrayList<Edge> mst = new ArrayList<>();  
  Node<T> source = vertices.get(ThreadLocalRandom.current().nextInt(vertices.size()));  
  source.setKey(0);  
  MinHeap<Node<T>> minHeap = new MinHeap<>();  
  for (Node<T> node: vertices) {  
    minHeap.add(node, node.getKey());  
  }  
  while (minHeap.size() > 0) {  
    source = minHeap.removeMin();  
    if (source.getPrevious() != null) mst.add(new Edge(source, source.getPrevious(), 0));  
    for (Node<T> node: adj(source)) {  
      int weight;  
      if ((weight = getWeight(source, node)) < node.getKey()) {  
        node.setKey(weight);  
        // The problem I was having is that nodes would have their keys changed, but the actual nodes in the minHeap would not have their priorities changed.  
        // My theory for no cycles being added is that since we only go for nodes still in the minHeap, therefore we cannot removeMin a node that has already been visited.        MinHeap.Node<Node<T>> node1 = minHeap.findNode(node);  
        if (node1 != null) node1.setPriority(weight);  
        node.setPrevious(source);  
      }  
      // Since some priorities have changed, the minHeap is not necessarily in a legal form.  
      // So we restructure it.      minHeap.recomputeMinHeap();  
    }  
  }  
  return mst;  
}
```



































#### Why:
#### How:









