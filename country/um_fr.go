//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsMinorOutlyingIslands.RegisterName(xlanguage.French, "Îles mineures éloignées des États-Unis")
	dataUsMinorOutlyingIslands.RegisterOfficialName(xlanguage.French, "Îles mineures éloignées des États-Unis")
}
