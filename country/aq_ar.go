//go:build (lang_ar || lang_all) && (country_all || country_antarctic || country_aq)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntarctica.RegisterName(xlanguage.Arabic, "القارة القطبية الجنوبية")
	dataAntarctica.RegisterOfficialName(xlanguage.Arabic, "القارة القطبية الجنوبية")
	dataAntarctica.RegisterCapital(xlanguage.Arabic, "—")
}
