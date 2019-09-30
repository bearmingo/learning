# Dynamic Progmming

$$
P(i,j)=
\begin{cases}
true, & \quad \text{if the substring } S_i...S_j \text{is a palindrome}\\
false, & \quad \text{otherwise.}
\end{cases}
$$

Therefore

$$
P(i, j)=(P(i+1, j-1) \text{ and } S_i == S_j)
$$

The base case are：

$$
P(i, i) = true
P(i, i+1) = (S_i == S_{i+1})
$$

## Complexity Analysis
The complexity: $O(n^2)$

Space complexity: $O(n^2)$