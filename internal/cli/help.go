package cli

import "github.com/fatih/color"

// PrintHelp 显示命令行帮助信息
func PrintHelp() {
	color.Blue("Usage: ghost [global flags] <command> [arguments]")
	color.Blue("Global Flags:")
	color.Blue("  -h                     Show help")
	color.Blue("  -v                     Print version")
	color.Blue("  -r                     Start REPL")
	color.Blue("Commands:")
	color.Blue("  repl                   Start REPL")
	color.Blue("  run <file>             Execute a .gh file")
	color.Blue("Examples:")
	color.Blue("  ghost -r               # Start REPL with flag")
	color.Blue("  ghost repl             # Start REPL with command")
	color.Blue("  ghost run main.gh      # Run a file")
}
