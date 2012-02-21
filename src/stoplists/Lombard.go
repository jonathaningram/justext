package gojustext
import ( "io"; "os"; "bytes"; "compress/gzip" )

func LombardStoplist() ([]byte, os.Error) {
var gz *gzip.Decompressor
var err os.Error
if gz, err = gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x00,0x00,0x00,0x00,0xff,0x4c,0x52,
0xbf,0x6a,0x1b,0x4f,0x10,0xee,0xbf,0xc7,0x70,0x23,0xff,0x40,
0x3f,0xa5,0x4f,0xe7,0x84,0x14,0x06,0x17,0x26,0x21,0xa4,0x5e,
0xad,0xc6,0xa7,0xc1,0x7b,0xbb,0xc7,0xee,0x9e,0x8a,0xab,0xf2,
0x1a,0xe9,0x54,0x4a,0xa0,0x2a,0x02,0x43,0xe0,0xaa,0xcc,0x9b,
0xe4,0x49,0xf2,0xed,0x19,0x43,0xba,0xdb,0x99,0xf9,0x66,0xbe,
0x3f,0xb7,0x13,0x04,0x87,0x9d,0x04,0xf0,0x63,0x65,0x27,0x38,
0x68,0xc4,0xa7,0x00,0x17,0xb0,0x8a,0x0e,0xde,0xe6,0xde,0x66,
0xb0,0xaf,0xf0,0x7b,0x81,0x1f,0xf9,0x8c,0xd8,0x11,0xa3,0x18,
0x39,0xb0,0x8a,0x78,0x70,0x90,0x88,0xe7,0xfe,0xf7,0x4f,0x14,
0x61,0x11,0xdd,0x6a,0xcf,0x46,0xc0,0x20,0x99,0xab,0x63,0xd1,
0x6a,0x47,0x64,0xe9,0xd4,0x7e,0xc1,0x27,0x2e,0x1a,0xe3,0x93,
0x5d,0x08,0x76,0xdb,0xcd,0xbb,0x86,0xdb,0x60,0xc8,0x52,0xdc,
0x60,0x57,0xbf,0x47,0x19,0x89,0x63,0xdf,0x0b,0xb4,0xba,0xa0,
0x76,0x5c,0xe3,0x31,0x8f,0x07,0x22,0xbc,0x36,0x4c,0x5b,0x17,
0xbd,0x5f,0xe3,0x8b,0x8b,0x9c,0xce,0xac,0x29,0x8a,0xf1,0xba,
0xcd,0x0d,0xaa,0xa5,0x80,0x4a,0xee,0x02,0x3e,0xba,0x58,0x53,
0x5c,0x20,0xfc,0xda,0xa0,0x66,0x7b,0x39,0x38,0x0c,0x2e,0x57,
0x54,0xe7,0x1d,0x55,0x50,0xa7,0x8b,0xde,0x21,0x08,0x7a,0xa9,
0x24,0x7c,0x4b,0x2d,0x41,0x0f,0x54,0xec,0x42,0xe5,0xa9,0x99,
0x0a,0x50,0x0e,0x3a,0x4d,0x6c,0xf2,0x60,0x68,0xe2,0xee,0xd1,
0xed,0x57,0xc2,0x55,0x09,0xf7,0x74,0x43,0x4b,0xcd,0x52,0x5b,
0x8d,0xbd,0xde,0x65,0xa4,0x00,0xee,0xd2,0x9a,0x32,0x99,0x55,
0xde,0x19,0xd2,0x90,0x82,0x9b,0x94,0x6c,0x9e,0x68,0x86,0x5f,
0x2c,0xe1,0x83,0xc3,0x1b,0x7c,0x6d,0x5b,0x85,0x3a,0xee,0x16,
0x17,0xd7,0x34,0x29,0x36,0xae,0x57,0xd2,0xfa,0xf3,0xfd,0x07,
0xb5,0x70,0x47,0xe5,0x6c,0xbb,0xc0,0x17,0xfd,0x7b,0xd3,0x84,
0x91,0x36,0x67,0x65,0x6b,0x27,0xda,0x84,0xb5,0x66,0xe5,0x82,
0x80,0x87,0x36,0x3a,0xa8,0xbd,0x60,0xaf,0x91,0x97,0xec,0xdc,
0xe4,0xbd,0x12,0x29,0x2d,0x8a,0x85,0xb6,0x9d,0x2a,0x06,0x9b,
0x4b,0x91,0xf6,0x03,0x20,0xa6,0x7e,0xc9,0x82,0x61,0x24,0x66,
0xe1,0x53,0xcf,0x5b,0xc4,0x49,0xb5,0x6b,0x96,0x25,0x07,0x5d,
0x62,0x08,0x12,0xbb,0xf1,0x95,0xd0,0xc5,0x35,0xe7,0x61,0xe7,
0xb7,0x00,0x96,0xec,0x58,0x29,0xcd,0x6a,0x29,0xd4,0xd9,0x06,
0x69,0xb0,0xc6,0x8e,0xfe,0x10,0x5d,0xc8,0xea,0x51,0xa5,0xb7,
0x6b,0xac,0x6b,0x92,0x94,0x3e,0xc5,0x6a,0xe7,0xf2,0x1e,0x1f,
0x1c,0xfd,0xfe,0xff,0x9b,0xcd,0xb9,0x56,0xe9,0xb7,0x92,0xbb,
0x35,0x85,0x77,0x9a,0xb2,0x8b,0xd3,0x3f,0xe9,0x6f,0xc8,0xc0,
0x4e,0x01,0x3b,0xbb,0xe0,0xf6,0xe6,0x33,0x7f,0x2e,0xc9,0x63,
0xec,0xca,0x56,0x26,0xcd,0xcf,0x37,0xff,0x21,0x73,0x85,0x1d,
0x03,0xfe,0x06,0x00,0x00,0xff,0xff,0x1e,0x17,0x78,0xa1,0xe2,
0x02,0x00,0x00,
})); err != nil {
	return nil, err
}

var b bytes.Buffer
io.Copy(&b, gz)
gz.Close()
return b.Bytes(), nil}