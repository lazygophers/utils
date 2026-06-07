//go:build country_all || country_europe || country_northern_europe || country_se

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSweden.RegisterName(xlanguage.Chinese, "瑞典")
	dataSweden.RegisterOfficialName(xlanguage.Chinese, "瑞典王国")
	dataSweden.RegisterCapital(xlanguage.Chinese, "斯德哥尔摩")
}
