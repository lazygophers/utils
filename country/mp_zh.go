//go:build country_all || country_micronesia || country_mp || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthernMarianaIslands.RegisterName(xlanguage.Chinese, "北马里亚纳群岛")
	dataNorthernMarianaIslands.RegisterOfficialName(xlanguage.Chinese, "北马里亚纳群岛自由联邦")
	dataNorthernMarianaIslands.RegisterCapital(xlanguage.Chinese, "塞班岛")
}
