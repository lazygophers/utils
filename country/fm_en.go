//go:build country_all || country_fm || country_micronesia || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMicronesia.RegisterName(xlanguage.English, "Micronesia")
	dataMicronesia.RegisterOfficialName(xlanguage.English, "Federated States of Micronesia")
	dataMicronesia.RegisterCapital(xlanguage.English, "Palikir")
}
