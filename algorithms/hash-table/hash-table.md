# 哈希表hash table
- hash map
- 散列表是实现字典操作的一种有效数据结构, 是普通数组概念的推广.
- 原理
  - 将多个键/值(key/value)对分散存储在buckets(桶)中. 给定一个键(key),哈希(hash)算法会计算出键值对存储的位置.

- go语言中的map又称哈希表,是使用极高的数据结构.

- 直接寻址表
  - 并不是把每个元素的关键字及其卫星数据都放在直接寻址表外部的一个对象中, 再由表中某个槽的指针指向该对象, 而是直接把该对象的卫星数据放在表的槽中, 从而节省空间.
- 散列函数
  - 除法散列法
    - h(k) = k mod m
  - 乘法散列法
    - h(k) = [m(kA mod 1)]
  - 全域散列
    - 任何一个特定的散列函数都可能出现将n个键全部散列到同一个槽中, 唯一有效的改进方法是随机地选择散列函数, 使之独立于要存储的关键字.
- 解决hash冲突的方法
  - 链接法
    - 把散列到同一槽中的所有元素都放在一个链表中
    - 单/双向链表, 双向链表删除快
  - 开放寻址法
    - 所有的元素都存放在散列表里,当查找某个元素时,要系统的检查所有的表项,直到找到所需的元素.
    - 不用指针,而是计算出要存储的槽序列
- 完全散列