package salsa2012

// #cgo pkg-config: libsodium
// #include <stdlib.h>
// #include <sodium.h>
import "C"
import "github.com/GoKillers/libsodium-go/support"

// KeyBytes returns the required length of a secret key
func KeyBytes() int {
	return int(C.crypto_stream_salsa2012_keybytes())
}

// NonceBytes returns the required length of a nonce
func NonceBytes() int {
	return int(C.crypto_stream_salsa2012_noncebytes())
}

// Random returns `clen` pseudo random bytes using a nonce `n` and a secret key `k`.
func Random(clen int, n []byte, k []byte) []byte {
	support.CheckSize(n, NonceBytes(), "nonce")
	support.CheckSize(k, KeyBytes(), "key")

	c := make([]byte, clen)

	C.crypto_stream_salsa2012(
		(*C.uchar)(support.BytePointer(c)),
		(C.ulonglong)(clen),
		(*C.uchar)(&n[0]),
		(*C.uchar)(&k[0]))

	return c
}

// XOR encrypts a message `m` using a nonce `n` and a secret key `k`.
func XOR(m []byte, n []byte, k []byte) []byte {
	support.CheckSize(n, NonceBytes(), "nonce")
	support.CheckSize(k, KeyBytes(), "key")

	c := make([]byte, len(m))

	C.crypto_stream_salsa2012_xor(
		(*C.uchar)(support.BytePointer(c)),
		(*C.uchar)(support.BytePointer(m)),
		(C.ulonglong)(len(m)),
		(*C.uchar)(&n[0]),
		(*C.uchar)(&k[0]))

	return c
}

// KeyGen generates a secret key
func KeyGen() []byte {
	k := make([]byte, KeyBytes())
	C.crypto_stream_salsa2012_keygen((*C.uchar)(&k[0]))
	return k
}
