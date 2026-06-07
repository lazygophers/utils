//go:build country_all || country_americas || country_ar || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArgentina.RegisterName(xlanguage.Chinese, "阿根廷")
	dataArgentina.RegisterOfficialName(xlanguage.Chinese, "阿根廷共和国")
	dataArgentina.RegisterCapital(xlanguage.Chinese, "布宜诺斯艾利斯")
}
