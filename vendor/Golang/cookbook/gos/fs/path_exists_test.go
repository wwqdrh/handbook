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

func Test_PathExistsFile(t *testing.T) {
  var (
    filePath = path.Join(TestDir, "pathExists_test.file")
    err      error
  )

  if isExist := PathExists(filePath); isExist {
    t.Errorf("path shoud not exist before create")
    return
  }

  if err = ioutil.WriteFile(filePath, []byte("hello world"), os.ModePerm); err != nil {
    t.Errorf("create file fail")
    return
  }

  defer func() {
    os.RemoveAll(filePath)
  }()

  if isExist := PathExists(filePath); !isExist {
    t.Errorf("path shoud exist after create")
    return
  }
}

func Test_PathExistsDir(t *testing.T) {
  var (
    filePath = path.Join(TestDir, "pathExists_test_dir")
    err      error
  )

  if isExist := PathExists(filePath); isExist {
    t.Errorf("path shoud not exist before create")
    return
  }

  if err = os.Mkdir(filePath, os.ModePerm); err != nil {
    t.Errorf("create dir fail")
    return
  }

  defer func() {
    os.RemoveAll(filePath)
  }()

  if isExist := PathExists(filePath); !isExist {
    t.Errorf("path shoud exist after create")
    return
  }
}
