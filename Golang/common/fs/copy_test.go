package fs

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func init() {
	err := EnsureDir(TestDir)
	if err != nil {
		panic(err)
	}
}

func Test_PathCopyFile(t *testing.T) {
	var (
		srcPath  = path.Join(TestDir, "copy_test.file1")
		distPath = path.Join(TestDir, "copy_test.file2")
		content  = "hello world"
		err      error
	)

	// before copy, the new file should not be exist
	if _, err := os.Stat(distPath); os.IsNotExist(err) == false {
		t.Errorf("path shoud not exist before create")
		return
	}

	if err = ioutil.WriteFile(srcPath, []byte(content), os.ModePerm); err != nil {
		t.Errorf("create file fail")
		return
	}

	defer func() {
		os.RemoveAll(srcPath)
	}()

	if err = Copy(srcPath, distPath); err != nil {
		t.Errorf("copy file fail")
		return
	}

	defer func() {
		os.RemoveAll(distPath)
	}()

	//after copy the new file should be exist
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		t.Errorf("src file should still there")
		return
	}

	//after copy the new file should be exist
	if _, err := os.Stat(distPath); os.IsNotExist(err) {
		t.Errorf("dist file should create")
		return
	}

	// and the content should be some
	if d, err := ioutil.ReadFile(distPath); err != nil {
		t.Errorf("dist file can not be read")
		return
	} else {
		if string(d[:]) != content {
			t.Errorf("dist file shoud have same content with src file")
			return
		}
	}
}

func Test_PathCopyFileIfItNotExist(t *testing.T) {
	var (
		srcPath  = path.Join(TestDir, "copy_not_exist_test.file1")
		distPath = path.Join(TestDir, "copy_not_exist_test.file2")
		err      error
	)
	if err = Copy(srcPath, distPath); err == nil {
		t.Error("Copy an unexist file, it should throw an error")
		return
	}
}

func Test_PathCopyDir(t *testing.T) {
	var (
		srcPath   = path.Join(TestDir, "copy_test_dir1")
		distPath  = path.Join(TestDir, "copy_test_dir2")
		file1     = path.Join(srcPath, "copy_test_dir_child.file1")
		file2     = path.Join(distPath, "copy_test_dir_child.file1")
		file1Byte []byte
		file2Byte []byte
		content   = "hello world"
		err       error
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

	// remove src dir
	defer func() {
		os.RemoveAll(srcPath)
	}()

	// write file in the dir
	if err = ioutil.WriteFile(file1, []byte(content), os.ModePerm); err != nil {
		t.Errorf("create file fail %v", err.Error())
		return
	}

	// remove src file
	defer func() {
		os.RemoveAll(file1)
	}()

	if err = Copy(srcPath, distPath); err != nil {
		t.Errorf("copy file fail %v", err.Error())
		return
	}

	// remove dist dir
	defer func() {
		os.RemoveAll(distPath)
	}()
	// remove dist file
	defer func() {
		os.RemoveAll(file2)
	}()

	//after copy the new file should be exist
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		t.Errorf("src file should still there")
		return
	}

	//after copy the new file should be exist
	if _, err := os.Stat(distPath); os.IsNotExist(err) {
		t.Errorf("dist file should create")
		return
	}

	if file1Byte, err = ioutil.ReadFile(file1); err != nil {
		t.Errorf("read file1 fail %v", err.Error())
		return
	}

	if file2Byte, err = ioutil.ReadFile(file2); err != nil {
		t.Errorf("read file2 fail %v", err.Error())
		return
	}

	file1Content := string(file1Byte[:])
	file2Content := string(file2Byte[:])

	if file1Content != file2Content {
		t.Errorf("the copy file should be the same")
		return
	}

}
