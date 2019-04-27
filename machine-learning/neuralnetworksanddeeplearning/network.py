#!/usr/bin/env python
# -*- coding=utf-8 -*-

"""
"""

import random

import numpy as np

# 注释使用[784, 30, 10]进行说明

class Network(object):

    def __init__(self, sizes):
        """The list `size` contains the number of nerons in the
        respective layers of the network. For example if the list
        was [2, 3, 1] then it would be a three-layer network, with the
        first layer contains 2 neurons, the second layer 3 neurons
        and the third layer 1 neuron.
        The biases and weights for the network are initialized randomly,
        using a Gaussian distribution with mean 0 and variance 1.
        Note that the first layer is assumed to be an input layer, and 
        conviention we won't se any biases for those neurons, since
        biases are only ever used in computing the outputs from later
        layers.
        """
        self.num_layers = len(sizes)
        self.sizes = sizes
        # 除了第一层，其他节点都有一个偏离量
        # biases[0]的形状是(30,) biases[1] 是(10,)
        self.biases = [np.random.randn(y, 1) for y in sizes[1:]]
        # 每个节点的输入项与上一层中的节点个数相同
        # weights[0]的形状是(30, 784), weights[1] 是(10, 30)
        self.weights = [np.random.randn(y, x) for x, y in zip(sizes[:-1], sizes[1:])]

    def feedforward(self, a):
        """Return the output of the network if `a` is input"""
        for b , w in zip(self.biases, self.weights):
            a = sigmoid(np.dot(w, a) + b)
        return a
    
    def SGD(self, training_data, epochs, mini_batch_size, eta,
            test_data=None):
        """Train the neural network using mini-batch stochastic
        gradient desceent. The `training_data` is a list of tuples
        `(x, y)` representing the training inputs and the desired
        output. 
        """
        if test_data:
            n_test = len(test_data)
        n = len(training_data)
        for j in range(epochs):
            random.shuffle(training_data)
            mini_batches = [
                training_data[k:k+mini_batch_size]
                    for k in range(0, n, mini_batch_size)]
            for mini_batch in mini_batches:
                self.update_mini_batch(mini_batch, eta)
            if test_data:
                print('Epoch {0}: {1}/{2}'.format(
                    j, self.evaluate(test_data), n_test))
            else:
                print('Epoch {0} complete'.format(j))

    def update_mini_batch(self, mini_batch, eta):
        
        # 微分
        nabla_b = [np.zeros(b.shape) for b in self.biases]
        nabla_w = [np.zeros(w.shape) for w in self.weights]
        for x, y in mini_batch:
            delta_nabla_b, delta_nabla_w = self.backprop(x, y)
            nabla_b = [nb + dnb for nb, dnb in zip(nabla_b, delta_nabla_b)]
            nabla_w = [nw + dnw for nw, dnw in zip(nabla_w, delta_nabla_w)]
        self.weights = [w - (eta/len(mini_batch)) * nw
                        for w, nw in zip(self.weights, nabla_w)]
        self.biases = [b - (eta/len(mini_batch)) * nb
                       for b, nb in zip(self.biases, nabla_b)]

    def backprop(self, x, y):
        nabla_b = [np.zeros(b.shape) for b in self.biases]
        nabla_w = [np.zeros(w.shape) for w in self.weights]
        activation = x
        activations = [x]
        zs = []
        for b, w in zip(self.biases, self.weights):
            z = np.dot(w, activation) + b
            zs.append(z)
            activation = sigmoid(z)
            activations.append(activation)
        
        delta = self.const_derivative(activations[-1], y) * \
            sigmoid_prime(zs[-1])
        nabla_b[-1] = delta
        nabla_w[-1] = np.dot(delta, activations[-2].transpose())

        for l in range(2, self.num_layers):
            z = zs[-l]
            sp = sigmoid_prime(z)
            delta = np.dot(self.weights[-l+1].transpose(), delta) * sp
            nabla_b[-l] = delta
            nabla_w[-l] = np.dot(delta, activations[-l-1].transpose())
        return (nabla_b, nabla_w)

    def evaluate(self, test_data):
        test_results = [(np.argmax(self.feedforward(x)), np.argmax(y))
                        for (x, y) in test_data]
        return sum(int(x == y) for (x, y) in test_results)

    def const_derivative(self, output_activations, y):
        """Return the vector of partial derivatives \\partial C_x
        \\partial a for the output activations."""
        return (output_activations-y)


def sigmoid(z):
    return 1.0/(1.0 + np.exp(-z))


def sigmoid_prime(z):
    """Derivative of the sigmoid functions"""
    return sigmoid(z)*(1 - sigmoid(z))


def main():
    import mnist
    training_data, test_data = mnist.load_data()
    net = Network([784, 30, 10])
    net.SGD(training_data, 30, 10, 3.0, test_data=test_data)


if __name__ == "__main__":
    main()