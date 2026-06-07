package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalestine.RegisterName(xlanguage.English, "Palestine")
	dataPalestine.RegisterOfficialName(xlanguage.English, "State of Palestine")
	dataPalestine.RegisterCapital(xlanguage.English, "Ramallah")
}
