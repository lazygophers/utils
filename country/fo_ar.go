//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFaroeIslands.RegisterName(xlanguage.Arabic, "جزر فارو")
	dataFaroeIslands.RegisterOfficialName(xlanguage.Arabic, "جزر فارو")
	dataFaroeIslands.RegisterCapital(xlanguage.Arabic, "تورشافن")
}
