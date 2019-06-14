package bin2array

import (
	"io"
	"fmt"
	"errors"
	"encoding/binary"
)

type Options struct {
	OutStream io.Writer
	InStream io.Reader
	Size int64
}

type Converter struct {
	opt *Options
}



func New(o *Options) (c *Converter, err error) {
	if (o == nil) {
		return nil, errors.New("Bad Argument") 
	}
	return &Converter{opt: o}, nil
}

func (c *Converter) Run() error {
	b := make([]byte, 16)
	r := c.opt.Size

	fmt.Fprintf(c.opt.OutStream, "{\n")
	for {
		if (r < 16) {
			b = make([]byte, r)
			binary.Read(c.opt.InStream, binary.LittleEndian, b)

			for _, x := range b {
				fmt.Fprintf(c.opt.OutStream, "0x%02x, ", x)
			}
			fmt.Fprintf(c.opt.OutStream, "\n")
			
			break;
		} else {
			binary.Read(c.opt.InStream, binary.LittleEndian, b)
			
			for _, x := range b {
				fmt.Fprintf(c.opt.OutStream, "0x%02x, ", x)
			}
			fmt.Fprintf(c.opt.OutStream, "\n")
			r -= 16
		}
		
	}
	fmt.Fprintf(c.opt.OutStream, "}")

	return nil
}
