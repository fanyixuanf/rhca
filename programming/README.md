## 结构化编程(Structured Programming)
## 函数式编程(Functional Programming)
## 面向对象编程(OOP:Object-Oriented Programming)
## 面向切面编程(AOP:Aspect Oriented Programming)
- AOP可以说是OOP (Object-Oriented Programming,面向对象编程)的补充和完善
### 实现AOP的技术，主要分为两大类:
- 一是采用动态代理技术， 利用截取消息的方式，对该消息进行装饰，以取代原有对象行为的执行;
- 二是采用静态织入的方式，引入特定的语法创建“方面”，从而使得编译器可以在编译期间织入有关“方面”的代码;
### AOP的特性
- 可扩展性
- 可重用性
- 易理解性
- 易维护性
### AOP程序设计步骤
（1）将系统需求进行功能性分解，区分出普通的关注点与横切关注点，确定哪些功能是组件必须实现的，哪些可以以aspect的形式动态的加入到组件中.

（2）单独完成每一个关注点和编码实现。

（3）用联结器制定的重组规则，将组建代码与aspect代码进行组合，形成最终的系统。
