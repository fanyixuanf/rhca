# Graph

- 在计算机科学中, 图(graph) 是一种抽象数据类型, 旨在实现数学中的无向图和有向图概念，特别是图论领域。
- 一个图数据结构是一个(由有限个或者可变数量的)顶点/节点/点和边构成的有限集。

- 如果顶点对之间是无序的,称为无序图,否则称为有序图;
- 如果顶点对之间的边是有权重的,该图可称为加权图。
- 有向图：图G中的每条边都是有方向的，则称G为有向图（Digraph）。在有向图中，一条有向边是由两个顶点组成的有序对，有序对通常用尖括号表示。有向边也被称为弧，将边的始点称为弧尾，将终点称为弧头。
- 无向图：如果图中的每条边都是没有方向的，这种图被称为无向图。无向图中的边都是顶点的无序对，通常用圆括号来表示无序对。
- 顶点：通常将图中的数据元素称为顶点（Vertex），通常用V来表示顶点的集合。
- 完全图：如果无向图中的任意两个顶点之间都存在着一条边，则将此无向图称为无向完全图。如果有向图中的任意两个顶点之间都存在着方向相反的两条弧，则将此有向图称为有向完全图
- 稠密图与稀疏图：当一个图接近完全图时被称为稠密图，反之将含有较少的边数（即当e«n(n-1)）的图称为稀疏图。
- 权和网：图中的每一条边（弧）都可以有一个相关的数值，将这种与边相关的数值称为权。权能够表示从一个顶点到另一个顶点的距离或花费的代价。边上带有权的图称为带权图，也称作网。
- 子图：假设存在两个图G=(V，E)和G'=(V'，E')，如果V’是V的子集（即V' V），并且E’是E的子集（即E' E），则称G’是G的子图。
- 邻接点：在无向图G=(V，E)中，如果边(vi ，vj )∈E，则称顶点vi 和vj 互为邻接点（Adjacent）；边（vi ，vj ）依附于顶点vi 和vj ，即vi 和vj 相关联。
- 顶点的度：顶点的度是指和顶点相关联的边的数量。在有向图中，以顶点vi 为弧尾的弧的值称为顶点vi 的出度，以顶点vi 为弧头的弧的值称为顶点vi 的入度，顶点vi 的入度与出度的和是顶点vi 的度。
- 路径：如果图中存在一个从顶点vi 到顶点vj 的顶点序列，则这个顶点序列被称为路径。在图中有如下两种路径。 -（1）简单路径：指路径中的顶点不重复出现。 -（2）回路或环：是指如果路径中除第一个顶点和最后一个顶点相同以外，其余顶点不重复。一条路径上经过的边的数目称为路径长度

- 邻接矩阵表示法
  - 邻接矩阵表示顶点直接的相邻关系，推出邻接矩阵的目的是表示一种关系
    - 用一个一维数组存放顶点信息。 
    - 用一个二维数组表示n个顶点之间的关系
- 邻接表表示法
  - 邻接表是由邻接矩阵改造而来的一种链接结构，因为它只考虑非零元素，所以就节省了零元素所占的存储空间。邻接矩阵的每一行都有一个线性链接表，链接表的表头对应着邻接矩阵该行的顶点，链接表中的每个节点对应着邻接矩阵中该行的一个非零元素

1. 在邻接表中，每个线性链接表中各个节点的顺序是任意的。
2. 只使用邻接表中的各个线性链接表，不能说明它们顶点之间的邻接关系。
3. 在无向图中，某个顶点的度数等于该顶点对应的线性链接表的节点数；在有向图中，某个顶点的“出度”数等于该顶点对应的线性链表的节点数。

- 深度优先遍历(DFS)
  - 深度优先搜索的过程，是对每一个可能的分支路径深入到不能再深入为止的过程，并且每个节点只能访问一次。
  - 图的深度优先遍历类似于二叉树的深度优先遍历，其基本思想是：从图中某个顶点v出发，访问此顶点，然后从v的未被访问的邻接点出发深度优先遍历图，直至图中所有和v有路径相通的顶点都被访问到。显然，这是一个递归的搜索过程
- 广度优先遍历(BFS)
  - 图的广度优先遍历算法是一个分层遍历的过程，和二叉树的广度优先遍历类似，其基本思想在于：从图中的某一个顶点Vi触发，访问此顶点后，依次访问Vi的各个为层访问过的邻接点，然后分别从这些邻接点出发，直至图中所有顶点都被访问到。

- [code](Graph.go)