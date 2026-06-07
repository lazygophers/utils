//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUganda.RegisterName(xlanguage.Japanese, "ウガンダ")
	dataUganda.RegisterOfficialName(xlanguage.Japanese, "ウガンダ共和国")
	dataUganda.RegisterCapital(xlanguage.Japanese, "カンパラ")
}
