package fake

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

// UUIDv4 returns a randomly generated UUID string in the RFC 4122 v4
// canonical format (8-4-4-4-12 lowercase hex). The version nibble is set
// to 4 and the variant nibble is set to one of 8, 9, a, or b.
func (f *Faker) UUIDv4() string {
	var b [16]byte
	hi := f.uint64()
	lo := f.uint64()
	for i := 0; i < 8; i++ {
		b[i] = byte(hi >> (56 - 8*i))
		b[8+i] = byte(lo >> (56 - 8*i))
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return formatUUID(b)
}

// UUIDv7 returns a randomly generated UUID string in the RFC 9562 v7
// format. The leading 48 bits encode the current Unix time in
// milliseconds; the remaining bits are pseudo-random with the standard
// version and variant nibbles applied.
func (f *Faker) UUIDv7() string {
	var b [16]byte
	ms := uint64(time.Now().UnixMilli()) & 0x0000_ffff_ffff_ffff
	b[0] = byte(ms >> 40)
	b[1] = byte(ms >> 32)
	b[2] = byte(ms >> 24)
	b[3] = byte(ms >> 16)
	b[4] = byte(ms >> 8)
	b[5] = byte(ms)
	r1 := f.uint64()
	r2 := f.uint64()
	b[6] = byte(r1 >> 56)
	b[7] = byte(r1 >> 48)
	b[8] = byte(r1 >> 40)
	b[9] = byte(r1 >> 32)
	b[10] = byte(r1 >> 24)
	b[11] = byte(r1 >> 16)
	b[12] = byte(r2 >> 56)
	b[13] = byte(r2 >> 48)
	b[14] = byte(r2 >> 40)
	b[15] = byte(r2 >> 32)
	b[6] = (b[6] & 0x0f) | 0x70
	b[8] = (b[8] & 0x3f) | 0x80
	return formatUUID(b)
}

// formatUUID renders a 16-byte UUID as the canonical 8-4-4-4-12 hex
// string with hyphen separators.
func formatUUID(b [16]byte) string {
	var out [36]byte
	hex.Encode(out[0:8], b[0:4])
	out[8] = '-'
	hex.Encode(out[9:13], b[4:6])
	out[13] = '-'
	hex.Encode(out[14:18], b[6:8])
	out[18] = '-'
	hex.Encode(out[19:23], b[8:10])
	out[23] = '-'
	hex.Encode(out[24:36], b[10:16])
	return string(out[:])
}

// IPv4 returns a random IPv4 address in dotted-decimal notation
// (e.g. "192.168.1.1"). Each octet is independently sampled from
// [0, 255].
func (f *Faker) IPv4() string {
	v := f.uint64()
	return fmt.Sprintf("%d.%d.%d.%d", byte(v), byte(v>>8), byte(v>>16), byte(v>>24))
}

// IPv6 returns a random IPv6 address as eight colon-separated 16-bit
// hex groups without zero compression
// (e.g. "2001:0db8:85a3:0000:0000:8a2e:0370:7334").
func (f *Faker) IPv6() string {
	hi := f.uint64()
	lo := f.uint64()
	return fmt.Sprintf("%04x:%04x:%04x:%04x:%04x:%04x:%04x:%04x",
		uint16(hi>>48), uint16(hi>>32), uint16(hi>>16), uint16(hi),
		uint16(lo>>48), uint16(lo>>32), uint16(lo>>16), uint16(lo),
	)
}

// Mac returns a random MAC address in the IEEE 802 colon-separated
// lowercase hex format (e.g. "01:23:45:67:89:ab").
func (f *Faker) Mac() string {
	v := f.uint64()
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		byte(v), byte(v>>8), byte(v>>16), byte(v>>24), byte(v>>32), byte(v>>40))
}

// Md5Hex returns a random 32-character lowercase hex string with the
// same shape as an MD5 digest. The value is not a real digest of any
// input; it is only intended for fake data.
func (f *Faker) Md5Hex() string {
	var b [16]byte
	fillRandom(f, b[:])
	return hex.EncodeToString(b[:])
}

// Sha1Hex returns a random 40-character lowercase hex string with the
// same shape as a SHA-1 digest. The value is not a real digest of any
// input; it is only intended for fake data.
func (f *Faker) Sha1Hex() string {
	var b [20]byte
	fillRandom(f, b[:])
	return hex.EncodeToString(b[:])
}

// Sha256Hex returns a random 64-character lowercase hex string with the
// same shape as a SHA-256 digest. The value is not a real digest of any
// input; it is only intended for fake data.
func (f *Faker) Sha256Hex() string {
	var b [32]byte
	fillRandom(f, b[:])
	return hex.EncodeToString(b[:])
}

// fillRandom writes pseudo-random bytes into dst using the Faker's RNG,
// consuming one uint64 per 8 bytes.
func fillRandom(f *Faker, dst []byte) {
	n := len(dst)
	i := 0
	for ; i+8 <= n; i += 8 {
		binary.LittleEndian.PutUint64(dst[i:i+8], f.uint64())
	}
	if i < n {
		v := f.uint64()
		for j := 0; i < n; i, j = i+1, j+1 {
			dst[i] = byte(v >> (8 * j))
		}
	}
}
