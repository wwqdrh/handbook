#include "os/file.h"
#include "gtest/gtest.h" // Google Test框架

// 测试用例类
class FileTest : public ::testing::Test {
protected:
  // 测试前执行
  void SetUp() override {
    // 创建一个File对象，传递当前目录作为路径
    f = new cppkit::os::File("./");
  }

  // 测试后执行
  void TearDown() override {
    // 删除File对象
    delete f;
  }

  // File对象指针
  cppkit::os::File *f;
};

// 测试get_folders方法
TEST_F(FileTest, GetFolders) {
  // 调用get_folders方法，获取当前目录下的所有文件夹，并断言结果不为空
  ASSERT_FALSE(f->get_folders().empty());
}

// 测试get_files方法
TEST_F(FileTest, GetFiles) {
  // 调用get_files方法，获取当前目录下的所有文件，并断言结果不为空
  ASSERT_FALSE(f->get_files().empty());
}

// 测试get_content方法
TEST_F(FileTest, GetContent) {
  // 调用get_content方法，获取当前目录下的README.md文件的内容，并断言结果不为空
  ASSERT_FALSE(f->get_content("README.md").empty());
}

int main(int argc, char *argv[]) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}