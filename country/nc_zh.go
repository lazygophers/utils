//go:build country_all || country_melanesia || country_nc || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewCaledonia.RegisterName(xlanguage.Chinese, "新喀里多尼亚")
	dataNewCaledonia.RegisterOfficialName(xlanguage.Chinese, "新喀里多尼亚")
	dataNewCaledonia.RegisterCapital(xlanguage.Chinese, "努美阿")
}
