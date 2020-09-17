package replacer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
)

type ReaderReplacer struct {
	source      io.Reader
	from        []byte
	to          []byte
	readyBuffer bytes.Buffer
	eof         bool
}

func NewReaderReplacer(reader io.Reader, from, to []byte) io.Reader {
	return &ReaderReplacer{
		source: bufio.NewReader(reader),
		from:   from,
		to:     to,
	}
}

// ust read 64k at a time and run the regex on the entire block.
// The only thing to keep in mind is that you'd want to ensure you have overlapping bytes so that if a number happens to be split on a 64k boundary you still find it. This is pretty easy to do though. just need to
//     make a 64k+32 byte buffer
//     read the first 64k into it
//     attempt match
//     copy the last 32 bytes to the beginning
//     read the next 64k into the buffer[32:]
//     attempt match
// https://www.reddit.com/r/golang/comments/9htl34/streaming_regex_with_ioreader/

func (rr *ReaderReplacer) Read(p []byte) (n int, err error) {
	var needToRead = len(p)
	if rr.readyBuffer.Len() >= needToRead {
		var n, err = rr.readyBuffer.Read(p)
		// returing it as it is, since there should not be err and n should be eq to needToRead
		return n, err
	}

	// to find token reading buffer should be at least as len of replacing token
	var tokenLen = len(rr.from)
	var readBuffer = make([]byte, needToRead+tokenLen*2)
	var readCnt, readErr = rr.source.Read(readBuffer)
	if readErr == io.EOF {
		// return what left in ready buffer
		return rr.readyBuffer.Read(p)
	} else if readErr != nil {
		// this should not happen
		log.Printf("some shit happened - %v", readErr)
		return readCnt, readErr
	}

	var temp bytes.Buffer
	var _, werr = temp.ReadFrom(&rr.readyBuffer)
	if werr != nil {
		return 0, fmt.Errorf("some shit happened - %w", werr)
	}

	_, werr = temp.Write(readBuffer[:readCnt])
	if werr != nil {
		return 0, fmt.Errorf("some shit happened - %w", werr)
	}

	var total = temp.Bytes()
	total = bytes.ReplaceAll(total, rr.from, rr.to)

	_, werr = rr.readyBuffer.Write(total)
	if werr != nil {
		return 0, fmt.Errorf("some shit happened - %w", werr)
	}

	return rr.readyBuffer.Read(p)
}
