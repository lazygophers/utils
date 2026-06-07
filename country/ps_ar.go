package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalestine.RegisterName(xlanguage.Arabic, "فلسطين")
	dataPalestine.RegisterOfficialName(xlanguage.Arabic, "دولة فلسطين")
	dataPalestine.RegisterCapital(xlanguage.Arabic, "القدس")
}
