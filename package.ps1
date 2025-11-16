# package.ps1 - 为每个平台生成独立 zip 包
# 依赖：先运行 .\build.ps1 生成 bin/

param(
    [string]$Version = "v0.1.0"
)

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

# 遍历 bin 下所有平台目录
$PlatformDirs = Get-ChildItem -Path "bin" -Directory

foreach ($platformDir in $PlatformDirs) {
    $GOOS = $platformDir.Name
    $ArchFiles = Get-ChildItem -Path $platformDir.FullName -File

    foreach ($file in $ArchFiles) {
        # 解析 arch：ghost-windows-amd64.exe → amd64
        $BaseName = [System.IO.Path]::GetFileNameWithoutExtension($file.Name)
        if ($BaseName -match "ghost-[^-]+-(.+)") {
            $GOARCH = $matches[1]
        } else {
            $GOARCH = "unknown"
        }

        # 生成 zip 名称：ghost-v0.1.0-windows-amd64.zip
        $ZipName = "ghost-$Version-$GOOS-$GOARCH.zip"
        $ZipPath = Join-Path $ReleasesDir $ZipName

        # 临时工作目录
        $TempDir = Join-Path $env:TEMP "ghost-pack-$(Get-Random)"
        New-Item -ItemType Directory -Path $TempDir | Out-Null

        # 复制二进制 + 文档
        $BinaryDest = if ($GOOS -eq "windows") { "ghost.exe" } else { "ghost" }
        Copy-Item $file.FullName (Join-Path $TempDir $BinaryDest)
        Copy-Item "README.md", "LICENSE" -Destination $TempDir -ErrorAction SilentlyContinue

        # 打包（使用 PowerShell Compress-Archive，兼容 Win10+）
        try {
            Compress-Archive -Path "$TempDir\*" -DestinationPath $ZipPath -Force
            Write-Host "✅ $ZipName" -ForegroundColor Green
        } catch {
            Write-Host "❌ 打包失败: $ZipName" -ForegroundColor Red
        } finally {
            Remove-Item -Recurse -Force $TempDir -ErrorAction SilentlyContinue
        }
    }
}

Write-Host "`n🎉 全部 $(Get-ChildItem $ReleasesDir -Filter *.zip | Measure-Object).Count 个平台打包完成！" -ForegroundColor Magenta
Write-Host "📁 输出目录: $(Resolve-Path $ReleasesDir)" -ForegroundColor Cyan