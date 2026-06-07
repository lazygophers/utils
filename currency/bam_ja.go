//go:build (lang_ja || lang_all) && (country_all || country_ba || country_europe || country_southern_europe || currency_all || currency_bam)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BAM.RegisterName(xlanguage.Japanese, "兌換マルク")
}
