//go:build country_all || country_ba || country_europe || country_southern_europe || currency_all || currency_bam

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BAM.RegisterName(xlanguage.English, "Convertible Mark")
}
