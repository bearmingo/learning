# GAN

[Generative Adversarial Networks](https://arxiv.org/abs/1406.2661)

## G（Generator)和D(Discriminator)

- G是一个生产图片的网络，它接受一个随机的噪声`z`，通过这个噪声生产图片，记做$G(z)$

- D是一个判别网络，判别一张图片是不是“真实的”。它的输入参数是$x$，$x$代表一张图片，输出$D(x)$代表$x$为真实图片的概率。如果为1，就代表100%是真实的图片，而输出为0，代表不能是真实的图片

在训练过程中，生成网络$G$的目标就是尽量生成真实的图片去欺骗判别网络$D$。而$D$的目标就是尽量把$G$生成的图片和真实的图片区分开来。

这样，G和D构成来一个动态的“博弈过程”。

最后博弈的结束是：在最理想的状态下，$G$可一生成足以“以假乱真”的图片$G(z)$。对于$D$来说。它难以判定$G$生成的图片究竟是不是真实的，因此$D(G(z))=0.5$

最终就得到一个生成式的模型G, 可以用来生成图片。

## GAN的核心原理：
$$
\min _{G} \max _{D} V(D, G)=\mathbb{E}_{\boldsymbol{x} \sim p_{\text { data }}(\boldsymbol{x})}[\log D(\boldsymbol{x})]+\mathbb{E}_{\boldsymbol{z} \sim p_{\boldsymbol{z}}(\boldsymbol{z})}[\log (1-D(G(\boldsymbol{z})))]
$$

公式分析：

- 整个式子由两项构成。x表示真实图片，z表示输入G网络的噪声，而G(z)表示G网络生成的图片。
- $D(x)$表示$D$网络判断真实图片是否真实的概率（因为x就是真实的，所以对于$D$来说，这个值越接近1越好）。而$D(G(z))$是$D$网络判断$G$生成的图片的是否真实的概率。
- $G$的目的：上面提到过，$D(G(z))$是$D$网络判断$G$生成的图片是否真实的概率，$G$应该希望自己生成的图片“越接近真实越好”。也就是说，$G$希望$D(G(z))$尽可能得大，这时$V(D, G)$会变小。因此我们看到式子的最前面的记号是$\min _{G}$。
- $D$的目的：$D$的能力越强，$D(x)$应该越大，$D(G(x))$应该越小。这时$V(D,G)$会变大。因此式子对于D来说是求最大$\max _D$

![](./GAN-adversarial-nets-framework.jpg)

随机梯度下降法训练D和G？论文中也给出了算法：

![](./GAN-gradient-descent-training.jpg)



## DCGAN

图片处理应用最好的模型CNN 与GAN结合

## Ref

[GAN学习指南：从原理入门到制作生成Demo](https://zhuanlan.zhihu.com/p/24767059)

