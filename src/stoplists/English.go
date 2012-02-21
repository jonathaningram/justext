package gojustext
import ( "io"; "os"; "bytes"; "compress/gzip" )

func EnglishStoplist() ([]byte, os.Error) {
var gz *gzip.Decompressor
var err os.Error
if gz, err = gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x00,0x00,0x00,0x00,0xff,0x4c,0x97,
0xcd,0x72,0xdc,0x38,0x12,0x84,0xef,0xf5,0x18,0x7b,0x9a,0x83,
0x43,0xef,0xd0,0xe3,0x3f,0x49,0x3b,0x6b,0x6f,0x4c,0xdb,0xa1,
0xd8,0x23,0x48,0x82,0xdd,0xb0,0x48,0x80,0x81,0x1f,0xd1,0x7c,
0xfb,0xfd,0x12,0xdd,0x76,0xcc,0xc5,0x62,0x93,0x00,0xaa,0x2a,
0x2b,0x33,0x0b,0xae,0x57,0x6f,0x69,0x36,0x17,0x27,0x0b,0xd1,
0x6a,0x32,0x67,0xbb,0x2b,0x16,0x8a,0x7d,0xe3,0xcb,0x9c,0xb2,
0xf1,0x2b,0x45,0xdb,0x43,0xbd,0xda,0x70,0x58,0xbd,0xba,0x6a,
0x73,0x4e,0xab,0xf1,0xf7,0xca,0x32,0x17,0x8d,0x85,0x4f,0xd1,
0x5c,0xf6,0xb6,0x7b,0xfd,0x73,0x0d,0x23,0x6b,0xbd,0x5d,0xd9,
0xfa,0xe8,0x2d,0x54,0xd3,0x31,0x4b,0x49,0xbc,0x99,0x6c,0x0e,
0xb9,0x54,0x7b,0xaa,0x1c,0xe5,0x43,0xb6,0x98,0xaa,0x0d,0x8d,
0xb3,0xdc,0x9b,0xb6,0x26,0x96,0x2b,0xa0,0xe7,0x33,0xa7,0x0f,
0xde,0xeb,0xfc,0x6c,0x75,0x4f,0xda,0x70,0x58,0xe2,0xdf,0x4c,
0xae,0x4a,0x75,0xae,0x3c,0xba,0x65,0x61,0x1f,0xcb,0xd6,0x44,
0xf0,0x6f,0xda,0x95,0xe2,0x72,0xd8,0x9e,0xda,0x32,0xd9,0xc9,
0x0a,0xe9,0x7d,0xf1,0x3b,0x9f,0x09,0x3b,0x92,0x6e,0x7a,0x63,
0xd7,0xd4,0x72,0x88,0x17,0x6d,0x64,0x53,0xe4,0x73,0x2b,0x7e,
0xb2,0xd2,0xc8,0xbc,0x6d,0x84,0xad,0xbb,0x22,0xaf,0x2e,0x1e,
0xfc,0x33,0x79,0x2b,0x69,0x55,0x4a,0xda,0x4e,0xb2,0xdf,0x63,
0xa8,0x2c,0x7f,0x8d,0x69,0xa7,0xf0,0x41,0xaf,0x6a,0xe8,0x0b,
0xd8,0x34,0xf8,0xd1,0xf1,0xdc,0xe2,0x44,0x9c,0x7f,0x09,0xc6,
0xc1,0x2b,0xd6,0xe6,0x72,0x2f,0x9a,0x80,0xd7,0xb0,0xda,0xe1,
0x5d,0x2e,0xfc,0xce,0x5e,0xfb,0x72,0x6a,0x97,0xab,0x7d,0x8d,
0x54,0x36,0x2e,0x6d,0xd2,0xfa,0xc5,0xa9,0xbc,0x3d,0x50,0xdf,
0x69,0xf5,0x39,0x28,0xf7,0x81,0xf2,0xed,0xd4,0xeb,0x6e,0xb1,
0x86,0x85,0xa3,0xe9,0x91,0xb7,0x33,0x51,0x76,0xcf,0xca,0x08,
0x2c,0x17,0x17,0x22,0xb5,0xd2,0x87,0xc5,0x53,0xf1,0xb2,0xa8,
0x32,0x3f,0x26,0x7a,0x7c,0x2a,0x3c,0x51,0xbf,0x5b,0x54,0x03,
0x0f,0x25,0xd4,0xc3,0x62,0x5b,0x07,0x4e,0x8c,0x4a,0x9b,0x04,
0x8b,0xb7,0x6d,0x71,0x07,0xbb,0x48,0x11,0x24,0x57,0x77,0x88,
0x0d,0x87,0xbd,0xa4,0x0c,0xa4,0x8f,0x20,0xbc,0xa4,0x91,0xec,
0x26,0xfb,0xe2,0x6a,0x48,0x91,0xd3,0xca,0x7d,0xef,0x6a,0xd9,
0x2f,0xde,0x09,0x4c,0xb6,0x90,0x19,0xd8,0xc1,0x0c,0x27,0x7c,
0x69,0x49,0x7e,0x85,0x30,0x87,0x95,0xf1,0x9a,0x12,0x9b,0x28,
0x96,0x5d,0xde,0xc1,0xa6,0x8b,0x30,0xb6,0xe7,0x74,0x85,0x6b,
0xf0,0x6d,0xf3,0x69,0x23,0xfb,0xc1,0x5f,0xa8,0xda,0x3b,0xda,
0x22,0xbc,0x6c,0x4b,0x5b,0x03,0x17,0x62,0x52,0xe9,0x6e,0xb3,
0x5b,0x03,0x09,0xce,0x61,0x59,0xe1,0x2a,0x88,0xdb,0xa8,0x7a,
0xfe,0xcc,0xa1,0x86,0x72,0xd5,0x2b,0xf1,0x63,0x68,0xab,0x8d,
0x9d,0x0d,0x14,0x7c,0xd8,0x63,0xda,0x05,0xc1,0x3b,0x3b,0xd3,
0xb7,0x6b,0xaf,0x1a,0x60,0x2a,0x69,0xf6,0x13,0x6a,0x4a,0xaf,
0x22,0x3e,0x90,0xf7,0x56,0x2a,0xf1,0x02,0xfc,0xbe,0xd8,0x27,
0x78,0x7c,0xae,0x14,0x5e,0x6c,0x0a,0x53,0xd7,0x04,0xa2,0x29,
0x7a,0x63,0x9e,0xad,0x43,0xaf,0xfb,0x3f,0xc0,0xf5,0xd4,0x11,
0x5a,0xa0,0x2e,0x51,0xcf,0x7e,0xab,0xbe,0x23,0x5c,0x68,0xd9,
0x82,0xde,0x2a,0x34,0xa9,0xd7,0xc4,0xc9,0xeb,0xfd,0xc3,0x2a,
0x1a,0xd7,0x4e,0xa7,0x25,0xd1,0xfc,0xc1,0x8d,0xaf,0x76,0xbe,
0xc1,0xb4,0xb8,0x0c,0x3a,0xcf,0x2e,0x36,0x47,0xfa,0xcf,0x0d,
0x71,0x5c,0xc8,0x75,0x53,0x5a,0x6f,0x3d,0x60,0x06,0x9f,0x6b,
0x80,0x40,0x3a,0xe0,0xc3,0x8d,0xda,0x5f,0x52,0xa6,0xba,0xe7,
0x06,0x3e,0x5f,0xc7,0x9a,0x14,0x25,0xcc,0xb6,0x84,0x57,0xc9,
0x7a,0x59,0xd2,0xae,0x45,0x43,0x0b,0x4b,0xb5,0x53,0xbb,0x34,
0x08,0x73,0xda,0x32,0x8c,0x5a,0x5b,0x09,0x23,0x44,0xcb,0xd1,
0xde,0xc8,0x56,0x8d,0xb9,0xa8,0xb9,0x53,0xf3,0x64,0xc2,0x32,
0xb8,0x41,0xd7,0xae,0x12,0x44,0x07,0xc3,0x16,0x3f,0x57,0x48,
0xf2,0x03,0x74,0x8a,0xaf,0x77,0x06,0x7b,0xfb,0xfe,0x70,0x7e,
0xe0,0x34,0x52,0xfb,0xe0,0xc7,0x5b,0x99,0x5f,0xd0,0x5e,0x7f,
0xc8,0xbc,0x81,0x7e,0x93,0xbd,0x48,0x31,0xff,0x13,0x2f,0x56,
0x58,0x6b,0x2f,0xea,0x71,0x1b,0x16,0x65,0x20,0x4f,0xfa,0xa3,
0xe7,0x51,0x20,0x94,0x38,0xa1,0x0f,0xc8,0x19,0x5e,0xbe,0xb1,
0x6b,0x0a,0xf3,0x0c,0xc3,0xa2,0xec,0x25,0x5c,0x82,0x58,0x78,
0x43,0x13,0x8e,0x93,0x97,0x38,0x02,0xc1,0x2e,0xb0,0xe8,0x22,
0xc9,0xc7,0x55,0x4b,0x1d,0x7e,0x23,0x03,0x89,0x62,0xd3,0x2e,
0x33,0x9b,0x02,0x07,0xae,0xac,0xa0,0x7f,0x88,0x84,0xaa,0x7e,
0x08,0x8c,0x4f,0x7e,0xc8,0x1d,0x6f,0x5e,0x6d,0xe2,0xec,0xbd,
0x2c,0xa4,0xa4,0xfe,0x74,0x89,0x4b,0x58,0x52,0x69,0xa7,0x0b,
0x44,0xba,0x6b,0x61,0x75,0x60,0xbc,0x2b,0x1a,0x12,0xa0,0xc7,
0x69,0x9e,0x71,0x21,0xe8,0xfa,0x89,0x6c,0x81,0x63,0xa6,0x72,
0x2b,0x47,0x81,0x14,0x36,0x63,0x3d,0x6f,0x2e,0x87,0xd4,0x8a,
0x5d,0x78,0x2f,0xfb,0x20,0xfa,0x47,0xf2,0x16,0x83,0xdf,0x8b,
0xce,0x9d,0x10,0x50,0x2d,0x4f,0xe6,0x56,0x3d,0x8b,0xbe,0xa8,
0x51,0x0b,0x3f,0xfb,0x8c,0x49,0xa1,0x2e,0x75,0x73,0x4a,0xe8,
0x10,0x62,0x52,0x7e,0x96,0x3a,0xc7,0x96,0x85,0x0f,0x0c,0xc0,
0x5d,0x6f,0x0b,0xe0,0x47,0xf4,0x3f,0x05,0x99,0xdc,0xe9,0x2b,
0x54,0x9a,0x3b,0x72,0x55,0x39,0x97,0xcd,0x8f,0xe2,0x39,0x07,
0xbb,0x5f,0x50,0x8f,0x88,0x57,0x67,0x2d,0x61,0xbe,0xd9,0x02,
0x4f,0xec,0x12,0x71,0xba,0x4b,0xe1,0x06,0x35,0xe5,0xe3,0x01,
0x69,0x2e,0xa8,0x4e,0xa4,0x97,0x2e,0x64,0x89,0xa5,0x6d,0x1b,
0x24,0xb4,0x1d,0x39,0x8a,0xf1,0x13,0x4b,0xb1,0x31,0x18,0x44,
0xe8,0x2d,0xa7,0xa9,0x8d,0x1c,0xfc,0x28,0xde,0xde,0x74,0x9d,
0xed,0x2f,0xef,0x2e,0xf0,0x4c,0xc4,0x0e,0x90,0x6c,0x4e,0xa9,
0x0e,0x52,0x06,0xe6,0x55,0x02,0x19,0xb3,0xfe,0x5c,0x1f,0x18,
0x51,0x07,0x14,0xaa,0x2d,0x47,0x5e,0x3c,0x45,0xec,0xf0,0x37,
0xf8,0x83,0xe4,0xdb,0x2d,0x10,0x25,0x29,0x0f,0x4e,0x0f,0xab,
0xf2,0x70,0x34,0x64,0x09,0x1d,0x87,0x9c,0x60,0xc5,0xdf,0xf2,
0x3f,0xb0,0xa2,0xa9,0x91,0x99,0xb3,0xba,0x8c,0xc8,0x7b,0x7b,
0x11,0xe2,0xd6,0xcd,0x4a,0xe1,0x40,0x97,0xe8,0x35,0xc4,0xc6,
0x2f,0xba,0x78,0xf3,0x8e,0xee,0x36,0xbc,0xf8,0x91,0x40,0x83,
0xbe,0x6c,0x1b,0x84,0xe2,0xa1,0xa6,0x4a,0x12,0x1b,0xcb,0xb2,
0xfd,0xa9,0x24,0x47,0x90,0xb6,0xf7,0x08,0xce,0xcb,0xe1,0xa8,
0xc1,0x5e,0x50,0x55,0xc0,0xf2,0x6a,0xa8,0x64,0x31,0x01,0x2f,
0x7d,0xe6,0xb8,0x7a,0xbc,0xbb,0x0d,0x85,0x07,0xd8,0xc5,0x54,
0x10,0x49,0x95,0xe1,0x27,0x4d,0xdb,0xd8,0xd5,0x0c,0x03,0x87,
0xfb,0x5a,0xc3,0xf9,0x02,0x7d,0x3e,0x64,0x6a,0x45,0x64,0xff,
0xd8,0x72,0x22,0x8b,0x68,0xa7,0x51,0x41,0x55,0x67,0xc2,0x7d,
0x5e,0xfa,0x18,0x28,0xe1,0xa7,0x4d,0xa0,0xc6,0x94,0xcd,0x92,
0x74,0xc5,0xa5,0xdf,0x42,0xdf,0x57,0x02,0x47,0x49,0x0d,0xdd,
0xdd,0x2f,0x3e,0xf6,0xf1,0xc0,0xbb,0x45,0x5f,0x5f,0x34,0xf8,
0xcf,0x92,0xc5,0xde,0xe7,0x51,0x77,0x46,0x5d,0x16,0xfe,0x09,
0x7a,0x2b,0x8d,0x2e,0xa1,0x95,0x1b,0xe1,0xe4,0x6d,0xc2,0xbd,
0x74,0x83,0xfd,0x7c,0x3f,0x51,0x13,0xf2,0x9d,0xe4,0xb4,0x36,
0xc6,0xe7,0x61,0x1f,0x65,0x26,0x8f,0x10,0x5f,0xbe,0xa2,0xd9,
0xf4,0xab,0x5d,0x9f,0x7d,0x92,0xe1,0x21,0xa3,0x3e,0x36,0x35,
0x8a,0xe0,0x0b,0x7f,0xf4,0xf3,0xbf,0xd9,0x0b,0x45,0x82,0x4c,
0xa0,0xb3,0xa4,0xad,0xcb,0xfa,0x19,0x39,0x62,0xca,0xb7,0x37,
0x62,0x2d,0x51,0x54,0x1a,0x07,0xaa,0x26,0x71,0xd8,0x46,0x16,
0x36,0xf0,0x9a,0x12,0x4b,0xe7,0x06,0x9e,0x32,0xa2,0xa6,0x79,
0x93,0x7b,0xa4,0xee,0xb3,0xe4,0x74,0x6b,0x9a,0x86,0x02,0xcc,
0x5b,0x74,0x6d,0xc0,0x13,0xab,0x81,0xd8,0x28,0x25,0x70,0xc3,
0x10,0xe3,0x21,0x29,0x9c,0x86,0x50,0x08,0x7c,0xa3,0x44,0x98,
0x73,0x15,0xdb,0xa5,0x24,0x6e,0x09,0x9a,0x6c,0x70,0x7c,0xf6,
0x45,0x20,0x8b,0x13,0x9d,0xf0,0xdd,0x96,0x36,0x48,0xc3,0x8d,
0xc7,0x73,0x1b,0x1a,0xe9,0xce,0x04,0x66,0xbf,0xae,0x00,0xd2,
0x8d,0xec,0xb6,0x48,0x24,0x21,0x29,0xda,0x5a,0xfc,0x32,0xff,
0xf6,0x38,0x50,0x06,0x79,0x0d,0xb4,0x58,0xa1,0x31,0x0d,0x75,
0x74,0x6f,0xa0,0xc5,0x0b,0xa1,0xec,0xdf,0x2a,0xe4,0x49,0x74,
0x4e,0xc8,0x45,0xa1,0x98,0xf0,0x8a,0x29,0xcb,0xac,0xb7,0x2e,
0xb9,0x31,0x27,0x96,0xfe,0x9d,0x0e,0x89,0x95,0xe2,0x31,0x4c,
0xee,0x33,0x76,0x40,0x2c,0xae,0x44,0x39,0xf1,0x8c,0xdb,0xa9,
0xe7,0xe4,0x43,0x7b,0x3e,0xe0,0x1e,0x93,0x7d,0xed,0x80,0x9d,
0xfb,0xcc,0x96,0x9b,0xf5,0xa0,0x73,0xd3,0xf5,0x0b,0xf6,0x9c,
0x80,0x88,0x2e,0x07,0xf1,0x8f,0x6b,0xdd,0x5f,0xdc,0x36,0x88,
0x7a,0xca,0xeb,0x6f,0xe7,0xe4,0xd6,0xe0,0x76,0x7a,0x54,0xc2,
0xa5,0xa3,0x25,0x72,0xce,0x34,0xa6,0x49,0x37,0x08,0x48,0xa0,
0xc8,0x84,0xe5,0x65,0x2a,0xe5,0x14,0x21,0x23,0xb0,0x43,0x40,
0xe2,0x3e,0x3b,0x5c,0xb8,0x5f,0x4d,0x72,0xba,0x64,0xe4,0xf3,
0xfd,0x6c,0xef,0x5d,0x74,0x93,0x22,0xce,0x68,0x88,0x35,0x5c,
0x3c,0xfa,0x49,0x39,0xfd,0x0c,0x2b,0x6c,0xa5,0x56,0x50,0xbc,
0x33,0x56,0x1f,0xbb,0xf6,0xf9,0x93,0xb8,0xe4,0x91,0x5d,0xf6,
0x1a,0x3a,0x7c,0xe3,0x1e,0x21,0x03,0x7e,0x8a,0xfd,0xb4,0x1d,
0x6f,0xd0,0xbc,0x64,0x08,0xf4,0xbb,0x2a,0x43,0x1c,0x3b,0x10,
0x33,0x2e,0xba,0xa4,0xba,0x18,0x01,0x4a,0xd6,0x05,0xf5,0x38,
0xa3,0x34,0x46,0xa8,0xc8,0x0d,0xab,0x82,0xe6,0x3d,0x38,0x1a,
0xdd,0xba,0x92,0xbb,0x28,0xee,0x96,0xdb,0xd5,0x53,0xb2,0x15,
0x6d,0x3a,0x79,0xde,0x5f,0x9b,0x86,0xf6,0x6d,0x12,0x73,0xd4,
0x25,0xd1,0x6b,0x04,0xaa,0x56,0x74,0x78,0xee,0xc8,0x70,0x43,
0x4c,0xa0,0xa0,0xea,0x40,0x6e,0xcc,0x61,0xe0,0x1b,0xaf,0xf1,
0x09,0xb6,0x97,0x3e,0xfd,0x60,0xdf,0xcd,0xba,0x69,0xe7,0x08,
0x5f,0xe5,0x9f,0xdc,0x84,0x13,0xe4,0x72,0xd9,0xc1,0x59,0x5c,
0x2e,0xf5,0xaf,0x88,0x9a,0x13,0xf3,0x1d,0x2d,0x1c,0xfd,0xa1,
0xbb,0xc9,0xc8,0xd8,0xd7,0x34,0x11,0x61,0xfe,0xd0,0xff,0x08,
0x74,0x55,0xe5,0xaa,0x3d,0xcf,0xa1,0x6f,0xeb,0xae,0x84,0x74,
0x6f,0x39,0xff,0x3f,0x00,0x00,0xff,0xff,0x12,0xcb,0x57,0x30,
0x33,0x0c,0x00,0x00,
})); err != nil {
	return nil, err
}

var b bytes.Buffer
io.Copy(&b, gz)
gz.Close()
return b.Bytes(), nil}