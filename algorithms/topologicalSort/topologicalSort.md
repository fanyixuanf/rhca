# Topological Sort

- 对于一些有前后依赖关系的排序算法，是利用有向无环图进行实现，通过局部依赖关系确定全局顺序的算法
- 应用:
  1. The canonical application of topological sorting is in scheduling a sequence of jobs or tasks based on their dependencies. The jobs are represented by vertices, and there is an edge from x to y if job x must be completed before job y can be started (for example, when washing clothes, the washing machine must finish before we put the clothes in the dryer). Then, a topological sort gives an order in which to perform the jobs. 
  2. Other application is dependency resolution. Each vertex is a package and each edge is a dependency of package a on package 'b'. Then topological sorting will provide a sequence of installing dependencies in a way that every next dependency has its dependent packages to be installed in prior.

## Kahn算法
- 利用贪心算法，如果两个顶点，顶点b依赖于顶点a,就将a指向b,当一个顶点的入度为零，将这个顶点就是最优排序点， 并且将顶点从图中移除，将可达顶点的入度减一


## DFS算法
1. 使用深度算法，产生逆向邻接表先输出其他依赖，最后输出自己。


- [code](topologicalSort.go)