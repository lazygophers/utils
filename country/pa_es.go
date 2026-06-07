//go:build country_all || country_americas || country_central_america || country_pa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPanama.RegisterName(xlanguage.Spanish, "Panamá")
	dataPanama.RegisterOfficialName(xlanguage.Spanish, "República de Panamá")
	dataPanama.RegisterCapital(xlanguage.Spanish, "Ciudad de Panamá")
}
