//go:build country_africa || country_all || country_eg || country_northern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEgypt.RegisterName(xlanguage.Chinese, "埃及")
	dataEgypt.RegisterOfficialName(xlanguage.Chinese, "阿拉伯埃及共和国")
	dataEgypt.RegisterCapital(xlanguage.Chinese, "开罗")
}
