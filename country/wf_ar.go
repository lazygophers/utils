//go:build (lang_ar || lang_all) && (country_all || country_oceania || country_polynesia || country_wf)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWallisAndFutuna.RegisterName(xlanguage.Arabic, "والس وفوتونا")
	dataWallisAndFutuna.RegisterOfficialName(xlanguage.Arabic, "إقليم والس وفوتونا")
	dataWallisAndFutuna.RegisterCapital(xlanguage.Arabic, "ماتا أوتو")
}
