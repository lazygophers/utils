package fake

func init() {
	registerDataSet("de", "addresses", "cities", dsDe_0)
	registerDataSet("de", "addresses", "streets", dsDe_1)
	registerDataSet("de", "companies", "names", dsDe_2)
	registerDataSet("de", "companies", "suffixes", dsDe_3)
	registerDataSet("de", "names", "first_female", dsDe_4)
	registerDataSet("de", "names", "first_male", dsDe_5)
	registerDataSet("de", "names", "last", dsDe_6)
	registerDataSet("de", "texts", "lorem", dsDe_7)
}

var dsDe_0 = &DataSet{
	Language: "de",
	Country:  "DE",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Berlin", Weight: 2.0, Tags: []string{"capital", "major"}, Meta: map[string]string{"state": "Berlin"}},
		{Value: "Hamburg", Weight: 1.8, Tags: []string{"major", "port"}, Meta: map[string]string{"state": "Hamburg"}},
		{Value: "München", Weight: 1.6, Tags: []string{"major", "bavarian"}, Meta: map[string]string{"state": "Bayern"}},
		{Value: "Köln", Weight: 1.4, Tags: []string{"major", "rhine"}, Meta: map[string]string{"state": "Nordrhein-Westfalen"}},
		{Value: "Frankfurt am Main", Weight: 1.2, Tags: []string{"major", "financial"}, Meta: map[string]string{"state": "Hessen"}},
		{Value: "Stuttgart", Weight: 1.1, Tags: []string{"major", "automotive"}, Meta: map[string]string{"state": "Baden-Württemberg"}},
		{Value: "Düsseldorf", Tags: []string{"major", "business"}, Meta: map[string]string{"state": "Nordrhein-Westfalen"}},
		{Value: "Dortmund", Weight: 0.9, Tags: []string{"major", "ruhr"}, Meta: map[string]string{"state": "Nordrhein-Westfalen"}},
		{Value: "Essen", Weight: 0.8, Tags: []string{"major", "industrial"}, Meta: map[string]string{"state": "Nordrhein-Westfalen"}},
		{Value: "Leipzig", Weight: 0.7, Tags: []string{"major", "eastern"}, Meta: map[string]string{"state": "Sachsen"}},
	},
}

var dsDe_1 = &DataSet{
	Language: "de",
	Country:  "DE",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Hauptstraße", Weight: 2.0, Tags: []string{"common", "main"}},
		{Value: "Bahnhofstraße", Weight: 1.8, Tags: []string{"common", "transport"}},
		{Value: "Kirchstraße", Weight: 1.6, Tags: []string{"common", "religious"}},
		{Value: "Marktstraße", Weight: 1.4, Tags: []string{"commercial", "central"}},
		{Value: "Schulstraße", Weight: 1.2, Tags: []string{"education"}},
		{Value: "Gartenstraße", Weight: 1.1, Tags: []string{"residential", "green"}},
		{Value: "Lindenstraße", Tags: []string{"residential", "tree"}},
		{Value: "Bergstraße", Weight: 0.9, Tags: []string{"elevation"}},
		{Value: "Mühlenstraße", Weight: 0.8, Tags: []string{"historical", "mill"}},
		{Value: "Friedhofstraße", Weight: 0.7, Tags: []string{"cemetery"}},
	},
}

var dsDe_2 = &DataSet{
	Language: "de",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Siemens", Weight: 2.0, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "BMW", Weight: 1.85, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Mercedes", Weight: 1.7, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Volkswagen", Weight: 1.55, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "SAP", Weight: 1.4, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Bayer", Weight: 1.25, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "BASF", Weight: 1.1, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Deutsche Bank", Weight: 0.95, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Allianz", Weight: 0.8, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Adidas", Weight: 0.6500000000000001, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
	},
}

var dsDe_3 = &DataSet{
	Language: "de",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "GmbH", Weight: 3.0, Tags: []string{"common"}},
		{Value: "AG", Weight: 2.7, Tags: []string{"common"}},
		{Value: "KG", Weight: 2.4, Tags: []string{"common"}},
		{Value: "Co. KG", Weight: 2.1, Tags: []string{"formal"}},
		{Value: "OHG", Weight: 1.8, Tags: []string{"formal"}},
		{Value: "e.V.", Weight: 1.5, Tags: []string{"formal"}},
		{Value: "mbH", Weight: 1.2000000000000002, Tags: []string{"formal"}},
		{Value: "SE", Weight: 0.8999999999999999, Tags: []string{"formal"}},
	},
}

var dsDe_4 = &DataSet{
	Language: "de",
	Country:  "DE",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Anna", Weight: 2.0, Tags: []string{"traditional", "biblical"}},
		{Value: "Emma", Weight: 1.8, Tags: []string{"modern", "popular"}},
		{Value: "Mia", Weight: 1.6, Tags: []string{"modern", "short"}},
		{Value: "Sophie", Weight: 1.4, Tags: []string{"traditional", "wisdom"}},
		{Value: "Marie", Weight: 1.4, Tags: []string{"traditional", "biblical"}},
		{Value: "Lena", Weight: 1.2, Tags: []string{"modern", "popular"}},
		{Value: "Laura", Weight: 1.1, Tags: []string{"classical", "nature"}},
		{Value: "Lea", Tags: []string{"biblical", "modern"}},
		{Value: "Julia", Weight: 0.9, Tags: []string{"classical", "noble"}},
		{Value: "Sarah", Weight: 0.8, Tags: []string{"biblical", "traditional"}},
	},
}

var dsDe_5 = &DataSet{
	Language: "de",
	Country:  "DE",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Max", Weight: 2.0, Tags: []string{"traditional", "popular"}},
		{Value: "Paul", Weight: 1.8, Tags: []string{"biblical", "traditional"}},
		{Value: "Leon", Weight: 1.6, Tags: []string{"modern", "strong"}},
		{Value: "Ben", Weight: 1.4, Tags: []string{"modern", "short"}},
		{Value: "Felix", Weight: 1.4, Tags: []string{"traditional", "happy"}},
		{Value: "Alexander", Weight: 1.2, Tags: []string{"classical", "strong"}},
		{Value: "Maximilian", Weight: 1.1, Tags: []string{"noble", "traditional"}},
		{Value: "Johannes", Tags: []string{"biblical", "traditional"}},
		{Value: "David", Weight: 0.9, Tags: []string{"biblical", "traditional"}},
		{Value: "Thomas", Weight: 0.8, Tags: []string{"biblical", "traditional"}},
	},
}

var dsDe_6 = &DataSet{
	Language: "de",
	Country:  "DE",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Müller", Weight: 2.0, Tags: []string{"occupational"}},
		{Value: "Schmidt", Weight: 1.8, Tags: []string{"occupational"}},
		{Value: "Schneider", Weight: 1.6, Tags: []string{"occupational"}},
		{Value: "Fischer", Weight: 1.4, Tags: []string{"occupational"}},
		{Value: "Weber", Weight: 1.2, Tags: []string{"occupational"}},
		{Value: "Meyer", Weight: 1.1, Tags: []string{"occupational"}},
		{Value: "Wagner", Tags: []string{"occupational"}},
		{Value: "Becker", Weight: 0.9, Tags: []string{"occupational"}},
		{Value: "Schulz", Weight: 0.8, Tags: []string{"occupational"}},
		{Value: "Hoffmann", Weight: 0.7, Tags: []string{"occupational"}},
	},
}

var dsDe_7 = &DataSet{
	Language: "de",
	Type:     "texts",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "text"},
		{Value: "wort", Weight: 0.95},
		{Value: "satz", Weight: 0.9},
		{Value: "absatz", Weight: 0.85},
		{Value: "artikel", Weight: 0.8},
		{Value: "bericht", Weight: 0.75},
		{Value: "studie", Weight: 0.7},
		{Value: "forschung", Weight: 0.6499999999999999},
		{Value: "analyse", Weight: 0.6},
		{Value: "inhalt", Weight: 0.55},
		{Value: "thema", Weight: 0.5},
		{Value: "begriff", Weight: 0.44999999999999996},
		{Value: "bedeutung", Weight: 0.3999999999999999},
		{Value: "sprache", Weight: 0.35},
		{Value: "literatur", Weight: 0.29999999999999993},
	},
}
