#!/bin/bash
# setup-macos.command - 为 macOS 创建带图标的 ghost.app

set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
BINARY="$SCRIPT_DIR/ghost"
ICON_PNG="$SCRIPT_DIR/image.png"
APP_NAME="ghost"
APP_PATH="$HOME/Applications/ghost.app"

# 检查文件
if [ ! -f "$BINARY" ]; then
    echo "❌ 未找到二进制文件 'ghost'"
    exit 1
fi
if [ ! -f "$ICON_PNG" ]; then
    echo "❌ 未找到图标文件 'image.png'"
    exit 1
fi

# 要求是 macOS
if [[ "$OSTYPE" != "darwin"* ]]; then
    echo "❌ 此脚本仅支持 macOS"
    exit 1
fi

echo "🔧 正在创建 $APP_NAME.app..."

# 清理旧版本
rm -rf "$APP_PATH"

# 创建目录结构
mkdir -p "$APP_PATH/Contents/MacOS"
mkdir -p "$APP_PATH/Contents/Resources"

# 复制二进制
cp "$BINARY" "$APP_PATH/Contents/MacOS/ghost"

# 将 PNG 转换为 ICNS（需要 sips 和 iconutil）
# 临时目录
TMP_ICONSET="$(mktemp -d)/Ghost.iconset"

# 创建标准尺寸（iconutil 需要特定命名）
sips -z 16 16   "$ICON_PNG" --out "$TMP_ICONSET/icon_16x16.png" > /dev/null
sips -z 32 32   "$ICON_PNG" --out "$TMP_ICONSET/icon_16x16@2x.png" > /dev/null
sips -z 32 32   "$ICON_PNG" --out "$TMP_ICONSET/icon_32x32.png" > /dev/null
sips -z 64 64   "$ICON_PNG" --out "$TMP_ICONSET/icon_32x32@2x.png" > /dev/null
sips -z 64 64   "$ICON_PNG" --out "$TMP_ICONSET/icon_64x64.png" > /dev/null
sips -z 128 128 "$ICON_PNG" --out "$TMP_ICONSET/icon_64x64@2x.png" > /dev/null
sips -z 128 128 "$ICON_PNG" --out "$TMP_ICONSET/icon_128x128.png" > /dev/null
sips -z 256 256 "$ICON_PNG" --out "$TMP_ICONSET/icon_128x128@2x.png" > /dev/null
sips -z 256 256 "$ICON_PNG" --out "$TMP_ICONSET/icon_256x256.png" > /dev/null
sips -z 512 512 "$ICON_PNG" --out "$TMP_ICONSET/icon_256x256@2x.png" > /dev/null
sips -z 512 512 "$ICON_PNG" --out "$TMP_ICONSET/icon_512x512.png" > /dev/null
sips -z 1024 1024 "$ICON_PNG" --out "$TMP_ICONSET/icon_512x512@2x.png" > /dev/null

# 转换为 .icns
iconutil -c icns -o "$APP_PATH/Contents/Resources/Icon.icns" "$TMP_ICONSET"

# 清理
rm -rf "$TMP_ICONSET"

# 创建 Info.plist
cat > "$APP_PATH/Contents/Info.plist" << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleExecutable</key>
    <string>ghost</string>
    <key>CFBundleIconFile</key>
    <string>Icon.icns</string>
    <key>CFBundleIdentifier</key>
    <string>com.ghostxiao.ghost</string>
    <key>CFBundleName</key>
    <string>$APP_NAME</string>
    <key>CFBundlePackageType</key>
    <string>APPL</string>
    <key>LSMinimumSystemVersion</key>
    <string>10.12</string>
</dict>
</plist>
EOF

# 设置权限
chmod +x "$APP_PATH/Contents/MacOS/ghost"

echo "✅ 安装成功！已将 $APP_NAME.app 放入「应用程序」文件夹。"
echo "💡 你可以通过 Launchpad 或 Finder 启动它。"