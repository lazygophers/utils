package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSyria.RegisterName(xlanguage.Arabic, "سوريا")
	dataSyria.RegisterOfficialName(xlanguage.Arabic, "الجمهورية العربية السورية")
	dataSyria.RegisterCapital(xlanguage.Arabic, "دمشق")
}
