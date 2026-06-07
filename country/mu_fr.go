//go:build country_africa || country_all || country_eastern_africa || country_mu

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritius.RegisterName(xlanguage.French, "Maurice")
	dataMauritius.RegisterOfficialName(xlanguage.French, "République de Maurice")
	dataMauritius.RegisterCapital(xlanguage.French, "Port-Louis")
}
