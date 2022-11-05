# Algorithms:stuck_out_tongue_winking_eye:

- golang实现

### 数据结构:yum:
- 链表
- 双向链表
- [队列](queue/queue.md)
- [双端队列](deque/deque.md)
- 优先队列
- [栈](stack/stack.md)
- 堆
- 哈希表
- 字典树
- 树
  - 二叉查找树
  - AVL树
  - 红黑树
  - 线段树
  - 树状数组
- 图
- 并查集
- 布隆过滤器
- [切片vs数组(golang)](slice/slice_test.go)

### 算法主题
- 数学
  - 位运算 - set/get/update/clear 位、乘以/除以二进制位 、变负等
  - 阶乘
  - 斐波那契数 - 经典 和 闭式 版本
  - 素数检测 (排除法)
  - 欧几里得算法 - 计算最大公约数 (GCD)
  - 最小公倍数 (LCM)
  - 素数筛 - 查找任意给定范围内的所有素数
  - 判断 2 次方数 - 检查数字是否为 2 的幂 (原生和按位算法)
  - 杨辉三角形
  - 复数 - 复数及其基本运算
  - 弧度和角 - 弧度与角的相互转换
  - 快速算次方
  - 整数拆分
  - 割圆术 - 基于 N-gons 的近似 π 计算
  - 离散傅里叶变换 - 把时间信号解析成构成它的频率
- 集合
  - 笛卡尔积 - 多集合结果
  - 洗牌算法 - 随机置换有限序列
  - 幂集 - 该集合的所有子集
  - 排列 (有/无重复)
  - 组合 (有/无重复)
  - 最长公共子序列 (LCS)
  - 最长递增子序列
  - 最短公共父序列 (SCS)
  - 背包问题 - 0/1 和 无边界 问题
  - 最大子数列问题 - BF 算法 和 动态规划
  - 组合求和 - 查找形成特定总和的所有组合
- 字符串
  - 汉明距离 - 符号不同的位置数
  - 莱温斯坦距离 - 两个序列之间的最小编辑距离
  - Knuth–Morris–Pratt 算法 KMP 算法 - 子串搜索 (模式匹配)
  - 字符串快速查找 - 子串搜索 (模式匹配)
  - Rabin Karp 算法 - 子串搜索
  - 最长公共子串
  - 正则表达式匹配
- 搜索
  - 线性搜索
  - 跳转搜索/块搜索 - 搜索有序数组
  - 二分查找 - 搜索有序数组
  - 插值搜索 - 搜索均匀分布的有序数组
- 排序
  -  冒泡排序
  -  选择排序
  -  插入排序
  -  堆排序
  -  归并排序
  -  快速排序 - in-place (原地) 和 non-in-place 版本
  -  希尔排序
  -  计数排序
  -  基数排序
- 链表
  - 正向遍历
  - 反向遍历
- 树
  - 深度优先搜索 (DFS)
  - 广度优先搜索 (BFS)
- 图
  - 深度优先搜索 (DFS)
  - 广度优先搜索 (BFS)
  - 克鲁斯克尔演算法 - 寻找加权无向图的最小生成树 (MST)
  - 戴克斯特拉算法 - 找到图中所有顶点的最短路径
  - 贝尔曼-福特算法 - 找到图中所有顶点的最短路径
  - 弗洛伊德算法 - 找到所有顶点对 之间的最短路径
  - 判圈算法 - 对于有向图和无向图 (基于 DFS 和不相交集的版本)
  - 普林演算法 - 寻找加权无向图的最小生成树 (MST)
  - 拓扑排序 - DFS 方法
  - 关节点 - Tarjan 算法 (基于 DFS)
  - 桥 - 基于 DFS 的算法
  - 欧拉回径与一笔画问题 - Fleury 的算法 - 一次访问每个边
  - 哈密顿图 - 恰好访问每个顶点一次
  - 强连通分量 - Kosaraju 算法
  - 旅行推销员问题 - 尽可能以最短的路线访问每个城市并返回原始城市
- 加密
  - 多项式 hash - 基于多项式的 rolling hash 函数
  - 机器学习
  - NanoNeuron -7个简单的JS函数，说明机器如何实际学习（向前/向后传播）
- 未分类
  - 汉诺塔
  - 旋转矩阵 - 原地算法
  - 跳跃游戏 - 回溯,、动态编程 (自上而下+自下而上) 和贪婪的例子
  - 独特(唯一) 路径 - 回溯、动态编程和基于 Pascal 三角形的例子
  - 雨水收集 - 诱捕雨水问题 (动态编程和暴力版本)
  - 递归楼梯 - 计算有共有多少种方法可以到达顶层 (4 种解题方案)
  - 八皇后问题
  - 骑士巡逻

### 算法范式:dizzy_face:

- BF算法 - `查找/搜索` 所有可能性并选择最佳解决方案
  - 线性搜索
  - 雨水收集 - 诱导雨水问题
  - 地柜楼梯 - 计算有共有多少种方法可以到达顶层 (4 种解题方案)
  - 最大子数列
  - 履行推销员 - 尽可能以最短的路线访问每个城市并返回原始城市
  - 离散傅里叶变换 - 把时间信号解析成构成它的频率
- 贪心法 - 在当前选择最佳选项，不考虑以后情况
  - 跳跃游戏
  - 背包问题
  - 戴克斯特拉算法 - 找到所有图顶点的最短路径
  - 普里姆算法 - 寻找加权无向图的最小生成树 (MST)
  - 克鲁斯卡尔算法 - 寻找加权无向图的最小生成树 (MST)
- 分治法 - 将问题分成较小的部分，然后解决这些部分
  - 二分查找
  - 汉诺塔
  - 杨辉三角形
  - 欧几里得算法 - 计算最大公约数 (GCD)
  - 归并排序
  - 快速排序
  - 树深度优先搜索 (DFS)
  - 图深度优先搜索 (DFS)
  - 跳跃游戏
  - 快速算次方
  - 排列 (有/无重复)
  - 组合 (有/无重复)
- 动态规划(Dynamic programming) - 使用以前找到的子解决方案构建解决方案
  - 斐波那契数
  - 跳跃游戏
  - 独特路径
  - 雨水收集 - 疏导雨水问题
  - 递归楼梯 - 计算有共有多少种方法可以到达顶层 (4 种解题方案)
  - 莱温斯坦距离 - 两个序列之间的最小编辑距离
  - 最长公共子序列 (LCS)
  - 最长公共子串
  - 最长递增子序列
  - 最短公共子序列
  - 0-1背包问题
  - 整数拆分
  - 最大子数列
  - 贝尔曼-福特算法 - 找到所有图顶点的最短路径
  - 弗洛伊德算法 - 找到所有顶点对之间的最短路径
  - 正则表达式匹配
- 回溯法 - 类似于 `BF 算法` 试图产生所有可能的解决方案，但每次生成解决方案测试如果它满足所有条件，那么只有继续生成后续解决方案。否则回溯并继续寻找不同路径的解决方案。
  - 跳跃游戏
  - 独特路径
  - 幂集 - 该集合的所有子集
  - 哈密顿图 - 恰好访问每个顶点一次
  - 八皇后问题
  - 骑士巡逻
  - 组合求和 - 从规定的总和中找出所有的组合