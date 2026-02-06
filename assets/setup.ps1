# 设置 .gh 文件的图标
# 确保文件以UTF-8 BOM格式保存
#Requires -RunAsAdministrator

# 设置输出编码为 UTF-8
$OutputEncoding = [System.Text.Encoding]::UTF8
[Console]::OutputEncoding = $OutputEncoding
[Console]::InputEncoding = $OutputEncoding

# ========== 核心修改：使用脚本同级目录的相对路径 ==========
# 获取当前脚本所在目录（无论在哪运行，都能定位到同级的ico文件）
$scriptDir = $PSScriptRoot
# ico文件名（请确保和脚本同目录，且文件名一致）
$iconFileName = "image.ico"
# 拼接ico文件的完整路径（相对路径转绝对路径）
$iconPath = Join-Path -Path $scriptDir -ChildPath $iconFileName

# 检查图标文件是否存在
if (-not (Test-Path -Path $iconPath -PathType Leaf)) {
    Write-Error "错误：图标文件不存在！请确保 $iconFileName 放在脚本同级目录，当前查找路径：$iconPath"
    exit 1
}

# 定义要设置的扩展名和文件类型标识
$extension = ".gh"
$fileTypeID = "ghost-lang-file"  # 自定义的文件类型标识（无空格）
$fileTypeDesc = "Ghost 语言文件" # 文件类型描述

# 1. 关联扩展名到自定义文件类型
$extensionKey = "HKCU:\Software\Classes\$extension"
if (-not (Test-Path -Path $extensionKey)) {
    New-Item -Path $extensionKey -Force | Out-Null
}
Set-ItemProperty -Path $extensionKey -Name "(Default)" -Value $fileTypeID -Force

# 2. 创建文件类型的基本信息
$fileTypeKey = "HKCU:\Software\Classes\$fileTypeID"
if (-not (Test-Path -Path $fileTypeKey)) {
    New-Item -Path $fileTypeKey -Force | Out-Null
}
Set-ItemProperty -Path $fileTypeKey -Name "(Default)" -Value $fileTypeDesc -Force

# 3. 设置默认图标（关键：路径格式必须是 "路径,0"）
$iconKey = "$fileTypeKey\DefaultIcon"
if (-not (Test-Path -Path $iconKey)) {
    New-Item -Path $iconKey -Force | Out-Null
}
# 拼接正确的图标路径格式："完整路径,0"（路径带引号，避免空格问题）
$iconValue = "`"$iconPath`",0"
Set-ItemProperty -Path $iconKey -Name "(Default)" -Value $iconValue -Force

# 4. 刷新系统图标缓存，让设置立即生效
Write-Host "正在刷新图标缓存..."
ie4uinit.exe -show
taskkill /f /im explorer.exe
start explorer.exe

Write-Host "成功为 .gh 文件设置图标！" -ForegroundColor Green
Write-Host "文件类型：$fileTypeDesc"
Write-Host "图标路径：$iconPath"

# 5. 询问用户是否将ghost添加到环境变量
$addToEnv = Read-Host -Prompt "是否是否将ghost添加到环境变量中？（y/n）"
if ($addToEnv -eq "y") {
    Write-Host "将ghost添加到环境变量中..."
    $env:GHOST_PATH = $scriptDir
}