//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalawi.RegisterName(xlanguage.Japanese, "マラウイ")
	dataMalawi.RegisterOfficialName(xlanguage.Japanese, "マラウイ共和国")
	dataMalawi.RegisterCapital(xlanguage.Japanese, "リロングウェ")
}
