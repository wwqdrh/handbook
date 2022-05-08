package fs

import (
  "testing"
  "path"
  "os"
  "fmt"
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
func Test_WriteJsonFile(t *testing.T) {
  var (
    filepath = path.Join(TestDir, "writeJson_test.json")
    err      error
    content  = "{\"name\": \"axetroy\"}"
  )

  if err = WriteJson(filepath, []byte(content)); err != nil {
    return
  }

  defer func() {
    os.RemoveAll(filepath)
  }()

  if m, err := ReadJson(filepath); err != nil {
    panic(err)
  } else {
    if m["name"] != "axetroy" {
      t.Errorf("name field should be axetroy, not %v", m["name"])
    }
    fmt.Println(m)
  }
}
