from typing import Generator


class BigFileRead:
    def __init__(self, file: str, block_size: int):
        """
        file: 文件名
        block_size: 每一次读取的字符数
        """
        self.file = file
        self.block_size = block_size

    def __enter__(self):
        def _f() -> Generator[str, None, None]:
            with open(self.file, mode="r") as fp:
                while chunk := fp.read(self.block_size):
                    yield chunk
        return _f()

    def __exit__(self, exc_type, exc_val, exc_tb):
        if exc_type is not None:
            print(exc_val)
        
        return True

def main():
    with BigFileRead("README.md", 8) as f:
        for i in f:
            print(i)
            

if __name__ == "__main__":
    main()