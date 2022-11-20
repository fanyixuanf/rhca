# 位运算

```go
&   bitwise AND
|   bitwise OR
^   bitwise XOR
&^   AND NOT
<<   left shift
>>   right shift
```

- The `&` Operator
  - ```go
      Given operands a, b
      AND(a, b) = 1; only if a = b = 1
                     else = 0
      ```

  - `&`运算符具有选择性的把整型数据的位清除为 0 的好的效果。 例如，我们可以使用 & 运算符去清除（设置）最后 4 个最低有效位（LSB）全部为 0 

  - 你可以用 & 操作去测试一个数字是奇数还是偶数。原因是当一个数字的二进制的最低位是 1 的时候，那他就是奇数。我们可以用一个数字和 1 进行 & 操作，然后在和 1 做 AND 运算，如果的到的结果是 1 ，那么这个原始的数字就是奇数

- The `|` Operator
  - ```go
      Given operands a, b
      OR(a, b) = 1; when a = 1 or b = 1
                   else = 0
     ```
  - 我们可以利用按位或操作符为给定的整数有选择地设置单个位。例如，在如下示例中我们使用按位或将示例数（从低位到高位（MSB））中的第 3 ，第 7 和第 8 位置为 1 。
  - 在使用位掩码技术为给定的整型数字设置任意位时，或运算非常有用


- The `^` Operator
  - ```go
        Given operands a, b
        XOR(a, b) = 1; only if a != b
             else = 0
    ```
  - `异或`运算的这个特性可以用来把二进制位的一个值变成另外一个值。举个例子，给到一个 16 进制的值，我们可以使用以下代码切换前 8 位（从 MSB 开始）的值。
  - 可以利用 `异或`运算去比较两个数字的符号是否一样。当 (a ^ b) ≥ 0 （或相反符号的 (a ^ b) < 0 ）为 true 的时候，两个整数 a，b 具有相同的符号
  - `^` 作为取反位运算符 （非）

- The `&^` Operator
  - ```go
        Given operands a, b
        AND_NOT(a, b) = AND(a, NOT(b))
    ```
  - 如果第二个操作数为 1 那么它则具有清除第一个操作数中的位的趣味特性

- The `<<` and `>>` Operators
  - ```go
        Given integer operands a and n,
         a << n; shifts all bits in a to the left n times
         a >> n; shifts all bits in a to the right n times
    ```
  - 可以利用左移和右移运算中，每次移动都表示一个数的 2 次幂这个特性，来作为某些乘法和除法运算的小技巧

- 当要位移的值（左操作数）是有符号值时，Go 自动应用算术位移。在右移操作期间，复制（或扩展）二进制补码符号位以填充位移的空隙。

- [Bit Hacking with Go](https://dev.to/vladimirvivien/bit-hacking-with-go)