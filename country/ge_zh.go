//go:build country_all || country_asia || country_ge || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGeorgia.RegisterName(xlanguage.Chinese, "格鲁吉亚")
	dataGeorgia.RegisterOfficialName(xlanguage.Chinese, "格鲁吉亚")
	dataGeorgia.RegisterCapital(xlanguage.Chinese, "第比利斯")
}
