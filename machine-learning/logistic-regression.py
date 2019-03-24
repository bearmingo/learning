import matplotlib.pyplot as plt
import numpy as np

def sigmoid(x):
    return 1.0 / (1.0 + np.exp(-x))

x = np.arange(-10, 10, 0.1)
h = sigmoid(x)

plt.plot(x, h)
plt.axvline(0.0, color='k')
plt.axhspan(0.0, 1.0, fac)