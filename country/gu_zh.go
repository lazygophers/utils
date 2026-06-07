//go:build country_all || country_gu || country_micronesia || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuam.RegisterName(xlanguage.Chinese, "关岛")
	dataGuam.RegisterOfficialName(xlanguage.Chinese, "关岛领地")
	dataGuam.RegisterCapital(xlanguage.Chinese, "阿加尼亚")
}
