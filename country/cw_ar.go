//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_cw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuracao.RegisterName(xlanguage.Arabic, "كوراساو")
	dataCuracao.RegisterOfficialName(xlanguage.Arabic, "كوراساو")
	dataCuracao.RegisterCapital(xlanguage.Arabic, "ويلمستاد")
}
