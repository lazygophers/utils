//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_qa || country_western_asia || currency_all || currency_qar)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	QAR.RegisterName(xlanguage.MustParse("zh-Hant"), "卡達里亞爾")
}
