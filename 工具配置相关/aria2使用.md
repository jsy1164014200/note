aria2使用

https://blog.csdn.net/wudi1107/article/details/80728891

firefox  chrome插件安装



开启服务 : sudo aria2c --conf-path=/etc/aria2/aria2.conf



基础笔记

```
package one.study.kotlin

fun main(args: Array<String>) {
    // 没有包装数据类型（与java相区别），kotlin会自动转换（如果需要使用包装数据类型）
    var bool: Boolean = false
    var byte: Byte = 10
    var int: Int = 20
    var long: Long = 40
    var char: Char = 'a'
    var double: Double = 1.234432424
    var float: Float = 1.34243234232344234f


    //智能类型推导
    var a = 10
    println(a)

    // java中 小类型能转换成大类型，kotlin 赋值强类型，只能通过to方法
    var b = 10
    var c = 10L
    b = c.toInt()
//    c = b

    // 可变变量，不可变变量
    // java中的 final 是编译器常量
    val con = 10   // 不可变变量，尽量多使用

    // 字符串
    val place = "xxxxxx"
    val place1 = """skdfsdfsd""".trimIndent()
    println(place1)

    // 字符串比较 == 等于 equal ，不同于java
    val path = "/path/yole/kotlin-book/sfd"

    // 元组类型
    val pair = Pair<String, Int>("xx", 10)
    // 简写
    var pair2 = "22" to 22

    val triple = Triple<String, Int, String>("xx", 10, "dsf")

    // null保护
    var str: String = "xxxx"     // 非空类型
    var str1: String? = null    // 可空类型

    str1!!.toInt()   // 告诉编译器 ，我一定是空，不用检查了(不建议使用）
    str1?.toInt() ?: -1 // 空安全调用符，  == if (str !=null ) {return str.toInt()是一个int?可空类型} else return null

    // println
    // readline
    readLine()
    // 模板字符串
    val result = "xxxxxx${place}sdfksdf${place}"
}

// kotlin 中 unit  类似 java void，但是前者是一个对象
fun getLength(name: String): Int {
    return 1
}

// Ctrl + alt + l

// java 是声明式语言，kotlin 是表达式语法，if when有返回值
fun max(a: Int, b: Int): Int {
    return if (a > b) a else b
}

fun forEach() {
    // for foreach 循环
    val str = "xxkjdsfkjs"
    tag@ for ((index, c) in str.withIndex()) {
        println("index=$index,$c")
        break@tag
    }
    // 高级foreach循环中不能使用 continue break
    str.forEachIndexed { index, c ->
        println("index=$index,$c")
    }
    str.forEach { it ->
        println(it)
    }
    // 标签处返回
    // while do while

}

fun range() {
    val range1 = 1..10
    val range2 = IntRange(1, 10)
    val range3 = 1.rangeTo(10)

    var longRange = 1L..10

    val charRange = 'a'.rangeTo('z') // 'a'..'z'

    for (c in charRange) {

    }
    charRange.forEach {

    }

    // 反向区间
    val ra = 100 downTo 1 step 2

    val reverseRange = 100 downTo 1 step 2
    reverseRange.forEach {
        println(it)
    }

    // 数组
    val arr = arrayOf("dsf", "dskjfk", "sdfjk") // Array<String>

    val arr1 = arrayOf("sdx", 10, 'x') // Array<Any>
    val arr2 = IntArray(10) // new [10]int
    val arr3 = IntArray(10) {
        30
    }

    // 调用
    arr1[2]
    arr1.set(2, 3)

    // 高阶函数
    val index = arr1.indexOfLast {
        it == 3
    }
}

// when 表达式

fun todo(age: Int): Int {
    when (age) {
        1, 2, 3, 4, 5, 6, 7 -> {
            return
        }
        in 1..8 -> {  // 支持 区间
            return
        }
        is Int -> {  // 支持判断
            return
        }
        else -> {
            return
        }
    }

    return when {
        age == 7 -> {
            111
        }
        age == 12 -> {
            222
        }
    }

}


// 函数表达式, 只能对一行的代码才能用 函数表达式
fun add(a: Int, b: Int = 10): Int = a + b

// 函数变量
val padd = ::add    // :: 叫做函数引用，类似c中的函数指针
// 为 KFunction2<int,int,int>
padd(10,10)   // 不能处理函数类型为空
padd?.invoke(10,10)     // 可以处理函数为空

// 匿名函数
val padd: (Int, Int) -> Int = { a, b -> a + b }

// 默认参数 和 具名参数
add(a=10,b=20);

// 可变参数

fun add(vararg a: Int): Int {
    // a 是 数组，java中是 int... a
}


// 异常处理

try {
    10 / 0
} catch (e:ArithmeticException) {

} finally {

}

// kotlin中 编译型异常
// kotlin 尾递归， 函数调用自身后没有调用其他东西

// 尾递归的原理，将递归 转化成 迭代
tailrec fun sum(n: Int, result: Int = 0):Int {
    if(n == 1) {
        return result + 1
    }else {
        return sum(n-1,result+n)
    }
}


// 面向对象
class Girl {
    var name:String = "xx"
    var age:Int = 22
    fun greeting() {
        println("hello")
    }
}

class Rect {
    var long:Int = 100
    var width:Int = 100
}


// 运算符重载，kotlin中每一个运算符 都是一个方法

class Boy {
    // 定义出对应的函数
    operator fun plus(boy:Boy):Boy {
        return this
    }
}

class Teacher {
    var level = 1
    private set // 相当于私有了 set 方法

    var age:Int = 1   // kotlin 会自动生成 私有 属性，然后给出 public 的set get方法
    set(value:Int) {
        if (value < 100) {
//            this.age  不能这样调用，不然会递归（因为 this相当于 调用了 setAge()
            field = value   // 只能用set
        }
    }

    operator fun inc():Teacher {

    }
}



// 构造函数
class Dats(name:String,age:Int) {
    var name = "xx"
    var age = 10
    init {
        this.name = name
        this.age = age
    }
}

// 或者
class Dat(var name:String,var age:Int) {
    init {
        // 已经默认自带了 this.age = age || this.name = name
        // 所以 var  val 相当 于自带了 定义属性和 定义构造函数的功能
        println("xxxx")
    }

    var phone = "xx"
    // 次构函数
    constructor(name:String,age:Int,phone:String):this(name,age) {
        this.phone = phone
    }

    private fun test() {

    }
}



// 继承 多态
open class Father {
    // kotlin 中 默认会加上 final 关键字， 所以必须加上 open 才能继承
    open var name = "zhangsan"
    open var age = 20

    open fun horbe() {
        println("xx")
    }
}

class Son:Father() {
    override var name = "xx"
    override var age = 10
    override fun horbe() {
        println("ss")
    }
}


// 抽象类 没有open 关键字
abstract class Person {
    abstract var name:String
    abstract var age:Int
    abstract fun eat()
    fun eat2() {}
    var a:String = "x"
}

class ZhPerson(override var name: String, override var age: Int) :Person() {
    override fun eat() {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}

// interface 接口   kotlin 中 接口的属性 不能有实现, java中接口属性可以有实现
// java中方法不能实现, kotlin中的 方法能构实现
interface RideBike {
    fun drive() {println("xxxx")}
    var name:String
}

// 多态,父类能调用子类 override 东西,但是不能调用 父类中没有的方法

// 智能类型转换
if (son is Son) {
    // val son = son as Son   这句话可以不要
//    kotlin 会自动将son 转换成 Son
}

// kotlin中内部类默认是static, 与外部类没有关系

// 如果加上 inner 就是 内部类了 ,可以访问外部属性
class OutClass {
    class InnerClass {

    }
    inner class RealClass{
        fun test() {
            println("xxx${this@OutClass.name}")
        }
    }
}


// 泛型类
open class Box<T>(thing:T) {

}

class FruitBox(thing:Fruit):Box<Fruit>(thing) {

}

class SonBox<T>(thing:T):Box<T>(thing)

abstract class Fruit


// 能过指定泛型的上限, T:thing  相关与 T继承thing
fun <T> parseType(thing:T) {
    String::class.java // 反射
}

// 泛型擦除 ,java 中有泛型擦除
// kotlin中 解决 泛型擦除 ,在fun前加 inline, 加reified
// 但是kotlin中的类型擦除只能解决 函数的, class 还是 有类型擦除
inline fun <reified T> parseType2(thing:T) {
    String::class.java // 反射
}


//      fun <? extend Fu>xxxx() {}  java中的 写法 要 继承 xxxx的 都能够传入 ,, ,,,,
// 在 kotin中Box<String> != Box<Int>

// 传父类   extend 改成 super
fun setFruitList(list:ArrayList<out Fruit>) {}
// * 类型  相当于 java中的 *
fun setFruitList(list:ArrayList<*>) {}
```