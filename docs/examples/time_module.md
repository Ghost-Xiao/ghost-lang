# 时间模块使用示例

以下是使用 Ghost Language 时间模块的完整示例：

## 基础使用

```ghost
import "time";

// 获取当前时间
var now = time.Time();
println("当前时间: ", now.format());
println("年: ", now.year());
println("月: ", now.month());
println("日: ", now.day());
println("小时: ", now.hour());
println("分钟: ", now.minute());
println("秒: ", now.second());
```

## 使用时间常量

```ghost
import "time";

// 使用时间常量创建 Duration
var oneSec = time.Duration(time.SECOND);
var oneMin = time.Duration(time.MINUTE);
var oneHour = time.Duration(time.HOUR);
var halfHour = time.Duration(time.MINUTE * 30);
var dayAndHalf = time.Duration(time.HOUR * 36);

println("1 秒: ", oneSec.second());
println("1 分钟: ", oneMin.minute());
println("1 小时: ", oneHour.hour());
println("30 分钟: ", halfHour.minute());
```

## 时间运算

```ghost
import "time";

var now = time.Time();
println("现在: ", now.format());

// 计算未来时间
var oneHourLater = now + time.Duration(time.HOUR);
println("1 小时后: ", oneHourLater.format());

// 计算过去时间
var twoMinutesAgo = now - time.Duration(time.MINUTE * 2);
println("2 分钟前: ", twoMinutesAgo.format());

// 计算时间差
var diff = oneHourLater - now;
println("时间差: ", diff.hour(), " 小时");
```

## 时间格式化

```ghost
import "time";

var now = time.Time();

// 默认格式
println("默认格式: ", now.format());

// 自定义格式
println("只显示日期: ", now.format("yyyy-MM-dd"));
println("只显示时间: ", now.format("HH:mm:ss"));
println("完整格式: ", now.format("yyyy年MM月dd日 HH时mm分ss秒"));
```

## 从字符串解析时间

```ghost
import "time";

// 从自定义格式的字符串解析时间
var birthday = time.Time("yyyy-MM-dd", "1990-01-15");
println("生日: ", birthday.format("yyyy-MM-dd"));
println("生日是星期: ", birthday.weekday());
println("生日是一年中的第几天: ", birthday.yearday());
```

## 设置时区

```ghost
import "time";

var now = time.Time();

// 设置为 UTC+8 (中国时区)
now.setZone(8);
println("UTC+8 时间: ", now.format());

// 设置为 UTC-5 (纽约时区)
now.setZone(-5);
println("UTC-5 时间: ", now.format());
```

## 暂停执行 (sleep)

```ghost
import "time";

println("开始计时...");
var start = time.Time();

// 暂停 2 秒
time.sleep(time.Duration(time.SECOND * 2));

var end = time.Time();
var duration = end - start;
println("暂停了 ", duration.second(), " 秒");
```

## 时间比较

```ghost
import "time";

var past = time.Time(2024, 1, 1, 0, 0, 0);
var present = time.Time();
var future = time.Time(2027, 1, 1, 0, 0, 0);

println("过去 < 现在: ", (past < present));
println("未来 > 现在: ", (future > present));
println("过去 == 未来: ", (past == future));
```

## 倒计时示例

```ghost
import "time";

// 简单的倒计时
func countdown(seconds) {
    for var i = seconds; i > 0; i-- {
        println(i, " 秒...");
        time.sleep(time.Duration(time.SECOND));
    };
    println("倒计时结束！");
};

countdown(5);
```

## 性能计时

```ghost
import "time";

// 测量函数执行时间
func measureTime(funcToMeasure) {
    var start = time.Time();
    funcToMeasure();
    var end = time.Time();
    var duration = end - start;
    return duration;
};

// 一个耗时的计算
func expensiveCalculation() {
    var sum = 0;
    var i = 0;
    for var [sum, i] = [0, 0]; i < 100000; {sum += i; i++} {
    };
};

var timeTaken = measureTime(expensiveCalculation);
println("计算耗时: ", timeTaken.millisecond(), " 毫秒");
```

这个示例展示了 Ghost Language 时间模块的以下特性：
- Time 类的创建和使用
- Duration 类和时间常量
- 时间的加减运算
- 时间格式化和解析
- 时区设置
- sleep 函数
- 时间比较
- 实际应用场景