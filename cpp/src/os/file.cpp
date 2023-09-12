#include <filesystem>
#include <fstream>
#include <iostream>
#include <string>
#include <vector>

#include "os/file.h"

namespace cppkit::os {
File::File(const std::string &p) : path(p) {
  // 检查路径是否存在
  if (!std::filesystem::exists(path)) {
    // 如果不存在，抛出异常
    throw std::runtime_error("Path does not exist: " + p);
  }
}

// @autodoc title 获取目录下所有文件
std::vector<std::string> File::get_folders() const {
  // 创建一个字符串向量用于存储结果
  std::vector<std::string> folders;
  // 遍历当前路径下的所有条目
  for (const auto &entry : std::filesystem::directory_iterator(path)) {
    // 如果条目是一个目录
    if (entry.is_directory()) {
      // 将其路径转换为字符串并添加到结果向量中
      folders.push_back(entry.path().string());
    }
  }
  // 返回结果向量
  return folders;
}

// @autodoc title 获取当前路径所有文件的方法
std::vector<std::string> File::get_files() const {
  // 创建一个字符串向量用于存储结果
  std::vector<std::string> files;
  // 遍历当前路径下的所有条目
  for (const auto &entry : std::filesystem::directory_iterator(path)) {
    // 如果条目是一个普通文件
    if (entry.is_regular_file()) {
      // 将其路径转换为字符串并添加到结果向量中
      files.push_back(entry.path().string());
    }
  }
  // 返回结果向量
  return files;
}

// @autodoc title 获取某个文件的内容
std::string File::get_content(const std::string &filename) const {
  // 创建一个字符串用于存储结果
  std::string content;
  // 创建一个输入文件流对象，打开指定的文件
  std::ifstream in(filename);
  // 检查文件是否打开成功
  if (in) {
    // 创建一个字符串用于存储每一行的内容
    std::string line;
    // 循环读取文件的每一行，直到文件结束
    while (std::getline(in, line)) {
      // 将每一行的内容添加到结果字符串中，并换行
      content += line + "\n";
    }
    // 关闭文件流对象
    in.close();
  } else {
    // 如果文件打开失败，抛出异常
    throw std::runtime_error("Failed to open file: " + filename);
  }
  // 返回结果字符串
  return content;
}
} // namespace cppkit::os
