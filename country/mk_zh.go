//go:build country_all || country_europe || country_mk || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthMacedonia.RegisterName(xlanguage.Chinese, "北马其顿")
	dataNorthMacedonia.RegisterOfficialName(xlanguage.Chinese, "北马其顿共和国")
	dataNorthMacedonia.RegisterCapital(xlanguage.Chinese, "斯科普里")
}
