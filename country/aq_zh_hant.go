//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntarctica.RegisterName(xlanguage.MustParse("zh-Hant"), "南極洲")
	dataAntarctica.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "南極洲")
}
