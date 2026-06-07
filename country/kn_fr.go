//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintKittsAndNevis.RegisterName(xlanguage.French, "Saint-Christophe-et-Niévès")
	dataSaintKittsAndNevis.RegisterOfficialName(xlanguage.French, "Fédération de Saint-Christophe-et-Niévès")
	dataSaintKittsAndNevis.RegisterCapital(xlanguage.French, "Basseterre")
}
