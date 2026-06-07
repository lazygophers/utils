//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalestine.RegisterName(xlanguage.French, "Palestine")
	dataPalestine.RegisterOfficialName(xlanguage.French, "État de Palestine")
	dataPalestine.RegisterCapital(xlanguage.French, "Ramallah")
}
