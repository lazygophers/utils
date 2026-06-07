//go:build (lang_ja || lang_all) && (country_all || country_americas || country_central_america || country_ni || currency_all || currency_nio)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Nio.RegisterName(xlanguage.Japanese, "ニカラグア・コルドバ")
}
