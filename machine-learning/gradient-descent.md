# 梯度

## 概念

- 步长(learning rate)：步长决定了在梯度下降迭代的过程中，每一步沿梯度负方向前进的长度
- 特征（feature)：指样本中的输入部分
- 假设函数（hypothesis function):在监督学习中，为来拟合输入样本，而使用的假设函数，记为$h_\theta(x)$。比如单特征的m个样本$(x^{(i)},y^{(i)})(i=1,2,3,...m)$,可以采用的拟合函数如下：$h_\theta(x)=\theta_0 + \theta_1x$
- 损失函数（loss function)：为来评估模型拟合的好坏，通常使用损失函数来度量拟合的程度。损失函数级小化，意味者拟合的程度最好，对应的模型参数即为最优参数。在线性回归函数中，损失函数通常为样本暑促和假设函数的差取平方。比如对于m个样本$(x_i,y_i)(i=1,2,3,...m)$，采用线性回归，损失函数为：
$$
J(\theta_0, \theta_1)=\sum_{x=1}^{m}{\left(h\theta(x_i)-y_i\right)^2}
$$

其中$x_i$表示第i个样本特征，$y_i$表示第i个样本对应的输出，$h_\theta(x_i)$为假设函数。

## 梯度下降的详细算法
### 代数描述
1. 先决条件：确认优化模型的假设函数和损失函数。
   
比如对于线性回归，假设函数表示为$h_\theta(x_1,x_2,...x_n)=\theta_0+\theta_1x_1+...+\theta_nx_n$，其中$\theta_i(i=0,1,2...n)$为每个样本的n个特征值。这个表示可以简化，增加一个特征$x_0=1$，这样$h_\theta(x_0,x_1,...x_n)=\sum^{n}_{i=0}{\theta_ix_i}$。

同样是线性回归函数，对于上面的假设函数，损失函数为：
$$
J(\theta_0, \theta_1,...\theta_n)=
\frac{1}{2m}\sum_{j=0}^m(h_\theta(x_0^{(j)},x_1^{(j)},...x_n^{(j)})- y_j)^2
$$

2. 算法相关参数初始化：主要式初始化$\theta_0, \theta_1, ...\theta2$，算法终止距离$\varepsilon$以及不长$\alpha$。在没有任何先验知识的情况下，将说有的$\theta$初始化为0，将步长初始化为1。在调优的时候在优化。
3. 算法过程
 
    1. 确定当前位置的损失函数的梯度，对于$\theta_i$其梯度表达式为：
   
    $$
    \frac{\partial}{\partial\theta_i}J(\theta_0, \theta_1,...\theta_n)
    $$

    2. 使用步长乘以损失函数的梯度，得到当前位置的下降的距离，即$\alpha\frac{\partial}{\partial\theta_i}J(\theta_0, \theta_1,...\theta_n)$
    3. 确定是否所有的$\theta_i$，梯度下降的距离都小于$\varepsilon$，如果小于$\varepsilon$则终止算法，当前所有的$\theta_i(i=0,1,...n)$即为最终结果。否则进入步骤`4`中
    4. 更新所有的$\theta_i$，对于$\theta_i$，其更新表达式如下。更新完毕后继续进入步骤`1`
    
    $$
    \theta_i=\theta_i-\alpha\frac{\partial}{\partial\theta_i}J(\theta_0, \theta_1,...\theta_n)
    $$

举例使用线行回归的例子具体描述梯度下降。假设有样本：
$
(x_0^{(0)},x_1^{(0)},...x_n^{(0)}, y_0),(x_0^{(1)},x_1^{(1)},...x_n^{(1)}, y_1)...(x_0^{(m)},x_1^{(m)},...x_n^{(m)}, y_m)
$，损失函数如先决条件所述：$J(\theta_0, \theta_1,...\theta_n)=
\frac{1}{2m}\sum_{j=0}^m(h_\theta(x_0^{(j)},x_1^{(j)},...x_n^{(j)})- y_j)^2$

则在算法过程`步骤1`中对$\theta_i$的偏导数如下：

$$

\frac{\partial}{\partial\theta_i}J(\theta_0, \theta_1,...\theta_n)=
\frac{1}{m}\sum_{j=0}^m(h_\theta(x_0^{(j)},x_1^{(j)},...x_n^{(j)})- y_j)x_i^{(j)}
$$


## 神经网络损失函数求导

$$
J(\Theta)=-\frac{1}{m} \sum_{i=1}^{m} \sum_{k=1}^{K}\left[y_{k}^{(i)} \log \left(\left(h_{\Theta}\left(x^{(i)}\right)\right)_{k}\right)+\left(1-y_{k}^{(i)}\right) \log \left(1-\left(h_{\Theta}\left(x^{(i)}\right)\right)_{k}\right)\right]+\frac{\lambda}{2 m} \sum_{l=1}^{L-1} \sum_{i=1}^{s l+1}\left(\Theta_{j, i}^{(l)}\right)^{2}
$$


https://www.cnblogs.com/pinard/p/5970503.html