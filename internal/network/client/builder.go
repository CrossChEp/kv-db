package client

type Option func(opt *options)

type options struct {
	bufferSize int
}

func WithBufferSize(bs int) Option {
	return func(opt *options) {
		opt.bufferSize = bs
	}
}
