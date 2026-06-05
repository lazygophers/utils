package fake

func init() {
	registerDataSet("fr", "addresses", "cities", dsFr_0)
	registerDataSet("fr", "addresses", "streets", dsFr_1)
	registerDataSet("fr", "companies", "names", dsFr_2)
	registerDataSet("fr", "companies", "suffixes", dsFr_3)
	registerDataSet("fr", "names", "first_female", dsFr_4)
	registerDataSet("fr", "names", "first_male", dsFr_5)
	registerDataSet("fr", "names", "last", dsFr_6)
	registerDataSet("fr", "texts", "lorem", dsFr_7)
}

var dsFr_0 = &DataSet{
	Language: "fr",
	Country:  "FR",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Paris", Weight: 2.0, Tags: []string{"capital", "major"}},
		{Value: "Marseille", Weight: 1.8, Tags: []string{"major", "port"}},
		{Value: "Lyon", Weight: 1.6, Tags: []string{"major", "gastronomy"}},
		{Value: "Toulouse", Weight: 1.4, Tags: []string{"major", "aerospace"}},
		{Value: "Nice", Weight: 1.2, Tags: []string{"coastal", "tourism"}},
		{Value: "Nantes", Weight: 1.1, Tags: []string{"major", "western"}},
		{Value: "Strasbourg", Tags: []string{"european", "eastern"}},
		{Value: "Montpellier", Weight: 0.9, Tags: []string{"university", "southern"}},
		{Value: "Bordeaux", Weight: 0.8, Tags: []string{"wine", "major"}},
		{Value: "Lille", Weight: 0.7, Tags: []string{"northern", "industrial"}},
	},
}

var dsFr_1 = &DataSet{
	Language: "fr",
	Country:  "FR",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Rue de la Paix", Weight: 2.0, Tags: []string{"famous", "upscale"}},
		{Value: "Avenue des Champs-Élysées", Weight: 1.8, Tags: []string{"famous", "tourist"}},
		{Value: "Rue Saint-Honoré", Weight: 1.6, Tags: []string{"commercial", "fashion"}},
		{Value: "Boulevard Saint-Germain", Weight: 1.4, Tags: []string{"cultural", "major"}},
		{Value: "Rue de Rivoli", Weight: 1.2, Tags: []string{"central", "shopping"}},
		{Value: "Avenue Montaigne", Weight: 1.1, Tags: []string{"luxury", "fashion"}},
		{Value: "Rue du Faubourg Saint-Antoine", Tags: []string{"artisan", "historical"}},
		{Value: "Boulevard Haussmann", Weight: 0.9, Tags: []string{"shopping", "major"}},
		{Value: "Rue de la République", Weight: 0.8, Tags: []string{"common", "central"}},
		{Value: "Avenue Victor Hugo", Weight: 0.7, Tags: []string{"residential", "prestigious"}},
	},
}

var dsFr_2 = &DataSet{
	Language: "fr",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Total", Weight: 2.0, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "L'Oréal", Weight: 1.85, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "LVMH", Weight: 1.7, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Carrefour", Weight: 1.55, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Peugeot", Weight: 1.4, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Renault", Weight: 1.25, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "EDF", Weight: 1.1, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Orange", Weight: 0.95, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Michelin", Weight: 0.8, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Danone", Weight: 0.6500000000000001, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
	},
}

var dsFr_3 = &DataSet{
	Language: "fr",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "SA", Weight: 3.0, Tags: []string{"common"}},
		{Value: "SARL", Weight: 2.7, Tags: []string{"common"}},
		{Value: "SAS", Weight: 2.4, Tags: []string{"common"}},
		{Value: "SASU", Weight: 2.1, Tags: []string{"formal"}},
		{Value: "SNC", Weight: 1.8, Tags: []string{"formal"}},
		{Value: "Société Anonyme", Weight: 1.5, Tags: []string{"formal"}},
		{Value: "EURL", Weight: 1.2000000000000002, Tags: []string{"formal"}},
	},
}

var dsFr_4 = &DataSet{
	Language: "fr",
	Country:  "FR",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Marie", Weight: 2.0, Tags: []string{"traditional", "biblical"}},
		{Value: "Françoise", Weight: 1.8, Tags: []string{"traditional", "french"}},
		{Value: "Monique", Weight: 1.6, Tags: []string{"traditional", "advisor"}},
		{Value: "Catherine", Weight: 1.4, Tags: []string{"traditional", "pure"}},
		{Value: "Nathalie", Weight: 1.2, Tags: []string{"modern", "birth"}},
		{Value: "Isabelle", Weight: 1.1, Tags: []string{"traditional", "oath"}},
		{Value: "Sylvie", Tags: []string{"modern", "forest"}},
		{Value: "Martine", Weight: 0.9, Tags: []string{"traditional", "war"}},
		{Value: "Nicole", Weight: 0.8, Tags: []string{"modern", "victory"}},
		{Value: "Brigitte", Weight: 0.7, Tags: []string{"traditional", "strength"}},
	},
}

var dsFr_5 = &DataSet{
	Language: "fr",
	Country:  "FR",
	Type:     "names",
	Version:  "2.0.0",
	Items: []DataItem{
		{Value: "Jean", Weight: 2.0, Tags: []string{"traditional", "biblical"}},
		{Value: "Pierre", Weight: 1.8, Tags: []string{"apostle", "rock"}},
		{Value: "Michel", Weight: 1.6, Tags: []string{"archangel", "biblical"}},
		{Value: "André", Weight: 1.4, Tags: []string{"apostle", "manly"}},
		{Value: "Philippe", Weight: 1.2, Tags: []string{"apostle", "royal"}},
		{Value: "Alain", Tags: []string{"breton", "noble"}},
		{Value: "Bernard", Tags: []string{"saint", "bear"}},
		{Value: "Robert", Weight: 0.9, Tags: []string{"royal", "bright"}},
		{Value: "Jacques", Weight: 0.9, Tags: []string{"apostle", "supplanter"}},
		{Value: "Daniel", Weight: 0.8, Tags: []string{"biblical", "prophet"}},
		{Value: "Henri", Weight: 0.8, Tags: []string{"royal", "ruler"}},
		{Value: "Paul", Weight: 0.7, Tags: []string{"apostle", "humble"}},
		{Value: "Christian", Weight: 0.7, Tags: []string{"religious", "christ"}},
		{Value: "Patrick", Weight: 0.6, Tags: []string{"saint", "noble"}},
		{Value: "Nicolas", Weight: 0.6, Tags: []string{"saint", "victory"}},
		{Value: "Dominique", Weight: 0.5, Tags: []string{"religious", "lord"}},
		{Value: "François", Weight: 0.5, Tags: []string{"saint", "free"}},
		{Value: "Christophe", Weight: 0.5, Tags: []string{"saint", "bearer"}},
		{Value: "Olivier", Weight: 0.4, Tags: []string{"nature", "peace"}},
		{Value: "Stéphane", Weight: 0.4, Tags: []string{"martyr", "crown"}},
		{Value: "Thierry", Weight: 0.4, Tags: []string{"germanic", "ruler"}},
		{Value: "Laurent", Weight: 0.4, Tags: []string{"martyr", "laurel"}},
		{Value: "Sébastien", Weight: 0.3, Tags: []string{"martyr", "venerable"}},
		{Value: "David", Weight: 0.3, Tags: []string{"biblical", "king"}},
		{Value: "Julien", Weight: 0.3, Tags: []string{"roman", "youthful"}},
		{Value: "Marc", Weight: 0.3, Tags: []string{"evangelist", "war"}},
		{Value: "Frédéric", Weight: 0.3, Tags: []string{"germanic", "peace"}},
		{Value: "Vincent", Weight: 0.3, Tags: []string{"martyr", "conquering"}},
	},
}

var dsFr_6 = &DataSet{
	Language: "fr",
	Country:  "FR",
	Type:     "names",
	Version:  "2.0.0",
	Items: []DataItem{
		{Value: "Martin", Weight: 3.0, Tags: []string{"common", "saint"}},
		{Value: "Bernard", Weight: 2.8, Tags: []string{"common", "germanic"}},
		{Value: "Dubois", Weight: 2.6, Tags: []string{"topographic", "woods"}},
		{Value: "Thomas", Weight: 2.4, Tags: []string{"common", "apostle"}},
		{Value: "Robert", Weight: 2.2, Tags: []string{"common", "germanic"}},
		{Value: "Richard", Weight: 2.0, Tags: []string{"royal", "germanic"}},
		{Value: "Petit", Weight: 1.8, Tags: []string{"descriptive", "size"}},
		{Value: "Durand", Weight: 1.6, Tags: []string{"common", "enduring"}},
		{Value: "Leroy", Weight: 1.4, Tags: []string{"royal", "king"}},
		{Value: "Moreau", Weight: 1.2, Tags: []string{"descriptive", "dark"}},
		{Value: "Simon", Tags: []string{"biblical", "apostle"}},
		{Value: "Laurent", Tags: []string{"roman", "laurel"}},
		{Value: "Lefebvre", Weight: 0.9, Tags: []string{"occupational", "smith"}},
		{Value: "Michel", Weight: 0.9, Tags: []string{"biblical", "archangel"}},
		{Value: "Garcia", Weight: 0.8, Tags: []string{"spanish", "bear"}},
		{Value: "Roux", Weight: 0.8, Tags: []string{"descriptive", "red"}},
		{Value: "David", Weight: 0.7, Tags: []string{"biblical", "king"}},
		{Value: "Bertrand", Weight: 0.7, Tags: []string{"germanic", "bright"}},
		{Value: "Vincent", Weight: 0.6, Tags: []string{"roman", "conquering"}},
		{Value: "Fournier", Weight: 0.6, Tags: []string{"occupational", "baker"}},
		{Value: "Morel", Weight: 0.5, Tags: []string{"descriptive", "dark"}},
		{Value: "Girard", Weight: 0.5, Tags: []string{"germanic", "spear"}},
		{Value: "André", Weight: 0.5, Tags: []string{"apostle", "manly"}},
		{Value: "Mercier", Weight: 0.4, Tags: []string{"occupational", "merchant"}},
		{Value: "Dupont", Weight: 0.4, Tags: []string{"topographic", "bridge"}},
		{Value: "Boyer", Weight: 0.4, Tags: []string{"occupational", "wood"}},
		{Value: "Blanc", Weight: 0.3, Tags: []string{"descriptive", "white"}},
		{Value: "Robin", Weight: 0.3, Tags: []string{"diminutive", "fame"}},
		{Value: "Lemoine", Weight: 0.3, Tags: []string{"religious", "monk"}},
		{Value: "Rousseau", Weight: 0.3, Tags: []string{"descriptive", "red"}},
	},
}

var dsFr_7 = &DataSet{
	Language: "fr",
	Type:     "texts",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "texte"},
		{Value: "mot", Weight: 0.95},
		{Value: "phrase", Weight: 0.9},
		{Value: "paragraphe", Weight: 0.85},
		{Value: "article", Weight: 0.8},
		{Value: "rapport", Weight: 0.75},
		{Value: "étude", Weight: 0.7},
		{Value: "analyse", Weight: 0.6499999999999999},
		{Value: "contenu", Weight: 0.6},
		{Value: "sujet", Weight: 0.55},
		{Value: "concept", Weight: 0.5},
		{Value: "signification", Weight: 0.44999999999999996},
		{Value: "langue", Weight: 0.3999999999999999},
		{Value: "littérature", Weight: 0.35},
		{Value: "écriture", Weight: 0.29999999999999993},
	},
}
