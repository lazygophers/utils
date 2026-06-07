//go:build (lang_ja || lang_all) && (country_all || country_americas || country_br || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrazil.RegisterName(xlanguage.Japanese, "ブラジル")
	dataBrazil.RegisterOfficialName(xlanguage.Japanese, "ブラジル連邦共和国")
	dataBrazil.RegisterCapital(xlanguage.Japanese, "ブラジリア")
}
