package gojustext
import ( "io"; "os"; "bytes"; "compress/gzip" )

func Low_SaxonStoplist() ([]byte, os.Error) {
var gz *gzip.Decompressor
var err os.Error
if gz, err = gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x00,0x00,0x00,0x00,0xff,0x34,0x53,
0x4b,0x72,0xd3,0x40,0x14,0xdc,0xf7,0x49,0x42,0x95,0xa3,0x3b,
0x24,0x38,0x45,0x42,0x02,0x45,0x61,0x08,0x55,0xec,0x46,0xd1,
0x93,0x34,0x68,0x3e,0xae,0x99,0x91,0x5d,0x64,0xc5,0x35,0xd8,
0x73,0x02,0x16,0x5e,0x79,0x85,0x6e,0xc2,0x49,0xe8,0x37,0x22,
0x55,0xb6,0x65,0xcd,0xa7,0x5f,0xbf,0xee,0x7e,0x9d,0xe0,0x30,
0x07,0xf0,0xd3,0x99,0x02,0x1b,0x60,0x33,0x24,0x60,0x2b,0xe8,
0xf8,0x18,0xa5,0x14,0x98,0x8c,0x12,0xe1,0x6d,0xe1,0x6b,0xdb,
0xea,0x2f,0xe2,0x84,0xbc,0x9c,0x43,0x87,0xa3,0x48,0x42,0xb6,
0x13,0xe6,0x82,0xbb,0x70,0x34,0x63,0x90,0x94,0x11,0xf7,0xd8,
0x12,0xcd,0x04,0x64,0xc1,0x3b,0x1b,0xf2,0xd3,0x48,0xb0,0xd6,
0xe2,0xef,0x8f,0x9f,0x01,0xfd,0x72,0x4a,0x08,0xf1,0x69,0xe4,
0x45,0x2e,0xdb,0x50,0x57,0x63,0x47,0xa4,0x60,0xb9,0x7a,0x34,
0x29,0x15,0x1c,0x63,0x4a,0xa1,0xc1,0x23,0x99,0xdd,0x05,0x42,
0x11,0x58,0x20,0x3c,0x4f,0x5e,0x01,0xb7,0x82,0x87,0xe5,0xbc,
0x9c,0x3b,0x04,0x83,0x1c,0x71,0x9f,0x84,0xc4,0xbb,0x98,0x70,
0xd5,0x67,0x37,0xe7,0x8c,0xaf,0x12,0xf2,0x9c,0xf1,0xde,0x4c,
0xc6,0xf3,0x16,0xd9,0x12,0x81,0xd8,0x05,0x57,0x99,0x54,0x0a,
0x3c,0xe9,0xcd,0x7b,0xb6,0x56,0xcb,0x77,0x4a,0xea,0xa3,0xb4,
0x22,0x54,0x21,0x37,0xb5,0xb1,0x0d,0x5e,0xc7,0x39,0x94,0xef,
0xac,0x5e,0xcf,0x98,0x03,0x11,0x06,0xdb,0xf7,0x05,0x6f,0xe3,
0x98,0xb0,0x13,0x4c,0xfe,0xcf,0xef,0x06,0xed,0x5c,0x58,0x62,
0xb7,0x4f,0xc6,0x4c,0xb8,0xd8,0x15,0x43,0x65,0x96,0x93,0x9e,
0xde,0x59,0x0a,0xf8,0x28,0xe9,0x68,0x9c,0x1e,0x39,0xb0,0xa6,
0xe7,0xa1,0xd2,0xe0,0xf3,0xee,0x0a,0x23,0x3b,0x65,0x2f,0x36,
0x97,0x98,0xac,0x34,0xab,0xbe,0x8d,0x1a,0x52,0xeb,0x6d,0xc5,
0xc7,0x21,0x99,0x5e,0xb7,0x96,0xb3,0x5f,0x95,0x2a,0x24,0x5f,
0x77,0x73,0x30,0x4f,0x13,0xdb,0x10,0x32,0x19,0x95,0x58,0xd2,
0x1e,0xb5,0x2f,0xe2,0x6f,0xf0,0xc9,0xb2,0x13,0x27,0x07,0x56,
0x95,0x20,0x3c,0xe1,0x04,0x6d,0x4c,0x7c,0xcd,0xc5,0x78,0x5f,
0x70,0x1f,0xbd,0xa7,0xb8,0x6f,0xd2,0xbc,0xdf,0x73,0xf5,0x3e,
0xee,0xf7,0x95,0xd0,0x7f,0xa9,0xd4,0x33,0xe9,0x1c,0x0d,0x11,
0x8a,0x68,0x3c,0xbe,0x09,0x75,0xc8,0x32,0xd2,0x14,0x76,0xd8,
0x31,0x07,0xcb,0x69,0x60,0xdd,0x2a,0x8e,0xe0,0x26,0x0c,0xe2,
0x78,0x49,0x83,0xe1,0xb4,0xea,0x4b,0x1c,0x8a,0x71,0x8e,0x87,
0x06,0xa5,0xb2,0xa1,0xae,0x9d,0x1a,0xbd,0xb6,0x58,0xa8,0xc2,
0xe5,0xf5,0x1c,0xba,0x4c,0x4e,0xd5,0x19,0x0d,0xc4,0x56,0x5d,
0x4c,0xad,0x58,0x4a,0xf1,0x48,0x5b,0x96,0x93,0x73,0x42,0x9d,
0x26,0x86,0x80,0xa8,0xab,0x9f,0x25,0x8e,0x33,0xed,0xa5,0x70,
0x32,0x49,0x50,0x3d,0x83,0xa5,0xe0,0xd9,0x96,0xe5,0x17,0xbb,
0xbf,0xb5,0xc3,0x48,0x32,0x31,0x3a,0xb4,0x2e,0x96,0xcc,0x7b,
0x9d,0x35,0xda,0x56,0x83,0x2f,0xf4,0x44,0x36,0x15,0xba,0xa7,
0x8d,0x04,0x23,0x8b,0x6d,0x74,0xce,0xa4,0x06,0x0f,0xa6,0x90,
0x03,0x4b,0x24,0x3b,0x19,0xd2,0xf0,0x31,0x12,0xfc,0xda,0x50,
0x06,0x17,0xd3,0xe5,0x4b,0xba,0xf8,0x54,0x15,0xe7,0x30,0xbc,
0x30,0xca,0x03,0x91,0x9e,0x71,0x63,0x9d,0x66,0x95,0x56,0x47,
0x7e,0x35,0x54,0xc7,0xb9,0x06,0x79,0xc7,0x58,0x6b,0x59,0xb5,
0xeb,0x62,0xcd,0xd5,0x2b,0x18,0xb7,0x6e,0xd7,0x3c,0xe9,0x2c,
0xe4,0x4d,0x9d,0xb9,0x06,0x9d,0xcd,0x9c,0x1d,0x67,0x87,0x81,
0x73,0x44,0x01,0xc5,0xd7,0x48,0x06,0x7c,0x30,0xcf,0xb6,0xb7,
0x13,0x13,0xaa,0x03,0x82,0xc2,0x55,0x06,0x45,0x1c,0x0d,0xc0,
0xb5,0x45,0x98,0x39,0x11,0x9c,0x21,0xa2,0xc2,0xf4,0xbc,0x43,
0xcd,0x0e,0xba,0x5d,0x63,0x51,0x27,0x8e,0x3e,0xb6,0xb4,0x80,
0xb2,0xaf,0xa3,0xa6,0x4d,0xeb,0xff,0xb8,0xce,0x74,0xc3,0x29,
0x4b,0xb9,0x08,0xfe,0x05,0x00,0x00,0xff,0xff,0x08,0xd0,0x27,
0x96,0x1b,0x04,0x00,0x00,
})); err != nil {
	return nil, err
}

var b bytes.Buffer
io.Copy(&b, gz)
gz.Close()
return b.Bytes(), nil}