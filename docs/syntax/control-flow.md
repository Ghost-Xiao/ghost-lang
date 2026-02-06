# 控制流语法说明

Ghost Language 支持多种控制流结构，以下是详细的控制流语法定义和示例。

## 条件表达式(IfExpression)
用于条件分支的表达式。

**语法定义：**
```
IfExpression ::= "if" Expression Statement ("else" Statement)?
```

**示例：**
```ghost
if x > 5 {
  x * 2;
} else {
  x / 2;
};
```

**注意事项：**
- 条件表达式的返回值是条件分支中最后一个语句的返回值。
- 如果没有 else 分支且条件表达式的条件为 false，条件表达式的返回值是 null。

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