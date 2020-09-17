package replacer

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestCorrectness(t *testing.T) {
	var bigFileName = "feed.xml"
	var src = []byte("mvideo.ru")
	var dst = []byte("kazanexpress.ru")

	var data, err = ioutil.ReadFile(bigFileName)
	if err != nil {
		t.Fatal(err)
	}

	var r1 = bytes.NewReader(data)
	var r2 = bytes.NewReader(data)

	sr, err := SlowReplace(r1, src, dst)
	if err != nil {
		t.Fatal(err)
	}

	changedSlow, err := ioutil.ReadAll(sr)
	if err != nil {
		t.Fatal(err)
	}

	var fastR = NewReaderReplacer(r2, src, dst)
	changedFast, err := ioutil.ReadAll(fastR)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(changedSlow, changedFast) {
		// ioutil.WriteFile("slow.xml", changedSlow, os.ModePerm)
		// ioutil.WriteFile("fast.xml", changedFast, os.ModePerm)
		t.Fatal("arrays are not equal")
	}
}
