package fs

import (
  "testing"
  "path"
  "io/ioutil"
  "os"
)

func init() {
  err := EnsureDir(TestDir)
  if err != nil {
    panic(err)
  }
}

func Test_RemoveFile(t *testing.T) {
  var (
    filepath = path.Join(TestDir, "remove_test.file")
    err      error
    content  = "hello world"
  )

  if err = ioutil.WriteFile(filepath, []byte(content), 0777); err != nil {
    t.Errorf("Write file fail %v", err.Error())
    return
  }

  if err = Remove(filepath); err != nil {
    t.Errorf("remove file fail %v", err.Error())
    return
  }

  if _, err := os.Stat(filepath); os.IsNotExist(err) == false {
    t.Errorf("the file should be remove")
    return
  }
}

func Test_RemoveDir(t *testing.T) {
  var (
    dirpath = path.Join(TestDir, "remove_test_dir")
    err     error
  )

  if err = os.Mkdir(dirpath, 0777); err != nil {
    t.Errorf("mkdir fail %v", err.Error())
    return
  }

  if err = Remove(dirpath); err != nil {
    t.Errorf("remove dir fail %v", err.Error())
    return
  }

  if _, err := os.Stat(dirpath); os.IsNotExist(err) == false {
    t.Errorf("the dir should be remove")
    return
  }
}
