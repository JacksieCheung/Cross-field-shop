package parse

import (
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

type Parser interface {
	ReadFile() error
	Parse()
	Process(userID int) error
	SaveFile() error
}

const (
	FileTypeMHT = ".mht"
)

const basePath = "./statics"

func ParserFactory(file *multipart.FileHeader) (Parser, error) {
	filetype := path.Ext(file.Filename)
	switch filetype {
	case FileTypeMHT:
		return NewMhtParser(file), nil
	default:
		return nil, errors.New("不存在该文件格式的解析器")
	}
}

type BaseParser struct {
	Data string
	File *multipart.FileHeader
}

func (bp *BaseParser) ReadFile() error {
	file, err := bp.File.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	var data []byte
	data, err = ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	bp.Data = string(data)
	return nil
}

func (bp *BaseParser) SaveFile() error {
	file, err := bp.File.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	filename := path.Join(basePath, bp.File.Filename)
	target, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(target, file)
	if err != nil {
		return err
	}

	return nil
}
