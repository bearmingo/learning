# 学习 Neural Network And Deep Learning 练习


## sigmoid

Sigmoid函数

$$
\sigma(z)=\frac{1}{1+e^{-x}}
$$

python实现代码

```python
def sigmoid(z):
    return 1.0/(1.0 + np.exp(-z))
```

Sigmoid的导数

$$
{\sigma}'(z)={(\frac{1}{1+e^{-z}})}'=\frac{e^{-z}}{(1+e^{-z})^2}=\frac{1+e^{-z}-1}{(1+e^{-z})^2}=\frac{1}{1+e^{-z}}(1-\frac{1}{1+e^{-z}})=\sigma(z)\cdot(1-\sigma(z))
$$

python实现代码
```python
def sigmoid_prime(z):
    return sigmoid(z)*(1-sigmoid(z))
```

## 训练方式

```python
import mnist
import network
import numpy as np

training_data, test_data = mnist.load_data()

net = network.Network([784, 30, 10])
net.SGD(training_data, 30, 10, 3.0, test_data=test_data)
```

## References

- [Neural Networks and Deep Learning](http://neuralnetworksanddeeplearning.com/)
- [THE MNIST DATABASE](http://yann.lecun.com/exdb/mnist/)

