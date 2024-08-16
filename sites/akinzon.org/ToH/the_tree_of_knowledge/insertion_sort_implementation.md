# Insertion Sort Implementation

Up: [Insertion Sort](insertion_sort)
Brother(s):
TARGET DECK

```Java
public static void sorted(int[] a) {  
  for (int i = 1; i < a.length; i++) {  
    int current = a[i]; // needed to keep track of current
    // since a[i] will change value 
    int count = i-1;  
    // current barrels down the list until it reaches its
    // correct position
    while (count >= 0 && a[count] > current) {  
      int temp;  
      temp = a[count+1];  
      a[count+1] = a[count];  
      a[count] = temp;  
      count--;  
    }  
  }  
}
```



































#### Why:
#### How:









