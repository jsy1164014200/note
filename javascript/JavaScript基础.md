# javascript

## 数据类型

1. 数值
2. 字符
3. 布尔
4. null undefined

> javascript 是基于 原型的脚本

## 函数

Javascript在创建函数的同时，会在函数内部创建一个arguments对象实例.

arguments是存储了函数传送过过来实参。arguments对象只有函数开始时才可用。函数的 arguments 对象并不是一个数组，访问单个参数的方式与访问数组元素的方式相同

arguments对象的长度是由实参个数而不是形参个数决定的

这个单词只在函数 内使用 ， 而且是正在执行的函数。  

```js
function fn(a,b) {
        //  console.log(fn.length); 返回的是 函数的 形参的个数
        // console.log(arguments.length);  返回的是正在执行的函数的 实参的个数
    // arguments  里面存放的是 [1,2]  
        if(fn.length == arguments.length)
        {
            console.log(a+b);
        }
        else
        {
            console.error("对不起，参数不匹配，参数正确的个数应该是" + fn.length);
        }
    }
```

var 声明的变量有 变量提升

