//go:build country_all || country_europe || country_no || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorway.RegisterName(xlanguage.Chinese, "挪威")
	dataNorway.RegisterOfficialName(xlanguage.Chinese, "挪威王国")
	dataNorway.RegisterCapital(xlanguage.Chinese, "奥斯陆")
}
