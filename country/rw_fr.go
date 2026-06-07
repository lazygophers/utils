package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRwanda.RegisterName(xlanguage.French, "Rwanda")
	dataRwanda.RegisterOfficialName(xlanguage.French, "République du Rwanda")
	dataRwanda.RegisterCapital(xlanguage.French, "Kigali")
}
