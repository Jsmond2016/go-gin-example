// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__f64toa = 48
)

const (
    _stack__f64toa = 72
)

const (
    _size__f64toa = 5072
)

var (
    _pcsp__f64toa = [][2]uint32{
        {0x1, 0},
        {0x6, 8},
        {0x8, 16},
        {0xa, 24},
        {0xc, 32},
        {0xd, 40},
        {0x11, 48},
        {0x136d, 72},
        {0x136e, 48},
        {0x1370, 40},
        {0x1372, 32},
        {0x1374, 24},
        {0x1376, 16},
        {0x1377, 8},
        {0x137b, 0},
        {0x13d0, 72},
    }
)

var _cfunc_f64toa = []loader.CFunc{
    {"_f64toa_entry", 0,  _entry__f64toa, 0, nil},
    {"_f64toa", _entry__f64toa, _size__f64toa, _stack__f64toa, _pcsp__f64toa},
}
