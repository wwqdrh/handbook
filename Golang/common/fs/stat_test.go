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

/**
Test Chmod of a file
 */
func Test_Stat(t *testing.T) {
  var (
    filepath = path.Join(TestDir, "stat_test.file")
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

  if _, err = Stat(filepath); err != nil {
    t.Errorf("test stat fail")
    return
  }

}

/**
Test Chmod of a file
 */
func Test_LStat(t *testing.T) {
  var (
    filepath = path.Join(TestDir, "lstat_test.file")
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

  if _, err = LStat(filepath); err != nil {
    t.Errorf("test stat fail")
    return
  }

}
