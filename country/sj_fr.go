//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSvalbardAndJanMayen.RegisterName(xlanguage.French, "Svalbard et Jan Mayen")
	dataSvalbardAndJanMayen.RegisterOfficialName(xlanguage.French, "Svalbard et Jan Mayen")
	dataSvalbardAndJanMayen.RegisterCapital(xlanguage.French, "Longyearbyen")
}
