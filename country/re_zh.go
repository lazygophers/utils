//go:build country_africa || country_all || country_eastern_africa || country_re

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataReunion.RegisterName(xlanguage.Chinese, "留尼汪")
	dataReunion.RegisterOfficialName(xlanguage.Chinese, "留尼汪")
	dataReunion.RegisterCapital(xlanguage.Chinese, "圣但尼")
}
