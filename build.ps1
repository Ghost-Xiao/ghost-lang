# build.ps1 - 交叉编译脚本 (Windows PowerShell)

# 确保文件以UTF-8 BOM格式保存
#Requires -RunAsAdministrator
# 设置输出编码为 UTF-8
$OutputEncoding = [System.Text.Encoding]::UTF8
[Console]::OutputEncoding = $OutputEncoding
[Console]::InputEncoding = $OutputEncoding

# 配置参数
$Version = "0.0.1<test>"                     # 应用版本
$OutputDir = "bin"                      # 输出目录
$MainPackage = "cmd/ghost/main.go"                  # 主程序入口

# 定义目标平台列表 (GOOS/GOARCH)
$TargetPlatforms = @(
    @{ GOOS = "windows"; GOARCH = "amd64"; Suffix = ".exe" },     # Windows 64位 (x86-64 / AMD64)
    @{ GOOS = "windows"; GOARCH = "386";   Suffix = ".exe" },     # Windows 32位 (x86)
    @{ GOOS = "windows"; GOARCH = "arm64"; Suffix = ".exe" },     # Windows ARM64

    @{ GOOS = "linux";   GOARCH = "amd64"; Suffix = "" },         # Linux 64位 (x86-64)
    @{ GOOS = "linux";   GOARCH = "386";   Suffix = "" },         # Linux 32位 (x86)
    @{ GOOS = "linux";   GOARCH = "arm";   Suffix = "" },         # Linux ARM32 (ARMv7)
    @{ GOOS = "linux";   GOARCH = "arm64"; Suffix = "" },         # Linux ARM64 (AArch64)
    @{ GOOS = "linux";   GOARCH = "mipsle"; Suffix = "" },        # Linux MIPS (小端)
    @{ GOOS = "linux";   GOARCH = "mips64le"; Suffix = "" },      # Linux MIPS64 (小端)
    @{ GOOS = "linux";   GOARCH = "ppc64le"; Suffix = "" },       # Linux PowerPC 64位 (小端)
    @{ GOOS = "linux";   GOARCH = "riscv64"; Suffix = "" },       # Linux RISC-V 64位
    @{ GOOS = "linux";   GOARCH = "s390x"; Suffix = "" },         # Linux IBM Z 大型机

    @{ GOOS = "darwin";  GOARCH = "amd64"; Suffix = "" },         # macOS Intel Mac (x86-64)
    @{ GOOS = "darwin";  GOARCH = "arm64"; Suffix = "" }          # macOS Apple Silicon (M1/M2/M3)
)

# 创建输出目录
if (-not (Test-Path -Path $OutputDir)) {
    New-Item -ItemType Directory -Path $OutputDir | Out-Null
}

# 获取构建时间 (ISO 8601格式)
$BuildTime = Get-Date -Format "yyyy-MM-ddTHH:mm:ss+08:00"

# 开始构建
Write-Host "🚀 开始交叉编译 ghost v$Version" -ForegroundColor Cyan
Write-Host "📅 构建时间: $BuildTime`n" -ForegroundColor Cyan

foreach ($platform in $TargetPlatforms) {
    $GOOS = $platform.GOOS
    $GOARCH = $platform.GOARCH
    $Suffix = $platform.Suffix
    
    # 生成输出文件名
    $OutputName = "ghost-$GOOS-$GOARCH$Suffix"
    $OutputPath = "$OutputDir\$GOOS\$OutputName"

    # 设置环境变量
    $env:GOOS = $GOOS
    $env:GOARCH = $GOARCH

    # 构建命令片段
    $ldflags = '-X github.com/Ghost-Xiao/ghost-lang/internal/cli.Version=' + $Version + ' -X github.com/Ghost-Xiao/ghost-lang/internal/cli.Platform=' + $GOOS + ' -X github.com/Ghost-Xiao/ghost-lang/internal/cli.Arch=' + $GOARCH + ' -X github.com/Ghost-Xiao/ghost-lang/internal/cli.BuildTime=' + $BuildTime

    $BuildCommand = @(
        "build",
        "-o", $OutputPath,
        "-ldflags", "`"$ldflags`"",
        $MainPackage
    )

    # 执行构建
    Write-Host "🔧 正在构建: $GOOS/$GOARCH ..." -ForegroundColor Yellow

    # 使用 Start-Process 执行 go build
    $Result = Start-Process -FilePath "go" -ArgumentList $BuildCommand -NoNewWindow -Wait -PassThru

    # 结果检查
    if ($Result.ExitCode -eq 0) {
        Write-Host "✅ 成功构建: $OutputName`n" -ForegroundColor Green
    } else {
        Write-Host "❌ 构建失败: $GOOS/$GOARCH (错误码: $LASTEXITCODE)`n" -ForegroundColor Red
    }
}

# 完成提示
Write-Host "🎉 全部构建任务完成! 输出目录: $PWD\$OutputDir" -ForegroundColor Magenta