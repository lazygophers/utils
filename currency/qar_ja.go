//go:build (lang_ja || lang_all) && (country_all || country_asia || country_qa || country_western_asia || currency_all || currency_qar)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	QAR.RegisterName(xlanguage.Japanese, "カタール・リヤル")
}
