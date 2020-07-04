package zip_in_memory

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
)

type EmptyZipFileError struct {}

func (e *EmptyZipFileError) Error() string {
	return "empty zip file"
}

func ExtractFirstFile(data []byte) ([]byte, error){
	readerAt := bytes.NewReader(data)
	zipReader, err := zip.NewReader(readerAt, int64(len(data)))
	if err != nil {
		return nil, err
	}

	if len(zipReader.File) > 0 {
		file, err := zipReader.File[0].Open()
		if err != nil {
			return nil, err
		}

		extractedData, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		if err := file.Close(); err != nil {
			return nil, err
		}

		return extractedData, nil
	}

	return nil, &EmptyZipFileError{}
}

func ExtractFiles(data []byte) ([][]byte, error){
	readerAt := bytes.NewReader(data)
	zipReader, err := zip.NewReader(readerAt, int64(len(data)))
	if err != nil {
		return nil, err
	}

	files := make([][]byte, 0, len(zipReader.File))
	for _, f := range zipReader.File {
		file, err := f.Open()
		if err != nil {
			return nil, err
		}

		extractedData, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		if err := file.Close(); err != nil {
			return nil, err
		}

		files = append(files, extractedData)
	}

	if len(files) == 0 {
		return nil, &EmptyZipFileError{}
	}

	return files, nil
}