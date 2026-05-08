package builtin_module

import (
	"fmt"
	"time"

	"github.com/Ghost-Xiao/ghost-lang/internal/errors"
	"github.com/Ghost-Xiao/ghost-lang/internal/frame"
	"github.com/Ghost-Xiao/ghost-lang/internal/object"
	"github.com/Ghost-Xiao/ghost-lang/internal/util"
	"github.com/Nomango/datefmt"
)

// TimeModule 表示 time 内置模块
var TimeModule = initTimeModule()

// initTimeModule 初始化 time 模块
//
// 返回值:
//
//	*object.Module - 初始化后的 time 模块
func initTimeModule() *object.Module {
	env := &object.Environment{
		Name:  "time",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	env.Set("Time", &object.Symbol{Name: "Time", Value: TimeClass, IsConst: true})
	env.Set("Duration", &object.Symbol{Name: "Duration", Value: DurationClass, IsConst: true})
	env.Set("sleep", &object.Symbol{Name: "sleep", Value: &TIME_SLEEP, IsConst: true})

	// 时间常量（纳秒）
	env.Set("NANOSECOND", &object.Symbol{Name: "NANOSECOND", Value: &object.Int{Value: 1}, IsConst: true})
	env.Set("MICROSECOND", &object.Symbol{Name: "MICROSECOND", Value: &object.Int{Value: 1000}, IsConst: true})
	env.Set("MILLISECOND", &object.Symbol{Name: "MILLISECOND", Value: &object.Int{Value: 1000000}, IsConst: true})
	env.Set("SECOND", &object.Symbol{Name: "SECOND", Value: &object.Int{Value: 1000000000}, IsConst: true})
	env.Set("MINUTE", &object.Symbol{Name: "MINUTE", Value: &object.Int{Value: 60000000000}, IsConst: true})
	env.Set("HOUR", &object.Symbol{Name: "HOUR", Value: &object.Int{Value: 3600000000000}, IsConst: true})
	env.Set("DAY", &object.Symbol{Name: "DAY", Value: &object.Int{Value: 86400000000000}, IsConst: true})

	return &object.Module{
		Name: "time",
		Env:  env,
	}
}

// TIME_SLEEP 表示 time.sleep 函数
// 暂停当前执行指定的持续时间
var TIME_SLEEP = object.BuiltinFunction{
	Name:         "sleep",
	Parameter:    []string{"duration"},
	DefaultValue: []object.Object{nil},
	HaveVariadic: false,
	Fn: func(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
		duration := args[0]
		if duration, ok := duration.(*object.Instance); ok {
			if duration.Class.Name == "Duration" {
				ns, err := getNs(duration, f, posStart, posEnd)
				if err != nil {
					return nil, err
				}
				time.Sleep(time.Duration(ns.Value) * time.Nanosecond)
				return nil, nil
			} else {
				return nil, &errors.TypeError{
					Frame:    f,
					Message:  "sleep() argument must be a Duration instance.",
					PosStart: posStart,
					PosEnd:   posEnd,
				}
			}
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "sleep() argument must be a Duration instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	},
}

// TimeClass 表示 Time 内置类的类定义
var TimeClass = initTimeClass()

// initTimeClass 初始化 Time 类
//
// 返回值:
//
//	*object.Class - 初始化后的 Time 类
func initTimeClass() *object.Class {
	member := &object.Environment{
		Name:  "Time",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	member.Set("timestampNs", &object.Symbol{Name: "timestampNs", Value: &object.Int{Value: 0}, IsConst: false})
	member.Set("zoneOffset", &object.Symbol{Name: "zoneOffset", Value: &object.Int{Value: 0}, IsConst: false})
	member.Set("init", &object.Symbol{Name: "init", Value: &TIME_INIT, IsConst: true})
	member.Set("timestamp", &object.Symbol{Name: "timestamp", Value: &TIME_TIMESTAMP, IsConst: true})
	member.Set("setZone", &object.Symbol{Name: "setZone", Value: &TIME_SETZONE, IsConst: true})
	member.Set("millisecond", &object.Symbol{Name: "millisecond", Value: &TIME_MILLISECOND, IsConst: true})
	member.Set("year", &object.Symbol{Name: "year", Value: &TIME_YEAR, IsConst: true})
	member.Set("month", &object.Symbol{Name: "month", Value: &TIME_MONTH, IsConst: true})
	member.Set("day", &object.Symbol{Name: "day", Value: &TIME_DAY, IsConst: true})
	member.Set("hour", &object.Symbol{Name: "hour", Value: &TIME_HOUR, IsConst: true})
	member.Set("minute", &object.Symbol{Name: "minute", Value: &TIME_MINUTE, IsConst: true})
	member.Set("second", &object.Symbol{Name: "second", Value: &TIME_SECOND, IsConst: true})
	member.Set("weekday", &object.Symbol{Name: "weekday", Value: &TIME_WEEKDAY, IsConst: true})
	member.Set("yearday", &object.Symbol{Name: "yearday", Value: &TIME_YEARDAY, IsConst: true})
	member.Set("format", &object.Symbol{Name: "format", Value: &TIME_FORMAT, IsConst: true})
	member.Set("__add__", &object.Symbol{Name: "__add__", Value: &TIME_ADD, IsConst: true})
	member.Set("__sub__", &object.Symbol{Name: "__sub__", Value: &TIME_SUB, IsConst: true})
	member.Set("__eq__", &object.Symbol{Name: "__eq__", Value: &TIME_EQ, IsConst: true})
	member.Set("__ne__", &object.Symbol{Name: "__ne__", Value: &TIME_NE, IsConst: true})
	member.Set("__gt__", &object.Symbol{Name: "__gt__", Value: &TIME_GT, IsConst: true})
	member.Set("__ge__", &object.Symbol{Name: "__ge__", Value: &TIME_GE, IsConst: true})
	member.Set("__lt__", &object.Symbol{Name: "__lt__", Value: &TIME_LT, IsConst: true})
	member.Set("__le__", &object.Symbol{Name: "__le__", Value: &TIME_LE, IsConst: true})

	return &object.Class{
		Name:   "Time",
		Parent: nil,
		Member: member,
	}
}

var (
	TIME_INIT = object.BuiltinFunction{
		Name:         "init",
		Parameter:    []string{"a"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: true,
		Fn:           TimeInit,
	}

	TIME_TIMESTAMP = object.BuiltinFunction{
		Name:         "timestamp",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           TimeTimestamp,
	}

	TIME_SETZONE = object.BuiltinFunction{
		Name:         "setZone",
		Parameter:    []string{"offset"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           TimeSetZone,
	}

	TIME_NANOSECOND = object.BuiltinFunction{
		Name:         "nanosecond",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           TimeNanoSecond,
	}

	TIME_MICROSECOND = object.BuiltinFunction{
		Name:         "microsecond",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           TimeMicroSecond,
	}

	TIME_MILLISECOND = object.BuiltinFunction{
		Name:         "millisecond",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           TimeMillisecond,
	}

	TIME_YEAR = object.BuiltinFunction{
		Name:         "year",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           TimeYear,
	}

	TIME_MONTH = object.BuiltinFunction{
		Name:         "month",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           TimeMonth,
	}

	TIME_DAY = object.BuiltinFunction{
		Name:         "day",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           TimeDay,
	}
	TIME_HOUR = object.BuiltinFunction{
		Name:         "hour",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           TimeHour,
	}

	TIME_MINUTE = object.BuiltinFunction{
		Name:         "minute",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           TimeMinute,
	}

	TIME_SECOND = object.BuiltinFunction{
		Name:         "second",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           TimeSecond,
	}

	TIME_WEEKDAY = object.BuiltinFunction{
		Name:         "weekday",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           TimeWeekday,
	}

	TIME_YEARDAY = object.BuiltinFunction{
		Name:         "yearday",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           TimeYearday,
	}

	TIME_FORMAT = object.BuiltinFunction{
		Name:         "format",
		Parameter:    []string{"layout"},
		DefaultValue: []object.Object{&object.String{Value: "yyyy-MM-dd HH:mm:ss"}},
		HaveVariadic: false,
		Fn:           TimeFormat,
	}

	TIME_ADD = object.BuiltinFunction{
		Name:         "__add__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           TimeAdd,
	}

	TIME_SUB = object.BuiltinFunction{
		Name:         "__sub__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           TimeSub,
	}

	TIME_EQ = object.BuiltinFunction{
		Name:         "__eq__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           TimeEq,
	}

	TIME_NE = object.BuiltinFunction{
		Name:         "__ne__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           TimeNe,
	}

	TIME_GT = object.BuiltinFunction{
		Name:         "__gt__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           TimeGt,
	}

	TIME_GE = object.BuiltinFunction{
		Name:         "__ge__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           TimeGe,
	}

	TIME_LT = object.BuiltinFunction{
		Name:         "__lt__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           TimeLt,
	}

	TIME_LE = object.BuiltinFunction{
		Name:         "__le__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           TimeLe,
	}
)

// TimeInit 实现 Time 类的 init 构造函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 构造函数参数，第一个为实例，第二个为可变参数列表
//
// 返回值:
//
//	object.Object - 构造函数返回值
//	error - 可能出现的错误
func TimeInit(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if this == nil || !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "method init() called without instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	arg := args[1].(*object.List).Elements
	argCount := len(arg)

	switch argCount {
	case 0:
		this.Member.Set("timestampNs",
			&object.Symbol{
				Name:    "timestampNs",
				Value:   &object.Int{Value: time.Now().UnixNano()},
				IsConst: false,
			})
		return nil, nil
	case 1:
		a := arg[0]
		switch v := a.(type) {
		case *object.Int:
			this.Member.Set("timestampNs",
				&object.Symbol{
					Name:    "timestampNs",
					Value:   &object.Int{Value: v.Value},
					IsConst: false,
				})
			return nil, nil
		default:
			return nil, &errors.ArgumentError{
				Frame:    f,
				Message:  "timestamp must be an integer.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	case 2:
		layoutArg, ok := arg[0].(*object.String)
		if !ok {
			return nil, &errors.ArgumentError{
				Frame:    f,
				Message:  "first argument must be a string pattern.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		valueArg, ok := arg[1].(*object.String)
		if !ok {
			return nil, &errors.ArgumentError{
				Frame:    f,
				Message:  "second argument must be a string value.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		layout := datefmt.GoLayout(layoutArg.Value)
		t, err := time.Parse(layout, valueArg.Value)
		if err != nil {
			return nil, &errors.ArgumentError{
				Frame:    f,
				Message:  "cannot parse time value with given pattern.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		this.Member.Set("timestampNs",
			&object.Symbol{
				Name:    "timestampNs",
				Value:   &object.Int{Value: t.UnixNano()},
				IsConst: false,
			})
		return nil, nil
	case 3:
		year, ok1 := arg[0].(*object.Int)
		month, ok2 := arg[1].(*object.Int)
		day, ok3 := arg[2].(*object.Int)
		if !ok1 || !ok2 || !ok3 {
			return nil, &errors.ArgumentError{
				Frame:    f,
				Message:  "year, month, day must be integers.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		this.Member.Set("timestampNs",
			&object.Symbol{
				Name:    "timestampNs",
				Value:   &object.Int{Value: time.Date(int(year.Value), time.Month(month.Value), int(day.Value), 0, 0, 0, 0, time.UTC).UnixNano()},
				IsConst: false,
			})
		return nil, nil
	case 6:
		year, ok1 := arg[0].(*object.Int)
		month, ok2 := arg[1].(*object.Int)
		day, ok3 := arg[2].(*object.Int)
		hour, ok4 := arg[3].(*object.Int)
		minute, ok5 := arg[4].(*object.Int)
		second, ok6 := arg[5].(*object.Int)
		if !ok1 || !ok2 || !ok3 || !ok4 || !ok5 || !ok6 {
			return nil, &errors.ArgumentError{
				Frame:    f,
				Message:  "all arguments must be integers.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		this.Member.Set("timestampNs",
			&object.Symbol{
				Name:    "timestampNs",
				Value:   &object.Int{Value: time.Date(int(year.Value), time.Month(month.Value), int(day.Value), int(hour.Value), int(minute.Value), int(second.Value), 0, time.UTC).UnixNano()},
				IsConst: false,
			})
		return nil, nil
	default:
		return nil, &errors.ArgumentError{
			Frame:    f,
			Message:  "invalid number of arguments.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// getTimeStampNs 获取时间戳的纳秒值
//
// 参数:
//
//	this - 时间实例
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//
// 返回值:
//
//	*object.Int - 时间戳的纳秒值
//	error - 可能出现的错误
func getTimeStampNs(this *object.Instance, f *frame.Frame, posStart, posEnd *util.Pos) (*object.Int, error) {
	timestampNs, ok := this.Member.Get("timestampNs")
	if !ok {
		return nil, &errors.VariableError{
			Frame:    f,
			Message:  "timestamp must be initialized.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	switch v := timestampNs.Value.(type) {
	case *object.Int:
		return v, nil
	default:
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "timestamp must be an integer.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// getZoneOffset 获取时区偏移值
//
// 参数:
//
//	this - 时间实例
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//
// 返回值:
//
//	*object.Int - 时区偏移值
//	error - 可能出现的错误
func getZoneOffset(this *object.Instance, f *frame.Frame, posStart, posEnd *util.Pos) (*object.Int, error) {
	zoneOffset, ok := this.Member.Get("zoneOffset")
	if !ok {
		return nil, &errors.VariableError{
			Frame:    f,
			Message:  "zoneOffset must be initialized.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	switch v := zoneOffset.Value.(type) {
	case *object.Int:
		return v, nil
	default:
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "zoneOffset must be an integer.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// TimeTimestamp 实现 timestamp 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 时间戳（纳秒）
//	error - 可能出现的错误
func TimeTimestamp(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "timestamp() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	return &object.Int{Value: timestampNs.Value}, nil
}

// TimeSetZone 实现 setZone 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 时区偏移
//	error - 可能出现的错误
func TimeSetZone(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "setZone() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	if offset, ok := args[1].(*object.Int); ok {
		this.Member.Set("zoneOffset", &object.Symbol{
			Name:    "zoneOffset",
			Value:   &object.Int{Value: offset.Value},
			IsConst: false,
		})
		return offset, nil
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "zoneOffset must be an integer.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// TimeNanoSecond 实现 nanosecond 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 纳秒部分
//	error - 可能出现的错误
func TimeNanoSecond(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "nanosecond() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	return &object.Int{Value: timestampNs.Value % 1000000000}, nil
}

// TimeMicroSecond 实现 microsecond 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 微秒部分
//	error - 可能出现的错误
func TimeMicroSecond(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "microsecond() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	return &object.Int{Value: timestampNs.Value / 1000}, nil
}

// TimeMillisecond 实现 millisecond 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 毫秒部分
//	error - 可能出现的错误
func TimeMillisecond(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "millisecond() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	return &object.Int{Value: timestampNs.Value / 1000000}, nil
}

// TimeYear 实现 year 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 年份
//	error - 可能出现的错误
func TimeYear(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "year() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	zoneOffset, err := getZoneOffset(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	// 创建时区对象
	offsetSeconds := int(zoneOffset.Value * 3600)
	loc := time.FixedZone(fmt.Sprintf("UTC%+d", zoneOffset.Value), offsetSeconds)
	t := time.Unix(0, timestampNs.Value).In(loc)
	return &object.Int{Value: int64(t.Year())}, nil
}

// TimeMonth 实现 month 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 月份
//	error - 可能出现的错误
func TimeMonth(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "month() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	zoneOffset, err := getZoneOffset(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	// 创建时区对象
	offsetSeconds := int(zoneOffset.Value * 3600)
	loc := time.FixedZone(fmt.Sprintf("UTC%+d", zoneOffset.Value), offsetSeconds)
	t := time.Unix(0, timestampNs.Value).In(loc)
	return &object.Int{Value: int64(t.Month())}, nil
}

// TimeDay 实现 day 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 日期
//	error - 可能出现的错误
func TimeDay(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "day() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	zoneOffset, err := getZoneOffset(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	// 创建时区对象
	offsetSeconds := int(zoneOffset.Value * 3600)
	loc := time.FixedZone(fmt.Sprintf("UTC%+d", zoneOffset.Value), offsetSeconds)
	t := time.Unix(0, timestampNs.Value).In(loc)
	return &object.Int{Value: int64(t.Day())}, nil
}

// TimeHour 实现 hour 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 小时
//	error - 可能出现的错误
func TimeHour(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "hour() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	zoneOffset, err := getZoneOffset(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	// 创建时区对象
	offsetSeconds := int(zoneOffset.Value * 3600)
	loc := time.FixedZone(fmt.Sprintf("UTC%+d", zoneOffset.Value), offsetSeconds)
	t := time.Unix(0, timestampNs.Value).In(loc)
	return &object.Int{Value: int64(t.Hour())}, nil
}

// TimeMinute 实现 minute 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 分钟
//	error - 可能出现的错误
func TimeMinute(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "minute() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	zoneOffset, err := getZoneOffset(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	// 创建时区对象
	offsetSeconds := int(zoneOffset.Value * 3600)
	loc := time.FixedZone(fmt.Sprintf("UTC%+d", zoneOffset.Value), offsetSeconds)
	t := time.Unix(0, timestampNs.Value).In(loc)
	return &object.Int{Value: int64(t.Minute())}, nil
}

// TimeSecond 实现 second 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 秒
//	error - 可能出现的错误
func TimeSecond(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "second() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	zoneOffset, err := getZoneOffset(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	// 创建时区对象
	offsetSeconds := int(zoneOffset.Value * 3600)
	loc := time.FixedZone(fmt.Sprintf("UTC%+d", zoneOffset.Value), offsetSeconds)
	t := time.Unix(0, timestampNs.Value).In(loc)
	return &object.Int{Value: int64(t.Second())}, nil
}

// TimeWeekday 实现 weekday 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 星期几
//	error - 可能出现的错误
func TimeWeekday(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "weekday() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	zoneOffset, err := getZoneOffset(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	// 创建时区对象
	offsetSeconds := int(zoneOffset.Value * 3600)
	loc := time.FixedZone(fmt.Sprintf("UTC%+d", zoneOffset.Value), offsetSeconds)
	t := time.Unix(0, timestampNs.Value).In(loc)
	return &object.Int{Value: int64(t.Weekday())}, nil
}

// TimeYearday 实现 yearday 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 一年中的第几天
//	error - 可能出现的错误
func TimeYearday(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "yearday() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	zoneOffset, err := getZoneOffset(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	// 创建时区对象
	offsetSeconds := int(zoneOffset.Value * 3600)
	loc := time.FixedZone(fmt.Sprintf("UTC%+d", zoneOffset.Value), offsetSeconds)
	t := time.Unix(0, timestampNs.Value).In(loc)
	return &object.Int{Value: int64(t.YearDay())}, nil
}

// TimeFormat 实现 format 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 格式化后的时间字符串
//	error - 可能出现的错误
func TimeFormat(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "format() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	layout := "2006-01-02 15:04:05"
	if layoutArg, ok := args[1].(*object.String); ok {
		layout = datefmt.GoLayout(layoutArg.Value)
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "layout must be a string.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	zoneOffset, err := getZoneOffset(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	// 创建时区对象
	offsetSeconds := int(zoneOffset.Value * 3600)
	loc := time.FixedZone(fmt.Sprintf("UTC%+d", zoneOffset.Value), offsetSeconds)
	t := time.Unix(0, timestampNs.Value).In(loc)
	return &object.String{Value: t.Format(layout)}, nil
}

// NewTimeInstance 创建一个新的 Time 实例
//
// 参数:
//
//	timeClass - Time 类
//	timestampNs - 时间戳（纳秒）
//	zoneOffset - 时区偏移
//	env - 执行环境
//
// 返回值:
//
//	*object.Instance - 新的 Time 实例
func NewTimeInstance(timeClass *object.Class, timestampNs int64, zoneOffset int64, env *object.Environment) *object.Instance {
	instance := &object.Instance{
		Class: timeClass,
	}
	// 创建实例环境
	member := &object.Environment{
		Name:  fmt.Sprintf("instance of class %s", timeClass.Name),
		Store: make(map[string]*object.Symbol),
		Outer: env,
	}
	// 绑定方法和属性到实例环境
	for name, symbol := range timeClass.Member.Store {
		value := symbol.Value
		if function, ok := value.(*object.BuiltinFunction); ok {
			// 绑定方法到实例环境
			member.Set(name, &object.Symbol{
				Name: name,
				Value: &object.Method{
					Name:     name,
					Function: function,
					Instance: instance,
				},
				IsConst: true,
			})
		} else {
			// 绑定属性到实例环境
			member.Set(name, &object.Symbol{
				Name:    name,
				Value:   value,
				IsConst: symbol.IsConst,
			})
		}
	}
	instance.Member = member
	// 绑定 timestampNs 和 zoneOffset 到实例环境
	instance.Member.Set("timestampNs", &object.Symbol{
		Name: "timestampNs",
		Value: &object.Int{
			Value: timestampNs,
		},
		IsConst: false,
	})
	instance.Member.Set("zoneOffset", &object.Symbol{
		Name: "zoneOffset",
		Value: &object.Int{
			Value: zoneOffset,
		},
		IsConst: false,
	})
	return instance
}

// TimeAdd 实现 __add__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 新的 Time 实例
//	error - 可能出现的错误
func TimeAdd(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__add__() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	zoneOffset, err := getZoneOffset(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if instance, ok := other.(*object.Instance); ok {
		if instance.Class.Name == "Duration" {
			ns, err := getNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			newTimeStampNs := timestampNs.Value + ns.Value
			return NewTimeInstance(this.Class, newTimeStampNs, zoneOffset.Value, env), nil
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be a Duration instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be a Duration instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// TimeSub 实现 __sub__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 新的 Time 或 Duration 实例
//	error - 可能出现的错误
func TimeSub(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__sub__() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	zoneOffset, err := getZoneOffset(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if instance, ok := other.(*object.Instance); ok {
		switch instance.Class.Name {
		case "Duration":
			ns, err := getNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			newTimeStampNs := timestampNs.Value - ns.Value
			return NewTimeInstance(this.Class, newTimeStampNs, zoneOffset.Value, env), nil
		case "Time":
			otherTimestampNs, err := getTimeStampNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			ns := timestampNs.Value - otherTimestampNs.Value
			return NewDurationInstance(DurationClass, ns, env), nil
		default:
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be a Time or Duration instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be a Time or Duration instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// TimeEq 实现 __eq__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 比较结果
//	error - 可能出现的错误
func TimeEq(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__eq__() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if instance, ok := other.(*object.Instance); ok {
		if instance.Class.Name == "Time" {
			otherTimestampNs, err := getTimeStampNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			return &object.Bool{Value: timestampNs.Value == otherTimestampNs.Value}, nil
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be a Time instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be a Time instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// TimeNe 实现 __ne__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 比较结果
//	error - 可能出现的错误
func TimeNe(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__ne__() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if instance, ok := other.(*object.Instance); ok {
		if instance.Class.Name == "Time" {
			otherTimestampNs, err := getTimeStampNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			return &object.Bool{Value: timestampNs.Value != otherTimestampNs.Value}, nil
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be a Time instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be a Time instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// TimeLt 实现 __lt__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 比较结果
//	error - 可能出现的错误
func TimeLt(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__lt__() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if instance, ok := other.(*object.Instance); ok {
		if instance.Class.Name == "Time" {
			otherTimestampNs, err := getTimeStampNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			return &object.Bool{Value: timestampNs.Value < otherTimestampNs.Value}, nil
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be a Time instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be a Time instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// TimeLe 实现 __le__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 比较结果
//	error - 可能出现的错误
func TimeLe(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__le__() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if instance, ok := other.(*object.Instance); ok {
		if instance.Class.Name == "Time" {
			otherTimestampNs, err := getTimeStampNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			return &object.Bool{Value: timestampNs.Value <= otherTimestampNs.Value}, nil
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be a Time instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be a Time instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// TimeGt 实现 __gt__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 比较结果
//	error - 可能出现的错误
func TimeGt(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__gt__() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if instance, ok := other.(*object.Instance); ok {
		if instance.Class.Name == "Time" {
			otherTimestampNs, err := getTimeStampNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			return &object.Bool{Value: timestampNs.Value > otherTimestampNs.Value}, nil
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be a Time instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be a Time instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// TimeGe 实现 __ge__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 比较结果
//	error - 可能出现的错误
func TimeGe(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__ge__() can only be called on Time instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	timestampNs, err := getTimeStampNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if instance, ok := other.(*object.Instance); ok {
		if instance.Class.Name == "Time" {
			otherTimestampNs, err := getTimeStampNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			return &object.Bool{Value: timestampNs.Value >= otherTimestampNs.Value}, nil
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be a Time instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be a Time instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// DurationClass 表示 Duration 内置类的类定义
var DurationClass = initDurationClass()

// initDurationClass 初始化 Duration 类
//
// 返回值:
//
//	*object.Class - 初始化后的 Duration 类
func initDurationClass() *object.Class {
	member := &object.Environment{
		Name:  "Duration",
		Store: map[string]*object.Symbol{},
		Outer: nil,
	}

	member.Set("ns", &object.Symbol{Name: "ns", Value: &object.Int{Value: 0}, IsConst: false})
	member.Set("init", &object.Symbol{Name: "init", Value: &DURATION_INIT, IsConst: true})
	member.Set("nanosecond", &object.Symbol{Name: "nanosecond", Value: &DURATION_NANOSECOND, IsConst: true})
	member.Set("microsecond", &object.Symbol{Name: "microsecond", Value: &DURATION_MICROSECOND, IsConst: true})
	member.Set("millisecond", &object.Symbol{Name: "millisecond", Value: &DURATION_MILLISECOND, IsConst: true})
	member.Set("second", &object.Symbol{Name: "second", Value: &DURATION_SECOND, IsConst: true})
	member.Set("minute", &object.Symbol{Name: "minute", Value: &DURATION_MINUTE, IsConst: true})
	member.Set("hour", &object.Symbol{Name: "hour", Value: &DURATION_HOUR, IsConst: true})
	member.Set("day", &object.Symbol{Name: "day", Value: &DURATION_DAY, IsConst: true})
	member.Set("__add__", &object.Symbol{Name: "__add__", Value: &DURATION_ADD, IsConst: true})
	member.Set("__sub__", &object.Symbol{Name: "__sub__", Value: &DURATION_SUB, IsConst: true})
	member.Set("__mul__", &object.Symbol{Name: "__mul__", Value: &DURATION_MUL, IsConst: true})
	member.Set("__div__", &object.Symbol{Name: "__div__", Value: &DURATION_DIV, IsConst: true})
	member.Set("__eq__", &object.Symbol{Name: "__eq__", Value: &DURATION_EQ, IsConst: true})
	member.Set("__ne__", &object.Symbol{Name: "__ne__", Value: &DURATION_NE, IsConst: true})
	member.Set("__gt__", &object.Symbol{Name: "__gt__", Value: &DURATION_GT, IsConst: true})
	member.Set("__ge__", &object.Symbol{Name: "__ge__", Value: &DURATION_GE, IsConst: true})
	member.Set("__lt__", &object.Symbol{Name: "__lt__", Value: &DURATION_LT, IsConst: true})
	member.Set("__le__", &object.Symbol{Name: "__le__", Value: &DURATION_LE, IsConst: true})
	member.Set("__neg__", &object.Symbol{Name: "__neg__", Value: &DURATION_NEG, IsConst: true})

	return &object.Class{
		Name:   "Duration",
		Parent: nil,
		Member: member,
	}
}

var (
	DURATION_INIT = object.BuiltinFunction{
		Name:         "init",
		Parameter:    []string{"nanoseconds"},
		DefaultValue: []object.Object{&object.Int{Value: 0}},
		HaveVariadic: false,
		Fn:           DurationInit,
	}

	DURATION_NANOSECOND = object.BuiltinFunction{
		Name:         "nanosecond",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           DurationNanosecond,
	}

	DURATION_MICROSECOND = object.BuiltinFunction{
		Name:         "microsecond",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           DurationMicrosecond,
	}

	DURATION_MILLISECOND = object.BuiltinFunction{
		Name:         "millisecond",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           DurationMillisecond,
	}

	DURATION_SECOND = object.BuiltinFunction{
		Name:         "second",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           DurationSecond,
	}

	DURATION_MINUTE = object.BuiltinFunction{
		Name:         "minute",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           DurationMinute,
	}

	DURATION_HOUR = object.BuiltinFunction{
		Name:         "hour",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           DurationHour,
	}

	DURATION_DAY = object.BuiltinFunction{
		Name:         "day",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           DurationDay,
	}

	DURATION_ADD = object.BuiltinFunction{
		Name:         "__add__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           DurationAdd,
	}

	DURATION_SUB = object.BuiltinFunction{
		Name:         "__sub__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           DurationSub,
	}

	DURATION_MUL = object.BuiltinFunction{
		Name:         "__mul__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           DurationMul,
	}

	DURATION_DIV = object.BuiltinFunction{
		Name:         "__div__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           DurationDiv,
	}

	DURATION_EQ = object.BuiltinFunction{
		Name:         "__eq__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           DurationEq,
	}

	DURATION_NE = object.BuiltinFunction{
		Name:         "__ne__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           DurationNe,
	}

	DURATION_GT = object.BuiltinFunction{
		Name:         "__gt__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           DurationGt,
	}

	DURATION_GE = object.BuiltinFunction{
		Name:         "__ge__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           DurationGe,
	}

	DURATION_LT = object.BuiltinFunction{
		Name:         "__lt__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           DurationLt,
	}

	DURATION_LE = object.BuiltinFunction{
		Name:         "__le__",
		Parameter:    []string{"other"},
		DefaultValue: []object.Object{nil},
		HaveVariadic: false,
		Fn:           DurationLe,
	}

	DURATION_NEG = object.BuiltinFunction{
		Name:         "__neg__",
		Parameter:    []string{},
		DefaultValue: []object.Object{},
		HaveVariadic: false,
		Fn:           DurationNeg,
	}
)

// DurationInit 实现 Duration 类的 init 构造函数
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 构造函数参数
//
// 返回值:
//
//	object.Object - 构造函数返回值
//	error - 可能出现的错误
func DurationInit(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "init() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}

	switch v := args[1].(type) {
	case *object.Int:
		this.Member.Set("ns", &object.Symbol{Name: "ns", Value: &object.Int{Value: v.Value}, IsConst: false})
		return nil, nil
	default:
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "duration must be an integer.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// getNs 获取 Duration 实例的纳秒值
//
// 参数:
//
//	this - Duration 实例
//	f - 当前调用栈
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//
// 返回值:
//
//	*object.Int - 纳秒值
//	error - 可能出现的错误
func getNs(this *object.Instance, f *frame.Frame, posStart, posEnd *util.Pos) (*object.Int, error) {
	nsSym, ok := this.Member.Get("ns")
	if !ok {
		return nil, &errors.VariableError{
			Frame:    f,
			Message:  "ns must be initialized.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	switch nsSym.Value.(type) {
	case *object.Int:
		return nsSym.Value.(*object.Int), nil
	default:
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "ns must be an integer.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// DurationNanosecond 实现 nanosecond 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 纳秒值
//	error - 可能出现的错误
func DurationNanosecond(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "nanosecond() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	return &object.Int{Value: ns.Value}, nil
}

// DurationMicrosecond 实现 microsecond 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 微秒值
//	error - 可能出现的错误
func DurationMicrosecond(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "microsecond() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	return &object.Float{Value: float64(ns.Value) / 1000.0}, nil
}

// DurationMillisecond 实现 millisecond 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 毫秒值
//	error - 可能出现的错误
func DurationMillisecond(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "millisecond() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	return &object.Float{Value: float64(ns.Value) / 1000000.0}, nil
}

// DurationSecond 实现 second 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 秒值
//	error - 可能出现的错误
func DurationSecond(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "second() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	return &object.Float{Value: float64(ns.Value) / 1000000000.0}, nil
}

// DurationMinute 实现 minute 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 分钟值
//	error - 可能出现的错误
func DurationMinute(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "minute() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	return &object.Float{Value: float64(ns.Value) / 60000000000.0}, nil
}

// DurationHour 实现 hour 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 小时值
//	error - 可能出现的错误
func DurationHour(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "hour() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	return &object.Float{Value: float64(ns.Value) / 3600000000000.0}, nil
}

// DurationDay 实现 day 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 天数值
//	error - 可能出现的错误
func DurationDay(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "day() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	return &object.Float{Value: float64(ns.Value) / 86400000000000.0}, nil
}

// NewDurationInstance 创建一个新的 Duration 实例
//
// 参数:
//
//	durationClass - Duration 类
//	ns - 纳秒值
//	env - 执行环境
//
// 返回值:
//
//	*object.Instance - 新的 Duration 实例
func NewDurationInstance(durationClass *object.Class, ns int64, env *object.Environment) *object.Instance {
	instance := &object.Instance{
		Class: durationClass,
	}
	// 创建实例环境
	member := &object.Environment{
		Name:  fmt.Sprintf("instance of class %s", durationClass.Name),
		Store: make(map[string]*object.Symbol),
		Outer: env,
	}
	// 绑定方法和属性到实例环境
	for name, symbol := range durationClass.Member.Store {
		value := symbol.Value
		if function, ok := value.(*object.BuiltinFunction); ok {
			// 绑定方法到实例环境
			member.Set(name, &object.Symbol{
				Name: name,
				Value: &object.Method{
					Name:     name,
					Function: function,
					Instance: instance,
				},
				IsConst: true,
			})
		} else {
			// 绑定属性到实例环境
			member.Set(name, &object.Symbol{
				Name:    name,
				Value:   value,
				IsConst: symbol.IsConst,
			})
		}
	}
	instance.Member = member
	// 绑定 ns 到实例环境
	instance.Member.Set("ns", &object.Symbol{
		Name: "ns",
		Value: &object.Int{
			Value: ns,
		},
		IsConst: false,
	})
	return instance
}

// DurationAdd 实现 __add__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 新的 Duration 实例
//	error - 可能出现的错误
func DurationAdd(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__add__() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if instance, ok := other.(*object.Instance); ok {
		if instance.Class.Name == "Duration" {
			otherNs, err := getNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			newNs := ns.Value + otherNs.Value
			return NewDurationInstance(this.Class, newNs, env), nil
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be a Duration instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be a Duration instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// DurationSub 实现 __sub__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 新的 Duration 实例
//	error - 可能出现的错误
func DurationSub(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__sub__() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if instance, ok := other.(*object.Instance); ok {
		if instance.Class.Name == "Duration" {
			otherNs, err := getNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			newNs := ns.Value - otherNs.Value
			return NewDurationInstance(this.Class, newNs, env), nil
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be a Duration instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be a Duration instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// DurationMul 实现 __mul__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 新的 Duration 实例
//	error - 可能出现的错误
func DurationMul(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__mul__() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if otherInt, ok := other.(*object.Int); ok {
		newNs := ns.Value * otherInt.Value
		return NewDurationInstance(this.Class, newNs, env), nil
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be an integer.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// DurationDiv 实现 __div__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 新的 Duration 实例或 Float
//	error - 可能出现的错误
func DurationDiv(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__div__() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if otherInt, ok := other.(*object.Int); ok {
		if otherInt.Value == 0 {
			return nil, &errors.MathError{
				Frame:    f,
				Message:  "division by zero.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
		newNs := ns.Value / otherInt.Value
		return NewDurationInstance(this.Class, newNs, env), nil
	} else if instance, ok := other.(*object.Instance); ok {
		if instance.Class.Name == "Duration" {
			otherNs, err := getNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			if otherNs.Value == 0 {
				return nil, &errors.MathError{
					Frame:    f,
					Message:  "division by zero.",
					PosStart: posStart,
					PosEnd:   posEnd,
				}
			}
			return &object.Float{Value: float64(ns.Value) / float64(otherNs.Value)}, nil
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be an integer or a Duration instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be an integer or a Duration instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// DurationEq 实现 __eq__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 比较结果
//	error - 可能出现的错误
func DurationEq(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__eq__() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if instance, ok := other.(*object.Instance); ok {
		if instance.Class.Name == "Duration" {
			otherNs, err := getNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			return &object.Bool{Value: ns.Value == otherNs.Value}, nil
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be a Duration instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be a Duration instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// DurationNe 实现 __ne__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 比较结果
//	error - 可能出现的错误
func DurationNe(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__ne__() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if instance, ok := other.(*object.Instance); ok {
		if instance.Class.Name == "Duration" {
			otherNs, err := getNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			return &object.Bool{Value: ns.Value != otherNs.Value}, nil
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be a Duration instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be a Duration instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// DurationGt 实现 __gt__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 比较结果
//	error - 可能出现的错误
func DurationGt(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__gt__() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if instance, ok := other.(*object.Instance); ok {
		if instance.Class.Name == "Duration" {
			otherNs, err := getNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			return &object.Bool{Value: ns.Value > otherNs.Value}, nil
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be a Duration instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be a Duration instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// DurationGe 实现 __ge__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 比较结果
//	error - 可能出现的错误
func DurationGe(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__ge__() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if instance, ok := other.(*object.Instance); ok {
		if instance.Class.Name == "Duration" {
			otherNs, err := getNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			return &object.Bool{Value: ns.Value >= otherNs.Value}, nil
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be a Duration instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be a Duration instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// DurationLe 实现 __le__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 比较结果
//	error - 可能出现的错误
func DurationLe(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__le__() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if instance, ok := other.(*object.Instance); ok {
		if instance.Class.Name == "Duration" {
			otherNs, err := getNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			return &object.Bool{Value: ns.Value <= otherNs.Value}, nil
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be a Duration instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be a Duration instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// DurationLt 实现 __lt__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 比较结果
//	error - 可能出现的错误
func DurationLt(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__lt__() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	other := args[1]
	if instance, ok := other.(*object.Instance); ok {
		if instance.Class.Name == "Duration" {
			otherNs, err := getNs(instance, f, posStart, posEnd)
			if err != nil {
				return nil, err
			}
			return &object.Bool{Value: ns.Value < otherNs.Value}, nil
		} else {
			return nil, &errors.TypeError{
				Frame:    f,
				Message:  "other must be a Duration instance.",
				PosStart: posStart,
				PosEnd:   posEnd,
			}
		}
	} else {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "other must be a Duration instance.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
}

// DurationNeg 实现 __neg__ 方法
//
// 参数:
//
//	f - 当前调用栈
//	env - 执行环境
//	posStart - 表达式起始位置
//	posEnd - 表达式结束位置
//	args - 方法参数
//
// 返回值:
//
//	object.Object - 新的 Duration 实例
//	error - 可能出现的错误
func DurationNeg(f *frame.Frame, env *object.Environment, posStart, posEnd *util.Pos, args ...object.Object) (object.Object, error) {
	this, ok := args[0].(*object.Instance)
	if !ok {
		return nil, &errors.TypeError{
			Frame:    f,
			Message:  "__neg__() can only be called on Duration instances.",
			PosStart: posStart,
			PosEnd:   posEnd,
		}
	}
	ns, err := getNs(this, f, posStart, posEnd)
	if err != nil {
		return nil, err
	}
	return NewDurationInstance(this.Class, -ns.Value, env), nil
}
