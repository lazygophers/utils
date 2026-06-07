//go:build (lang_ko || lang_all) && (country_all || country_americas || country_central_america || country_mx)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMexico.RegisterName(xlanguage.Korean, "멕시코")
	dataMexico.RegisterOfficialName(xlanguage.Korean, "멕시코 합중국")
	dataMexico.RegisterCapital(xlanguage.Korean, "멕시코시티")
}
