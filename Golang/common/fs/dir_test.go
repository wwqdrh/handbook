package fs

import (
  "testing"
  "path"
  "os"
  "io/ioutil"
)

const TestDir = ".temp"

func init() {
  err := EnsureDir(TestDir)
  if err != nil {
    panic(err)
  }
}

/**
Test mkdir
 */
func Test_MkDir(t *testing.T) {
  var (
    dirPath = path.Join(TestDir, "mkdir_test_dir")
    err     error
  )

  if err = Mkdir(dirPath); err != nil {
    t.Errorf("Create mkdir fial")
    return
  }

  defer func() {
    Remove(dirPath)
  }()

  if _, err := os.Stat(dirPath); os.IsNotExist(err) {
    t.Errorf("Test Mkdir fail")
    return
  }
}

/**
Ensure a dir
 */
func Test_EnsureDir(t *testing.T) {
  var (
    dirPath = path.Join(TestDir, "ensure_test_dir")
    err     error
  )

  if err = EnsureDir(dirPath); err != nil {
    t.Errorf("ensure dir fail")
  }

  defer func() {
    Remove(dirPath)
  }()

  if _, err := os.Stat(dirPath); os.IsNotExist(err) {
    t.Errorf("ensure dir fail")
    return
  }
}

/**
Ensure a dir
 */
func Test_EnsureDirIfParentDirNotExist(t *testing.T) {
  var (
    dirPath = path.Join(TestDir, "ensure_parent_test_dir/parent/nest")
    err     error
  )

  if err = EnsureDir(dirPath); err != nil {
    t.Errorf("ensure dir fail")
    return
  }

  defer func() {
    Remove(path.Join(TestDir, "ensure_parent_test_dir"))
  }()

  if _, err := os.Stat(dirPath); os.IsNotExist(err) {
    t.Errorf("ensure dir fail")
    return
  }
}

/**
Read dir
 */
func Test_ReadDir(t *testing.T) {
  var (
    distDir = path.Join(".github")
  )

  if files, err := Readdir(distDir); err != nil {
    t.Error("Readdir Fail.")
  } else {
    if len(files) != 4 {
      t.Error("Readdir Fail.")
      return
    }
  }
}

func Test_Mkemp(t *testing.T) {
  if dir, err := Mktemp("./.temp", "test_temp_dir"); err != nil {
    t.Errorf("create temp dir fail %v", err.Error())
    return
  } else {
    defer os.RemoveAll(dir)
    files, err := ioutil.ReadDir(dir)
    if err != nil {
      t.Errorf("read temp dir fail %v", err.Error())
      return
    }

    if len(files) != 0 {
      t.Errorf("temp dir should be empty")
      return
    }

  }
}
