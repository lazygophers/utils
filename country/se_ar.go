//go:build (lang_ar || lang_all) && (country_all || country_europe || country_northern_europe || country_se)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSweden.RegisterName(xlanguage.Arabic, "السويد")
	dataSweden.RegisterOfficialName(xlanguage.Arabic, "مملكة السويد")
	dataSweden.RegisterCapital(xlanguage.Arabic, "ستوكهولم")
}
