diesel_cli 安装需要有mysql、postgres、sqlite等后端程序

暂时不知道数据库在docker中如何处理

macos默认有sqlite那么就先用sqlite

`cargo install diesel_cli --no-default-features --features sqlite`

`sqliteurl格式  sqlite:./app.db`

```
diesel setup
diesel migration generate message
diesel migration run
```