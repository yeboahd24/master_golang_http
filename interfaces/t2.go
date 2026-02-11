package main

type Writer interface {
	Write(p []byte) (n int, err error)
}

type File struct {
	path string
}

func (f *File) Write(p []byte) (n int, err error) {
	return len(p), nil
}

type Buffer struct {
	buf []byte
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	b.buf = append(b.buf, p...)
	return len(p), nil
}

func SaveLog(w Writer, msg string) error {
	_, err := w.Write([]byte(msg))
	return err
}

func main() {
	f := &File{path: "log.txt"}
	b := &Buffer{}
	SaveLog(f, "hello")
	SaveLog(b, "world")
}
