//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChile.RegisterName(xlanguage.Japanese, "チリ")
	dataChile.RegisterOfficialName(xlanguage.Japanese, "チリ共和国")
	dataChile.RegisterCapital(xlanguage.Japanese, "サンティアゴ")
}
