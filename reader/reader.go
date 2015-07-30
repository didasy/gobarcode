package reader

type ReadWriter struct {
	OutChan chan string
	UseChan bool
}

func (r *ReadWriter) Write(d []byte) (int, error) {
	if r.UseChan {
		r.OutChan <- string(d)
	}
	return len(d), nil
}

func (r *ReadWriter) SetOutChannel(c chan string) {
	r.OutChan = c
	r.UseChan = true
}

func New() *ReadWriter {
	return &ReadWriter{}
}
