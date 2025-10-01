package main

import (
	"bytes"
	"io"

	"github.com/google/brotli/go/cbrotli"
)

func compressBrotli(data []byte, quality int) ([]byte, error) {
	var buf bytes.Buffer
	writer := cbrotli.NewWriter(&buf, cbrotli.WriterOptions{Quality: quality})
	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decompressBrotli(data []byte) ([]byte, error) {
	reader := cbrotli.NewReader(bytes.NewReader(data))
	var buf bytes.Buffer
	_, err := io.Copy(&buf, reader)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
