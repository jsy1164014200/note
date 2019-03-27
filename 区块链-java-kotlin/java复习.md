# Java复习

## 1. 代码块
- 普通的代码块每次在初始化的时候都会加载
- static 代码块只会在连接的时候加载

## 2. 内部类

- 外部类名.this.
- 内部类 无门槛访问外面的类
- 外面的类 访问内部的类 要new一个对象
- static 声明的 内部类 相当于 外部类

## 3. Java子类继承了父类的所有属性，而private只是访问控制符，控制的何处能够访问到属性。 但是子类访问不了父类的 private属性，就不可能重写

## 4. 

多态的使用
```java
public class Example extends Fu{
    public int i = 99;
    @Override
    public void printInfo(){
        System.out.println("zi nei function!");
    }

    public static void main(String[] args) {
        Fu e = new Example();
        System.out.println(e.i);
        System.out.println(e.j);
        e.printInfo();

        Example e1 = (Example)e;
        System.out.println(e1.i);

        e1.printInfo();
    }

}
class Fu{
    public int i = 2;
    public int j = 1;

    public void printInfo(){
        System.out.println("fu nei function!");
    }
}
// 结果
2
1
zi nei function!
99
zi nei function!
```

## 5. final的使用

- final加在类的前面就是不允许继承
- final加在 方法前面就是 不允许重写
- final加在变量前面就是不允许更改

## 6. instanceof

new的时候是什么对象 instanceof 就是什么对象 
它也是父类的 实例

## 7. abstract 

抽象类 ：不允许做抽象化 该类

## 8. 接口 interface 

接口所有的量 都是public static final 
public abstract 


javac -d . ./cn/itcast/store/  *.java

## 9. 对象包装器 跟 自动装箱

基本类型变量 int 比 Arraylist<Interger> 性能要高，但是不便于添加元素等操作

**包装类都是不可变的**
1. 自动装箱
如果 Arraylist<Integer> a = new Arraylist<>();
a.add(3) => a.add(Integer.valueOf(3))
2. 自动拆箱
int n = a.get(3) => int n = a.get(3).intValue();
3. 如果一个Integer 类型加上 Double类型，前一个类型就会自动的拆箱 然后装箱
4. int x = Integer.parseInt(string)


## 10. 枚举

枚举中的 枚举量就是 该枚举类的一个实例

所有的枚举都是Enum 的子类 
- Enum.valueOf(Class enumClass,String name) 返回 该枚举类的 实例
- 所有的枚举类 自带 toString
```java
public enum Size {
    SMALL("S"),MEDIUM("M"),LARGE("L"),EXTRA_LARGE("XL");

    private String abbreviation;

    private Size(String abbreviation) {
        this.abbreviation = abbreviation;
    }

    public static void main(String[] args) {
        Size a = Size.SMALL;
        System.out.println(a.toString());
    }
}
```

## 11. 反射

1. class.forName(String s)
2. class.newInstance()
3. instance.getClass()

## 12. 接口

这是它与抽象类的 最大区别，所以说接口是比抽象类更加抽象的一种类
1. 接口中的field 都是 public static final
2. 接口中的方法 都是 public abstract

Java8中 接口中的方法可以有方法体 

如果超类跟 接口冲突 
> extends Person implements Name
则会默认是 超类的方法

如果是 接口之间的冲出 ，会编译错误 子类要实现该方法 
> implements Person ,Name
> String getName() {return Person.super.getName()}

## 13. 方法引用，lambda表达式

> System.out::println

java中的>>>表示无符号的右移

1. float f=3.4;是否正确？
答:不正确。3.4是双精度数，将双精度型（double）赋值给浮点型（float）属于下转型（down-casting，也称为窄化）会造成精度损失，因此需要强制类型转换float f =(float)3.4; 或者写成float f =3.4F;。

2. short s1 = 1; s1 = s1 + 1;有错吗?short s1 = 1; s1 += 1;有错吗？
答：对于short s1 = 1; s1 = s1 + 1;由于1是int类型，因此s1+1运算结果也是int 型，需要强制转换类型才能赋值给short型。而short s1 = 1; s1 += 1;可以正确编译，因为s1+= 1;相当于s1 = (short)(s1 + 1);其中有隐含的强制类型转换。


3. 注意理解 Java中的 整数池（-128~127） 以及 字符串池
```java
Integer a = 1;
Integer b = 1;
Integer c = new Integer(1);
Integer d = new Integer(1);

System.out.println(a == b); // true
System.out.println(c == d); // false

// 但是 要注意
Integer f1 = 100, f2 = 100, f3 = 150, f4 = 150;

System.out.println(f1 == f2); // true
System.out.println(f3 == f4); // false

```
对于String 字符串，new , subString 都会在堆中分配空间， 所以他们的引用是不相同的，但是对于 
> String a = "dksjf"
> String b = "dksjf"  a == b // true
```java
String s1 = "Programming";
String s2 = new String("Programming");
String s3 = "P" + "rogramming";
String s4 = "Programming123".substring(0,11);

System.out.println(s1);
System.out.println(s2);
System.out.println(s3);
System.out.println(s4);

System.out.println(s1 == s3); // true
System.out.println(s1 == s4); // false 
```


同理就能解释上面的 自动拆箱 包装

4. Math.round(11.5) 等于多少？Math.round(-11.5)等于多少？
答：Math.round(11.5)的返回值是12，Math.round(-11.5)的返回值是-11。四舍五入的原理是在参数上加0.5然后进行下取整。

5. java继承 会继承所有的东西，但是除了构造函数。构造器（constructor）是否可被重写（override）？
答：构造器不能被继承，因此不能被重写，但可以被重载。


```java
class A {

    static {
        System.out.println("A static");
    }
    {
        System.out.println("A common");
    }
    public A() {
        System.out.println("A construct");
    }
}

class B extends A{

    static {
        System.out.println("B static");
    }
    {
        System.out.println("B common");
    }
    public B() {
        System.out.println("B construct");
    }
}

B b= new B()
B c= new B()
A static
B static
A common
A construct
B common
B construct
A common
A construct
B common
B construct
```


java 类中的变量都会被 初始化  但是在 main中的量不会被初始化 

常量必须手动初始化


# Java-ppt总结

+ OOP 面向对象 
    - 1.封装  （private ，然后提供访问API）
    - 2.继承  （Java是单继承，，抽象类，接口）
    - 3.多态  （非静态方法的重写，上级引用调用）

+ 修饰符
    - private 成员属性或方法只能在定义类的内部访问。
    - (没有修饰符) 同一个包内的类都可以访问
    - protected 只有继承的子类与同一包内的类可以访问 *这个地方的protected与c++的有些不同，Java的在一个包中也能够访问*
    - public 代表此对象的对外使用界面，其它类的方法可以访问此对象的成员属性或方法。

+ 重载 与 重写（覆盖重写）
    - 重载Overload ：在同一个类中，方法名字一样
    - 重写Override ：在子类中，要方法名字一样，参数类型返回类型也要一模一样

+ 子类（subclass）继承父类 superclass (parent class)所有的属性和方法
    - 不继承父类的构造方法
    - 子类不能访问父类声明成private的成员。
    - 子类用super()访问父类构造方法。

+ 隐式类型
    - t + "" 会调用 t的toString
    - 小转大

+ final一定要被 初始化  类中的成员非final（包括static）都会被自动初始化


+ 理解Java懒加载
```java

public class Test {
//    public static final int a;
    static{
        System.out.println("testA");
    }
    public static void main(String[] args) {
//        A a=new A();
        System.out.println("counter = " + A.counter);
        System.out.println("counter = " + B.counter);
        new B();
// output：：：：
// testA
// A father
// Aclass
// counter = 48
// counter = 48
// Bclass

    }
}

class  A_father{
    static {
        System.out.println("A father");
    }
}

class A extends A_father{
    public static int counter=47;
    static {
        counter++;
        System.out.println("Aclass");
    }
}

class B extends A {
    static {
        System.out.println("Bclass");
    }
}
```

+ 局部方法内部类能够访问外部类的任何 (包括私有)成员
    - 因为不能保证局部变量的存活期和方法本地内部类对象一样长
    - 把局部变量声明为final的情况下可访问 
    - 局部方法内部类与局部变量声明相同，不能被标识为public, private, protected, static, transient等 能应用于局部方法内部类的修饰符是abstract和final(这两个修饰符不能同时使用) 
    - 在一个静态方法内声明的内部类只能访问其外部类的静态成员



public class Example {
    public static void method() {
        try{
            wrench();
            System.out.println("a");
        } catch (ArithmeticException e){
            System.out.println("b");
        } catch (NullPointerException e){
            System.out.println("catch");
        }
        finally {
            System.out.println("c");
        }
        System.out.println("d");
    }
    public static void wrench(){
        throw new NullPointerException();
    }

    public static void main(String[] args) {
        try{
            method();
        }catch (Exception e){
            System.out.println("e");
        }
        System.out.println("f");
    }
}

catch 
c
d
f
public class Example {
    public static void method() {
        try{
            wrench();
            System.out.println("a");
        } catch (ArithmeticException e){
            System.out.println("b");
        }
        finally {
            System.out.println("c");
        }
        System.out.println("d");
    }
    public static void wrench(){
        throw new NullPointerException();
    }

    public static void main(String[] args) {
        try{
            method();
        }catch (Exception e){
            System.out.println("e");
        }
        System.out.println("f");
    }
}

c
e
f

2>1 ? 9:99.0  

9.0