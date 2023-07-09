#!/bin/bash

# 参数检查
if [ $# -ne 2 ]; then
  echo "Usage: $0 path suffix"
  exit 1
fi

# 获取路径和后缀
path=$1
suffix=$2

# 判断后缀是rs还是py，设置注释符号
if [ $suffix = "rs" ]; then
  comment="//"
  language="rust"
elif [ $suffix = "py" ]; then
  comment="#"
  language="python"
elif [ $suffix = "cpp" ]; then
  comment="//"
  language="cpp"
else
  echo "Unsupported suffix: $suffix"
  exit 2
fi

# 创建_sidebar.md文件
sidebar="docs/_sidebar.md"
if [ ! -f $sidebar ]; then
  echo "* [首页](/ \"简介\")" > $sidebar
fi

# 创建[文件名].md文件
mkdir -p "docs/assets/$suffix"
echo "* [$language](assets/$language/README.md)" >> $sidebar

# 遍历路径下的所有文件
for file in $(find $path -type f -name "*.$suffix"); do
  # 获取文件名和文件路径
  filename=$(basename -- "$file")
  filepath=$(dirname -- "$file")
  echo "** [$filename](assets/$suffix/$filename.md)" >> $sidebar

  mdindexfile="docs/assets/$suffix/README.md"
  echo "some code" > $mdindexfile
  mdfile="docs/assets/$suffix/$filename.md"
  echo "" > $mdfile

  # 设置一个标志位，表示是否开始读取代码块
  flag=0

  # 逐行读取文件内容
  while read line; do
    # 如果遇到@autodoc title注释，则提取title后的值，并加入到html和_sidebar.md中，并且把注释行下紧贴的代码也获取到并且加入到[文件名].md中
    if [[ $line == "$comment @autodoc title"* ]]; then
      title=$(echo "$line" | cut -d ' ' -f 3-)
      echo "## $title" >> $mdfile
      echo "\`\`\`$language" >> $mdfile
      flag=1
    # 如果遇到空行，则设置标志位为0，表示结束读取代码块，并在[文件名].md中结束代码块
    elif [[ $flag -eq 1 && -z $line ]]; then
      echo "\`\`\`" >> $mdfile
      flag=0
    # 如果标志位为1，并且不是以注释符号开头的行，则表示是代码块的一部分，将其写入到[文件名].md中，并转义特殊字符
    elif [[ $flag -eq 1 && $line != "$comment"* ]]; then
      # line=${line//&/&amp;}
      # line=${line//</&lt;}
      # line=${line//>/&gt;}
      echo "$line" >> $mdfile
    fi
  done < $file

  if [[ $flag -eq 1 ]]; then
    echo "\`\`\`" >> $mdfile
  fi

done
