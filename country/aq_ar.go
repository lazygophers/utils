//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntarctica.RegisterName(xlanguage.Arabic, "القارة القطبية الجنوبية")
	dataAntarctica.RegisterOfficialName(xlanguage.Arabic, "القارة القطبية الجنوبية")
	dataAntarctica.RegisterCapital(xlanguage.Arabic, "—")
}
