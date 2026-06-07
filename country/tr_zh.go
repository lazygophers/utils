//go:build country_all || country_asia || country_tr || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkey.RegisterName(xlanguage.Chinese, "土耳其")
	dataTurkey.RegisterOfficialName(xlanguage.Chinese, "土耳其共和国")
	dataTurkey.RegisterCapital(xlanguage.Chinese, "安卡拉")
}
