//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurundi.RegisterName(xlanguage.Japanese, "ブルンジ")
	dataBurundi.RegisterOfficialName(xlanguage.Japanese, "ブルンジ共和国")
	dataBurundi.RegisterCapital(xlanguage.Japanese, "ギテガ")
}
