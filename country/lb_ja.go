//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLebanon.RegisterName(xlanguage.Japanese, "レバノン")
	dataLebanon.RegisterOfficialName(xlanguage.Japanese, "レバノン共和国")
	dataLebanon.RegisterCapital(xlanguage.Japanese, "ベイルート")
}
