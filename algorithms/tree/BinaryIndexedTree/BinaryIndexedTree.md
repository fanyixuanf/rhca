# Binary Indexed Tree

- 树状数组或二元索引树（英语：Binary Indexed Tree），其初衷是解决数据压缩里的累积频率（Cumulative Frequency）的计算问题，现多用于高效计算数列的前缀和， 区间和。它可以以O(log n)的时间得到任意前缀和sum _{i=1}^{j}A[i],1<=j<=N，并同时支持在O(log n)时间内支持动态单点值的修改。空间复杂度O(n)。

- [code](BinaryIndexedTree.go)