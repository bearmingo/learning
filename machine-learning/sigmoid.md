# sigmoid

## 函数表达式

$$
\sigma = \frac{1}{1+e^{-1}}
$$

## 导数推导

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





