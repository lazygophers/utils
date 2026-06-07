//go:build country_africa || country_all || country_ng || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNigeria.RegisterName(xlanguage.Chinese, "尼日利亚")
	dataNigeria.RegisterOfficialName(xlanguage.Chinese, "尼日利亚联邦共和国")
	dataNigeria.RegisterCapital(xlanguage.Chinese, "阿布贾")
}
