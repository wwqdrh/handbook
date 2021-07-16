代码，发行，打包为独立可执行程序

# nuitka进行打包

## 包含第三方库
python -m nuitka --follow-imports --standalone simple.py  # 会生成 [].dist 文件夹，其中包含了可执行文件, 可以将这个分发给其他没有python环境的使用
python -m nuitka --follow-imports simple.py