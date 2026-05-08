# 面向对象编程示例

以下是使用 Ghost Language 面向对象编程的完整示例：

## 基础类定义

```ghost
// 定义一个简单的 Person 类
class Person {
    var [name, age] = ["", 0];

    // 构造函数
    func init(name, age) {
        this.name = name;
        this.age = age;
    };

    // 方法
    func greet() {
        return "你好，我是" + this.name + "，今年" + String(this.age) + "岁。";
    };

    func haveBirthday() {
        this.age = this.age + 1;
        println(this.name + "过生日了！现在" + String(this.age) + "岁。");
    };
};

// 创建实例
var alice = Person("Alice", 25);
println(alice.greet());
alice.haveBirthday();
```

## 继承

```ghost
// 基类
class Animal {
    var name = "";

    func init(name) {
        this.name = name;
    };

    func makeSound() {
        println("动物发出声音...");
    };
};

// 继承自 Animal 的 Dog 类
class Dog extends Animal {
    var breed = "";

    func init(name, breed) {
        // 调用父类构造函数
        super.init(name);
        this.breed = breed;
    };

    // 重写方法
    func makeSound() {
        println(this.name + " 汪汪叫！");
    };

    func fetch() {
        println(this.name + " 去捡球！");
    };
};

// 继承自 Animal 的 Cat 类
class Cat extends Animal {
    var color = "";

    func init(name, color) {
        super.init(name);
        this.color = color;
    };

    func makeSound() {
        println(this.name + " 喵喵叫！");
    };

    func climb() {
        println(this.name + " 爬上了树！");
    };
};

// 使用继承
var dog = Dog("旺财", "金毛");
dog.makeSound();
dog.fetch();

var cat = Cat("咪咪", "橙色");
cat.makeSound();
cat.climb();
```

## 封装：计算器类

```ghost
class Calculator {
    var result = 0;

    func init() {
        this.result = 0;
    };

    func add(n) {
        this.result = this.result + n;
    };

    func subtract(n) {
        this.result = this.result - n;
    };

    func multiply(n) {
        this.result = this.result * n;
    };

    func divide(n) {
        if n != 0 {
            this.result = this.result / n;
        } else {
            println("不能除以零！");
        }
    };

    func reset() {
        this.result = 0;
    };

    func getResult() {
        return this.result;
    };
};

// 使用计算器
var calc = Calculator();
calc.add(10);
calc.multiply(2);
calc.subtract(5);
println("计算结果: ", calc.getResult());
calc.reset();
println("重置后: ", calc.getResult());
```

## 多态示例

```ghost
class Shape {
    func init() {
    };

    func area() {
        return 0;
    };

    func perimeter() {
        return 0;
    };

    func describe() {
        return "这是一个形状";
    };
};

class Circle extends Shape {
    var radius = 0;

    func init(radius) {
        super.init();
        this.radius = radius;
    };

    func area() {
        return 3.14159 * this.radius * this.radius;
    };

    func perimeter() {
        return 2 * 3.14159 * this.radius;
    };

    func describe() {
        return "这是一个半径为" + String(this.radius) + "的圆";
    };
};

class Rectangle extends Shape {
    var [width, height] = [0, 0];

    func init(width, height) {
        super.init();
        this.width = width;
        this.height = height;
    };

    func area() {
        return this.width * this.height;
    };

    func perimeter() {
        return 2 * (this.width + this.height);
    };

    func describe() {
        return "这是一个" + String(this.width) + "x" + String(this.height) + "的矩形";
    };
};

// 多态使用
func printShapeInfo(shape) {
    println(shape.describe());
    println("面积: ", shape.area());
    println("周长: ", shape.perimeter());
    println("---");
};

var circle = Circle(5);
var rectangle = Rectangle(4, 6);

printShapeInfo(circle);
printShapeInfo(rectangle);
```

## 银行账户示例

```ghost
class BankAccount {
    var owner = "";
    var balance = 0;
    var transactionHistory = [];

    func init(owner, balance) {
        this.owner = owner;
        this.balance = balance;
        this.transactionHistory = [];
    };

    func deposit(amount) {
        if amount > 0 {
            this.balance = this.balance + amount;
            this.transactionHistory.append(["存款", amount]);
            println("存款成功：+", amount);
        }
    };

    func withdraw(amount) {
        if amount > 0 && amount <= this.balance {
            this.balance = this.balance - amount;
            this.transactionHistory.append(["取款", amount]);
            println("取款成功：-", amount);
        } else {
            println("取款失败：余额不足或金额无效");
        }
    };

    func getBalance() {
        return this.balance;
    };

    func printHistory() {
        println(this.owner + "的交易记录：");
        for var i = 0; i < len(this.transactionHistory); i++ {
            var tx = this.transactionHistory[i];
            println(tx[0], ": ", tx[1]);
        };
    };
};

// 使用银行账户
var account = BankAccount("张三", 1000);
println("初始余额: ", account.getBalance());
account.deposit(500);
account.withdraw(300);
account.deposit(200);
account.withdraw(1500);  // 应该失败
println("当前余额: ", account.getBalance());
account.printHistory();
```

这个示例展示了 Ghost Language 面向对象编程的以下特性：
- 类定义和实例化
- 构造函数
- 成员变量和成员方法
- this 关键字
- 继承和 super 关键字
- 方法重写（多态）
- 封装数据和行为
- 实际应用场景