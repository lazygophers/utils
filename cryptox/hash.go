package cryptox

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash/fnv"
	"io"
	"os"
)

func Md5[M string | []byte](s M) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func Sha256[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

func Sha224[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha256.Sum224([]byte(s)))
}

func Sha512[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(s)))
}

func Sha384[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha512.Sum384([]byte(s)))
}

func Sha512_256[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha512.Sum512_256([]byte(s)))
}

func Sha512_224[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha512.Sum512_224([]byte(s)))
}

func FileMd5(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return Md5WithReader(file)
}

func FileSha256(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return Sha256WithReader(file)
}

func FileSha224(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return Sha224WithReader(file)
}

func FileSha512(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return Sha512WithReader(file)
}

func FileSha384(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return Sha384WithReader(file)
}

func FileSha512_256(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return Sha512_256WithReader(file)
}

func FileSha512_224(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return Sha512_224WithReader(file)
}

func Md5WithReader(file io.Reader) (string, error) {
	hash := md5.New()
	_, err := io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func Sha256WithReader(file io.Reader) (string, error) {
	hash := sha256.New()
	_, err := io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func Sha224WithReader(file io.Reader) (string, error) {
	hash := sha256.New224()
	_, err := io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func Sha512WithReader(file io.Reader) (string, error) {
	hash := sha512.New()
	_, err := io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func Sha384WithReader(file io.Reader) (string, error) {
	hash := sha512.New384()
	_, err := io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func Sha512_256WithReader(file io.Reader) (string, error) {
	hash := sha512.New512_256()
	_, err := io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func Sha512_224WithReader(file io.Reader) (string, error) {
	hash := sha512.New512_224()
	_, err := io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func Hash32[M string | []byte](s M) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}

func Hash32a[M string | []byte](s M) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func Hash64[M string | []byte](s M) uint64 {
	h := fnv.New64()
	h.Write([]byte(s))
	return h.Sum64()
}

func Hash64a[M string | []byte](s M) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}
