# 函数语法说明

Ghost Language 支持函数定义和调用，以下是详细的函数语法定义和示例。

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

## 函数调用表达式(CallExpression)
表示函数调用的表达式节点。

**语法定义：**
```
CallExpression ::= Expression "(" (ArgumentList)? ")"
ArgumentList ::= Argument ("," Argument)*
Argument ::= Expression | ""
```

**示例：**
```ghost
println("hello");
add(1, , 2);
len([1, 2, 3]);
```

**注意事项：**
- 调用函数时，参数列表中的空参数（逗号分隔，空参数代表使用默认值）会被忽略。
- 调用函数时，如果参数数量少于函数定义的参数数量，未被赋值的参数会使用默认值。