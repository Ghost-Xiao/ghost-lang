# 数学模块使用示例

以下是使用 Ghost Language 数学模块的完整示例：

## 数学常量

```ghost
import "math";

println("圆周率 π = ", math.PI);
println("自然对数的底 e = ", math.E);
println("2π = ", math.TAU);
```

## 基础运算

```ghost
import "math";

// 绝对值
println("|-5| = ", math.abs(-5));
println("|-3.14| = ", math.abs(-3.14));

// 取整
println("floor(3.7) = ", math.floor(3.7));
println("floor(-3.7) = ", math.floor(-3.7));
println("ceil(3.2) = ", math.ceil(3.2));
println("ceil(-3.2) = ", math.ceil(-3.2));
println("round(3.5) = ", math.round(3.5));
println("round(3.4) = ", math.round(3.4));

// 保留小数位数
println("floor(3.14159, 2) = ", math.floor(3.14159, 2));
println("ceil(3.14159, 2) = ", math.ceil(3.14159, 2));
println("round(3.14159, 2) = ", math.round(3.14159, 2));
```

## 三角函数

```ghost
import "math";

println("sin(π/2) = ", math.sin(math.PI / 2));
println("cos(π) = ", math.cos(math.PI));
println("tan(π/4) = ", math.tan(math.PI / 4));

println("asin(1) = ", math.asin(1));
println("acos(-1) = ", math.acos(-1));
println("atan(1) = ", math.atan(1));
```

## 平方根和对数

```ghost
import "math";

// 平方根
println("sqrt(4) = ", math.sqrt(4));
println("sqrt(2) = ", math.sqrt(2));
println("sqrt(25) = ", math.sqrt(25));

// 对数
println("log(2, 8) = ", math.log(2, 8));
println("lg(100) = ", math.lg(100));
println("ln(e) = ", math.ln(math.E));
```

## 最大值、最小值、求和、乘积

```ghost
import "math";

// 最小值和最大值
println("min(5, 3, 7, 1) = ", math.min(5, 3, 7, 1));
println("max(5, 3, 7, 1) = ", math.max(5, 3, 7, 1));
println("min(-1, -5, -3) = ", math.min(-1, -5, -3));
println("max(0.5, 0.8, 0.3) = ", math.max(0.5, 0.8, 0.3));

// 求和和乘积
println("sum(1, 2, 3, 4, 5) = ", math.sum(1, 2, 3, 4, 5));
println("product(1, 2, 3, 4, 5) = ", math.product(1, 2, 3, 4, 5));

// 使用数组解包
var nums = [10, 20, 30, 40, 50];
println("sum(nums...) = ", math.sum(nums...));
println("product(nums...) = ", math.product(nums...));
```

## 统计函数

```ghost
import "math";

// 平均值（mean）
println("mean(1, 2, 3, 4, 5) = ", math.mean(1, 2, 3, 4, 5));

// 中位数（median）
println("median(1, 3, 5, 7, 9) = ", math.median(1, 3, 5, 7, 9));
println("median(1, 3, 5, 7) = ", math.median(1, 3, 5, 7));

// 方差（variance）
println("variance(1, 2, 3, 4, 5) = ", math.variance(1, 2, 3, 4, 5));

// 标准差（stdDev）
println("stdDev(1, 2, 3, 4, 5) = ", math.stdDev(1, 2, 3, 4, 5));
```

## 随机数

```ghost
import "math";

// 生成 0-1 之间的随机浮点数
println("随机浮点数: ", math.rand());
println("随机浮点数: ", math.rand());
println("随机浮点数: ", math.rand());

// 生成指定范围的随机整数
println("1-10 随机整数: ", math.randInt(1, 10));
println("1-10 随机整数: ", math.randInt(1, 10));
println("100-200 随机整数: ", math.randInt(100, 200));
```

## 实用工具：圆形计算

```ghost
import "math";

// 计算圆的周长
func circleCircumference(radius) {
    return 2 * math.PI * radius;
};

// 计算圆的面积
func circleArea(radius) {
    return math.PI * radius * radius;
};

println("半径 5 的圆：");
println("周长 = ", circleCircumference(5));
println("面积 = ", circleArea(5));
```

## 实用工具：温度转换

```ghost
import "math";

// 摄氏温度转华氏温度
func celsiusToFahrenheit(c) {
    return c * 9 / 5 + 32;
};

// 华氏温度转摄氏温度
func fahrenheitToCelsius(f) {
    return (f - 32) * 5 / 9;
};

println("0°C = ", celsiusToFahrenheit(0), "°F");
println("100°C = ", celsiusToFahrenheit(100), "°F");
println("32°F = ", fahrenheitToCelsius(32), "°C");
println("212°F = ", fahrenheitToCelsius(212), "°C");
```

## 实用工具：距离计算

```ghost
import "math";

// 计算两点之间的欧氏距离
func distance(x1, y1, x2, y2) {
    var dx = x2 - x1;
    var dy = y2 - y1;
    return math.sqrt(dx * dx + dy * dy);
};

println("点 (0,0) 到 (3,4) 的距离 = ", distance(0, 0, 3, 4));
println("点 (1,1) 到 (4,5) 的距离 = ", distance(1, 1, 4, 5));
```

## 实用工具：数组统计

```ghost
import "math";

// 分析数组的统计信息
func analyzeArray(arr) {
    println("数组: ", arr);
    println("数量: ", len(arr));
    println("最小值: ", math.min(arr...));
    println("最大值: ", math.max(arr...));
    println("总和: ", math.sum(arr...));
    println("乘积: ", math.product(arr...));
    println("平均值: ", math.mean(arr...));
    println("中位数: ", math.median(arr...));
    println("方差: ", math.variance(arr...));
    println("标准差: ", math.stdDev(arr...));
};

var data = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
analyzeArray(data);
```

## 掷骰子游戏

```ghost
import "math";

// 掷骰子函数
func rollDice(sides, times) {
    var results = [];
    for var i = 0; i < times; i++ {
        var roll = math.randInt(1, sides);
        results.append(roll);
    };
    return results;
};

// 模拟掷骰子

println("掷 3 个 6 面骰子:");
var rolls = rollDice(6, 3);
println("结果: ", rolls);
println("总和: ", math.sum(rolls...));
println("最大: ", math.max(rolls...));
println("最小: ", math.min(rolls...));

println("\n掷 5 个 10 面骰子:");
var rolls2 = rollDice(10, 5);
println("结果: ", rolls2);
println("总和: ", math.sum(rolls2...));
```

## 数值积分（简单示例）

```ghost
import "math";

// 计算函数在区间 [a, b] 的定积分（矩形法）
func integrate(f, a, b, steps) {
    var dx = (b - a) / steps;
    var total = 0;
    for var i = 0; i < steps; i++ {
        var x = a + i * dx;
        total = total + f(x) * dx;
    };
    return total;
};

// 计算函数值
func f(x) {
    return x * x;
};

println("∫x² dx 从 0 到 2 = ", integrate(f, 0, 2, 1000));
println("理论值 = 8/3 ≈ ", (8 / 3));
```

## 简单的统计分析器

```ghost
import "math";

// 成绩分析器
class GradeAnalyzer {
    var grades = [];

    func init() {
        this.grades = [];
    };

    func addGrade(grade) {
        this.grades.append(grade);
    };

    func getAverage() {
        if len(this.grades) == 0 {
            return 0;
        };
        return math.mean(this.grades...);
    };

    func getHighest() {
        if len(this.grades) == 0 {
            return 0;
        };
        return math.max(this.grades...);
    };

    func getLowest() {
        if len(this.grades) == 0 {
            return 0;
        };
        return math.min(this.grades...);
    };

    func getTotal() {
        if len(this.grades) == 0 {
            return 0;
        };
        return math.sum(this.grades...);
    };

    func getStandardDeviation() {
        if len(this.grades) == 0 {
            return 0;
        };
        return math.stdDev(this.grades...);
    };

    func printReport() {
        println("=== 成绩报告 ===");
        println("成绩数量: ", len(this.grades));
        println("成绩列表: ", this.grades);
        println("最高分: ", this.getHighest());
        println("最低分: ", this.getLowest());
        println("平均分: ", this.getAverage());
        println("总分: ", this.getTotal());
        println("标准差: ", this.getStandardDeviation());
    };
};

// 使用成绩分析器
var analyzer = GradeAnalyzer();
analyzer.addGrade(85);
analyzer.addGrade(92);
analyzer.addGrade(78);
analyzer.addGrade(88);
analyzer.addGrade(95);
analyzer.addGrade(76);
analyzer.addGrade(89);
analyzer.addGrade(91);
analyzer.printReport();
```

## 简单的几何计算器

```ghost
import "math";

class GeometryCalculator {
    // 圆
    func circleArea(radius) {
        return math.PI * radius * radius;
    };

    func circleCircumference(radius) {
        return 2 * math.PI * radius;
    };

    // 矩形
    func rectangleArea(width, height) {
        return width * height;
    };

    func rectanglePerimeter(width, height) {
        return 2 * (width + height);
    };

    // 三角形
    func triangleArea(base, height) {
        return base * height / 2;
    };

    func trianglePerimeter(a, b, c) {
        return a + b + c;
    };

    // 球体
    func sphereVolume(radius) {
        return 4 / 3 * math.PI * radius * radius * radius;
    };

    func sphereSurfaceArea(radius) {
        return 4 * math.PI * radius * radius;
    };

    // 勾股定理求斜边
    func hypotenuse(a, b) {
        return math.sqrt(a * a + b * b);
    };
};

var calc = GeometryCalculator();

println("--- 圆形 ---");
println("半径 5: 面积 = ", calc.circleArea(5), " 周长 = ", calc.circleCircumference(5));

println("\n--- 矩形 ---");
println("3x4: 面积 = ", calc.rectangleArea(3, 4), " 周长 = ", calc.rectanglePerimeter(3, 4));

println("\n--- 三角形 ---");
println("底 5 高 4: 面积 = ", calc.triangleArea(5, 4));
println("边 3,4,5: 周长 = ", calc.trianglePerimeter(3, 4, 5));

println("\n--- 球体 ---");
println("半径 3: 体积 = ", calc.sphereVolume(3), " 表面积 = ", calc.sphereSurfaceArea(3));

println("\n--- 勾股定理 ---");
println("3,4 的斜边 = ", calc.hypotenuse(3, 4));
```

这个示例展示了 Ghost Language 数学模块的以下特性：

- 数学常量（PI, E, TAU）
- 基础运算（abs, floor, ceil, round）
- 三角函数（sin, cos, tan, asin, acos, atan）
- 平方根和对数（sqrt, log, lg, ln）
- 聚合函数（min, max, sum, product）
- 统计函数（mean, median, variance, stdDev）
- 随机数（rand, randInt）
- 实际应用场景

