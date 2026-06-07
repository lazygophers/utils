//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSanMarino.RegisterName(xlanguage.Japanese, "サンマリノ")
	dataSanMarino.RegisterOfficialName(xlanguage.Japanese, "サンマリノ共和国")
	dataSanMarino.RegisterCapital(xlanguage.Japanese, "サンマリノ")
}
