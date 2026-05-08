package cli

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	Version   string // 版本号，通过编译参数注入
	BuildTime string // 构建时间，通过编译参数注入
	Platform  string // 目标平台，通过编译参数注入
	Arch      string // 目标架构，通过编译参数注入
)

// PrintVersion 打印ghost-lang的版本信息
func PrintVersion() {
	color.Blue(fmt.Sprintf("ghost-lang: ghost %s.", Version))
}
