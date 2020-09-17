package replacer

import (
	"bytes"
	"io"
	"io/ioutil"
)

func SlowReplace(data io.Reader, src, dst []byte) (io.Reader, error) {
	var bs, err = ioutil.ReadAll(data)
	if err != nil {
		return nil, err
	}
	bs = bytes.ReplaceAll(bs, src, dst)
	return bytes.NewBuffer(bs), nil
}
