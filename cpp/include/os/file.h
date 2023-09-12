#ifndef FILE_H
#define FILE_H

#include <string>
#include <vector>

// 命名空间
namespace cppkit::os {
// 文件操作类
class File {
private:
  std::string path; // 文件路径
public:
  // 构造函数，接受一个字符串作为路径
  File(const std::string &p);

  // 获取当前路径所有文件夹的方法
  std::vector<std::string> get_folders() const;

  // 获取当前路径所有文件的方法
  std::vector<std::string> get_files() const;

  // 获取某个文件的内容
  std::string get_content(const std::string &filename) const;
};

} // namespace cppkit::os

#endif // FILE_H
