package utils

import "io"

func TeeReaderAt(r io.ReaderAt, w io.WriterAt) io.ReaderAt {
	return &teeReaderAt{r, w}
}

type teeReaderAt struct {
	r io.ReaderAt
	w io.WriterAt
}

func (t *teeReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	n, err = t.r.ReadAt(p, off)

	if n > 0 {
		if n, err := t.w.WriteAt(p[:n], off); err != nil {
			return n, err
		}
	}
	return
}
