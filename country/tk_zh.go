//go:build country_all || country_oceania || country_polynesia || country_tk

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTokelau.RegisterName(xlanguage.Chinese, "托克劳")
	dataTokelau.RegisterOfficialName(xlanguage.Chinese, "托克劳")
}
