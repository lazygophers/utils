//go:build country_all || country_americas || country_co || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataColombia.RegisterName(xlanguage.Chinese, "哥伦比亚")
	dataColombia.RegisterOfficialName(xlanguage.Chinese, "哥伦比亚共和国")
	dataColombia.RegisterCapital(xlanguage.Chinese, "波哥大")
}
