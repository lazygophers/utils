//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_sx)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSintMaarten.RegisterName(xlanguage.Arabic, "سينت مارتن")
	dataSintMaarten.RegisterOfficialName(xlanguage.Arabic, "سينت مارتن")
	dataSintMaarten.RegisterCapital(xlanguage.Arabic, "فيليبسبورغ")
}
