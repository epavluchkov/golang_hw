package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")

func Copy(fromPath, toPath string, offset, limit int64) error {
	fileInfo, err := os.Stat(from)
	if err != nil {
		return err
	}

	if fileInfo.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	var qtyBytes int64
	if offset+limit <= fileInfo.Size() && limit != 0 {
		qtyBytes = limit
	} else {
		qtyBytes = fileInfo.Size() - offset
	}

	p := make([]byte, qtyBytes)

	fileFrom, err := os.OpenFile(fromPath, os.O_RDONLY, 0655)
	if err != nil {
		return err
	}
	defer fileFrom.Close()

	_, err = fileFrom.Seek(offset, 0)
	if err != nil {
		return err
	}

	_, err = fileFrom.ReadAt(p, offset)
	if err != nil {
		return err
	}

	fileTo, err := os.OpenFile(toPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer fileTo.Close()

	bar := pb.Full.Start64(qtyBytes)
	wr := bar.NewProxyWriter(fileTo)

	_, err = io.CopyN(wr, fileFrom, qtyBytes)
	if err != nil {
		return err
	}

	bar.Finish()

	return nil
}
