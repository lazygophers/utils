//go:build (lang_fr || lang_all) && (country_all || country_asia || country_ps || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalestine.RegisterName(xlanguage.French, "Palestine")
	dataPalestine.RegisterOfficialName(xlanguage.French, "État de Palestine")
	dataPalestine.RegisterCapital(xlanguage.French, "Ramallah")
}
