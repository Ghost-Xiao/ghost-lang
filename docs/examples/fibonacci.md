# 斐波那契数列示例

以下是一个使用 Ghost Language 实现的斐波那契数列函数示例：

```ghost
func fib(n) {
    if n <= 1 {
        return n;
    };
    return fib(n - 1) + fib(n - 2);
};

println(fib(10)); // print 55
```

这个示例展示了 Ghost Language 的以下特性：
- 函数定义和递归调用
- 条件表达式（if 语句）
- 返回语句
- 函数调用和参数传递
- 打印输出功能