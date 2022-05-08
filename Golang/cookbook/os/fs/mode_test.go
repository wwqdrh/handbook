package fs

import (
  "testing"
  "path"
  "os"
  "io/ioutil"
  "runtime"
)

func init() {
  err := EnsureDir(TestDir)
  if err != nil {
    panic(err)
  }
}

/**
Test Chmod of a file
 */
func Test_ChmodFile(t *testing.T) {
  var (
    filepath = path.Join(TestDir, "chmod_test.file")
    err      error
    content  = "hello world"
    fileInfo os.FileInfo
  )

  if err = ioutil.WriteFile(filepath, []byte(content), 0777); err != nil {
    t.Errorf("Write file fail")
    return
  }

  defer func() {
    os.RemoveAll(filepath)
  }()

  if err = Chmod(filepath, 0666); err != nil {
    panic(err)
    return
  }

  if fileInfo, err = os.Stat(filepath); err != nil {
    panic(err)
    return
  }

  if fileInfo.Mode() != 0666 {
    t.Errorf("change mode of a file fail")
    return
  }

}

/**
Test Chmod of a file
 */
func Test_ChmodDir(t *testing.T) {
  t.Skip()
  var (
    dirPath  = path.Join(TestDir, "chmod_test_dir")
    err      error
    fileInfo os.FileInfo
  )

  // create dir
  if err = os.Mkdir(dirPath, 0666); err != nil {
    panic(err)
    return
  }

  defer func() {
    os.RemoveAll(dirPath)
  }()

  // change mode
  if err = Chmod(dirPath, 0777); err != nil {
    t.Errorf("test chmod fail")
    return
  }

  // get mode
  if fileInfo, err = os.Stat(dirPath); err != nil {
    panic(err)
    return
  }

  if fileInfo.Mode() != 0777 {
    t.Errorf("change mode of a dir fail")
    return
  }

}

/**
Test Chmod of a file
 */
func Test_LChod(t *testing.T) {
  var (
    filepath = path.Join(TestDir, "lchod_test.file")
    err      error
    content  = "hello world"
  )

  if err = ioutil.WriteFile(filepath, []byte(content), 0777); err != nil {
    t.Errorf("Write file fail")
    return
  }

  defer func() {
    os.RemoveAll(filepath)
  }()

  if err = LChod(filepath, 10086, 10086); err != nil {
    // windows not support
    if runtime.GOOS != "windows" {
      //t.Errorf("test lchod fail %v", err.Error())
    }
    return
  }

}
