# kotlin

> 百度网盘下载工具 aria2
>
> https://addons.mozilla.org/zh-CN/firefox/addon/baidu-pan-exporter/ firefox插件

jetbrain发布的基于jvm的编程规范语言

Andrey Breslav，kotlin科特林岛，java 爪哇岛

1. 简洁
2. 空值安全
3. 兼容 java scala groov
4. 函数式
5. thread
6. dsl领域特定语言



.kt文件， 编译，*Kt.class字节码文件，jvm运行

可以从字节码反编译成 java源码





## java接口回调

```kotlin
public class JavaTest {
//    ArrayList list = new ArrayList<>()

    class Mother {
        void make() throws InterruptedException {
            SuperMarket.buySoy(new Son());
            System.out.println("做甜点");
        }
    }
    class Son implements FeedBack {

        @Override
        public void feed(Soy soy) {
            System.out.printf("买到了酱油");
            System.out.println("妈妈开始做饭");
        }
    }

    static class SuperMarket {

        public static void buySoy(FeedBack feedBack) throws InterruptedException {
            Thread.sleep(5000L);
            Soy soy = new Soy();
            feedBack.feed(soy);
        }
    }
    static class Soy {}



}

interface FeedBack {
    void feed(JavaTest.Soy soy);
}
```



## kotlin函数回调

kotlin中 不用像java那样 ,使用对象来 调用 接口

```kotlin
fun main(args: Array<String>) {
    val superMarket = SuperMarket()
    superMarket.buySoy { println(it) }
}


class SuperMarket {
    fun buySoy(block: (Soy) -> Unit) {
        Thread {
            Thread.sleep(5000)
            block(Soy())
        }.start()
    }
}

class Soy {}
```



## 基础数据类型

```kotlin
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





## 特殊类

```kotlin
// 中缀表达式, 可以自定义一些 操作符, 让代码简单易懂   , 必须是成员函数 ,必须只有一个参数, 参数不能可变或者默认参数
class Test {
    infix fun sayHelloTo(name: String) {
        println("xx")
    }
}

// 设计模式
// 类委托
// 洗碗能力
interface WashPower {
    fun wash()
}

class BigHeadSon : WashPower {
    override fun wash() {
        println("大头儿子洗碗")
    }
}

class SmallHeadFather : WashPower by BigHeadSon() {  // by 关键字实现 委托

}

// 或者
class SmallHeadFather2(val washPower: WashPower) : WashPower by washPower {
    override fun wash() {
        println("fuqian")
        washPower.wash()
        println("xiwan")
    }
}

// 属性委托
class Son {
    // 儿子的前 交给 妈妈
    var money: Int by Mother()
}

class Mother {
    operator fun getValue(son: Son, property: KProperty<*>): Int {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
        return sonMoney
    }

    operator fun setValue(son: Son, property: KProperty<*>, i: Int) {
        sonMoney = i / 2
        selfMoney = i / 2
    }

    var sonMoney = 0
    var selfMoney = 0
}

// by lazy 懒加载
// 延迟加载

class Person {
    // 用的时候再加载
    val name: String by lazy { println("xx"); "zhangsan" }  // 字段必须是 val 不可变, val 可以单独存在 ,或者在类中;;; by lazy 是线程安全
    val age: Int = 10

    lateinit var name2: String   // 用的时候再赋值,步赋值不能用 ;; var  ,
}

// 扩展函数 ,类似 动态语言中的 添加字段
fun String?.myIsEmpty(): Boolean {
    return this == null || this.length == 0
}

fun kuozhanghanshu() {
    val str: String? = "zhangsan"
    val myIsEmpty = str?.myIsEmpty()
}


// 恶汉式单例模式 只有一个对象
object SingleInstance {
    // java
    // 第一步 ,私有化 构造函数

//    private Static SingleInstance single = new SingleInstance()
//    private SingleInstance() {}

    // 第二步: 一个返回当前对象的函数

//    public Static SingleInstance getInstance() {return single}

    // kotlin
// kotlin 将 class 改成 Object  然后通过 SingleInstance2.xxxx 调用就可以了
    val name = "xx"  // kotlin 会把它变成静态的

    fun say() {  //  方法不会变成静态的
        println("x")
    }
}
/*Object 单例在有少数 属性的适用,因为他的属性都是 static*/


// 懒汉式单例模式
class SingleInstance2 {
    // java
    // 第一步 ,私有化 构造函数

//    private Static SingleInstance single = null   // 关键
//    private SingleInstance() {}

    // 第二步: 一个返回当前对象的函数

//    public Static SingleInstance getInstance() {
//     synchronized (SingleInstance.class)  // 解决 多线程 不安全的问题
//      if(single == null ) {
//          single = new SingleInstance2()
//      }
//      return single}


}


// 伴生对象 , 实现 饱汉单例模式
class CompanionObject private constructor() {
    var age = 20

    companion object {  // 在里面的都是 static
        var name = "xx"
        fun test() {}
        val instance by lazy { CompanionObject() }
    }
}

// 枚举
// 枚举区间 WEEK.a..WEEK.e
enum class WEEK {
    a, b, c, d, e
}

enum class COLOR(val r: Int, val g: Int, val b: Int) {
    RED(255, 0, 0), BLUE(0, 255, 0), GREEN(0, 0, 255)
}

// 数据类
// 只保存数据,没有任何其他操作的类
data class News(var title: String, var desc: String, var content: String, var author: String)
// 生成 get set toSting hashcode equals copy
// 使用 val a = News()  val (title,desc,content,author) = a  相当于析构

// 密封类
sealed class NedStark {
    class RobStark : NedStark()
    class SansaStark : NedStark()
    class ArayStarK : NedStark()
}
// 使用 is 判断类型

class Nex : NedStark()

// java中 ArrayList 相比 vector 是线程不安全的  hashMap 对比hashtable 线程不安全

fun Collection() {
    val list = listOf("xx","xx")   // 不可变 不能添加 不能修改
    val list2 = mutableListOf("xx","xx") // 可变集合  ,, 最后的也是 ArrayList

    val set = setOf<String>("xx","xx")
    val set2 = mutableSetOf("xx","xx")
    val set3 = hashSetOf("xx")
    val treeSet = TreeSet<String>()

    val map = mapOf("xx" to "xx","xxksjdf" to "xx")
//    hashMapOf<>()

}
```



## kotlin函数式编程

```kotlin
// kotlin oop fp
// 闭包, lambda 表达式
fun test(): () -> Unit {
    var a = 10
    return {
        println("xx")
        a++
    }
}

// 高阶函数
fun add(a: Int, b: Int): Int {
    return a + b
}

fun cacl(a: Int, b: Int, block: (Int, Int) -> Int): Int {  // block 是 函数引用类型
    return block(a, b)   // 或者 block.invoke(a,b)
}

fun test2() {
    cacl(1, 1, ::add)
    cacl(1, 1, { m, n -> m + n })
    // lambda表达式 括号可以前移
    cacl(1, 1) { m, n -> m - n }
}

//  {} 里面的都是 函数  ,JavaScript中 {} 都是 对象
fun test3() {
    { m: Int, n: Int -> println("xx") }(1, 1)
}

fun test4() {
    val girls = arrayListOf("xxx", "xxx", "xxxxx")
//    girls.filterIndexed { index, s -> println("xx") }
    girls.sort()
    girls.sortDescending()
//    girls.sortBy { it.xxx }
    girls.groupBy { "xxx" }
    girls.distinctBy { "x" }  // 去重

    // let 表达式
    girls.let {
                it.add("xx")
    }

    // with 函数
    with(girls){
        this.add("xx")
    }
    // run 函数
    girls.run { add("xx") }
    // apply 函数
    girls.apply { add("xx") }
}

// 接口回调, 函数回调
// 1. 定义某种能力的接口
// 2. 接受能力对象
// 3. 通过能力对象 将数据传回去
// 4. 在具有能力的 对象方法下接受
```





## DSL

> 领域特定语言,不同于GPL(通用计算机语言)

TODO: 学习