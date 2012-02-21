package gojustext
import ( "io"; "os"; "bytes"; "compress/gzip" )

func CatalanStoplist() ([]byte, os.Error) {
var gz *gzip.Decompressor
var err os.Error
if gz, err = gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x00,0x00,0x00,0x00,0xff,0x4c,0x94,
0xbd,0x6e,0xeb,0x46,0x10,0x85,0xfb,0x79,0x11,0xa5,0x10,0xf4,
0x00,0xe9,0x0c,0xc8,0x09,0x2e,0x20,0x04,0x01,0x6e,0x92,0x7e,
0x44,0xae,0xa8,0xbd,0xd9,0x1f,0x7a,0x77,0x49,0x38,0x8f,0x90,
0xb7,0x70,0x69,0x17,0x6e,0xec,0x36,0x55,0xf8,0x62,0xf9,0x86,
0x94,0x9d,0xdb,0x08,0xdc,0xbf,0x33,0xe7,0x9c,0x39,0xa3,0xde,
0x49,0x50,0xf1,0xa2,0xe2,0x82,0x3c,0x4c,0x4e,0x5c,0x92,0x9e,
0xcf,0x59,0x65,0x4a,0x32,0xba,0x22,0xc1,0x55,0x3e,0x55,0x34,
0x9e,0x45,0x03,0xd7,0xaa,0xb0,0xb3,0xbc,0x54,0xbb,0x57,0xe5,
0x3e,0x48,0xb4,0x45,0x97,0xa3,0x9c,0x54,0xb2,0x54,0xde,0x54,
0x37,0x89,0x2b,0x2a,0x29,0x03,0x94,0x58,0x02,0x77,0xc9,0x93,
0x5c,0x3d,0xf8,0xad,0x38,0x19,0xf3,0x39,0x68,0xe7,0x97,0x77,
0xde,0x83,0x91,0x44,0xd3,0x5f,0x15,0xf8,0x6a,0x25,0x97,0x37,
0x99,0xfd,0xec,0x61,0x72,0xe5,0x99,0x4f,0xeb,0x66,0xcd,0x89,
0xb2,0x77,0xd2,0xe0,0xb1,0xbc,0x48,0x73,0xc9,0x2b,0xfb,0x54,
0xf7,0xed,0x9b,0xc2,0xaf,0x2e,0xef,0x49,0x3a,0xed,0x79,0xa2,
0x71,0x79,0x0d,0x1c,0x17,0x97,0x58,0x66,0x74,0x68,0x69,0xf2,
0xeb,0x46,0x0c,0xaa,0x3a,0x82,0x3c,0x73,0xa1,0x77,0x75,0x2c,
0xab,0x94,0x1d,0x6a,0xfb,0xb5,0xc4,0x1f,0x40,0xe5,0x33,0x1c,
0x39,0x94,0x7b,0x48,0xe1,0x4a,0x6d,0x12,0x73,0x68,0xd2,0x72,
0xc3,0x24,0x5c,0xf8,0xa6,0x32,0x14,0x94,0x75,0x7e,0x6a,0xda,
0xe4,0xc4,0x55,0x83,0x30,0xc5,0x71,0x7d,0x3b,0x71,0xda,0x28,
0x37,0x04,0x67,0x7b,0x60,0x57,0x89,0xce,0x44,0x5e,0xa0,0xa1,
0x01,0x13,0xaa,0x7c,0x75,0x43,0x46,0x9d,0x69,0xdf,0xcb,0x58,
0x7c,0xe4,0x08,0x78,0x60,0x1d,0xdb,0x61,0xc7,0x81,0xb4,0x92,
0xcf,0x2a,0xcb,0xdf,0xac,0xbd,0xe1,0x2d,0xff,0xe0,0x20,0xe6,
0xe2,0x8e,0x29,0x5f,0x9e,0xe8,0x86,0x51,0xd8,0x68,0xc2,0xdd,
0xcd,0x20,0x8f,0x37,0xa2,0x6b,0xd1,0xd5,0x70,0xeb,0xd5,0x55,
0xcf,0x9e,0xbb,0x03,0x17,0x3e,0x3f,0xa5,0xe1,0xe6,0x57,0x23,
0x7b,0x6f,0x38,0xf0,0xcb,0x25,0x92,0x86,0xda,0x80,0xee,0x32,
0x8f,0x55,0xee,0x02,0xec,0x32,0x4d,0x70,0xf5,0x20,0x51,0x9b,
0xf3,0x8f,0x12,0xa7,0xe4,0x3b,0x3f,0x7a,0x53,0x67,0x0e,0xba,
0xd4,0x29,0x57,0x37,0x15,0x0a,0xed,0x95,0x45,0x6f,0x71,0x82,
0x18,0xad,0x80,0xbd,0xbf,0x95,0x4d,0x0d,0x9c,0x10,0x72,0xf7,
0xff,0x1a,0xde,0xf4,0x8d,0xdf,0xc4,0x47,0xd8,0x6d,0x9a,0x6a,
0x46,0x90,0xb9,0x73,0xa0,0xdf,0x71,0x84,0x9b,0x0f,0x41,0x7e,
0x4f,0x9e,0xfb,0xc7,0x8f,0xd6,0x59,0x12,0x8a,0x7c,0x91,0xdf,
0x3e,0x82,0x51,0xa2,0x23,0xb0,0xd8,0xea,0x1f,0x97,0xd7,0x95,
0x43,0x59,0x3d,0x9a,0x2d,0x50,0x39,0x42,0xc6,0x9e,0xb0,0x00,
0x61,0x79,0xee,0xbc,0x03,0xd1,0x3a,0x6b,0x24,0x07,0x8b,0xe4,
0x95,0x0e,0x8c,0xa6,0xf1,0xe2,0x1a,0x6d,0x9e,0x46,0x34,0xa7,
0x2d,0xce,0x93,0xdc,0x6d,0x71,0xb8,0x65,0x0c,0xb0,0x98,0xc9,
0x56,0xcc,0xd5,0x8c,0xc2,0x3b,0x70,0x8b,0x9f,0x90,0x90,0x4b,
0x8f,0xef,0x1a,0x06,0xe3,0x82,0x59,0xab,0xb9,0x89,0x5d,0xe9,
0x2d,0xd2,0x38,0x53,0x6d,0x8f,0x40,0xe3,0xfd,0xa7,0x0f,0x7b,
0x4e,0x67,0x4b,0x3b,0x59,0xbf,0x75,0x34,0x59,0x18,0xb4,0x14,
0x7f,0xd6,0xb2,0x09,0x38,0x88,0x9e,0x95,0xcd,0xe3,0x96,0xb2,
0x87,0x69,0x79,0xc6,0xdc,0x36,0x81,0x66,0xbd,0x38,0xc8,0x90,
0xc1,0x48,0xf2,0x67,0xfc,0xf7,0x0d,0xeb,0xfc,0x48,0xe8,0xab,
0xaf,0x78,0x48,0x83,0x18,0x05,0xf3,0xeb,0x8b,0x0c,0x93,0x2b,
0x70,0x3e,0xba,0x98,0xc9,0xf2,0xc5,0xeb,0xe1,0x63,0x2c,0x2d,
0x94,0x08,0xf4,0xcc,0x50,0x01,0xad,0x4e,0xbd,0x9c,0xd6,0x30,
0x1e,0x2d,0x45,0x08,0x06,0xe0,0x07,0x84,0x82,0xe5,0x82,0xfd,
0x3f,0x84,0xe0,0x99,0xb4,0x42,0x5e,0x0a,0x7d,0x1f,0xb3,0xcd,
0x66,0xaf,0xb3,0xd2,0x2d,0xd7,0x9b,0x6e,0x9b,0x67,0x98,0x9e,
0x76,0xb6,0xb4,0x32,0x5c,0xa3,0xde,0xe6,0xeb,0x8f,0x18,0xbf,
0xd2,0xd2,0xef,0x12,0xba,0x97,0x9f,0x70,0xdb,0xa6,0xae,0xca,
0x2f,0xf9,0xe6,0x0a,0xa5,0xba,0xe5,0x69,0xe4,0x0a,0xd9,0xb0,
0xc2,0xbd,0x67,0x9c,0x48,0x38,0xfa,0x6c,0x0d,0x57,0x1e,0x58,
0x4e,0xf7,0x4c,0x79,0x55,0x46,0x18,0x43,0x77,0xb7,0x11,0xfe,
0xec,0x19,0x06,0x61,0x5c,0x95,0x9f,0x37,0x07,0x88,0x83,0x89,
0x29,0x79,0x5e,0x5e,0x53,0xb7,0xb2,0xa0,0x18,0x9c,0xf9,0xf3,
0x18,0x38,0x2b,0xf6,0xb7,0x67,0x01,0xb2,0x7c,0xae,0x13,0xfb,
0x30,0xf9,0xd5,0x2b,0x27,0xff,0x05,0x00,0x00,0xff,0xff,0x1e,
0x54,0xa0,0x63,0x40,0x05,0x00,0x00,
})); err != nil {
	return nil, err
}

var b bytes.Buffer
io.Copy(&b, gz)
gz.Close()
return b.Bytes(), nil}