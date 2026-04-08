# build.ps1 - 交叉编译脚本 (Windows PowerShell)

# 确保文件以UTF-8 BOM格式保存
#Requires -RunAsAdministrator

# 设置输出编码为 UTF-8
$OutputEncoding = [System.Text.Encoding]::UTF8
[Console]::OutputEncoding = $OutputEncoding
[Console]::InputEncoding = $OutputEncoding

# 自动切换到项目根目录（即 go.mod 所在目录）
Set-Location -Path "$PSScriptRoot/.."

# 配置参数
$Version = "0.4.0"
$OutputDir = "bin"
$MainPackage = "./cmd/ghost"  # 相对于项目根目录

# 图标和资源配置
$IconPath = "assets/image.ico"          # ← 请根据实际路径调整！
$SysoPath = "cmd/ghost/ghost.syso"     # .syso 输出路径

# 定义目标平台列表
$TargetPlatforms = @(
    @{ GOOS = "windows"; GOARCH = "amd64"; Suffix = ".exe" },
    @{ GOOS = "windows"; GOARCH = "386";   Suffix = ".exe" },
    @{ GOOS = "windows"; GOARCH = "arm64"; Suffix = ".exe" },

    @{ GOOS = "linux";   GOARCH = "amd64"; Suffix = "" },
    @{ GOOS = "linux";   GOARCH = "386";   Suffix = "" },
    @{ GOOS = "linux";   GOARCH = "arm";   Suffix = "" },
    @{ GOOS = "linux";   GOARCH = "arm64"; Suffix = "" },
    @{ GOOS = "linux";   GOARCH = "mipsle"; Suffix = "" },
    @{ GOOS = "linux";   GOARCH = "mips64le"; Suffix = "" },
    @{ GOOS = "linux";   GOARCH = "ppc64le"; Suffix = "" },
    @{ GOOS = "linux";   GOARCH = "riscv64"; Suffix = "" },
    @{ GOOS = "linux";   GOARCH = "s390x"; Suffix = "" },

    @{ GOOS = "darwin";  GOARCH = "amd64"; Suffix = "" },
    @{ GOOS = "darwin";  GOARCH = "arm64"; Suffix = "" }
)

# 创建输出目录
if (-not (Test-Path -Path $OutputDir)) {
    New-Item -ItemType Directory -Path $OutputDir | Out-Null
}

# 获取构建时间
$BuildTime = Get-Date -Format "yyyy-MM-ddTHH:mm:ss+08:00"

Write-Host "🚀 开始交叉编译 ghost v$Version" -ForegroundColor Cyan
Write-Host "📅 构建时间: $BuildTime`n" -ForegroundColor Cyan

# 确保 rsrc 已安装（可选，提升体验）
$hasRsrc = $null -ne (Get-Command rsrc -ErrorAction SilentlyContinue)
if (-not $hasRsrc) {
    Write-Host "⚠️  未找到 'rsrc' 工具，尝试自动安装..." -ForegroundColor Yellow
    go install github.com/akavel/rsrc@latest
    # 将 GOPATH/bin 加入 PATH（如果不在）
    $goBin = Join-Path ([Environment]::GetFolderPath("UserProfile")) "go/bin"
    if ($env:PATH -notlike "*$goBin*") {
        $env:PATH += ";$goBin"
    }
}

foreach ($platform in $TargetPlatforms) {
    $GOOS = $platform.GOOS
    $GOARCH = $platform.GOARCH
    $Suffix = $platform.Suffix

    # 输出路径
    $OutputName = "ghost-$GOOS-$GOARCH$Suffix"
    $OutputSubDir = "$OutputDir\$GOOS"
    $OutputPath = "$OutputSubDir\$OutputName"

    # 创建子目录
    if (-not (Test-Path -Path $OutputSubDir)) {
        New-Item -ItemType Directory -Path $OutputSubDir | Out-Null
    }

    # 清理旧的 .syso（安全起见）
    if (Test-Path $SysoPath) {
        Remove-Item $SysoPath
    }

    # 仅 Windows 平台生成 .syso
    if ($GOOS -eq "windows") {
        if (-not (Test-Path $IconPath)) {
            Write-Host "❌ 错误: 未找到图标文件 '$IconPath'，请创建或修改路径！" -ForegroundColor Red
            exit 1
        }

        Write-Host "🖼️  为 $GOOS/$GOARCH 生成图标资源..." -ForegroundColor Blue
        # 使用 rsrc 生成 .syso（支持多架构）
        rsrc -arch $GOARCH -ico $IconPath -o $SysoPath
        if ($LASTEXITCODE -ne 0) {
            Write-Host "❌ rsrc 生成失败！" -ForegroundColor Red
            exit 1
        }
    }

    # 设置环境变量
    $env:GOOS = $GOOS
    $env:GOARCH = $GOARCH

    # ldflags
    $ldflags = "-X github.com/Ghost-Xiao/ghost-lang/internal/cli.Version=$Version -X github.com/Ghost-Xiao/ghost-lang/internal/cli.Platform=$GOOS -X github.com/Ghost-Xiao/ghost-lang/internal/cli.Arch=$GOARCH -X github.com/Ghost-Xiao/ghost-lang/internal/cli.BuildTime=$BuildTime"

    $BuildCommand = @(
        "build",
        "-buildvcs=false",
        "-o", $OutputPath,
        "-ldflags", "`"$ldflags`"",
        $MainPackage
    )

    Write-Host "🔧 正在构建: $GOOS/$GOARCH ..." -ForegroundColor Yellow

    $Result = Start-Process -FilePath "go" -ArgumentList $BuildCommand -NoNewWindow -Wait -PassThru

    # 构建后立即清理 .syso（即使失败也清理）
    if (Test-Path $SysoPath) {
        Remove-Item $SysoPath
    }

    if ($Result.ExitCode -eq 0) {
        Write-Host "✅ 成功构建: $OutputName`n" -ForegroundColor Green
    } else {
        Write-Host "❌ 构建失败: $GOOS/$GOARCH (退出码: $($Result.ExitCode))`n" -ForegroundColor Red
    }
}

Write-Host "🎉 全部构建任务完成! 输出目录: $PWD\$OutputDir" -ForegroundColor Magenta