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

/**
Test mkdir
 */
func Test_WriteFile(t *testing.T) {
  var (
    filepath = path.Join(TestDir, "write_test.file")
    err      error
    content  = "hello world"
    data     []byte
  )

  if err = WriteFile(filepath, []byte(content)); err != nil {
    t.Errorf("Write file fail")
    return
  }

  defer func() {
    os.RemoveAll(filepath)
  }()

  if data, err = ioutil.ReadFile(filepath); err != nil {
    t.Errorf("Write file fail")
    return
  }

  s := string(data[:])

  if s != content {
    t.Errorf("Write file fail")
    return
  }

}

/**
Test read file
 */
func Test_ReadFile(t *testing.T) {
  var (
    filepath = path.Join(TestDir, "read_test.file")
    err      error
    content  = "hello world"
    data     []byte
  )

  if err = ioutil.WriteFile(filepath, []byte(content), os.ModePerm); err != nil {
    panic(err)
    return
  }

  defer func() {
    os.RemoveAll(filepath)
  }()

  if data, err = ReadFile(filepath); err != nil {
    t.Errorf("Write file fail")
    return
  }

  s := string(data[:])

  if s != content {
    t.Errorf("Write file fail")
    return
  }

}

func Test_EnsureFile(t *testing.T) {
  var (
    filepath = path.Join(TestDir, "ensure_test.file")
    err      error
    content  string
    data     []byte
  )

  if err = EnsureFile(filepath); err != nil {
    t.Errorf("ensure file fail")
    return
  }

  defer func() {
    os.RemoveAll(filepath)
  }()

  // read file
  if data, err = ioutil.ReadFile(filepath); err != nil {
    panic(err)
  }

  content = string(data[:])

  if content != "" {
    t.Errorf("ensure file fail")
    return
  }

}

func Test_EnsureFileIfParentDirNotExist(t *testing.T) {
  var (
    filepath = path.Join(TestDir, "ensureFile_nest_dir", "ensure_test.file")
    err      error
    content  string
    data     []byte
  )

  if err = EnsureFile(filepath); err != nil {
    t.Errorf("ensure file fail")
    return
  }

  defer func() {
    os.RemoveAll(path.Dir(filepath))
  }()

  // read file
  if data, err = ioutil.ReadFile(filepath); err != nil {
    panic(err)
  }

  content = string(data[:])

  if content != "" {
    t.Errorf("ensure file fail")
    return
  }

}

func Test_AppendFile(t *testing.T) {
  var (
    filepath = path.Join(TestDir, "append_test.file")
    err      error
    content  = "hello"
    data     []byte
  )

  if err = ioutil.WriteFile(filepath, []byte(content), os.ModePerm); err != nil {
    panic(err)
    return
  }

  defer func() {
    os.RemoveAll(filepath)
  }()

  if err = AppendFile(filepath, []byte(" world")); err != nil {
    panic(err)
  }

  // read file
  if data, err = ioutil.ReadFile(filepath); err != nil {
    panic(err)
  }

  content = string(data[:])

  if content != "hello world" {
    t.Errorf("append file fail")
    return
  }

}
