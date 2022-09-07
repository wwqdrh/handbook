package fs

import (
  "testing"
  "path"
  //"os"
  "io/ioutil"
  "fmt"
  "os"
)

func init() {
  err := EnsureDir(TestDir)
  if err != nil {
    panic(err)
  }
}

func Test_Link(t *testing.T) {
  var (
    filePath    = path.Join(TestDir, "link_test.file")
    linkPath    = path.Join(TestDir, "link_test_link.file")
    err         error
    content     = "hello world"
    linkContent string
  )

  if err = ioutil.WriteFile(filePath, []byte(content), 0777); err != nil {
    t.Errorf("Write file fail")
    return
  }

  defer os.RemoveAll(filePath)

  if err = Link(filePath, linkPath); err != nil {
    panic(err)
    return
  }

  defer os.RemoveAll(linkPath)

  if linkContent, err = ReadLink(linkPath); err != nil {
    // if your computer is windows, it will throw an error
  }

  fmt.Println(linkContent)
}

func Test_Symlink(t *testing.T) {
  var (
    filePath    = path.Join(TestDir, "symlink_test.file")
    linkPath    = path.Join(TestDir, "symlink_test_link.file")
    err         error
    content     = "hello world"
    linkContent string
  )

  if err = ioutil.WriteFile(filePath, []byte(content), 0777); err != nil {
    t.Errorf("Write file fail")
    return
  }

  defer os.RemoveAll(filePath)

  if err = Symlink(filePath, linkPath); err != nil {
    // if your computer is windows, it will throw an error
    return
  }

  defer os.RemoveAll(linkPath)

  if linkContent, err = ReadLink(linkPath); err != nil {
    // if your computer is windows, it will throw an error
  }

  fmt.Println(linkContent)
}
