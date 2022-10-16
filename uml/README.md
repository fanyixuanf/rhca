### OOA-UML图

#### 数据流图(DFD)

#### 程序流程图,盒图 

#### 静态图(结构图)
- 类图
  - 一组类,接口,写作和它们之间的关系
- 对象图
  - 一组对象及它们之间的关系
- 构件图
  - 一个封装的类和它的接口
- 部署图
  - 软硬件之间映射
- 制品图
  - 系统的物理结构
- 包图
  - 由模型本身分解而成的组织单元,以及它们之间的依赖关系
- 组合结构图

#### 动态图(行为图)
- 用例图
  - 系统与外部参与者的交互
- 顺序图
  - 强调按时间顺序
- 通信图
  - 协作图
- 状态图
  - 状态转换变迁
- 活动图
  - 类似程序流程图,并行行为
- 定时图
  - 强调实际时间
- 交互概览图

#### 4+1视图
1. 逻辑视图
   1. 系统分析,设计人员
   2. 类与对象
2. 实现视图
   1. 程序员
   2. 物理代码文件和组件
3. 部署视图
   1. 系统和网络工程师
   2. 软件到硬件的映射
4. 进程视图
   1. 系统集成人员
   2. 线程, 进程, 并发
5. 用例视图
   1. 最终用户
   2. 需求分析模型

#### 包含, 扩展, 泛化(箭头指向共性部分)
- 包含关系
    - ![img.png](img/baohan.png)
    - 公共部分,每次都用到
- 扩展关系
  - ![img.png](img/kuozhan.png)
  - 有时用,有时不用
- 泛化关系
  - ![img.png](img/fanhua-yongli.png)
  - 父子关系

- 案例
- ![img.png](img/guanxitu.png)
  - 车的类图结构为abstract, 表示车是一个抽象类;
  - 它有两个继承类: 汽车和自行车; 它们之间的关系为实现关系;
  - 汽车与SUV之间也是继承关系，它们之间的关系为泛化关系;
  - 汽车与发动机之间是聚合关系;
  - 汽车与轮胎之间是聚合关系;
  - 学生与班级是聚合关系;
  - 班级与学校是组合关系(公司与部门);
  - 学生与身份证之间为关联关系;
  - 学生上学用自行车，与自行车是一种依赖关系;

### 类与对象
- 依赖关系
  - ![img.png](img/yilai.png)
  - 一个事物发生变化影响另一个事物
- 泛化关系
  - ![img.png](img/fanhua.png)
  - 特殊(子)/一般(父)关系
- 关联关系
  - ![img.png](img/img.png)
  - 描述一组链, 链是对象之间的链接
- 聚合关系
  - ![img.png](img/juhe.png)
  - 整体与部分生命周期不同(汽车/轮子)
- 组合关系
  - ![img.png](img/zuhe.png)
  - 整体与部分生命周期相同(公司/部门)
- 实现关系
  - ![img.png](img/shixian.png)
  - 接口与类的关系

<div class="mxgraph" style="max-width:100%;border:1px solid transparent;" data-mxgraph="{&quot;highlight&quot;:&quot;#0000ff&quot;,&quot;nav&quot;:true,&quot;resize&quot;:true,&quot;toolbar&quot;:&quot;zoom layers tags lightbox&quot;,&quot;edit&quot;:&quot;_blank&quot;,&quot;xml&quot;:&quot;&lt;mxfile host=\&quot;drawio-plugin\&quot; modified=\&quot;2022-10-16T08:15:12.282Z\&quot; agent=\&quot;5.0 (Windows NT 6.2; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36\&quot; etag=\&quot;poQFJLkiXPpUoP8-I5Xi\&quot; version=\&quot;15.5.4\&quot; type=\&quot;embed\&quot;&gt;&lt;diagram id=\&quot;1czklNrwi9EsNagj-m8T\&quot; name=\&quot;第 1 页\&quot;&gt;&lt;mxGraphModel dx=\&quot;1828\&quot; dy=\&quot;809\&quot; grid=\&quot;1\&quot; gridSize=\&quot;10\&quot; guides=\&quot;1\&quot; tooltips=\&quot;1\&quot; connect=\&quot;1\&quot; arrows=\&quot;1\&quot; fold=\&quot;1\&quot; page=\&quot;1\&quot; pageScale=\&quot;1\&quot; pageWidth=\&quot;827\&quot; pageHeight=\&quot;1169\&quot; background=\&quot;#114B5F\&quot; math=\&quot;0\&quot; shadow=\&quot;0\&quot;&gt;&lt;root&gt;&lt;mxCell id=\&quot;0\&quot;/&gt;&lt;mxCell id=\&quot;1\&quot; parent=\&quot;0\&quot;/&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-30\&quot; value=\&quot;&amp;lt;span style=&amp;quot;color: rgb(64, 64, 64); font-family: Lato, proxima-nova, &amp;amp;quot;Helvetica Neue&amp;amp;quot;, Arial, sans-serif; font-size: 16px; font-weight: 400; text-align: left; background-color: rgb(252, 252, 252);&amp;quot;&amp;gt;&amp;amp;lt;&amp;amp;lt;abstract&amp;amp;gt;&amp;amp;gt; 车&amp;lt;/span&amp;gt;\&quot; style=\&quot;swimlane;fontStyle=1;align=center;verticalAlign=middle;childLayout=stackLayout;horizontal=1;startSize=29;horizontalStack=0;resizeParent=1;resizeParentMax=0;resizeLast=0;collapsible=0;marginBottom=0;html=1;fillColor=#F45B69;strokeColor=#028090;fontColor=#E4FDE1;\&quot; parent=\&quot;1\&quot; vertex=\&quot;1\&quot;&gt;&lt;mxGeometry x=\&quot;214\&quot; y=\&quot;170\&quot; width=\&quot;349\&quot; height=\&quot;60\&quot; as=\&quot;geometry\&quot;/&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-33\&quot; value=\&quot;汽车\&quot; style=\&quot;swimlane;fontStyle=1;align=center;verticalAlign=middle;childLayout=stackLayout;horizontal=1;startSize=29;horizontalStack=0;resizeParent=1;resizeParentMax=0;resizeLast=0;collapsible=0;marginBottom=0;html=1;fillColor=#F45B69;strokeColor=#028090;fontColor=#E4FDE1;\&quot; parent=\&quot;1\&quot; vertex=\&quot;1\&quot;&gt;&lt;mxGeometry x=\&quot;220\&quot; y=\&quot;290\&quot; width=\&quot;140\&quot; height=\&quot;54\&quot; as=\&quot;geometry\&quot;/&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-39\&quot; value=\&quot;\&quot; style=\&quot;endArrow=block;dashed=1;endFill=0;endSize=12;html=1;curved=1;strokeColor=#028090;fontColor=#E4FDE1;labelBackgroundColor=#114B5F;\&quot; parent=\&quot;AhttPOFJjhMUD6R0sSNN-33\&quot; edge=\&quot;1\&quot;&gt;&lt;mxGeometry width=\&quot;160\&quot; relative=\&quot;1\&quot; as=\&quot;geometry\&quot;&gt;&lt;mxPoint x=\&quot;60\&quot; as=\&quot;sourcePoint\&quot;/&gt;&lt;mxPoint x=\&quot;60\&quot; y=\&quot;-60\&quot; as=\&quot;targetPoint\&quot;/&gt;&lt;/mxGeometry&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-34\&quot; value=\&quot;自行车\&quot; style=\&quot;swimlane;fontStyle=1;align=center;verticalAlign=middle;childLayout=stackLayout;horizontal=1;startSize=29;horizontalStack=0;resizeParent=1;resizeParentMax=0;resizeLast=0;collapsible=0;marginBottom=0;html=1;fillColor=#F45B69;strokeColor=#028090;fontColor=#E4FDE1;\&quot; parent=\&quot;1\&quot; vertex=\&quot;1\&quot;&gt;&lt;mxGeometry x=\&quot;420\&quot; y=\&quot;290\&quot; width=\&quot;140\&quot; height=\&quot;54\&quot; as=\&quot;geometry\&quot;/&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-40\&quot; value=\&quot;\&quot; style=\&quot;endArrow=block;dashed=1;endFill=0;endSize=12;html=1;curved=1;strokeColor=#028090;fontColor=#E4FDE1;labelBackgroundColor=#114B5F;\&quot; parent=\&quot;AhttPOFJjhMUD6R0sSNN-34\&quot; edge=\&quot;1\&quot;&gt;&lt;mxGeometry width=\&quot;160\&quot; relative=\&quot;1\&quot; as=\&quot;geometry\&quot;&gt;&lt;mxPoint x=\&quot;70\&quot; as=\&quot;sourcePoint\&quot;/&gt;&lt;mxPoint x=\&quot;70\&quot; y=\&quot;-60\&quot; as=\&quot;targetPoint\&quot;/&gt;&lt;/mxGeometry&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-41\&quot; value=\&quot;学生\&quot; style=\&quot;swimlane;fontStyle=1;align=center;verticalAlign=middle;childLayout=stackLayout;horizontal=1;startSize=29;horizontalStack=0;resizeParent=1;resizeParentMax=0;resizeLast=0;collapsible=0;marginBottom=0;html=1;fillColor=#F45B69;strokeColor=#028090;fontColor=#E4FDE1;\&quot; parent=\&quot;1\&quot; vertex=\&quot;1\&quot;&gt;&lt;mxGeometry x=\&quot;420\&quot; y=\&quot;400\&quot; width=\&quot;140\&quot; height=\&quot;54\&quot; as=\&quot;geometry\&quot;/&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-43\&quot; value=\&quot;\&quot; style=\&quot;html=1;verticalAlign=bottom;labelBackgroundColor=none;endArrow=open;endFill=0;dashed=1;entryX=0.5;entryY=1;entryDx=0;entryDy=0;curved=1;strokeColor=#028090;fontColor=#E4FDE1;\&quot; parent=\&quot;AhttPOFJjhMUD6R0sSNN-41\&quot; target=\&quot;AhttPOFJjhMUD6R0sSNN-34\&quot; edge=\&quot;1\&quot;&gt;&lt;mxGeometry width=\&quot;160\&quot; relative=\&quot;1\&quot; as=\&quot;geometry\&quot;&gt;&lt;mxPoint x=\&quot;70\&quot; as=\&quot;sourcePoint\&quot;/&gt;&lt;mxPoint x=\&quot;230\&quot; as=\&quot;targetPoint\&quot;/&gt;&lt;/mxGeometry&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-44\&quot; value=\&quot;身份证\&quot; style=\&quot;swimlane;fontStyle=1;align=center;verticalAlign=middle;childLayout=stackLayout;horizontal=1;startSize=29;horizontalStack=0;resizeParent=1;resizeParentMax=0;resizeLast=0;collapsible=0;marginBottom=0;html=1;fillColor=#F45B69;strokeColor=#028090;fontColor=#E4FDE1;\&quot; parent=\&quot;1\&quot; vertex=\&quot;1\&quot;&gt;&lt;mxGeometry x=\&quot;660\&quot; y=\&quot;400\&quot; width=\&quot;140\&quot; height=\&quot;54\&quot; as=\&quot;geometry\&quot;/&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-47\&quot; value=\&quot;\&quot; style=\&quot;line;strokeWidth=1;fillColor=none;align=left;verticalAlign=middle;spacingTop=-1;spacingLeft=3;spacingRight=3;rotatable=0;labelPosition=right;points=[];portConstraint=eastwest;strokeColor=#028090;fontColor=#E4FDE1;labelBackgroundColor=#114B5F;\&quot; parent=\&quot;1\&quot; vertex=\&quot;1\&quot;&gt;&lt;mxGeometry x=\&quot;560\&quot; y=\&quot;430\&quot; width=\&quot;100\&quot; height=\&quot;8\&quot; as=\&quot;geometry\&quot;/&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-48\&quot; value=\&quot;班级\&quot; style=\&quot;swimlane;fontStyle=1;align=center;verticalAlign=middle;childLayout=stackLayout;horizontal=1;startSize=29;horizontalStack=0;resizeParent=1;resizeParentMax=0;resizeLast=0;collapsible=0;marginBottom=0;html=1;fillColor=#F45B69;strokeColor=#028090;fontColor=#E4FDE1;\&quot; parent=\&quot;1\&quot; vertex=\&quot;1\&quot;&gt;&lt;mxGeometry x=\&quot;420\&quot; y=\&quot;525\&quot; width=\&quot;140\&quot; height=\&quot;54\&quot; as=\&quot;geometry\&quot;/&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-51\&quot; value=\&quot;轮胎\&quot; style=\&quot;swimlane;fontStyle=1;align=center;verticalAlign=middle;childLayout=stackLayout;horizontal=1;startSize=29;horizontalStack=0;resizeParent=1;resizeParentMax=0;resizeLast=0;collapsible=0;marginBottom=0;html=1;fillColor=#F45B69;strokeColor=#028090;fontColor=#E4FDE1;\&quot; parent=\&quot;1\&quot; vertex=\&quot;1\&quot;&gt;&lt;mxGeometry x=\&quot;20\&quot; y=\&quot;400\&quot; width=\&quot;140\&quot; height=\&quot;54\&quot; as=\&quot;geometry\&quot;/&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-53\&quot; value=\&quot;发动机\&quot; style=\&quot;swimlane;fontStyle=1;align=center;verticalAlign=middle;childLayout=stackLayout;horizontal=1;startSize=29;horizontalStack=0;resizeParent=1;resizeParentMax=0;resizeLast=0;collapsible=0;marginBottom=0;html=1;fillColor=#F45B69;strokeColor=#028090;fontColor=#E4FDE1;\&quot; parent=\&quot;1\&quot; vertex=\&quot;1\&quot;&gt;&lt;mxGeometry x=\&quot;20\&quot; y=\&quot;490\&quot; width=\&quot;140\&quot; height=\&quot;54\&quot; as=\&quot;geometry\&quot;/&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-54\&quot; value=\&quot;SUV\&quot; style=\&quot;swimlane;fontStyle=1;align=center;verticalAlign=middle;childLayout=stackLayout;horizontal=1;startSize=29;horizontalStack=0;resizeParent=1;resizeParentMax=0;resizeLast=0;collapsible=0;marginBottom=0;html=1;fillColor=#F45B69;strokeColor=#028090;fontColor=#E4FDE1;\&quot; parent=\&quot;1\&quot; vertex=\&quot;1\&quot;&gt;&lt;mxGeometry x=\&quot;216\&quot; y=\&quot;586\&quot; width=\&quot;140\&quot; height=\&quot;54\&quot; as=\&quot;geometry\&quot;/&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-55\&quot; value=\&quot;\&quot; style=\&quot;endArrow=block;html=1;align=center;verticalAlign=bottom;endFill=0;labelBackgroundColor=none;endSize=8;curved=1;strokeColor=#028090;fontColor=#E4FDE1;\&quot; parent=\&quot;1\&quot; edge=\&quot;1\&quot;&gt;&lt;mxGeometry relative=\&quot;1\&quot; as=\&quot;geometry\&quot;&gt;&lt;mxPoint x=\&quot;280\&quot; y=\&quot;585\&quot; as=\&quot;sourcePoint\&quot;/&gt;&lt;mxPoint x=\&quot;280\&quot; y=\&quot;345\&quot; as=\&quot;targetPoint\&quot;/&gt;&lt;/mxGeometry&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-56\&quot; value=\&quot;\&quot; style=\&quot;resizable=0;html=1;align=center;verticalAlign=top;labelBackgroundColor=none;fillColor=#F45B69;strokeColor=#028090;fontColor=#E4FDE1;\&quot; parent=\&quot;AhttPOFJjhMUD6R0sSNN-55\&quot; connectable=\&quot;0\&quot; vertex=\&quot;1\&quot;&gt;&lt;mxGeometry relative=\&quot;1\&quot; as=\&quot;geometry\&quot;/&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-58\&quot; value=\&quot;\&quot; style=\&quot;endArrow=diamondThin;endFill=0;endSize=24;html=1;entryX=0;entryY=0.5;entryDx=0;entryDy=0;exitX=0.5;exitY=0;exitDx=0;exitDy=0;curved=1;strokeColor=#028090;fontColor=#E4FDE1;labelBackgroundColor=#114B5F;\&quot; parent=\&quot;1\&quot; source=\&quot;AhttPOFJjhMUD6R0sSNN-51\&quot; target=\&quot;AhttPOFJjhMUD6R0sSNN-33\&quot; edge=\&quot;1\&quot;&gt;&lt;mxGeometry width=\&quot;160\&quot; relative=\&quot;1\&quot; as=\&quot;geometry\&quot;&gt;&lt;mxPoint x=\&quot;80\&quot; y=\&quot;380\&quot; as=\&quot;sourcePoint\&quot;/&gt;&lt;mxPoint x=\&quot;240\&quot; y=\&quot;380\&quot; as=\&quot;targetPoint\&quot;/&gt;&lt;Array as=\&quot;points\&quot;&gt;&lt;mxPoint x=\&quot;90\&quot; y=\&quot;317\&quot;/&gt;&lt;/Array&gt;&lt;/mxGeometry&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-59\&quot; value=\&quot;\&quot; style=\&quot;endArrow=diamondThin;endFill=0;endSize=24;html=1;entryX=0.25;entryY=1;entryDx=0;entryDy=0;curved=1;strokeColor=#028090;fontColor=#E4FDE1;labelBackgroundColor=#114B5F;\&quot; parent=\&quot;1\&quot; target=\&quot;AhttPOFJjhMUD6R0sSNN-33\&quot; edge=\&quot;1\&quot;&gt;&lt;mxGeometry width=\&quot;160\&quot; relative=\&quot;1\&quot; as=\&quot;geometry\&quot;&gt;&lt;mxPoint x=\&quot;160\&quot; y=\&quot;520\&quot; as=\&quot;sourcePoint\&quot;/&gt;&lt;mxPoint x=\&quot;320\&quot; y=\&quot;520\&quot; as=\&quot;targetPoint\&quot;/&gt;&lt;Array as=\&quot;points\&quot;&gt;&lt;mxPoint x=\&quot;255\&quot; y=\&quot;520\&quot;/&gt;&lt;/Array&gt;&lt;/mxGeometry&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-60\&quot; value=\&quot;\&quot; style=\&quot;endArrow=diamondThin;endFill=0;endSize=24;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;curved=1;strokeColor=#028090;fontColor=#E4FDE1;labelBackgroundColor=#114B5F;\&quot; parent=\&quot;1\&quot; source=\&quot;AhttPOFJjhMUD6R0sSNN-41\&quot; target=\&quot;AhttPOFJjhMUD6R0sSNN-48\&quot; edge=\&quot;1\&quot;&gt;&lt;mxGeometry width=\&quot;160\&quot; relative=\&quot;1\&quot; as=\&quot;geometry\&quot;&gt;&lt;mxPoint x=\&quot;480\&quot; y=\&quot;480\&quot; as=\&quot;sourcePoint\&quot;/&gt;&lt;mxPoint x=\&quot;640\&quot; y=\&quot;480\&quot; as=\&quot;targetPoint\&quot;/&gt;&lt;/mxGeometry&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-61\&quot; value=\&quot;学校\&quot; style=\&quot;swimlane;fontStyle=1;align=center;verticalAlign=middle;childLayout=stackLayout;horizontal=1;startSize=29;horizontalStack=0;resizeParent=1;resizeParentMax=0;resizeLast=0;collapsible=0;marginBottom=0;html=1;fillColor=#F45B69;strokeColor=#028090;fontColor=#E4FDE1;\&quot; parent=\&quot;1\&quot; vertex=\&quot;1\&quot;&gt;&lt;mxGeometry x=\&quot;420\&quot; y=\&quot;630\&quot; width=\&quot;140\&quot; height=\&quot;54\&quot; as=\&quot;geometry\&quot;/&gt;&lt;/mxCell&gt;&lt;mxCell id=\&quot;AhttPOFJjhMUD6R0sSNN-62\&quot; value=\&quot;\&quot; style=\&quot;endArrow=diamondThin;endFill=1;endSize=24;html=1;entryX=0.429;entryY=0;entryDx=0;entryDy=0;entryPerimeter=0;curved=1;strokeColor=#028090;fontColor=#E4FDE1;labelBackgroundColor=#114B5F;\&quot; parent=\&quot;1\&quot; edge=\&quot;1\&quot;&gt;&lt;mxGeometry width=\&quot;160\&quot; relative=\&quot;1\&quot; as=\&quot;geometry\&quot;&gt;&lt;mxPoint x=\&quot;488\&quot; y=\&quot;580\&quot; as=\&quot;sourcePoint\&quot;/&gt;&lt;mxPoint x=\&quot;488.05999999999995\&quot; y=\&quot;630\&quot; as=\&quot;targetPoint\&quot;/&gt;&lt;/mxGeometry&gt;&lt;/mxCell&gt;&lt;/root&gt;&lt;/mxGraphModel&gt;&lt;/diagram&gt;&lt;/mxfile&gt;&quot;}"></div>
<script type="text/javascript" src="https://viewer.diagrams.net/js/viewer-static.min.js"></script>
