# 表达式语法说明

Ghost Language 支持多种表达式类型，以下是详细的表达式语法定义和示例。

## 整数字面量(IntegerLiteral)
表示整数值的表达式节点，支持十进制、二进制、八进制和十六进制。

**语法定义：**
```
IntegerLiteral ::= DecimalLiteral | BinaryLiteral | OctalLiteral | HexLiteral
DecimalLiteral ::= [0-9]+
BinaryLiteral ::= ("0b" | "0B") [01]+
OctalLiteral ::= ("0o" | "0O" | "0") [0-7]+
HexLiteral ::= ("0x" | "0X") [0-9a-fA-F]+
```

**示例：**
```ghost
42;
-10;
0b1010;
0B1101;
0o52;
0O67;
040;
0x2a;
0X1F;
```

## 浮点数字面量(FloatLiteral)
表示浮点数值的表达式节点。

**语法定义：**
```
FloatLiteral ::= [0-9]+ "." [0-9]* (("e" | "E") ("+" | "-")? [0-9]+)?
               | [0-9]+ (("e" | "E") ("+" | "-")? [0-9]+)
```

**示例：**
```ghost
3.14;
-2.718;
1.2e3;
1.2E-4;
5e2;
```

## 布尔字面量(BooleanLiteral)
表示布尔值(true/false)的表达式节点。

**语法定义：**
```
BooleanLiteral ::= "true" | "false"
```

**示例：**
```ghost
true;
false;
```

## 空字面量(NullLiteral)
表示空值的表达式节点。

**语法定义：**
```
NullLiteral ::= "null"
```

**示例：**
```ghost
null;
```

## 字符串字面量(StringLiteral)
表示字符串值的表达式节点。

**语法定义：**
```
StringLiteral ::= (""" .*? """) | ("'" .*? "'") | ("`" .*? "`")
```

**示例：**
```ghost
"Hello, Ghost!";
'Hello, Ghost!';
`Hello, Ghost!`;
```

**注意事项：**
- 字符串字面量支持使用双引号、单引号和反引号。
- 反引号内的转义字符不会被解析，直接输出。

## 列表字面量(ListLiteral)
表示列表值的表达式节点。

**语法定义：**
```
ListLiteral ::= "[" Expression ("," Expression)* "]"
```

**示例：**
```ghost
[1, 2, 3, 4, 5];
["apple", "banana", "orange"];
```

**注意事项：**
- 列表字面量的每个元素的类型必须相同。

## 标识符(Identifier)
表示变量名或函数名的表达式节点。

**语法定义：**
```
Identifier ::= [a-zA-Z_][a-zA-Z0-9_]*
```

**示例：**
```ghost
x;
myVariable;
add;
```

## 前缀表达式(PrefixExpression)
表示一元操作符表达式，如负号、逻辑非、按位取反等。

**语法定义：**
```
PrefixExpression ::= ("-" | "!" | "~") Expression
```

**示例：**
```ghost
-x;
!true;
~i;
```

## 中缀表达式(InfixExpression)
表示二元操作符表达式，如加法、减法、比较、位运算等。

**语法定义：**
```
InfixExpression ::= (Expression Operator Expression) | (Expression "[" Expression "]")
Operator ::= "+" | "-" | "*" | "/" | "%" | "==" | "!=" | "<" | ">" | "<=" | ">=" | "&&" | "||" | "&" | "|" | "^" | "<<" | ">>"
```

**示例：**
```ghost
x + 10;
y * 2;
a > b;
x == 10;
```

## 分组表达式(GroupExpression)
用于改变运算优先级的括号表达式。

**语法定义：**
```
GroupExpression ::= "(" Expression ")"
```

**示例：**
```ghost
(5 + 3) * 2;
```

## 变量初始化表达式(VarInitializationExpression)
用于变量初始化的表达式。

**语法定义：**
```
VarInitializationExpression ::= ("var" | "const") LvalueList "=" Expression
LvalueList ::= Identifier | ListExpression
```

**示例：**
```ghost
var x = 20;
const PI = 3.14159;
var [a, b] = [1, 2];
const [x, y, z] = [10, 20, 30];
```

**注意事项：**
- 当使用列表表达式作为左值时，进行解构赋值
- 解构赋值要求右侧必须是列表表达式
- 左右两侧的元素数量必须一致

## 变量赋值表达式(VarAssignmentExpression)
用于给已声明的变量重新赋值的表达式。

**语法定义：**
```
VarAssignmentExpression ::= LvalueList "=" Expression
LvalueList ::= Identifier | IndexExpression | NamespaceAccessExpression | MemberAccessExpression | ListExpression
```

**示例：**
```ghost
x = 30;
a[0] = 40;
[a, b] = [1, 2];
[x, y, z] = [10, 20, 30];
[l[0], b, A::a] = [4, 5, 6];

import math;
math.PI = 3.14;
```

**注意事项：**
- 当使用列表表达式作为左值时，进行解构赋值
- 解构赋值要求右侧必须是列表表达式
- 左右两侧的元素数量必须一致
- 支持索引表达式、命名空间访问表达式和成员访问表达式作为左值

## 复合赋值表达式(CompoundAssignmentExpression)
用于复合赋值操作的表达式。

**语法定义：**
```
CompoundAssignmentExpression ::= Identifier CompoundOperator Expression
CompoundOperator ::= "+=" | "-=" | "*=" | "/=" | "%=" | "&=" | "|=" | "^=" | "<<=" | ">>="
```

**示例：**
```ghost
a += 3;
b -= 2;
c *= 4;
d /= 2;
```

## 前缀自增 / 自减表达式(PrefixUnaryIncDecExpression)
用于前缀自增 / 自减表达式

**语法定义：**
```
PrefixUnaryIncDecExpression ::= ("++" | "--") Lvalue
```

**示例：**
```ghost
++x;
--y;
```

## 后缀自增 / 自减表达式(PostfixUnaryIncDecExpression)
用于后缀自增 / 自减表达式

**语法定义：**
```
PostfixUnaryIncDecExpression ::= Lvalue ("++" | "--")
```

**示例：**
```ghost
x++;
y--;
```

## 块表达式(BlockExpression)
用于包含多个表达式的代码块。

**语法定义：**
```
BlockExpression ::= "{" Statement* "}"
```

**示例：**
```ghost
{
  var c = 10;
  c + 5;
};
```

**注意事项：**
- 块表达式有自己的作用域，其中声明的变量在块表达式结束后会被销毁。
- 块表达式的返回值是最后一个语句的返回值。

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

### 解包参数(Unpack Argument)
用于将列表中的元素解包为多个参数传递给函数。

**语法定义：**
```
UnpackArgument ::= Expression "..."
```

**示例：**
```ghost
import "math";
var args = [1, 3, 5, 7, 9];
math.sum(args...);  // 等价于 math.sum(1, 3, 5, 7, 9)

import "fmt";
fmt.print("Hello", ["World", "!"]...);  // 等价于 fmt.print("Hello", "World", "!")
```

**注意事项：**
- 解包参数只能用于函数调用中
- 解包参数后面必须跟随一个列表表达式
- 解包参数会将列表中的每个元素作为独立的参数传递给函数
- 解包参数支持与普通参数混合使用
- 解包参数支持内置函数和自定义函数

## 索引表达式(IndexExpression)
表示列表索引访问的表达式节点。

**语法定义：**
```
IndexExpression ::= Expression "[" Expression "]"
```

**示例：**
```ghost
list[0];
matrix[1][2];
```

## 命名空间访问表达式(NamespaceAccessExpression)
用于访问命名空间中的成员的表达式节点。

**语法定义：**
```
NamespaceAccessExpression ::= Identifier "::" Identifier
```

**示例：**
```ghost
Array::array;
Utils::add;
```

**注意事项：**
- 命名空间访问表达式用于访问命名空间中定义的变量或函数
- 命名空间访问表达式可以作为左值使用，用于修改命名空间中的变量

## 成员访问表达式(MemberAccessExpression)
用于访问模块中的成员的表达式节点。

**语法定义：**
```
MemberAccessExpression ::= Identifier "." Identifier
```

**示例：**
```ghost
import math;
math.PI;
math.sqrt(2);
```

**注意事项：**
- 成员访问表达式用于访问导入模块中定义的变量或函数
- 成员访问表达式可以作为左值使用，用于修改模块中的变量

## 范围表达式(RangeExpression)
用于创建整数范围的表达式节点。

**语法定义：**
```
RangeExpression ::= Expression ".." Expression
```

**示例：**
```ghost
1..5;
(1+2)..(10-3);
```

**注意事项：**
- 范围运算符是闭区间，包含起始值和结束值
- 起始值和结束值必须是整数类型
- 如果起始值大于结束值，返回空列表

## 包含检查表达式(ContainsExpression)
用于检查元素是否存在于容器中的表达式节点。

**语法定义：**
```
ContainsExpression ::= Expression "contains" Expression
```

**示例：**
```ghost
[1, 2, 3] contains 2;
[1, 2, 3] contains 4;
"hello" contains "ell";
```

**注意事项：**
- 左操作数必须是列表或字符串类型
- 对于列表，检查元素是否存在于列表中
- 对于字符串，检查子字符串是否存在
- 不支持检查子列表是否存在