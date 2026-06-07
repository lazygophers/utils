//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSintMaarten.RegisterName(xlanguage.Arabic, "سينت مارتن")
	dataSintMaarten.RegisterOfficialName(xlanguage.Arabic, "سينت مارتن")
	dataSintMaarten.RegisterCapital(xlanguage.Arabic, "فيليبسبورغ")
}
