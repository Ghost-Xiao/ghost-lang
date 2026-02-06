# package.ps1 - 为每个平台生成独立 zip 包
# 脚本文件保留原始名称：setup.ps1 / setup.sh / setup.command

# 确保文件以UTF-8 BOM格式保存
#Requires -RunAsAdministrator
# 设置输出编码为 UTF-8
$OutputEncoding = [System.Text.Encoding]::UTF8
[Console]::OutputEncoding = $OutputEncoding
[Console]::InputEncoding = $OutputEncoding

$Version = "v0.2.0"

# 确保 bin 目录存在
if (-not (Test-Path "bin")) {
    Write-Error "❌ bin 目录不存在！请先运行 .\build.ps1"
    exit 1
}

# 创建 releases 目录
$ReleasesDir = "releases"
if (-not (Test-Path $ReleasesDir)) {
    New-Item -ItemType Directory -Path $ReleasesDir | Out-Null
}

Write-Host "📦 开始为每个平台打包..." -ForegroundColor Cyan
Write-Host "🔖 版本: $Version`n" -ForegroundColor Green

# 定位项目根目录
$ScriptDir = $PSScriptRoot
$ProjectRoot = Split-Path -Parent $ScriptDir

# 遍历 bin 下所有平台目录
$PlatformDirs = Get-ChildItem -Path "bin" -Directory

foreach ($platformDir in $PlatformDirs) {
    $GOOS = $platformDir.Name
    $ArchFiles = Get-ChildItem -Path $platformDir.FullName -File

    foreach ($file in $ArchFiles) {
        # 解析 arch
        $BaseName = [System.IO.Path]::GetFileNameWithoutExtension($file.Name)
        if ($BaseName -match "ghost-[^-]+-(.+)") {
            $GOARCH = $matches[1]
        } else {
            $GOARCH = "unknown"
        }

        $ZipName = "ghost-$Version-$GOOS-$GOARCH.zip"
        $ZipPath = Join-Path $ReleasesDir $ZipName

        $TempDir = Join-Path $env:TEMP "ghost-pack-$(Get-Random)"
        New-Item -ItemType Directory -Path $TempDir | Out-Null

        # 复制二进制
        $BinaryDest = if ($GOOS -eq "windows") { "ghost.exe" } else { "ghost" }
        Copy-Item $file.FullName (Join-Path $TempDir $BinaryDest)

        # 公共文件
        @("README.md", "LICENSE") | ForEach-Object {
            $src = Join-Path $ProjectRoot $_
            if (Test-Path $src) { Copy-Item $src $TempDir }
        }

        # 平台特定资源（保留原始文件名）
        switch ($GOOS) {
            "windows" {
                $iconSrc = Join-Path $ProjectRoot "assets/image.ico"
                $setupSrc = Join-Path $ProjectRoot "assets/setup.ps1"

                if (Test-Path $iconSrc) {
                    Copy-Item $iconSrc $TempDir
                } else { Write-Warning "⚠️ 未找到 image.ico" }

                if (Test-Path $setupSrc) {
                    Copy-Item $setupSrc $TempDir
                } else { Write-Warning "⚠️ 未找到 setup.ps1" }
            }

            "linux" {
                $iconSrc = Join-Path $ProjectRoot "assets/image.png"
                $setupSrc = Join-Path $ProjectRoot "assets/setup.sh"

                if (Test-Path $iconSrc) {
                    Copy-Item $iconSrc $TempDir
                } else { Write-Warning "⚠️ 未找到 image.png" }

                if (Test-Path $setupSrc) {
                    Copy-Item $setupSrc $TempDir
                } else { Write-Warning "⚠️ 未找到 setup.sh" }
            }

            "darwin" {
                $iconSrc = Join-Path $ProjectRoot "assets/image.png"
                $setupSrc = Join-Path $ProjectRoot "assets/setup.command"

                if (Test-Path $iconSrc) {
                    Copy-Item $iconSrc $TempDir
                } else { Write-Warning "⚠️ 未找到 image.png" }

                if (Test-Path $setupSrc) {
                    Copy-Item $setupSrc $TempDir
                } else { Write-Warning "⚠️ 未找到 setup.command" }
            }

            default {
                Write-Host "ℹ️ 未知平台 $GOOS" -ForegroundColor Yellow
            }
        }

        # 打包
        try {
            Compress-Archive -Path "$TempDir\*" -DestinationPath $ZipPath -Force
            Write-Host "✅ $ZipName" -ForegroundColor Green
        } catch {
            Write-Host "❌ 打包失败: $ZipName - $_" -ForegroundColor Red
        } finally {
            Remove-Item -Recurse -Force $TempDir -ErrorAction SilentlyContinue
        }
    }
}

Write-Host "`n🎉 全部打包完成！" -ForegroundColor Magenta
Write-Host "📁 输出目录: $(Resolve-Path $ReleasesDir)" -ForegroundColor Cyan