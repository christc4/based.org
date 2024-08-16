![Pasted image 20240223085449.png](pasted_image_20240223085449.png)

![Pasted image 20240223085920.png](pasted_image_20240223085920.png)

E differente

main branch most likely the outdated one
needs to be updated via reverse integration
Here. We. Go.


```python
import pandas as pd
import matplotlib.pyplot as plt
import os
import glob
import math
user_key= 25468

def plot_all_csv_pressure():
    path = os.getcwd()
    csv_files = glob.glob(os.path.join(path, '*.csv'))

    fr f in csv_files:
        storm = pd.read_csv(f)
        storm['Pressure'].plot()
        plt.show()

def plot_all_csv_intensity():
    path = os.getcwd()
    csv_files = glob.glob(os.path.join(path, '*.csv'))

    for f in csv_files:
        storm = pd.read_csv(f)
        storm['Intensity'].plot()
        plt.show()
```

imports in wrong order
spelling mistake in the for loop
`user_key` wrong formatting



```python
import pandas as pd
import matplotlib.pyplot as plt
import os
import glob
import math
user_key = 8073

def plot_all_csv_pressure():
    path = os.getcwd()
    csv_files = glob.glob(os.path.join(path, '*.csv'))

    for f in csv_files:
        storm = pd.read_csv(f)
        storm['Pressure'].plot()
        plt.show()
```


![Pasted image 20240223095808.png](pasted_image_20240223095808.png)


![Pasted image 20240223095832.png](pasted_image_20240223095832.png)