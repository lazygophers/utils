//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJersey.RegisterName(xlanguage.French, "Jersey")
	dataJersey.RegisterOfficialName(xlanguage.French, "Bailliage de Jersey")
	dataJersey.RegisterCapital(xlanguage.French, "Saint-Hélier")
}
