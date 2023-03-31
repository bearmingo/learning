## 数据结构
Series: 一维数组，与[NumPy](./numpy)的一维`array`类似。
Time-Series：以时间为索引的`Seriie`。
DataFrame：二维的表格数据结构。

## 操作

### 模块导入
```python
import numpy as np
import pandas as pd
```

### Series

创建一个Series：
```python
s = pd.Series(data, index)
```

|参数|意义|
|--|--|
|data|字典、ndarray或者标量（标量就是只有大小，没有方向的量）|
|index|对data的索引值，类似字典的key，index参数事可以省略的。如果不带index参数，pandas会自动用默认的index进行索引，比如ndarray数组，索引值就是\[0,...len(data)-1\]|

可以从以下几种方式创建
**1. 从python的list创建**
```python
contries = ['USA', 'China']
data = [1, 200]
pd.Series(data, countries)
```

**2. 从numpy ndarry创建Series**
```python
np_arr = np.array(data)
pd.Series(np_arr)
```

**3. 从python的dict创建Series**
```python
d = {'a': 1, 'b': 2}
pd.Series(d)
```
但data为dict时，且没有传入索引，索引就为dict的key，顺序为dict的插入顺序。

**4. 从标量中创建Series**
```python
pd.Series(5, index['a', 'b'])
```

**对 Series 进行算术运算操作**
基于 index 进行的。我们可以用加减乘除（+ - * /）这样的运算符对两个 Series 进行运算，Pandas 将会根据索引 index，对响应的数据进行计算，结果将会以浮点数的形式存储，以避免丢失精度。如果 Pandas 在两个 Series 里找不到相同的 index，对应的位置就返回一个空值 NaN。

### DataFrame

1. 使用Series创建一个DataFrame:
```python
df = {'Name': pd.Series(['Jon', 'Aaron', 'Todd'], index=['a', 'b', 'c']),
	  'Age': pd.Series(['30', '30', '40'], index=['a', 'b', 'c', 'd']),
	  'Nationality': pd.Series(['US', 'Chine', 'US'], ['a', 'b', 'c'])
}
pd.DataFrame(df)
```

2. 使用字典来创建一个DataFrame：
```python
df = {'Name': pd.Series(['Jon', 'Aaron', 'Todd']),
	  'Age': pd.Series(['30', '30', '40']),
	  'Nationality': pd.Series(['US', 'Chine', 'US'])
}
pd.DataFrame(df, , index=['a', 'b', 'c'])
```

