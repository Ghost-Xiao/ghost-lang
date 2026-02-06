#!/bin/bash
# setup-linux.sh - 为 Linux 桌面环境安装 Ghost 应用图标，并为.gh文件添加图标

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BINARY="$SCRIPT_DIR/ghost"
ICON_PNG="$SCRIPT_DIR/image.png"
APP_NAME="ghost"
EXEC_NAME="ghost"

# 检查依赖
if ! command -v desktop-file-validate &> /dev/null; then
    echo "⚠️  推荐安装 'desktop-file-utils' 以验证 .desktop 文件（非必需）"
fi

# 检查文件是否存在
if [ ! -f "$BINARY" ]; then
    echo "❌ 错误: 未找到二进制文件 'ghost'"
    exit 1
fi
if [ ! -f "$ICON_PNG" ]; then
    echo "❌ 错误: 未找到图标文件 'image.png'"
    exit 1
fi

# 目标路径
ICON_DIR="/usr/share/icons"
DESKTOP_DIR="/usr/share/applications"
DESKTOP_FILE="$DESKTOP_DIR/ghost.desktop"

# 创建目录
mkdir -p "$ICON_DIR"
mkdir -p "$DESKTOP_DIR"

# 安装图标
echo "🖼️  安装图标..."
cp "$ICON_PNG" "$ICON_DIR/ghost.png"

# 创建 .desktop 文件
echo "📝 创建 .desktop 文件..."
cat > "$DESKTOP_FILE" << EOF
[Desktop Entry]
Name=$APP_NAME
Comment=Ghost Language
Exec="$BINARY" -r
Icon=ghost
Terminal=true
Type=Application
Categories=Development;Utility;
StartupNotify=false
EOF

# 设置权限
chmod +x "$DESKTOP_FILE"
chmod +x "$BINARY"

# 刷新桌面数据库（可选）
if command -v update-desktop-database &> /dev/null; then
    update-desktop-database "$DESKTOP_DIR"
fi

echo "✅ 安装成功！请在应用菜单中搜索 '$APP_NAME'。"
echo "💡 提示：如果看不到图标，请注销或重启桌面环境。"

# 配置后缀图标
echo "🖼️  开始配置Ubuntu全用户后缀图标"

# 判断是否存在gh-mime.xml文件
if [ ! -f "/usr/share/mime/packages/gh-mime.xml" ]; then
    
fi
# 写入XML配置内容
sudo tee "/usr/share/mime/packages/gh-mime.xml" > /dev/null << EOF
<?xml version="1.0" encoding="UTF-8"?>
<mime-info xmlns="http://www.freedesktop.org/standards/shared-mime-info">
  <mime-type type="text/x-gh">
  <comment>Ghost 语言文件</comment>
    <glob pattern="*.gh"/>
    <icon name="ghost"/>
  </mime-type>
</mime-info>
EOF

# 向系统注册新的MIME类型
xdg-mime install /usr/share/mime/packages/gh-mime.xml

# 安装图标文件
xdg-icon-resource install --context mimetypes --size 1024 "$ICON_DIR/ghost.png" ghost

echo "✅ 后缀图标配置完成！"

# 询问用户是否将ghost添加到环境变量
echo -n "是否将ghost添加到环境变量中？（y/n）"
# 保存当前终端设置
stty_state=$(stty -g)
# 直接从终端读取（不是从sudo的stdin）
if [ -t 0 ]; then
    stty raw -echo
    answer=$(head -c 1)
    stty "$stty_state"
    echo
else
    # 如果不是交互式终端，使用默认值
    answer="n"
fi

if [ "$answer" = "y" ] || [ "$answer" = "Y" ]; then
    echo "将ghost添加到环境变量中..."
    BASH_RC="$REAL_HOME/.bashrc"
    PATH_ADD="export PATH=\"$SCRIPT_DIR:\$PATH\""
    
    # 检查是否已添加
    if ! grep -q "$PATH_ADD" "$BASH_RC" 2>/dev/null; then
        echo "$PATH_ADD" >> "$BASH_RC"
        echo "✅ 已添加到 $BASH_RC"
        echo "💡 提示：需要重新打开终端或运行 'source $BASH_RC' 来生效"
    else
        echo "⚠️  环境变量已配置，无需重复添加"
    fi
fi

echo "✨ 安装完成！"