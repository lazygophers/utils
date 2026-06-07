//go:build (lang_fr || lang_all) && (country_all || country_americas || country_py || country_south_america || currency_all || currency_pyg)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Pyg.RegisterName(xlanguage.French, "Guarani")
}
