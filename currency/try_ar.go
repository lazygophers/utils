//go:build (lang_ar || lang_all) && (country_all || country_asia || country_tr || country_western_asia || currency_all || currency_try)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Try.RegisterName(xlanguage.Arabic, "ليرة تركية")
}
