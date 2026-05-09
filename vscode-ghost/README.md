# Ghost 语言 VS Code 插件

这个插件为 VS Code 提供 Ghost 编程语言的支持。

## 功能特点

- **语法高亮**：为 Ghost 语言文件 (.gh) 提供语法高亮
- **代码补全**：支持关键字、内置函数、错误类和模块的代码补全
- **错误检查**：检查常见的语法问题，如未闭合的字符串和不匹配的括号

## 安装方法

1. 在 VS Code 中打开 vscode-ghost 文件夹
2. 运行 `npm install` 安装依赖
3. 运行 `npm run compile` 编译 TypeScript 代码
4. 按 F5 启动扩展开发主机窗口

## 使用说明

- 创建一个扩展名为 `.gh `的文件
- 开始编写 Ghost 代码吧！

## 示例代码

```ghost
// Hello World 示例
println("你好，Ghost！");

// 变量
var x = 42;
const PI = 3.14159;

// 函数
func add(a, b=0) {
    return a + b;
};

var result = add(5, 3);
println(result);

// 循环
foreach var i in 1..10 {
    println(i);
}

// 类
class Person {
    var name = "";
    var age = 0;

    func init(name, age) {
        this.name = name;
        this.age = age;
    };

    func greet() {
        println("你好，我叫" + this.name + "，今年" + this.age + "岁。");
    };
};

var alice = Person("爱丽丝", 30);
alice.greet();

// 错误处理
try {
    var x = 10 / 0;
} catch e {
    println("错误: " + e.message);
};
```

