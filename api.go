package pinata

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type PinResponse struct {
	IpfsHash string `json:"IpfsHash,omitempty"`
	PinSize string `json:"PinSize,omitempty"`
	Timestamp string `json:"Timestamp,omitempty"`
	Error string `json:"error,omitempty"`
}

func (c *Client)PinFile(filepath string) (PinResponse, error) {
	b, w, err := createMultipartFormData(filepath)
	req, _ := http.NewRequest("POST", c.Node +ApiPinFile, &b)
	req.Header.Set("Authorization", "Bearer " + c.JWT)
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return PinResponse{}, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return PinResponse{},err
	}
	var out PinResponse
	if err = json.Unmarshal(content, &out); err != nil {
		return PinResponse{},err
	}
	return out, nil
}

func createMultipartFormData(fileName string) (bytes.Buffer, *multipart.Writer, error) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)
	var fw io.Writer
	file, err := os.Open(fileName)
	if err != nil {
		return b, w, err
	}
	if fw, err = w.CreateFormFile("file", file.Name()); err != nil {
		return b, w, err
	}
	if _, err = io.Copy(fw, file); err != nil {
		return b, w, err
	}
	w.Close()
	return b, w, nil
}

