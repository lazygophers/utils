//go:build (lang_ja || lang_all) && (country_all || country_americas || country_bm || country_northern_america || currency_all || currency_bmd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BMD.RegisterName(xlanguage.Japanese, "バミューダ・ドル")
}
