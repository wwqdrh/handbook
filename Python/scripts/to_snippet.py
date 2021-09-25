"""
将python脚本转为vscode snippet格式

1、一行行读取为每一行加上双引号
2、为行内的引号做转义
3、制表符使用\t转义  默认将开头的四个为一组替换
"""
import io
import pathlib


def trans_snippet(path: pathlib.Path) -> None:
    output_buffer = io.StringIO()

    with open(path, mode="r", encoding="utf8") as f:
        while line := f.readline():
            output_buffer.write('"')

            while line.startswith("    "):
                output_buffer.write(r"\t")
                line = line[4:]

            line = line.strip()
            for ch in line:
                if ch in ('"', "'"):
                    output_buffer.write(f"\{ch}")
                    continue

                output_buffer.write(ch)

            output_buffer.write('",\n')

    print(output_buffer.getvalue())
    output_buffer.close()


if __name__ == "__main__":
    if len(argv := __import__("sys").argv) > 1 and argv[1] == "test":
        import doctest

        doctest.testmod()

    if len(argv := __import__("sys").argv) > 1:
        file = argv[1]
        trans_snippet(pathlib.Path(file))