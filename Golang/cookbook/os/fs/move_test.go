package fs

import (
  "testing"
  "path"
  "os"
  "io/ioutil"
)

func init() {
  err := EnsureDir(TestDir)
  if err != nil {
    panic(err)
  }
}

func Test_PathMoveFile(t *testing.T) {
  var (
    srcPath      = path.Join(TestDir, "move_test.file1")
    distPath     = path.Join(TestDir, "move_test.file2")
    content      = "hello world"
    err          error
    mode         os.FileMode
    fileInfo     os.FileInfo
    distFileInfo os.FileInfo
  )

  if err = ioutil.WriteFile(srcPath, []byte(content), os.ModePerm); err != nil {
    t.Errorf("create file fail")
    return
  }

  // before copy, the new file should not be exist
  if fileInfo, err = os.Stat(srcPath); os.IsNotExist(err) {
    t.Errorf("path shoud be exist after create")
    return
  }

  mode = fileInfo.Mode()

  if err = Move(srcPath, distPath); err != nil {
    t.Errorf("move file fail %v", err.Error())
    return
  }

  defer func() {
    os.RemoveAll(distPath)
  }()

  //after copy the new file should be exist
  if distFileInfo, err = os.Stat(distPath); os.IsNotExist(err) {
    t.Errorf("dist file should create")
    return
  }

  // mode same
  if distFileInfo.Mode() != mode {
    t.Errorf("mode not same after move file")
    return
  }

  // and file should be some
  if d, err := ioutil.ReadFile(distPath); err != nil {
    t.Errorf("dist file can not be read %v", err.Error())
    return
  } else {
    // content same
    if string(d[:]) != content {
      t.Errorf("dist file shoud have same content with src file")
      return
    }
  }
}

func Test_PathMoveDir(t *testing.T) {
  var (
    srcPath      = path.Join(TestDir, "copy_test_dir1")
    distPath     = path.Join(TestDir, "copy_test_dir2")
    file1        = path.Join(srcPath, "copy_test_dir_child.file1")
    file2        = path.Join(distPath, "copy_test_dir_child.file1")
    file2Byte    []byte
    content      = "hello world"
    err          error
    mode         os.FileMode
    srcFileInfo  os.FileInfo
    distFileInfo os.FileInfo
  )

  // before copy, the new file should not be exist
  if _, err := os.Stat(distPath); os.IsNotExist(err) == false {
    t.Errorf("path shoud not exist before create")
    return
  }

  if err = os.Mkdir(srcPath, os.ModePerm); err != nil {
    t.Errorf("create dir fail %v", err.Error())
    return
  }

  // write file in the dir
  if err = ioutil.WriteFile(file1, []byte(content), os.ModePerm); err != nil {
    t.Errorf("create file fail %v", err.Error())
    return
  }

  if srcFileInfo, err = os.Stat(file1); err != nil {
    t.Errorf("stat file1 fail %v", err.Error())
    return
  }

  mode = srcFileInfo.Mode()

  if err = Move(srcPath, distPath); err != nil {
    t.Errorf("copy file fail %v", err.Error())
    return
  }

  // remove dist dir
  defer func() {
    os.RemoveAll(distPath)
  }()

  // after move, the src dir should be remove
  if _, err := os.Stat(srcPath); os.IsNotExist(err) == false {
    t.Errorf("src file should be move")
    return
  }

  // after move, the dist dir should be exist
  if _, err := os.Stat(distPath); os.IsNotExist(err) {
    t.Errorf("dist file should create")
    return
  }

  // and the new file should readable
  if file2Byte, err = ioutil.ReadFile(file2); err != nil {
    t.Errorf("read file2 fail %v", err.Error())
    return
  }

  file2Content := string(file2Byte[:])

  // same content
  if file2Content != content {
    t.Errorf("the copy file should be the same")
    return
  }

  // same mode
  if distFileInfo, err = os.Stat(file2); err != nil {
    t.Errorf("stat file2 fail %v", err.Error())
    return
  }

  if distFileInfo.Mode() != mode {
    t.Errorf("the copy file should be the same")
    return
  }

}
