//go:build country_all || country_americas || country_caribbean || country_jm || currency_all || currency_jmd

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	JMD.RegisterName(xlanguage.English, "Jamaican Dollar")
}
