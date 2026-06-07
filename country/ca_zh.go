//go:build country_all || country_americas || country_ca || country_northern_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCanada.RegisterName(xlanguage.Chinese, "加拿大")
	dataCanada.RegisterOfficialName(xlanguage.Chinese, "加拿大")
	dataCanada.RegisterCapital(xlanguage.Chinese, "渥太华")
}
