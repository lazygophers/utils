//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuracao.RegisterName(xlanguage.Arabic, "كوراساو")
	dataCuracao.RegisterOfficialName(xlanguage.Arabic, "كوراساو")
	dataCuracao.RegisterCapital(xlanguage.Arabic, "ويلمستاد")
}
