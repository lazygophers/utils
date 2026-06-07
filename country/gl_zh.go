//go:build country_all || country_americas || country_gl || country_northern_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreenland.RegisterName(xlanguage.Chinese, "格陵兰")
	dataGreenland.RegisterOfficialName(xlanguage.Chinese, "格陵兰")
	dataGreenland.RegisterCapital(xlanguage.Chinese, "努克")
}
