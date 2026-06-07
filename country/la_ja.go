//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLaos.RegisterName(xlanguage.Japanese, "ラオス")
	dataLaos.RegisterOfficialName(xlanguage.Japanese, "ラオス人民民主共和国")
	dataLaos.RegisterCapital(xlanguage.Japanese, "ヴィエンチャン")
}
