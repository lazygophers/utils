//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDjibouti.RegisterName(xlanguage.Japanese, "ジブチ")
	dataDjibouti.RegisterOfficialName(xlanguage.Japanese, "ジブチ共和国")
	dataDjibouti.RegisterCapital(xlanguage.Japanese, "ジブチ市")
}
