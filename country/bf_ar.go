//go:build (lang_ar || lang_all) && (country_africa || country_all || country_bf || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurkinaFaso.RegisterName(xlanguage.Arabic, "بوركينا فاسو")
	dataBurkinaFaso.RegisterOfficialName(xlanguage.Arabic, "بوركينا فاسو")
	dataBurkinaFaso.RegisterCapital(xlanguage.Arabic, "واغادوغو")
}
