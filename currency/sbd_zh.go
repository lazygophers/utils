//go:build country_all || country_melanesia || country_oceania || country_sb || currency_all || currency_sbd

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Sbd.RegisterName(xlanguage.Chinese, "所罗门群岛元")
}
