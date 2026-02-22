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
ParameterList ::= NonDefaultParameter ("," NonDefaultParameter)* ("," DefaultParameter ("," DefaultParameter)*)?
NonDefaultParameter ::= Identifier
DefaultParameter ::= Identifier "=" Expression
```

**示例：**
```ghost
func add(x, y=0) {
  return x + y;
};
```

**注意事项：**
- 函数参数可以是非默认参数或默认参数。
- 默认参数必须在参数列表的末尾。

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