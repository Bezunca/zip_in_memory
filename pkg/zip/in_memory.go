package zip

import (
	"archive/zip"
	"bytes"
	"errors"
	"io/ioutil"
)

func ExtractInMemory(data []byte) ([]byte, error){
	readerAt := bytes.NewReader(data)
	r, err := zip.NewReader(readerAt, int64(len(data)))
	if err != nil {
		return nil, err
	}

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}

		extractedData, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, err
		}
		rc.Close()
		return extractedData, nil
	}
	return nil, errors.New("empty zip file")
}