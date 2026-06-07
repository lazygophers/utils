//go:build country_africa || country_all || country_gw || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuineaBissau.RegisterName(xlanguage.Chinese, "几内亚比绍")
	dataGuineaBissau.RegisterOfficialName(xlanguage.Chinese, "几内亚比绍共和国")
	dataGuineaBissau.RegisterCapital(xlanguage.Chinese, "比绍")
}
