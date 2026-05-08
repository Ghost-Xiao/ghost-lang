# 异常处理示例

以下是使用 Ghost Language 异常处理的完整示例：

## 基础 try-catch

```ghost
// 简单的 try-catch 示例
func divide(a, b) {
    if b == 0 {
        throw Error("不能除以零！");
    };
    return a / b;
};

try {
    var result = divide(10, 2);
    println("结果: ", result);
} catch error {
    println("捕获到错误: ", error);
};

try {
    var result = divide(10, 0);
    println("结果: ", result);
} catch error {
    println("捕获到错误: ", error);
};
```

## try-catch-finally

```ghost
// 带 finally 的异常处理
func readFile(filename) {
    println("尝试打开文件: ", filename);
    if filename == "" {
        throw Error("文件名不能为空");
    };
    println("文件打开成功");
};

try {
    readFile("data.txt");
} catch error {
    println("错误: ", error);
} finally {
    println("清理资源...");
};

try {
    readFile("");
} catch error {
    println("错误: ", error);
} finally {
    println("清理资源...");
};
```

## 不同类型的异常

```ghost
import "math";

// 使用不同类型的异常
func checkNumber(n) {
    if !(isInstanceOf(n, Int) || isInstanceOf(n, Float)) {
        throw TypeError("必须是数字类型");
    };
    if n < 0 {
        throw MathError("不能处理负数");
    };
    if n > 1000 {
        throw OperationError("数值超出范围");
    };
    return math.sqrt(n);
};

try {
    var result = checkNumber(16);
    println("结果: ", result);
} catch error {
    if isInstanceOf(error, TypeError) {
        println("类型错误: ", error);
    } else if isInstanceOf(error, MathError) {
        println("数学错误: ", error);
    } else
        println("其他错误: ", error);
};
```

## 异常传递

```ghost
// 异常在函数调用链中传递
func level3() {
    println("进入 level3");
    throw Error("level3 出错了！");
};

func level2() {
    println("进入 level2");
    level3();
    println("离开 level2");  // 这行不会执行
};

func level1() {
    println("进入 level1");
    level2();
    println("离开 level1");  // 这行不会执行
};

try {
    level1();
} catch error {
    println("在顶层捕获: ", error);
};
```

## 在 finally 中确保清理

```ghost
// 使用 finally 确保资源释放
class Resource {
    var name = "";

    func init(name) {
        this.name = name;
        println("资源 " + name + " 已打开");
    };

    func close() {
        println("资源 " + this.name + " 已关闭");
    };
};

func useResource(name) {
    var res = Resource(name);
    try {
        if name == "bad" {
            throw Error("坏资源！");
        };
        println("使用资源: ", name);
    } finally {
        res.close();
    };
};

useResource("good");
useResource("bad");
```

## 自定义异常类

```ghost
// 自定义异常类
class DatabaseError extends Error {
    var sql = "";

    func init(message, sql) {
        super.init(message);
        this.sql = sql;
    };

    func getSql() {
        return this.sql;
    };
};

func executeQuery(sql) {
    if sql == "" {
        throw DatabaseError("SQL 为空", sql);
    };
    if sql == "BAD SQL" {
        throw DatabaseError("SQL 执行失败", sql);
    };
    println("执行 SQL: ", sql);
};

try {
    executeQuery("SELECT * FROM users");
} catch error {
    println("数据库错误: ", error);
    println("问题 SQL: " + error.getSql());
};

try {
    executeQuery("BAD SQL");
} catch error {
    println("数据库错误: ", error);
    println("问题 SQL: " + error.getSql());
};

try {
    executeQuery("");
} catch error {
    println("数据库错误: ", error);
    println("问题 SQL: ", error.getSql());
};
```

## 验证用户输入

```ghost
// 实际应用：验证用户输入
func validateEmail(email) {
    if !isInstanceOf(email, String) {
        throw TypeError("邮箱必须是字符串");
    };
    if email == "" {
        throw ArgumentError("邮箱不能为空");
    };
    // 简单的验证（实际应用中应该更完善）
    if email.indexOf("@") == -1 {
        throw ArgumentError("邮箱格式不正确");
    };
    println("邮箱验证通过: " + email);
};

var emails = [
    "user@example.com",
    "",
    123,
    "invalid-email"
];

var i = 0;
foreach var email in emails {
    try {
        validateEmail(email);
    } catch error {
        println("错误: ", error);
    };
};
```

## 安全计算

```ghost
import "math";

// 安全的数学计算
func safeCalculate(operation, a, b) {
    try {
        if operation == "add" {
            return a + b;
        } else if operation == "subtract" {
            return a - b;
        } else if operation == "multiply" {
            return a * b;
        } else if operation == "divide" {
            if b == 0 {
                throw MathError("除数不能为零");
            };
            return a / b;
        } else if operation == "sqrt" {
            if a < 0 {
                throw MathError("不能对负数开平方");
            };
            return math.sqrt(a);
        } else {
            throw OperationError("未知操作: " + operation);
        };
    } catch error {
        if isInstanceOf(error, MathError) {
            println("数学错误: ", error);
            return 0;
        };
        if isInstanceOf(error, OperationError) {
            println("操作错误: ", error);
            return 0;
        };
        if isInstanceOf(error, Error) {
            println("其他错误: ", error);
            return 0;
        };
        println("未知错误: ", error);
        return 0;
    };
};

println("2 + 3 = ", safeCalculate("add", 2, 3));
println("5 - 3 = ", safeCalculate("subtract", 5, 3));
println("4 * 5 = ", safeCalculate("multiply", 4, 5));
println("10 / 2 = ", safeCalculate("divide", 10, 2));
println("10 / 0 = ", safeCalculate("divide", 10, 0));
println("sqrt(16) = ", safeCalculate("sqrt", 16, 0));
println("sqrt(-4) = ", safeCalculate("sqrt", -4, 0));
println("unknown = ", safeCalculate("unknown", 5, 5));
```

## 文件操作错误处理

```ghost
import "io";

// 安全的文件操作
func safe_readFile(path) {
    try {
        if !io.exists(path) {
            throw OperationError("文件不存在: " + path);
        };
        if !io.isFile(path) {
            throw OperationError("不是文件: " + path);
        };
        var content = io.read(path);
        println("读取成功，文件大小: " + io.fileSize(path) + " 字节");
        return content;
    } catch error {
        println("文件操作失败: ", error);
        return "";
    };
};

var content = safe_readFile("existing_file.txt");
if content != "" {
    println("文件内容: " + content);
};

var nonexistent = safe_readFile("nonexistent_file.txt");
```

## 重试机制

```ghost
import "time";

// 带重试的函数
func retry(operation, maxRetries) {
    var attempts = 0;
    for ...; attempts < maxRetries; attempts++ {
        try {
            println("尝试 ", attempts, "/", maxRetries);
            return operation();
        } catch error {
            if attempts == maxRetries {
                println("所有尝试都失败了！");
                throw error;
            };
            println("失败，等待后重试...");
            time.sleep(time.Duration(time.SECOND));
        };
    };
};

// 一个可能失败的操作
var failCount = 0;
func flakyOperation() {
    failCount = failCount + 1;
    if failCount < 3 {
        throw Error("临时错误");
    };
    println("操作成功！");
    return "成功结果";
};

try {
    var result = retry(flakyOperation, 5);
    println("最终结果: " + result);
} catch error {
    println("最终失败: ", error);
};
```

这个示例展示了 Ghost Language 异常处理的以下特性：
- try-catch 语句
- try-catch-finally 语句
- throw 语句抛出异常
- 不同类型的异常处理
- 异常在函数调用链中传递
- finally 块确保资源清理
- 自定义异常类
- 实际应用场景