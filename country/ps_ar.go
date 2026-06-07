//go:build country_all || country_asia || country_ps || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalestine.RegisterName(xlanguage.Arabic, "فلسطين")
	dataPalestine.RegisterOfficialName(xlanguage.Arabic, "دولة فلسطين")
	dataPalestine.RegisterCapital(xlanguage.Arabic, "القدس")
}
