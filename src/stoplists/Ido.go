package gojustext
import ( "io"; "os"; "bytes"; "compress/gzip" )

func IdoStoplist() ([]byte, os.Error) {
var gz *gzip.Decompressor
var err os.Error
if gz, err = gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x00,0x00,0x00,0x00,0xff,0x1c,0x8d,
0xc1,0x71,0xc3,0x30,0x0c,0x04,0xff,0xd7,0x49,0x1e,0x51,0x17,
0xf9,0xe5,0x91,0x99,0x54,0x70,0x1a,0x22,0x36,0x22,0x12,0x90,
0x49,0x42,0x33,0x76,0x59,0x2e,0xc1,0x95,0x05,0xca,0xef,0xc8,
0xbb,0x5d,0x54,0x42,0x06,0x07,0x8a,0x42,0x50,0x04,0x9f,0xf9,
0x36,0x38,0x9a,0x3c,0x98,0xdd,0x2d,0x88,0x2b,0x8f,0x1c,0x44,
0x5f,0x1d,0x23,0xd6,0xfc,0x52,0x74,0x39,0xc4,0xc2,0x21,0x87,
0x63,0x6a,0xc5,0xee,0x1d,0x3f,0x6c,0x5a,0x55,0x33,0xef,0x51,
0x3d,0xa9,0x55,0x67,0x72,0xdf,0x72,0x09,0xcb,0x61,0xc6,0x3b,
0xbb,0x62,0xfa,0x24,0xf6,0x1a,0xd8,0xda,0xeb,0x89,0x8f,0xf3,
0xf6,0xd5,0x5b,0x54,0x5d,0x70,0xe9,0xb4,0x22,0x23,0x59,0x69,
0xb4,0x73,0xfc,0x95,0xd9,0x9b,0xda,0x69,0x3d,0x7a,0x56,0xef,
0x55,0x4d,0x7c,0x39,0x05,0x03,0xdb,0xbf,0xd8,0x1f,0x5c,0xc0,
0x2e,0x0e,0xde,0x22,0x2b,0x19,0x3a,0xd0,0xf4,0xf5,0x7c,0xc3,
0x5d,0x50,0xfd,0x97,0x36,0x15,0x25,0xc0,0x3a,0x3b,0xb1,0xb9,
0x4d,0x16,0xc7,0x5f,0x00,0x00,0x00,0xff,0xff,0xa7,0xc4,0x4a,
0x2e,0xf9,0x00,0x00,0x00,
})); err != nil {
	return nil, err
}

var b bytes.Buffer
io.Copy(&b, gz)
gz.Close()
return b.Bytes(), nil}