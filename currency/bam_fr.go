//go:build (lang_fr || lang_all) && (country_all || country_ba || country_europe || country_southern_europe || currency_all || currency_bam)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BAM.RegisterName(xlanguage.French, "Mark convertible")
}
