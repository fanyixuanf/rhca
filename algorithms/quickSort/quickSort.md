# QuickSort

- 快速排序是一种分而治之的算法。快速排序首先将一个大数组分成两个较小的子数组：比某个数小的元素和比某个数大的元素。然后快速排序可以递归地对子数组进行排序。

- 步骤:
  1. 从数组中选择一个元素，称为基点
  2. 分区：对数组重新排序，使所有值小于基点的元素都在它左边，而所有值大于基点的元素都在它右边（相等的值可以放在任何一边）。在此分区之后，基点处于其最终位置（左边和右边的中间位置）。这称为分区操作。
  3. 递归地将上述步骤应用于左边的数组和右边的数组。

- 性能
  - 时间复杂度: O(nlgn)

## 快速排序的核心算法思想是分治法:
- 分治法（Divide&Conquer），把一个大的问题，转化为若干个子问题（Divide），每个子问题“都”解决，大的问题便随之解决（Conquer）。这里的关键词是“都”。从伪代码里可以看到，快速排序递归时，先通过partition把数组分隔为两个部分，两个部分“都”要再次递归。
- 减治法（Reduce&Conquer），把一个大的问题，转化为若干个子问题（Reduce），这些子问题中“只”解决一个，大的问题便随之解决（Conquer）。这里的关键词是“只”。
  - 二分查找binary_search，BS，是一个典型的运用减治法思想的算法.

- [code](quickSort.go)