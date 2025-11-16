package cli

import "fmt"

var (
	Version   string // 版本号，通过编译参数注入
	BuildTime string // 构建时间，通过编译参数注入
	Platform  string // 目标平台，通过编译参数注入
	Arch      string // 目标架构，通过编译参数注入
)

func PrintVersion() {
	printInfo(fmt.Sprintf("ghost-lang: ghost %s.", Version))
}
