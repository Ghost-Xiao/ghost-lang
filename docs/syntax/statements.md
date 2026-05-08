# 语句语法说明

Ghost Language 支持多种语句类型，以下是详细的语句语法定义和示例。

## 表达式语句(ExpressionStatement)
将表达式作为语句执行。

**语法定义：**
```
ExpressionStatement ::= Expression
```

**示例：**
```ghost
x = 10;
println("Hello");
```

## For 循环语句(ForStatement)
用于循环执行代码块的语句。

**语法定义：**
```
ForStatement ::= "for" Statement ";" Expression ";" Statement Statement
```

**示例：**
```ghost
for var i = 0; i < 10; i = i + 1 {
  println(i);
};
```

## 函数声明语句(FunctionDeclarationStatement)
用于声明函数的语句。

**语法定义：**
```
FunctionDeclarationStatement ::= "func" Identifier "(" (ParameterList)? ")" Statement
ParameterList ::= NonDefaultParameter ("," NonDefaultParameter)* ("," DefaultParameter ("," DefaultParameter)*)? ("," VariadicParameter)?
NonDefaultParameter ::= Identifier
DefaultParameter ::= Identifier "=" Expression
VariadicParameter ::= "..." Identifier
```

**示例：**
```ghost
func add(x, y=0) {
  return x + y;
};

func foo(...args) {
  println(args);
};

func printAll(...args) {
  foreach var val in args {
    println(val);
  };
};

printAll(1, "hello", true, 3.14);
// 输出:
// 1
// hello
// true
// 3.14

func sum(...nums) {
  var total = 0;
  foreach var n in nums {
    total = total + n;
  };
  return total;
};

sum(1, 2, 3, 4, 5);  // 返回 15
```

**注意事项：**
- 函数参数可以是非默认参数或默认参数。
- 默认参数必须在参数列表的末尾。
- 可变参数必须是最后一个参数。
- 可变参数在函数内部作为一个列表使用。
- 可变参数可以接受任意数量、任意类型的参数。
- 一个函数只能有一个可变参数。

## 返回语句(ReturnStatement)
用于从函数中返回值的语句。

**语法定义：**
```
ReturnStatement ::= "return" Expression
```

**示例：**
```ghost
return 42;
return x + y;
```

## 命名空间声明语句(NamespaceDeclarationStatement)
用于声明命名空间的语句，命名空间可以包含变量、函数等成员。

**语法定义：**
```
NamespaceDeclarationStatement ::= "namespace" Identifier Statement
```

**示例：**
```ghost
namespace Array {
    var array = [1, 2, 3, 4, 5];
};

namespace Utils {
    func add(x, y) {
        return x + y;
    };
};
```

**注意事项：**
- 命名空间中的成员通过 `命名空间名::成员名` 的方式访问
- 命名空间中的变量和函数在命名空间声明后即可访问

## Foreach 循环语句(ForeachStatement)
用于遍历可索引对象（如列表）的循环语句。

**语法定义：**
```
ForeachStatement ::= "foreach" ("var")? (("[" Identifier "," Identifier "]") | Identifier) "in" Expression ("step" Expression)? Statement
```

**示例：**
```ghost
foreach var val in [1, 2, 3, 4, 5] {
    println(val);
};

foreach var [i, val] in [1, 2, 3] {
    println(i, val);
};

var item = 0;
foreach item in [1, 2, 3] {
    println(item);
};

var index = 0;
var value = 0;
foreach [index, value] in [1, 2, 3] {
    println(index, value);
};

foreach var i in 1..10 step 2 {
    println(i);
};
```

**注意事项：**
- 使用 `var` 关键字时，会在循环内声明新变量
- 不使用 `var` 时，变量必须已在外部作用域中定义
- 可同时获取索引和值，使用方括号和逗号分隔
- 可选 `step` 关键字用于指定遍历步长
- `break` 和 `continue` 语句可用于控制循环流程

## 导入语句(ImportStatement)
用于导入内置模块的语句。

**语法定义：**
```
ImportStatement ::= "import" Expression
```

**示例：**
```ghost
import "math";
import "fmt";
import "io";
```

**注意事项：**
- 导入语句必须在程序的顶级作用域中
- 导入语句可以重复导入同一模块，但不会重复加载
- 导入语句的模块名必须是字符串常量
- 导入语句用于导入内置模块或自定义模块
- 导入的模块可以通过成员访问表达式访问
- 支持的内置模块包括：math、fmt、io
- 模块导入后可以使用 `模块名.变量名` 或 `模块名.函数名()` 的方式访问

**使用示例：**
```ghost
import "math";
math.PI;
math.sqrt(4);
```

## Try 语句(TryStatement)

用于异常处理的语句，支持 try、catch 和 finally 块。

**语法定义：**
```
TryStatement ::= "try" Statement ("catch" Identifier Statement)? ("finally" Statement)?
```

**示例：**
```ghost
try {
    // 可能抛出异常的代码
    var x = 10 / 0;
} catch e {
    println("捕获到异常:", e);
} finally {
    println("清理工作");
};

try {
    throw Error("自定义错误");
} catch err {
    println(err);
};
```

**注意事项：**
- `try` 块是必需的，用于包裹可能抛出异常的代码
- `catch` 块是可选的，用于捕获和处理异常
- `finally` 块是可选的，用于执行无论是否发生异常都需要执行的清理代码
- `catch` 块必须有一个参数，用于捕获异常实例
- finally 块始终会执行，即使 try 或 catch 块中有 return 语句

## Throw 语句(ThrowStatement)

用于抛出异常的语句。

**语法定义：**
```
ThrowStatement ::= "throw" Expression
```

**示例：**
```ghost
throw Error("发生了错误");
throw TypeError("类型不匹配");

func divide(x, y) {
    if y == 0 {
        throw Error("除数不能为零");
    };
    return x / y;
};
```

**注意事项：**
- throw 语句必须抛出 Error 类或其子类的实例
- 异常可以在 try-catch 块中被捕获和处理

## 类声明语句(ClassDeclarationStatement)

用于声明类的语句，类可以包含成员变量和方法。

**语法定义：**
```
ClassDeclarationStatement ::= "class" Identifier ("extends" Identifier)? Statement
```

**示例：**
```ghost
class Person {
    var name = "";
    var age = 0;
    
    func init(name, age) {
        this.name = name;
        this.age = age;
    };
    
    func greet() {
        println("Hello, my name is", this.name);
    };
};

class Student extends Person {
    var grade = 0;
    
    func init(name, age, grade) {
        super.init(name, age);
        this.grade = grade;
    };
    
    func study() {
        println(this.name, "is studying in grade", this.grade);
    };
};
```

**注意事项：**
- 类声明使用 `class` 关键字
- 可以使用 `extends` 关键字继承另一个类
- 类中可以包含变量声明和函数声明
- 类中的函数称为方法
- 使用 `this` 关键字访问当前实例的成员
- 使用 `super` 关键字访问父类的成员
- 类的构造函数通常命名为 `init`
- 类声明后可以通过 `类名()` 或 `new 类名()` 创建实例
