//go:build (lang_ko || lang_all) && (country_all || country_ba || country_europe || country_southern_europe || currency_all || currency_bam)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BAM.RegisterName(xlanguage.Korean, "보스니아 헤르체고비나 태환 마르카")
}
