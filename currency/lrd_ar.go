//go:build (lang_ar || lang_all) && (country_africa || country_all || country_lr || country_western_africa || currency_all || currency_lrd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	LRD.RegisterName(xlanguage.Arabic, "دولار ليبيري")
}
