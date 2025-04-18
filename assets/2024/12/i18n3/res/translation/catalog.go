// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package translation

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"de_DE": &dictionary{index: de_DEIndex, data: de_DEData},
		"en_US": &dictionary{index: en_USIndex, data: en_USData},
		"hr_HR": &dictionary{index: hr_HRIndex, data: hr_HRData},
	}
	fallback := language.MustParse("en-US")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"Welcome to i18n\n": 0,
}

var de_DEIndex = []uint32{ // 2 elements
	0x00000000, 0x00000000,
} // Size: 32 bytes

const de_DEData string = ""

var en_USIndex = []uint32{ // 2 elements
	0x00000000, 0x00000015,
} // Size: 32 bytes

const en_USData string = "\x04\x00\x01\n\x10\x02Welcome to i18n"

var hr_HRIndex = []uint32{ // 2 elements
	0x00000000, 0x00000009,
} // Size: 32 bytes

const hr_HRData string = "\x04\x00\x01\n\x04\x02qwe"

// Total table size 126 bytes (0KiB); checksum: 6BDA8438
