# css3

## 选择器

### 1. 属性选择器

 1、E[attr] 表示存在attr属性即可；

 2、E[attr=val] 表示属性值完全等于val；

 3、E[attr*=val] 表示的属性值里包含val字符并且在“任意”位置；

 4、E[attr^=val] 表示的属性值里包含val字符并且在“开始”位置；

 5、E[attr$=val] 表示的属性值里包含val字符并且在“结束”位置；

### 2. 伪类选择器

 E:first-child第一个子元素

 E:last-child最后一个子元素

 E:nth-child(n) 第n个子元素，计算方法是E元素的全部兄弟元素；

 E:nth-last-child(n) 同E:nth-child(n) 相似，只是倒着计算；

 n遵循线性变化，其取值0、1、2、3、4、... 但是当n<=0时，选取无效。

 n可是多种形式：nth-child(2n)、nth-child(2n+1)、nth-child(-1n+5)等；

 E:empty 选中没有任何子节点的E元素；（使用不是非常广泛）

### 3. 伪元素选择器

 E::first-letter文本的第一个单词或字（如中文、日文、韩文等）；

 E::first-line 文本第一行；

 重点：E::before、E::after

## 颜色

 新增了RGBA、HSLA模式，其中的A 表示透明度通道，即可以设置颜色值的透明度，相较opacity，它们不具有继承性，即不会影响子元素的透明度。

 Red、Green、Blue、Alpha即RGBA

 Hue、Saturation、Lightness、Alpha即HSLA

 R、G、B 取值范围0~255

 H 色调 取值范围0~360，0/360表示红色、120表示绿色、240表示蓝色

 S 饱和度 取值范围0%~100%

 L 亮度 取值范围0%~100%

 A 透明度 取值范围0~1

 关于透明度：

 1、opacity只能针对整个盒子设置透明度，子盒子及内容会继承父盒子的透明度；

 2 、transparent 不可调节透明度，始终完全透明

 RGBA、HSLA可应用于所有使用颜色的地方。

## 文本

 text-shadow，可分别设置偏移量、模糊度、颜色（可设透明度）。

 1、水平偏移量 正值向右 负值向左；

 2、垂直偏移量 正值向下 负值向上；

 3、模糊度是不能为负值；

## 边框

 border-radius

 圆角处理时，脑中要形成圆、圆心、横轴、纵轴的概念，正圆是椭圆的一种特殊情况。

 可分别设置长、短半径，以“/”进行分隔，遵循“1，2，3，4”规则，“/”前面的1~4个用来设置横轴半径（分别对应横轴1、2、3、4位置 ），“/”后面1~4个参数用来设置纵轴半径（分别对应纵轴1、2、3、4位置 ）



 box-shadow

 1、水平偏移量 正值向右 负值向左；

 2、垂直偏移量 正值向下 负值向上；

 3、模糊度是不能为负值；

 4、inset可以设置内阴影；

 设置边框阴影不会改变盒子的大小，即不会影响其兄弟元素的布局。

 可以设置多重边框阴影，实现更好的效果，增强立体感。



 border-image

 设置的图片将会被“切割”成九宫格形式，然后进行设置。如下图

 ![1548311613400](/home/jsy/.config/Typora/typora-user-images/1548311613400.png)

 “切割”完成后生成虚拟的9块图形，然后按对应位置设置背景，

 其中四个角位置、形状保持不变，中心位置水平垂直两个方向平铺。如下图

![1548311629223](/home/jsy/.config/Typora/typora-user-images/1548311629223.png)

## 盒子模型

 CSS3中可以通过box-sizing 来指定盒模型，即可指定为content-box、border-box，这样我们计算盒子大小的方式就发生了改变。

 可以分成两种情况：

 1、box-sizing: border-box  计算方式为width = border + padding + content

 2、box-sizing: content-box  计算方式为width = content



## 背景

 **1**、*background-size*设置背景图片的尺寸

 cover会自动调整缩放比例，保证图片始终填充满背景区域，如有溢出部分则会被隐藏。

 contain会自动调整缩放比例，保证图片始终完整显示在背景区域。

 也可以使用长度单位或百分比  

 **2**、background-origin设置背景定位的原点

 border-box以边框做为参考原点；

 padding-box以内边距做为参考原点；

 content-box以内容区做为参考点；

 **3**、background-clip设置背景区域裁切

 border-box裁切边框以内为背景区域；

 padding-box裁切内边距以内为背景区域；

 content-box裁切内容区做为背景区域；

 **4**、以逗号分隔可以设置多背景，可用于自适应局



## 渐变

 radial-gradient径向渐变指从一个中心点开始沿着四周产生渐变效果 

 **1**、必要的元素：

 a、辐射范围即圆半径  

 b、中心点 即圆的中心

 c、渐变起始色

 d、渐变终止色

 **2**、关于中心点：中心位置参照的是盒子的左上角

 **3**、关于辐射范围：其半径可以不等，即可以是椭圆



## 过渡

 特点：当前元素只要有“属性”发生变化时，可以平滑的进行过渡。

 transition-property设置过渡属性

 transition-duration设置过渡时间

 transition-timing-function设置过渡速度

 transition-delay设置过渡延时

 以上四属性重在理解  



## 2D转换

 转换是CSS3中具有颠覆性的特征之一，可以实现元素的位移、旋转、变形、缩放，甚至支持矩阵方式，配合即将学习的过渡和动画知识，可以取代大量之前只能靠Flash才可以实现的效果。

 1、移动 translate(x, y) 可以改变元素的位置，x、y可为负值；

 2、缩放 scale(x, y) 可以对元素进行水平和垂直方向的缩放，x、y的取值可为小数，不可为负值；

 4、旋转 rotate(deg) 可以对元素进行旋转，正值为顺时针，负值为逆时针；

 5、倾斜 skew(deg, deg) 可以使元素按一定的角度进行倾斜

## 3D 转换

 **1**、****3D****坐标轴

 用X、Y、Z分别表示空间的3个维度，三条轴互相垂直。如下图 

 **2**、左手坐标系

 伸出左手，让拇指和食指成“L”形，大拇指向右，食指向上，中指指向前方。这样我们就建立了一个左手坐标系，拇指、食指和中指分别代表X、Y、Z轴的正方向。如下图

 

 **3**、左手法则

 左手握住旋转轴，竖起拇指指向旋转轴正方向，正向就是其余手指卷曲的方向。



## 动画

 动画是CSS3中具有颠覆性的特征之一，可通过设置多个节点来精确控制一个或一组动画，常用来实现复杂的动画效果。

 **1**、必要元素：

 a、通过@keyframes指定动画序列；

 b、通过百分比将动画序列分割成多个节点；

 c、在各节点中分别定义各属性

 d、通过animation将动画应用于相应元素；

 **2**、关键属性

 a、animation-name设置动画序列名称

 b、animation-duration动画持续时间

 c、animation-delay动画延时时间

 d、animation-timing-function动画执行速度，linear、ease等

 e、animation-play-state动画播放状态，play、paused等

 f、animation-direction动画逆播，alternate等

 g、animation-fill-mode动画执行完毕后状态，forwards、backwards等

 h、animation-iteration-count动画执行次数，inifinate等



## flex伸缩布局