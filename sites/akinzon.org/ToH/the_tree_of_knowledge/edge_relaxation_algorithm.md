# Edge Relaxation Algorithm

Up: [Edge Relaxation](edge_relaxation)
Brother(s):
TARGET DECK

```Java
public void relax(Edge edge) {  
  if (edge.getU().getKey() + edge.getWeight() < edge.getV().getKey()) {  
    edge.getV().setPrevious(edge.getU());  
    edge.getV().setKey(edge.getU().getKey() + edge.getWeight());  
  }  
}
```

```java
public void relax(Node<T> u, Node<T> v) {  
  int weight;  
  if (u.getKey() + (weight = getWeight(u.getValue(), v.getValue())) < v.getKey()) {  
    v.setPrevious(u);  
    v.setKey(u.getKey() + weight);  ``
  }  
}
```


































#### Why:
#### How:









