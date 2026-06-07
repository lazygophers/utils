//go:build country_all || country_melanesia || country_nc || country_oceania || country_pf || country_polynesia || country_wf || currency_all || currency_xpf

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Xpf.RegisterName(xlanguage.English, "CFP Franc")
}
