package fake

import "strconv"

// fileExts is the static pool of file extensions sampled by [Faker.FileExt]
// and used to suffix [Faker.FileName]. Indices align with [mimeTypes] so the
// two pools can be sampled jointly when needed.
var fileExts = []string{
	"pdf", "doc", "docx", "xls", "xlsx", "ppt", "pptx", "txt", "csv", "json",
	"xml", "html", "md", "jpg", "png", "gif", "svg", "mp4", "mp3", "wav",
	"zip", "tar", "gz",
}

// mimeTypes is the static pool of MIME identifiers returned by
// [Faker.MimeType]. The slice is index-aligned with [fileExts].
var mimeTypes = []string{
	"application/pdf",
	"application/msword",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"application/vnd.ms-excel",
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"application/vnd.ms-powerpoint",
	"application/vnd.openxmlformats-officedocument.presentationml.presentation",
	"text/plain",
	"text/csv",
	"application/json",
	"application/xml",
	"text/html",
	"text/markdown",
	"image/jpeg",
	"image/png",
	"image/gif",
	"image/svg+xml",
	"video/mp4",
	"audio/mpeg",
	"audio/wav",
	"application/zip",
	"application/x-tar",
	"application/gzip",
}

// fileNameBases is the small built-in word pool used to compose
// [Faker.FileName]. It is kept local to avoid coupling misc helpers to the
// lorem word tables in text.go.
var fileNameBases = []string{
	"report", "document", "photo", "image", "data", "note", "draft", "final",
	"summary", "backup", "invoice", "receipt", "screenshot", "proposal",
	"contract", "chart", "log", "archive",
}

const hexDigits = "0123456789abcdef"

// HexColor returns a random CSS-style color in #RRGGBB form using lowercase
// hexadecimal digits.
func (f *Faker) HexColor() string {
	buf := make([]byte, 7)
	buf[0] = '#'
	v := f.uint64()
	for i := 0; i < 6; i++ {
		buf[1+i] = hexDigits[v&0xF]
		v >>= 4
	}
	return string(buf)
}

// RgbColor returns a random color formatted as rgb(R, G, B) where each
// channel is an integer in [0, 255].
func (f *Faker) RgbColor() string {
	r := f.intN(256)
	g := f.intN(256)
	b := f.intN(256)
	return "rgb(" + strconv.Itoa(r) + ", " + strconv.Itoa(g) + ", " + strconv.Itoa(b) + ")"
}

// HslColor returns a random color formatted as hsl(H, S%, L%) with hue in
// [0, 360) and saturation/lightness in [0, 100].
func (f *Faker) HslColor() string {
	h := f.intN(360)
	s := f.intN(101)
	l := f.intN(101)
	return "hsl(" + strconv.Itoa(h) + ", " + strconv.Itoa(s) + "%, " + strconv.Itoa(l) + "%)"
}

// FileName returns a plausible file name composed of a base word, an
// optional `_<n>` numeric suffix (n in [1, 99], added roughly half the
// time), and a random extension drawn from [fileExts].
func (f *Faker) FileName() string {
	base := f.pickString(fileNameBases)
	ext := f.pickString(fileExts)
	if f.intN(2) == 0 {
		suffix := f.intN(99) + 1
		return base + "_" + strconv.Itoa(suffix) + "." + ext
	}
	return base + "." + ext
}

// FileExt returns a random file extension (without leading dot) drawn from
// [fileExts].
func (f *Faker) FileExt() string {
	return f.pickString(fileExts)
}

// MimeType returns a random MIME type identifier drawn from [mimeTypes].
func (f *Faker) MimeType() string {
	return f.pickString(mimeTypes)
}
