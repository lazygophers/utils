//go:build (lang_ja || lang_all) && (country_all || country_oceania || country_polynesia || country_ws)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSamoa.RegisterName(xlanguage.Japanese, "サモア")
	dataSamoa.RegisterOfficialName(xlanguage.Japanese, "サモア独立国")
	dataSamoa.RegisterCapital(xlanguage.Japanese, "アピア")
}
