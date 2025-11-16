# Ghost Language

一个用 Go 编写的简洁、易学的解释型编程语言实现，专注于教学和学习编程语言设计。

## 快速开始

### 安装

有两种方式可以使用 Ghost：

1. 使用 PowerShell 脚本构建（推荐）：
```bash
# 克隆项目
git clone https://github.com/Ghost-Xiao/ghost-lang.git
cd ghost-lang

# 使用 PowerShell 脚本构建（Windows）
.\build.ps1

# 构建后的可执行文件位于 bin/ 目录下
```

2. 使用 Go 工具直接构建：
```bash
go build -o ghost cmd/ghost/main.go
```

### 运行 REPL

```bash
./ghost repl
# 或
./ghost -r
```

### 执行脚本文件

```bash
./ghost run script.gh
```

## 语言语法说明

Ghost Lang 支持多种语法结构，包括表达式、语句和控制结构。以下是基于 AST 节点的详细语法说明。

### 程序(Program)

程序是 AST 的根节点，表示整个源代码文件，由一系列按顺序执行的语句组成。

**语法定义：**
```
Program ::= Statement*
```

**示例：**
```ghost
var x = 10;
println(x);
```

### 表达式(Expression)

#### 整数字面量(IntegerLiteral)
表示整数值的表达式节点。

**语法定义：**
```
IntegerLiteral ::= [0-9]+
```

**示例：**
```ghost
42;
-10;
```

#### 浮点数字面量(FloatLiteral)
表示浮点数值的表达式节点。

**语法定义：**
```
FloatLiteral ::= [0-9]+ "." [0-9]+
```

**示例：**
```ghost
3.14;
-2.718;
```

#### 布尔字面量(BooleanLiteral)
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

#### 空字面量(NullLiteral)
表示空值的表达式节点。

**语法定义：**
```
NullLiteral ::= "null"
```

**示例：**
```ghost
null;
```

#### 字符串字面量(StringLiteral)
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

#### 列表字面量(ListLiteral)
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

#### 标识符(Identifier)
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

#### 前缀表达式(PrefixExpression)
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

#### 中缀表达式(InfixExpression)
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

#### 分组表达式(GroupExpression)
用于改变运算优先级的括号表达式。

**语法定义：**
```
GroupExpression ::= "(" Expression ")"
```

**示例：**
```ghost
(5 + 3) * 2;
```

#### 变量初始化表达式(VarInitializationExpression)
用于变量初始化的表达式。

**语法定义：**
```
VarInitializationExpression ::= ("var" | "const") Identifier "=" Expression
```

**示例：**
```ghost
var x = 20;
const PI = 3.14159;
```

#### 变量赋值表达式(VarAssignmentExpression)
用于给已声明的变量重新赋值的表达式。

**语法定义：**
```
VarAssignmentExpression ::= Lvalue "=" Expression
Lvalue ::= Identifier | IndexExpression
```

**示例：**
```ghost
x = 30;
a[0] = 40;
```

#### 复合赋值表达式(CompoundAssignmentExpression)
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

#### 前缀自增 / 自减表达式(PrefixUnaryIncDecExpression)
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

#### 后缀自增 / 自减表达式(PostfixUnaryIncDecExpression)
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

#### 块表达式(BlockExpression)
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

#### 条件表达式(IfExpression)
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

#### 函数调用表达式(CallExpression)
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

#### 索引表达式(IndexExpression)
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

### 语句(Statement)

#### 表达式语句(ExpressionStatement)
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

#### For 循环语句(ForStatement)
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

#### 函数声明语句(FunctionDeclarationStatement)
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

#### 返回语句(ReturnStatement)
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

## 代码示例

```ghost
func fib(n) {
    if n <= 1 {
        return n;
    };
    return fib(n - 1) + fib(n - 2);
};

println(fib(10)); // print 55
```

## 项目架构

```
main.go → cli包 → 词法分析器(lexer)
               → 语法分析器(parser) → AST抽象语法树
               → 解释执行器(evaluator) → 运行时对象(object)
                                    → 执行环境(frame)
               → REPL交互模块
```

## 如何贡献

我们欢迎任何形式的贡献！无论是报告 bug、提出新功能建议，还是提交代码改进。

1. Fork 本项目
2. 创建您的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交您的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启一个 Pull Request

## 测试

项目包含单元测试，可以使用以下命令运行：

```bash
go test ./...
```

这将运行所有包中的测试，包括词法分析器、语法分析器和解释执行器的测试。