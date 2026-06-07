//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBouvetIsland.RegisterName(xlanguage.French, "Île Bouvet")
	dataBouvetIsland.RegisterOfficialName(xlanguage.French, "Île Bouvet")
}
