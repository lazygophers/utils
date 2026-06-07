//go:build (lang_ar || lang_all) && (country_africa || country_all || country_mr || country_western_africa || currency_all || currency_mru)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MRU.RegisterName(xlanguage.Arabic, "أوقية موريتانية")
}
