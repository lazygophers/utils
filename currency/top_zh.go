//go:build country_all || country_oceania || country_polynesia || country_to || currency_all || currency_top

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	TOP.RegisterName(xlanguage.Chinese, "汤加潘加")
}
