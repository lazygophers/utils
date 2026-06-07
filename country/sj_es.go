//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSvalbardAndJanMayen.RegisterName(xlanguage.Spanish, "Svalbard y Jan Mayen")
	dataSvalbardAndJanMayen.RegisterOfficialName(xlanguage.Spanish, "Svalbard y Jan Mayen")
	dataSvalbardAndJanMayen.RegisterCapital(xlanguage.Spanish, "Longyearbyen")
}
