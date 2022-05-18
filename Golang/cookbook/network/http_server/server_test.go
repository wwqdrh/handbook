package main

func TestUploadFile(t *testing.T) {
	path := "/home/ubuntu/test.go" //要上传文件所在路径
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("my_file", filepath.Base(path))
	if err != nil {
		writer.Close()
		t.Error(err)
	}
	io.Copy(part, file)
	writer.Close()

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res := httptest.NewRecorder()

	upload(res, req)

	if res.Code != http.StatusOK {
		t.Error("not 200")
	}

	t.Log(res.Body.String())
	// t.Log(io.read)
}
