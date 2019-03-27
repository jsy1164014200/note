



![1548304921509](/home/jsy/.config/Typora/typora-user-images/1548304921509.png)

# 基本的css

## 基础

### 选择器：选择页面上的某一个元素。

- 标签选择器：p ...

- 类选择器： .类名

- id选择器:   #id名

- 后代选择器

    选择器 选择器 选择器 ｛

   	属性名1：属性值1;

   	属性名2：属性值2;

   ｝

-  子元素选择器

  选择器>选择器｛

   	属性名1：属性值1;

   	属性名2：属性值2;

   		.......
   ｝

- 并集选择器

   选择器 ，选择器｛

   属性名1：属性值1;

   属性名2：属性值2;

   ｝

- 交集选择器

   选择器选择器 {
   属性名1：属性值1;

   属性名2：属性值2;

   }

### css的属性

1. 继承

    作用：给父元素设置一个属性，然后子元素可以使用

    应用：如果页面上有很多的文件都是红色，并且大小都是20px，那么这个时候给每个元素单独设置会很麻烦，所以可以考虑继承。

    注意：将来在写代码的时候，我们的css初始化，与页面文字的整体颜色一般会先设置。

    什么样的属性才可以继承呢？

    凡是以line-,text-,font-开头的属性都是可以继承。

   特殊点：

   - a标签的颜色是不会从父元素继承过来，如果要设置a标签的颜色，必须要直接指定。
   - h标签的字体的大小不会从父元素继承过来，如果要设置直接指定就行了。
   -  如果div不设置宽高，那么这个div的宽从父元素直接继承，高度为0；div的宽是可以继承的，高是不可以继承的。 让div居中：margin: 0 auto;

2. 层叠

    是页面处理冲突属性的一个能力。

    如果多个选择器为同一个元素设置了不同的属性它们会同时作用于这个元素。

    如果多个选择器为同一个元素调协了相同的属性它们会发生层叠。

    层叠的最终结果跟优先级有关系 

   - 优先级从大到小

   > !important>行内样式>id>class>标签>通配符>继承>默认
   >
   > 注意：！Important属性是不能继承的

   - 权重： 

     如何计算权重：

     数标签：先数id,如果id相等再数类如果id不相等id多的选择器权重高，权重越高，优先级越高。如果id选择器数量相同，再数类选择器，最后数标签。

   > 注意：比较权重的时候一定要注意：我们的选择器一定要命中对应的标签才可能让标签拥有对应的属性。



### 元素的显示方式

1. 块级元素：

 特点：单独占一行，可以给这个元素设置宽高。

 display(显示方式):block;

 所属标签：div,p,h1,ul,li,ol,dl

1. 行内元素：

 特点：可以多个标签放在同一行，但是给标签设置宽高没有作用。

 display(显示方式):inline;

 所属标签：

 span,b,u,i,s,ins,del,strong,em

1. 行内块级元素：

 特点：可以多个标签放在同一行，并且可以标签设置宽高。

 display(显示方式):inline-block;

 img

### 背景

1. Background-color:背景颜色;

 作用：设置背景颜色。

1. Background-image:背景图片

 作用：设置背景图片，默认情况下是平铺（如果图片不能占满整个容器，所以会让图片铺满整个容器）

1. Background-repeat:是否平铺

   no-repeat,repeat-x,repeat-y

2. background-position

    top :上   bottom:下  left:左  right: 右   center:中

    两类：

    水平方向：left：左   center:中   right: 右

    垂直方向：top:  上  center: 中   bottom:下

   > 连写 ：background : basckground-color background-image  background-repeat background-postion;

### 字体样式相关的属性：

 Font:Font-style font-weight font-size font-family;

### 文本的相关属性

1. text-align:

   作用：用于设置文本的排序方式. 取值：center:居中（在父容器居中）left:靠左 right:靠右

2. text-decoration:

   作用：用于将a标签的下划线去掉。text-decoration:none;

3. text-indent 

   作用 ：设置缩进text-indent : 2em;

4. line-height 行高，设置为 容器的高度时，居中显示![选区_086](/home/jsy/%E5%9B%BE%E7%89%87/%E9%80%89%E5%8C%BA_086.png)

单位：

1. px像素
2. % 百分比
3. em  一行em的大小 相当于 是当前标签中的font-size的大小

### 伪类

 伪类选择器也是给对应的标签设置样式，伪类选择器的样式要显示出来，那么这个标签必须具体对应的条件。

1. a:link  设置元素没有被访问过的样式。link:没有被访问过的，如果已经点击了一次，那么将来这个伪类中的样式将不会再被显示。
2. a: visited 设置元素被访问过的样式.如果这个元素已经被点了一次，那么将来这个伪类中的样式会一直显示。
3. a:honver 设置鼠标悬停时元素的样式
4. a: active 设置鼠标点击时的样式

> 注意 ：起作用的 大小lv ----->ha



## 盒子模型

![1548307258519](/home/jsy/.config/Typora/typora-user-images/1548307258519.png)

### 外壳 - border

1. border-width
2. border-style  solid等
3. border-color

> border:border-width border-style border-color.

![1548307546399](/home/jsy/.config/Typora/typora-user-images/1548307546399.png)

### 填充泡沫padding

 padding: 50px 40px 30px 20px;会将盒子的上右下左的顺序分别给盒子设置填充（padding）为50（上），40（右），30（下），20（左）按照顺时针方式来。

![1548307640949](/home/jsy/.config/Typora/typora-user-images/1548307640949.png)

### margin 

跟上面的一样

> margin
>
> margin-top
>
> margin-right
>
> margin-bottom
>
> margin-left

如果两个盒子是并列关系，给上面盒子设置margin-bottom,给下面的盒子设置margin-top，它们两个的值不会相加，而是发生了合并现象：

 margin-bottom:50;

 margin-top:25;

 margin-bottom+ margin-top=50;

 由于是合并现象，所以将来在取值的时候会取得两者之间比较大的数。



如果两个盒子是包含关系，如果让子盒子在父盒子之内向下平移100px：（margin塌陷现象）

 如果要实现：

1. 给父盒子设置padding(麻烦，给父盒子设置了padding之后将来如果要父盒子的大小保持不变，还必须把padding值减掉。)
2. 给子盒子设置margin-top(这里有一个bug，如果父盒子没有边框，那么将来给子盒子设置以后父盒子也会随着子盒子一起向下掉)
3. 给父盒子设置边框
4. 给父盒子设置属性：overflow(溢出)：hidden(隐藏)

## 浮动

> float: left
>
> float: right

1. 标准流 

   标准流就是浏览器在解析标签时候默认遵守的一些约定，说白就是浏览器解析标签规则，也叫做文档流. 一般情况下（如果不浮动的情况），所有的标签都是标准流，也就是说浏览器在解析的时候不会按其它特殊情况来解析，而浮动就是一种**脱离了标准流**的技术。

2. 浮动排序规则

   浮动找浮动，不浮动找不浮动，浮动会覆盖在 标准流上面

3. 浮动会将元素的显示方式改为：inline-block;

清除浮动

 使用当前最主流的清除方式：使用伪元素来清除浮动

```css
.clearfix:after {
 	content:””;
    height: 0;
    line-height: 0;
    display: block;
    visibility:hidden;
    clear:both;
}
.clearfix {
 	zoom: 1;/*用来兼容ie浏览器*/
}
双伪元素法
.clearfix:after , .clearfix:before {
 		content:””;
    display: table;
    clear:both;
}
.clearfix {
   	zoom: 1;
}
```

## 定位

1. 静态定位：

   position:static.所有标准流都是这个定位，这个属性一般情况下是不会单独使用，一般情况下是通过js来动态给元素清除定位。

2. 绝对定位 可以将一个盒子移动到页面的任意位置。

    作用：可以将一个盒子移动到页面的任意位置。

    步骤：

   1. 将需要移动的元素的postion属性设置为absolute.

   2. 将需要移动的元素的位置的坐标记录下来，并通过关键字来设置。

      关键字： top:设置元素与浏览器顶部的距离left：设置元素与浏览器左边的距离 right：设置元素与浏览器右边的距离 bottom：设置元素与浏览器底边的距离。

      注意点：

    a) 如果盒子没有父元素，那么将来在定位的时候，我们trbl是相对于body元素的。

    b) 如果定位的盒子有父元素，但是父元素没有定位，那么这个子元素定位的trbl还是相对于body.

    c )  如果定位的盒子有父元素，并且父元素有定位，那么这个元素的定位是相对于它的父元素（子元素定位的trbl是相对于父元素。）。

    d ) **绝对定位的盒子不占页面上的位置（脱离标准流）**

    e ) 如果没有设置top，left那么这个元素就算设置了有position属性位置也不会改变。

3. 相对定位 position: relative

   1. 如果设置了相对定位但没有设置值，我们的盒子的位置不会改变。
   2. 相对定位是占据标准流的位置。

> 总结：如果将来在页面上使用定位，绝对要遵守这个规律：子绝父相的规律（子元素使用绝对定位，父元素使用相对定位。）

## overflow的使用

 hidden:将超出的内容隐藏起来;

 scroll:将超出的部分放在滚动条之内;

 auto:如果超出放在滚动条之内，如果没有超出没有滚动条。

## 关于页面元素的显示和隐藏

 overflow: hidden;//将超出的部分隐藏。

 dispaly :none;//将元素直接隐藏，并且不会在页面上占据位置

 (可见性)visibility:hidden;//将元素直接隐藏，但是会保留元素所在的位置

## 图片和文字显示在一行上

 如果文字和图片显示在一行中，那么默认情况下图片底部跟文字的基线对齐。

 如果文字和图片在一行，给图片设置margin-top，文字也会跟随图片一起下移。

 如果要让文字的中线和图片的中线设置对齐，使用vertical-align:middel;

 vertical-align与dispaly:inline-block使用是最好的。

> z-index : 2  用来设置覆盖的层次

































# ajax

![协商缓存流程图](/home/jsy/desktop/%E7%AC%94%E8%AE%B0/html/%E5%8D%8F%E5%95%86%E7%BC%93%E5%AD%98%E6%B5%81%E7%A8%8B%E5%9B%BE.png)

浏览器对象：XMLHttpRequest

```js
let xhr = new XMLHttpRequest()
xhr.open('post','index.php')
xhr.setRequestHeader('Content-Type','text/html')
xhr.send('name=jsy&age=10')
xhr.onreadystatechange = () => {
    if(xhr.readyState == 4 && xhr.status == 200){
        let result = document.querSelector('.result')
        result.innerHTML = xhr.responseText
    }
}
```

api

```js
xhr.open()
xhr.setRequestHeader()
xhr.send()

xhr.onreadystatechange
xhr.readyState
xhr.responseText
xhr.responseXML
xhr.status
xhr.statusText
xhr.getResponseHeader()
xhr.getAllResponseHeaders()
```

js解析xml，json （有内建对象）

> eavl()、JSON对象 JSON.parse()、JSON.stringify()；

1. 跨域的分类：

![1545967574859](/home/jsy/.config/Typora/typora-user-images/1545967574859.png)

解决方案：

- document.domain + iframe
- window.name + iframe
- location.hash + iframe
- window.postMessage()
- JSONP