//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaymanIslands.RegisterName(xlanguage.French, "Îles Caïmans")
	dataCaymanIslands.RegisterOfficialName(xlanguage.French, "Îles Caïmans")
	dataCaymanIslands.RegisterCapital(xlanguage.French, "George Town")
}
