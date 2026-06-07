//go:build (lang_ja || lang_all) && (country_all || country_asia || country_qa || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataQatar.RegisterName(xlanguage.Japanese, "カタール")
	dataQatar.RegisterOfficialName(xlanguage.Japanese, "カタール国")
	dataQatar.RegisterCapital(xlanguage.Japanese, "ドーハ")
}
