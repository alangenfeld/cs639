// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Bob Jenkins Hash 
// http://burtleburtle.net/bob/c/lookup3.c 
package bob

import (
	"hash"
	"os"
)

type H struct {
	a, b, c uint32
}

func New() hash.Hash32 {
	x := new(H)
	x.Reset()
	return x
}

func (x *H) Reset() {
	x.a = 0xCafeBaba
	x.b = x.a
	x.c = x.a
}

func (x *H) Sum() []byte {
	p := make([]byte, 4)
	s := x.Sum32()
	p[0] = byte(s >> 24)
	p[1] = byte(s >> 16)
	p[2] = byte(s >> 8)
	p[3] = byte(s)
	return p
}

func (x *H) Sum32() uint32 {
	return x.c
}

func (x *H) Size() int {
	return 4
}

func (x *H) Write(b []byte) (n int, err os.Error) {
	i := 0
	n = len(b)
	if n == 0 {
		return
	}
	n = (n-1)>>2 + 1 // number of uint32

	// 3 unit32 at a time 
	for ; n > 3; n -= 3 {
		x.a += pack(b, i)
		i += 4
		x.b += pack(b, i)
		i += 4
		x.c += pack(b, i)
		i += 4
		x.mix()
	}

	// remaining uint32 
	switch n {
	case 3:
		x.c += pack(b, i+8)
		fallthrough
	case 2:
		x.b += pack(b, i+4)
		fallthrough
	case 1:
		x.a += pack(b, i)
		x.final()
	}
	return len(b), nil
}

// internal 
func pack(b []byte, i int) (v uint32) { // 0-4 bytes to a little endian 
	n := len(b) - i
	if n <= 0 {
		return
	}
	if n > 4 {
		n = 4
	}
	switch n {
	case 4:
		v = uint32(b[3+i])
		fallthrough
	case 3:
		v = v<<8 | uint32(b[2+i])
		fallthrough
	case 2:
		v = v<<8 | uint32(b[1+i])
		fallthrough
	case 1:
		v = v<<8 | uint32(b[0+i])
	}
	return
}

func rot(x, k uint32) uint32 {
	return x<<k | x>>(32-k)
}

func (x *H) mix() {
	x.a -= x.c
	x.a ^= rot(x.c, 4)
	x.c += x.b
	x.b -= x.a
	x.b ^= rot(x.a, 6)
	x.a += x.c
	x.c -= x.b
	x.c ^= rot(x.b, 8)
	x.b += x.a
	x.a -= x.c
	x.a ^= rot(x.c, 16)
	x.c += x.b
	x.b -= x.a
	x.b ^= rot(x.a, 19)
	x.a += x.c
	x.c -= x.b
	x.c ^= rot(x.b, 4)
	x.b += x.a
}

func (x *H) final() {
	x.c ^= x.b
	x.c -= rot(x.b, 14)
	x.a ^= x.c
	x.a -= rot(x.c, 11)
	x.b ^= x.a
	x.b -= rot(x.a, 25)
	x.c ^= x.b
	x.c -= rot(x.b, 16)
	x.a ^= x.c
	x.a -= rot(x.c, 4)
	x.b ^= x.a
	x.b -= rot(x.a, 14)
	x.c ^= x.b
	x.c -= rot(x.b, 24)
}