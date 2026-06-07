//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontserrat.RegisterName(xlanguage.Japanese, "モントセラト")
	dataMontserrat.RegisterOfficialName(xlanguage.Japanese, "モントセラト")
	dataMontserrat.RegisterCapital(xlanguage.Japanese, "プリマス")
}
