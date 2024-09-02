package progress

import (
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"golang.org/x/term"
	"io"
	"os"
)

type Percent struct {
	enabled bool
	Writer  io.Writer
	total   int
	current int
	prefix  string
}

func NewReader(r io.Reader, prefix string, total int, enabled bool) (io.Reader, *Percent) {
	progressBar := newPercent(prefix, total, enabled)
	return io.TeeReader(r, progressBar), progressBar
}

func NewReaderAt(r io.ReaderAt, prefix string, total int, enabled bool) (io.ReaderAt, *Percent) {
	progressBar := newPercent(prefix, total, enabled)
	return utils.TeeReaderAt(r, progressBar), progressBar
}

func newPercent(prefix string, total int, enabled bool) *Percent {
	writer := os.Stderr
	return &Percent{
		prefix:  prefix,
		total:   total,
		enabled: enabled && term.IsTerminal(int(writer.Fd())),
		Writer:  writer,
	}
}

func (b *Percent) reset() {
	b.current = 0
}

func (b *Percent) Write(p []byte) (n int, err error) {
	return b.tick(p)
}

func (b *Percent) WriteAt(p []byte, _ int64) (n int, err error) {
	return b.tick(p)
}

func (b *Percent) tick(p []byte) (n int, err error) {
	if !b.enabled {
		return len(p), nil
	}

	b.current += len(p)

	err = clearCurrentLine(b.Writer)
	if err != nil {
		return 0, err
	}

	fmt.Fprintf(b.Writer, "%s%d%%", b.prefix, int(float32(b.current)/float32(b.total)*100))
	return len(p), nil
}

func (b *Percent) Done() (err error) {
	if !b.enabled {
		return nil
	}

	err = clearCurrentLine(b.Writer)
	if err != nil {
		return err
	}

	fmt.Fprintf(b.Writer, "%s100%%\n", b.prefix)
	return
}

func (b *Percent) Stop() (err error) {
	if !b.enabled {
		return nil
	}

	fmt.Fprintln(b.Writer)
	return
}
