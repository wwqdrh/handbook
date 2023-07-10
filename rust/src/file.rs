use std::fs::File;
use std::fs;
use std::io::prelude::*;
use std::path::Path;

// @autodoc title 创建文件
fn createFile() -> std::io::Result<()> {
    let mut file = File::create("foo.txt")?;
    file.write_all(b"Hello, world!")?;
    Ok(())
}


// @autodoc title 读取文件
fn cat(path: &Path) -> std::io::Result<String> {
    let mut f = File::open(path)?;
    let mut s = String::new();
    f.read_to_string(&mut s)?;
    Ok(s)
}

// @autodoc title 删除文件
fn deleteFile() -> std::io::Result<()> {
    fs::remove_file("foo.txt")?;
    Ok(())
}

// @autodoc title 创建文件夹
fn createFolder() -> std::io::Result<()> {
    fs::create_dir("bar")?;
    fs::create_dir_all("a/b/c")?;
    Ok(())
}

// @autodoc title 删除文件夹
fn deleteFolder() -> std::io::Result<()> {
    fs::remove_dir("bar")?;
    fs::remove_dir_all("a")?;
    Ok(())
}