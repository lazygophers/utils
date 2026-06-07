//go:build country_all || country_dk || country_europe || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDenmark.RegisterName(xlanguage.Chinese, "丹麦")
	dataDenmark.RegisterOfficialName(xlanguage.Chinese, "丹麦王国")
	dataDenmark.RegisterCapital(xlanguage.Chinese, "哥本哈根")
}
