// Package exif implements ...
package exif

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// Reader is the structure useful for reading exif data
type Reader struct {
	file os.File
	app1 []byte
}

// Open creates a new reader from the image path given in parameter
func Open(name string) (*Reader, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var r = &Reader{file: *f}
	Read(r)
	return r, nil
}

// All returns an array of all tags in this reader
func (r *Reader) All() []int {
	return []int{0, 1}
}

// Get returns the value for the corresponding tag. Otherwise returns an error
func (r *Reader) Get(tag int) (interface{}, error) {
	return "", nil
}

//-----------------------------------------------------------------------------

// Writer is the structure useful for write exif data
type Writer struct {
	//
}

// NewWriter create an empty new writer to be filled with exif data.
func NewWriter() *Writer {
	return &Writer{}
}

//-----------------------------------------------------------------------------

func get(f *os.File, s uint) []byte {
	b := make([]byte, s)
	if _, err := f.Read(b); err != nil {
		fmt.Println(err)
	}
	return b
}

func extract(bo binary.ByteOrder, info []byte, data []byte, casio bool, prefix string) {
	i := Tag(toUint16(bo, info[0:2]))
	t := exifFormat(toUint16(bo, info[2:4]))
	c := toUint32(bo, info[4:8])
	e := toUint32(bo, info[8:12])
	//fmt.Printf(hex.Dump(info))

	if i == 34665 {
		//readIFD(bo, data, e, "SubIFD")

	} else if i == 34853 {
		fmt.Println("-- GPS --")
		fmt.Println("-- End GPS --")

	} else if i == 40965 {
		fmt.Println("-- Interop --")
		fmt.Printf(hex.Dump(info))
		fmt.Println("-- End Interop --")

	} else if i == 37500 {

		// Markernote, specific to manufacturer
		if string(data[e:e+6]) == "QVC\x00\x00\x00" { // CASIO
			readMarkerNote(bo, data, e+6)
		}

	} else {

		switch t {
		case exifFormatASCII:

			var value string
			if uint32(t.size())*c <= 4 {
				value = strings.Trim(string(info[8:12]), " \r\n\x00")
			} else {
				value = strings.Trim(string(data[e:e+c]), " \r\n\x00")
			}
			fmt.Printf("[%s] (0x%04x) %s: %s\n", prefix, i, i.string(casio), value)

		case exifFormatShort:
			value := toUint16(bo, []byte{info[8], info[9]})
			fmt.Printf("[%s] (0x%04x) %s: %d\n", prefix, i, i.string(casio), value)

		case exifFormatLong:
			fmt.Printf("[%s] (0x%04x) %s: %d\n", prefix, i, i.string(casio), toUint32(bo, info[8:12]))

		case exifFormatSlong:
			fmt.Printf("[%s] (0x%04x) %s: %d\n", prefix, i, i.string(casio), toInt32(bo, info[8:12]))

		case exifFormatByte:
			if uint32(t.size())*c <= 4 {
				fmt.Printf("[%s] (0x%04x) %s : %x\n", prefix, i, i.string(casio), info[8:12])
			} else {
				fmt.Printf("[%s] (0x%04x) %s: %x\n", prefix, i, i.string(casio), data[e:e+c])
			}

		case exifFormatRational:
			n := toUint32(bo, data[e:e+(c*4)])
			d := toUint32(bo, data[e+(c*4):e+(c*8)])
			fmt.Printf("[%s] (0x%04x) %s: %d/%d\n", prefix, i, i.string(casio), n, d)

		case exifFormatSrational:
			n := toInt32(bo, data[e:e+(c*4)])
			d := toInt32(bo, data[e+(c*4):e+(c*8)])
			fmt.Printf("[%s] (0x%04x) %s: %d/%d\n", prefix, i, i.string(casio), n, d)

		case exifFormatUndefined:
			var value string
			if uint32(t.size())*c <= 4 {
				value = strings.Trim(string(info[8:12]), " \r\n\x00")
			} else {
				value = strings.Trim(string(data[e:e+c]), " \r\n\x00")
			}
			fmt.Printf("[%s] (0x%04x) %s: %s\n", prefix, i, i.string(casio), value)

		default:
			fmt.Printf(hex.Dump(info))
			fmt.Printf("[%s]555 (0x%04x) %s: %s - %d at %d\n", prefix, i.string(casio), i, t.string(), c, e)
		}
	}
}

func readMarkerNote(bo binary.ByteOrder, data []byte, off uint32) {
	var n int16
	b := bytes.NewReader(data[off : off+2])
	err := binary.Read(b, bo, &n)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	pos := off + 2
	for i := 0; i < int(n); i++ {
		extract(bo, data[pos:pos+12], data, true, "Makernote")
		pos += 12
	}
}

func readIFD(bo binary.ByteOrder, data []byte, off uint32, prefix string) {
	var n int16
	b := bytes.NewReader(data[off : off+2])
	err := binary.Read(b, bo, &n)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	pos := off + 2
	for i := 0; i < int(n); i++ {
		extract(bo, data[pos:pos+12], data, false, prefix)
		pos += 12
	}
}

// Read displays exif information about the file
func Read(reader *Reader) {

	APP1, _ := hex.DecodeString("FFE1")
	APP1Header, _ := hex.DecodeString("457869660000")
	APP2, _ := hex.DecodeString("FFE2")
	//FPXR, _ := hex.DecodeString("465058520000")
	DQT, _ := hex.DecodeString("FFDB")
	inExif := false

	buf := make([]byte, 2)

	for {
		_, err := reader.file.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}

		if bytes.Equal(buf, APP1) {
			length := make([]byte, 2)
			if _, err := reader.file.Read(length); err != nil {
				fmt.Println(err)
				break
			}

			var l int16
			b := bytes.NewReader(length)
			err := binary.Read(b, binary.BigEndian, &l) // Always big endian
			if err != nil {
				fmt.Println("binary.Read failed:", err)
			}
			//fmt.Printf("Size: %d\n", l)

			exifTag := make([]byte, 6)
			if _, err := reader.file.Read(exifTag); err != nil {
				fmt.Println(err)
				break
			}

			if bytes.Equal(exifTag, APP1Header) {

				app1Info := make([]byte, l)
				if _, err := reader.file.Read(app1Info); err != nil {
					fmt.Println(err)
					break
				}

				// Compute endianess of the block
				var bo binary.ByteOrder
				switch string(app1Info[0:2]) {
				case "II":
					bo = binary.LittleEndian
				case "MM":
					bo = binary.BigEndian
				default:
					fmt.Printf("Wrong header \"%s\", abort!\n", string(app1Info[0:2]))
				}

				readIFD(bo, app1Info, 8, "IFD0")
			}
		}
		if bytes.Equal(buf, APP2) {
			fmt.Println("Welcome APP2")
		}
		if inExif {
			//
		}
		if bytes.Equal(buf, DQT) {
			return
			//inExif = false
		}
	}

}
