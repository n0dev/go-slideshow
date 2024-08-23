package exif

type exifTagDefinition struct {
	name        string
	description string
}

// Tag is the key attribute for the value
type Tag uint16

// Hence the values
const (
	ProcessingSoftware Tag = iota
	NewSubfileType
	SubfileType
	ImageWidth
	ImageLength
	BitsPerSample
	Compression
	PhotometricInterpretation
	Orientation
	SamplesPerPixel
	PlanarConfiguration
	XResolution
	YResolution
	ResolutionUnit
	DateTime
	ImageDescription
	Make
	Model
	Software
	Artist
	JPEGProc
	JPEGInterchangeFormat
	JPEGInterchangeFormatLength
	JPEGRestartInterval
	JPEGLosslessPredictors
	JPEGPointTransforms
	JPEGQTables
	JPEGDCTables
	JPEGACTables
	YCbCrCoefficients
	YCbCrSubSampling
	YCbCrPositioning
	ReferenceBlackWhite
	XMLPacket
	Copyright
	ExifIFDpointer
	ExifVersion
	FlashpixVersion
	ColorSpace
	ComponentsConfiguration
	CompressedBitsPerPixel
	PixelXDimension
	PixelYDimension
	MakerNote
	UserComment
	RelatedSoundFile
	DateTimeOriginal
	DateTimeDigitized
	SubSecTime
	SubSecTimeOriginal
	SubSecTimeDigitized
	ImageUniqueID
	ExposureTime
	FNumber
	ExposureProgram
	SpectralSensitivity
	ISOSpeedRatings
	OECF
	SensitivityType
	ShutterSpeedValue
	ApertureValue
	BrightnessValue
	ExposureBiasValue
	MaxApertureValueS
	ubjectDistance
	MeteringMode
	LightSource
	Flash
	FocalLength
	SubjectArea
	FlashEnergy
	SpatialFrequencyResponse
	FocalPlaneXResolution
	FocalPlaneYResolution
	FocalPlaneResolutionUnit
	SubjectLocation
	ExposureIndex
	SensingMethod
	FileSource
	SceneType
	CFAPattern
	CustomRendered
	ExposureMode
	WhiteBalance
	DigitalZoomRatio
	FocalLengthIn35mmFilm
	SceneCaptureType
	GainControl
	Contrast
	Saturation
	Sharpness
	DeviceSettingDescription
	SubjectDistanceRange
	LensMake
	LensModel
	PrintImageMatching // http://www.exiv2.org/tags.html
)

var exifTags = map[Tag]exifTagDefinition{
	/////////////////////////////////////
	////////// IFD 0 ////////////////////
	/////////////////////////////////////

	0x00b0: {"ProcessingSoftware", ""},
	0x00fe: {"NewSubfileType", ""},
	0x00ff: {"SubfileType", ""},

	// image data structure for the thumbnail
	0x0100: {"ImageWidth", ""},
	0x0101: {"ImageLength", ""},
	0x0102: {"BitsPerSample", ""},
	0x0103: {"Compression", ""},
	0x0106: {"PhotometricInterpretation", ""},
	0x0112: {"Orientation", ""},
	0x0115: {"SamplesPerPixel", ""},
	0x011C: {"PlanarConfiguration", ""},
	0x011A: {"XResolution", ""},
	0x011B: {"YResolution", ""},
	0x0128: {"ResolutionUnit", ""},

	// Other tags
	0x0132: {"DateTime", ""},
	0x010E: {"ImageDescription", ""},
	0x010F: {"Make", ""},
	0x0110: {"Model", ""},
	0x0131: {"Software", ""},
	0x013B: {"Artist", ""},

	0x0200: {"JPEGProc", "This field indicates the process used to produce the compressed data"},
	0x0201: {"JPEGInterchangeFormat", "The offset to the start byte (SOI) of JPEG compressed thumbnail data. This is not used for primary image JPEG data."},
	0x0202: {"JPEGInterchangeFormatLength", "The number of bytes of JPEG compressed thumbnail data. This is not used for primary image JPEG data. JPEG thumbnails are not divided but are recorded as a continuous JPEG bitstream from SOI to EOI. Appn and COM markers should not be recorded. Compressed thumbnails must be recorded in no more than 64 Kbytes, including all other data to be recorded in APP1."},
	0x0203: {"JPEGRestartInterval", "This Field indicates the length of the restart interval used in the compressed image data."},                  //	Short
	0x0205: {"JPEGLosslessPredictors", "This Field points to a list of lossless predictor-selection values, one per component."},                   //	Short
	0x0206: {"JPEGPointTransforms", "This Field points to a list of point transform values, one per component."},                                   //	Short
	0x0207: {"JPEGQTables", "This Field points to a list of offsets to the quantization tables, one per component."},                               //	Long
	0x0208: {"JPEGDCTables", "This Field points to a list of offsets to the DC Huffman tables or the lossless Huffman tables, one per component."}, //	Long
	0x0209: {"JPEGACTables", "This Field points to a list of offsets to the Huffman AC tables, one per component."},                                //	Long
	0x0211: {"YCbCrCoefficients", ""},                                                                                                              //
	0x0212: {"YCbCrSubSampling", "Short	The sampling ratio of chrominance components in relation to the luminance component. In JPEG compressed data a JPEG marker is used instead of this tag."}, //
	0x0213: {"YCbCrPositioning", ""},    //	Short	The position of chrominance components in relation to the luminance component. This field is designated only for JPEG compressed data or uncompressed YCbCr data. The TIFF default is 1 (centered); but when Y:Cb:Cr = 4:2:2 it is recommended in this standard that 2 (co-sited) be used to record data, in order to improve the image quality when viewed on TV systems. When this field does not exist, the reader shall assume the TIFF default. In the case of Y:Cb:Cr = 4:2:0, the TIFF default (centered) is recommended. If the reader does not have the capability of supporting both kinds of <YCbCrPositioning>, it shall follow the TIFF default regardless of the value in this field. It is preferable that readers be able to support both centered and co-sited positioning.
	0x0214: {"ReferenceBlackWhite", ""}, //	Rational	The reference black point value and reference white point value. No defaults are given in TIFF, but the values below are given as defaults here. The color space is declared in a color space information tag, with the default being the value that gives the optimal image characteristics Interoperability these conditions.
	0x02bc: {"XMLPacket", ""},           //	Byte	XMP Metadata (Adobe technote 9-14-02)

	0x8298: {"Copyright", ""},

	0x8769: {"Exif IFD pointer", ""},

	0x9000: {"ExifVersion", ""},
	0xA000: {"FlashpixVersion", ""},

	0xA001: {"ColorSpace", ""},

	0x9101: {"ComponentsConfiguration", ""},
	0x9102: {"CompressedBitsPerPixel", ""},
	0xA002: {"PixelXDimension", ""},
	0xA003: {"PixelYDimension", ""},

	0x927C: {"MakerNote", ""},
	0x9286: {"UserComment", ""},

	0xA004: {"RelatedSoundFile", ""},
	0x9003: {"DateTimeOriginal", ""},
	0x9004: {"DateTimeDigitized", ""},
	0x9290: {"SubSecTime", ""},
	0x9291: {"SubSecTimeOriginal", ""},
	0x9292: {"SubSecTimeDigitized", ""},

	0xA420: {"ImageUniqueID", ""},

	// picture conditions
	0x829A: {"ExposureTime", ""},
	0x829D: {"FNumber", ""},
	0x8822: {"ExposureProgram", ""},
	0x8824: {"SpectralSensitivity", ""},
	0x8827: {"ISOSpeedRatings", ""},
	0x8828: {"OECF", ""},
	0x8830: {"SensitivityType", ""},
	0x9201: {"ShutterSpeedValue", ""},
	0x9202: {"ApertureValue", ""},
	0x9203: {"BrightnessValue", ""},
	0x9204: {"ExposureBiasValue", ""},
	0x9205: {"MaxApertureValue", ""},
	0x9206: {"SubjectDistance", ""},
	0x9207: {"MeteringMode", ""},
	0x9208: {"LightSource", ""},
	0x9209: {"Flash", ""},
	0x920A: {"FocalLength", ""},
	0x9214: {"SubjectArea", ""},
	0xA20B: {"FlashEnergy", ""},
	0xA20C: {"SpatialFrequencyResponse", ""},
	0xA20E: {"FocalPlaneXResolution", ""},
	0xA20F: {"FocalPlaneYResolution", ""},
	0xA210: {"FocalPlaneResolutionUnit", ""},
	0xA214: {"SubjectLocation", ""},
	0xA215: {"ExposureIndex", ""},
	0xA217: {"SensingMethod", ""},
	0xA300: {"FileSource", ""},
	0xA301: {"SceneType", ""},
	0xA302: {"CFAPattern", ""},
	0xA401: {"CustomRendered", ""},
	0xA402: {"ExposureMode", ""},
	0xA403: {"WhiteBalance", ""},
	0xA404: {"DigitalZoomRatio", ""},
	0xA405: {"FocalLengthIn35mmFilm", ""},
	0xA406: {"SceneCaptureType", ""},
	0xA407: {"GainControl", ""},
	0xA408: {"Contrast", ""},
	0xA409: {"Saturation", ""},
	0xA40A: {"Sharpness", ""},
	0xA40B: {"DeviceSettingDescription", ""},
	0xA40C: {"SubjectDistanceRange", ""},
	0xa430: {"CameraOwnerName", ""},
	0xa431: {"BodySerialNumber", ""},
	0xa432: {"LensSpecification", ""},
	0xA433: {"LensMake", ""},
	0xA434: {"LensModel", ""},

	0xc4a5: {"PrintImageMatching", ""}, // http://www.exiv2.org/tags.html
}

var casioMarkerTags = map[Tag]exifTagDefinition{ // http://owl.phy.queensu.ca/~phil/exiftool/TagNames/Casio.html
	0x0001: {"RecordingMode", ""},
	0x0002: {"PreviewImageSize", ""},
	0x0003: {"PreviewImageLength", ""},
	0x0004: {"PreviewImageStart", ""},
	0x0005: {"FlashIntensity", ""},
	0x0008: {"QualityMode", ""},
	0x0009: {"CasioImageSize", ""},
	0x000d: {"FocusMode", ""},
	0x0014: {"ISO", ""},
	0x0019: {"WhiteBalance", ""},
	0x001d: {"FocalLength", ""},
	0x001f: {"Saturation", ""},
	0x0020: {"Contrast", ""},
	0x0021: {"Sharpness", ""},
	0x0e00: {"PrintIM", ""}, //	-	--> PrintIM Tags
	0x2000: {"PreviewImage", ""},
	0x2001: {"FirmwareDate", ""},
	0x2011: {"WhiteBalanceBias", ""},
	0x2012: {"WhiteBalance", ""},
	0x2021: {"AFPointPosition", ""},
	0x2022: {"ObjectDistance", ""},
	0x2034: {"FlashDistance", ""},
	0x2076: {"SpecialEffectMode", ""},
	/*0x2089:	FaceInfo1
	  FaceInfo2
	  FaceInfoUnknown?	-
	  -
	  Y	--> Casio FaceInfo1 Tags
	  --> Casio FaceInfo2 Tags*/
	0x211c: {"FacesDetected", ""},
	0x3000: {"RecordMode", ""},
	0x3001: {"ReleaseMode", ""},
	0x3002: {"Quality", ""},
	0x3003: {"FocusMode", ""},
	0x3006: {"HometownCity", ""},
	0x3007: {"BestShot Mode", ""},
	0x3008: {"AutoISO", ""},
	0x3009: {"AFMode", ""},
	0x3011: {"Sharpness", ""},
	0x3012: {"Contrast", ""},
	0x3013: {"Saturation", ""},
	0x3014: {"ISO", ""},
	0x3015: {"ColorMode", ""},
	0x3016: {"Enhancement", ""},
	0x3017: {"ColorFilter", ""},
	0x301b: {"ArtMode", ""},
	0x301c: {"SequenceNumber", ""},
	0x301d: {"BracketSequence", ""},
	0x3020: {"ImageStabilization", ""},
	0x302a: {"LightingMode", ""},
	0x302b: {"PortraitRefiner", ""},
	0x3030: {"SpecialEffectLevel", ""},
	0x3031: {"SpecialEffectSetting", ""},
	0x3103: {"DriveMode", ""},
	0x310b: {"ArtModeParameters", ""},
	0x4001: {"CaptureFrameRate", ""},
	0x4003: {"VideoQuality", ""},
}

// string returns the name for the tag
func (t Tag) string(casio bool) string {
	if casio {
		if val, ok := casioMarkerTags[t]; ok {
			return val.name
		}
	}
	if val, ok := exifTags[t]; ok {
		return val.name
	}
	return "UnknownTag"
}
