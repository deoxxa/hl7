//+build gofuzz

package hl7

// If you want to fuzz, call "go generate -tags=gofuzz".
// Maybe you'll need the go-fuzz and go-fuzz-build binaries:
// go get github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build

//go:generate go-fuzz-build github.com/medbridge/hl7
//go:generate mkdir -p /tmp/fuzz-hl7/corpus
//go:generate sh -c "cp -a testdata/* /tmp/fuzz-hl7/corpus/"
//go:generate echo go-fuzz -bin=hl7-fuzz.zip -workdir=/tmp/fuzz-hl7

func Fuzz(input []byte) int {
	if _, _, err := ParseMessage(input); err != nil {
		return 0
	}

	return 1
}
