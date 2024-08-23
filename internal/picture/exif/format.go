package exif

// List all possible format for exif data
type exifFormat int

type formatDefinition struct {
	name string
	size uint
}

const (
	_ exifFormat = iota
	exifFormatByte
	exifFormatASCII
	exifFormatShort
	exifFormatLong
	exifFormatRational
	exifFormatSbyte
	exifFormatUndefined
	exifFormatSshort
	exifFormatSlong
	exifFormatSrational
	exifFormatFload
	exifFormatDouble
)

var exifFormats = []formatDefinition{
	exifFormatByte:      {"Byte", 1},
	exifFormatASCII:     {"ASCII", 1},
	exifFormatShort:     {"Short", 2},
	exifFormatLong:      {"Long", 4},
	exifFormatRational:  {"Rational", 8},
	exifFormatSbyte:     {"Sbyte", 1},
	exifFormatUndefined: {"Undefined", 1},
	exifFormatSshort:    {"Sshort", 2},
	exifFormatSlong:     {"Slong", 4},
	exifFormatSrational: {"Srational", 8},
	exifFormatFload:     {"Fload", 4},
	exifFormatDouble:    {"Double", 8},
}

// string returns the string name of the format field
func (f exifFormat) string() string { return exifFormats[f].name }

// size returns the size of the format field
func (f exifFormat) size() uint { return exifFormats[f].size }
