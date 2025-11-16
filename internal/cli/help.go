package cli

// PrintHelp 显示命令行帮助信息
func PrintHelp() {
	printInfo("Usage: ghost [global flags] <command> [arguments]")
	printInfo("Global Flags:")
	printInfo("  -h                     Show help")
	printInfo("  -v                     Print version")
	printInfo("  -r                     Start REPL")
	printInfo("Commands:")
	printInfo("  repl                   Start REPL")
	printInfo("  run <file>             Execute a .gh file")
	printInfo("Examples:")
	printInfo("  ghost -r               # Start REPL with flag")
	printInfo("  ghost repl             # Start REPL with command")
	printInfo("  ghost run main.gh      # Run a file")
}
