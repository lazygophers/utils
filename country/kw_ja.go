//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKuwait.RegisterName(xlanguage.Japanese, "クウェート")
	dataKuwait.RegisterOfficialName(xlanguage.Japanese, "クウェート国")
	dataKuwait.RegisterCapital(xlanguage.Japanese, "クウェート市")
}
