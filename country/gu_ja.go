//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuam.RegisterName(xlanguage.Japanese, "グアム")
	dataGuam.RegisterOfficialName(xlanguage.Japanese, "グアム")
	dataGuam.RegisterCapital(xlanguage.Japanese, "ハガニア")
}
