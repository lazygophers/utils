//go:build (lang_ar || lang_all) && (country_all || country_cz || country_eastern_europe || country_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCzechia.RegisterName(xlanguage.Arabic, "التشيك")
	dataCzechia.RegisterOfficialName(xlanguage.Arabic, "الجمهورية التشيكية")
	dataCzechia.RegisterCapital(xlanguage.Arabic, "براغ")
}
