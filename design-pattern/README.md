### 设计模式的六大原则👀️

#### 目的： 高内聚，低耦合

- 单一职责(Single Responsibility Principle)：类的职责单一，对外只提供一种功能，而引起类变化的原因都应该只有一个。
- 开闭原则(Open Close Principle)：对扩展开放,对修改关闭，多使用抽象类和接口。(类的改动是通过**增加**代码进行的， 而不是修改源代码)
- 里氏替换原则(Liskov Substitution Principle)：任何抽象类（interface接口）出现的地方都可以用他的实现类进行替换，实际就是虚拟机制，语言级别实现面向对象功能。
- 依赖倒转原则(Dependence Inversion Principle)：要依赖于抽象(接口),不要依赖于具体(类)，针对接口编程,不针对实现编程。
- 接口隔离原则(Interface Segregation Principle)：使用多个隔离的接口,比使用单个接口好，建立最小的接口。
- 迪米特法则(Demeter Principle)：一个软件实体应当尽可能少地与其他实体发生相互作用，通过中间类建立联系。
- 合成复用原则(Composite Reuse Principle)：如果使用继承，会导致父类的任何变换都可能影响到子类的行为。如果使用对象组合，就降低了这种依赖关系。对于继承和组合，优先使用组合。

### [创建型🚀️](Creational_Pattern/README.md)
- 简单工厂
- 工厂方法
- 抽象工厂
- 单例
- 原型模式
- 建造者模式

### [结构型🚀️](Structural_Pattern/README.md)
- 适配器模式
- 桥接模式
- 组合模式
- 装饰器模式
- 外观模式
- 享元模式
- 代理模式

### [行为型🚀️](Behavioral_Pattern/README.md)
- 职责链
- 命令
- 解释器
- 迭代器
- 中介者
- 备忘录
- 观察者
- 策略
- 状态
- 模板方法
- 访问者

### 总结🎉️

<escape>
<table>
  <tr>
    <th>类型</th>
    <th>模式</th>
    <th>En</th>
    <th>说明</th>
    <th>关键字</th>
    <th>应用范围</th>
  </tr>
  <tr>
    <td rowspan="6">创建型</td>
    <td>简单工厂模式</td>
    <td>Simple Factory</td>
    <td>(不属于GoF)通过专门定义一个类来负责创建其他类的实例，被创建的实例通常都具有共同的父类</td>
    <td >---</td>
    <td >---</td>
  </tr>
  <tr>
    <td>工厂方法模式</td>
    <td >Factory Method</td>
    <td>定义了创建对象的接口,允许子类决定实例化哪个类</td>
    <td>动态生产对象</td>
    <td>类模式,对象模式</td>
  </tr>
  <tr>
    <td>抽象工厂模式</td>
    <td >Abstract Factory</td>
    <td>提供一个接口,可以创建一系列相关或互相依赖的对象,而无需指定它们具体的类</td>
    <td>生成成系列对象</td>
    <td>对象模式</td>
  </tr>
  <tr>
      <td>单例模式</td>
      <td >Singleton</td>
      <td>保证一个类仅有一个实例，并提供一个访问它的全局访问点。</td>
      <td>单实例</td>
      <td>对象模式</td>
  </tr>
  <tr>
      <td>原型模式</td>
      <td >Prototype</td>
      <td>允许对象在不了解要创建对象的确切类以及如何创建等细节的情况下创建自定义对象,通过拷贝原型对象来创建新的对象</td>
      <td>克隆对象</td>
      <td>对象模式</td>
  </tr>
  <tr>
      <td>建造者模式</td>
      <td >Builder</td>
      <td>将一个复杂类的表示与构造相分离, 使得相同的构件过程能够得到不同的表示</td>
      <td>复杂对象构造</td>
      <td>对象模式</td>
  </tr>
  <tr>
    <td rowspan="7">结构型</td>
    <td >适配器模式</td>
    <td >Adapter</td>
    <td >将一个类的接口转换成用户希望得到的另一个接口, 使原本不相关的接口得以协同工作</td>
    <td >转换接口</td>
    <td >类模式, 对象模式</td>
  </tr>
  <tr>
    <td>桥接模式</td>
    <td >Bridge</td>
    <td>将一个复杂的组件拆分成两个独立的但又相关的继承层次结构, 将类的抽象部分和实现部分分离开, 使它们可以独立的变化</td>
    <td>继承树拆分</td>
    <td>对象模式</td>
  </tr>
  <tr>
    <td>组合模式</td>
    <td >Composite</td>
    <td>创建树型层次结构来改变复杂性, 同时允许结构中的每一个元素操作同一个接口, 用来表示"整体-部分"的层次结构</td>
    <td>树形目录结构</td>
    <td>对象模式</td>
  </tr>
  <tr>
    <td>装饰器模式</td>
    <td >Decorator</td>
    <td>在不修改对象外观和功能的情况下添加或者删除对象功能, 即动态的给一个对象添加一些额外的职责</td>
    <td>动态附加职责</td>
    <td>对象模式</td>
  </tr>
  <tr>
    <td>外观模式</td>
    <td>Facade</td>
    <td>子系统中的一组接口提供了一个统一的接口</td>
    <td>对外统一接口</td>
    <td>对象模式</td>
  </tr>
  <tr>
    <td>享元模式</td>
    <td >Flyweight</td>
    <td>可以通过共享对象减少系统中低等级的,详细的对象数目, 提供支持大量细粒度对象共享的有效方法</td>
    <td>汉子编码</td>
    <td>对象模式</td>
  </tr>
  <tr>
    <td>代理模式</td>
    <td >Proxy</td>
    <td>为控制对初始对象的访问提供了一个代理或者占位符对象</td>
    <td>快捷方式</td>
    <td>对象模式</td>
  </tr>
  <tr>
    <td rowspan="11">行为型</td>
    <td>解释器模式</td>
    <td>Interpreter</td>
    <td>可以解释定义其语法表示的语言, 还提供了用表示来解释语言中的语句的解释器</td>
    <td>虚拟机的机制</td>
    <td>类模式, 对象模式</td>
  </tr>
  <tr>
    <td>模板方法模式</td>
    <td >Template Method</td>
    <td>提供了在不重写方法的前提下允许子类重载部分方法的方法</td>
    <td>框架</td>
    <td>类模式, 对象模式</td>
  </tr>
  <tr>
    <td>职责链模式</td>
    <td >Chain of Responsibility</td>
    <td>可以在系统中建立一个链, 这样消息就可以在首先接收到它的级别处被处理, 或者可以定位到可以处理它的对象</td>
    <td>传递职责</td>
    <td>对象模式</td>
  </tr>
  <tr>
    <td>命令模式</td>
    <td >Command</td>
    <td>在对象中封装了请求, 这样就可以保存命令, 将该命令传递给方法以及像任何其他对象一样返回该命令</td>
    <td>日志记录, 可撤销</td>
    <td>对象模式</td>
  </tr>
  <tr>
    <td>迭代器模式</td>
    <td >Iterator</td>
    <td>为集合中的有序访问提供了一致的方法, 而该集合是独立于基础集合, 并与之相分离</td>
    <td>数据集</td>
    <td>对象模式</td>
  </tr>
  <tr>
    <td>中介者模式</td>
    <td >Mediator</td>
    <td>通过引入一个能够管理对象间消息分布的对象, 简化了系统中对象间的通信</td>
    <td>不直接引用</td>
    <td>对象模式</td>
  </tr>
  <tr>
    <td>备忘录模式</td>
    <td >Memento</td>
    <td>保持对象状态的快照(snapshot), 这样对象可以在不向外界公开其内容的情况下返回到它最初状态</td>
    <td>游戏存档</td>
    <td>对象模式</td>
  </tr>
  <tr>
    <td>观察者模式</td>
    <td >Observer</td>
    <td>为组件向相关接收方广播消息提供了灵活的方法, 定义对象间的一种一对多的依赖关系</td>
    <td>订阅,广播,联动</td>
    <td>对象模式</td>
  </tr>
  <tr>
    <td>状态模式</td>
    <td >State</td>
    <td>允许一个对象在其内部改变状态时改变它的行为</td>
    <td>状态变成类</td>
    <td>对象模式</td>
  </tr>
  <tr>
    <td>策略模式</td>
    <td >Strategy</td>
    <td>定义一系列算法, 把它们一一封装起来,并且可以使它们之间可以相互替换, 从而让算法可以独立于使用它的用户变化</td>
    <td>多方案切换</td>
    <td>对象模式</td>
  </tr>
  <tr>
    <td>访问者模式</td>
    <td >Visitor</td>
    <td>提供了一种方便, 可维护的方法来表示在对象结构元素要进行的操作</td>
    <td>数据与操作分离</td>
    <td>对象模式</td>
  </tr>
</table>
</escape>

### 设计模式之间的关系
![设计模式之间的关系](/images/guanxi.svg)