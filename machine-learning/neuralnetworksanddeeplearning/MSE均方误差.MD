# MSE均方误差

## 公式及导数推导

### 损失函数

$$
C=\frac{1}{2n}\|y(x) - a^L(x)\|^2
$$
其中：$a=\sigma(z)$, $z=wx+b$

对$w$和$b$求导：
$$
\frac{\partial C}{\partial w}=(a-y)\sigma'(z)x \\
\frac{\partial C}{\partial b}=(a-y)\sigma'(z)
$$

### $\sigma$推导

$$
\begin{aligned}
\sigma'(z) & = (\frac{1}{1+e^{-z}})' \\
    & = (-1)(\frac{1}{1+e^{-z}})^{(-1)-1}\cdot(e^{-z})' \\
    & = \frac{1}{(1+e^{-z})^2}\cdot(e^{-z}) \\
    & = \frac{1}{1+e^{-z}}\cdot(\frac{e^{-z}}{1+e^{-z}}) \\
    & = \frac{1}{1+e^{-z}}\cdot(1-\frac{1}{1+e^{-z}}) \\
    & = \sigma(z)(1-\sigma(z))
\end{aligned}
$$
即
$$
\phi'(x)=\phi(x)(1-\phi(x))
$$
