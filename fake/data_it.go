package fake

func init() {
	registerDataSet("it", "addresses", "cities", dsIt_0)
	registerDataSet("it", "addresses", "streets", dsIt_1)
	registerDataSet("it", "companies", "names", dsIt_2)
	registerDataSet("it", "companies", "suffixes", dsIt_3)
	registerDataSet("it", "names", "first_female", dsIt_4)
	registerDataSet("it", "names", "first_male", dsIt_5)
	registerDataSet("it", "names", "last", dsIt_6)
	registerDataSet("it", "texts", "lorem", dsIt_7)
}

var dsIt_0 = &DataSet{
	Language: "it",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Roma", Weight: 2.0, Tags: []string{"major"}},
		{Value: "Milano", Weight: 1.8, Tags: []string{"major"}},
		{Value: "Napoli", Weight: 1.6, Tags: []string{"major"}},
		{Value: "Torino", Weight: 1.4, Tags: []string{"medium"}},
		{Value: "Palermo", Weight: 1.2, Tags: []string{"medium"}},
		{Value: "Genova", Tags: []string{"medium"}},
		{Value: "Bologna", Weight: 0.7999999999999998, Tags: []string{"medium"}},
		{Value: "Firenze", Weight: 0.5999999999999999, Tags: []string{"medium"}},
		{Value: "Bari", Weight: 0.5, Tags: []string{"medium"}},
		{Value: "Catania", Weight: 0.5, Tags: []string{"medium"}},
	},
}

var dsIt_1 = &DataSet{
	Language: "it",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Via Roma", Weight: 2.0, Tags: []string{"common"}},
		{Value: "Corso Italia", Weight: 1.7, Tags: []string{"common"}},
		{Value: "Via Nazionale", Weight: 1.4, Tags: []string{"common"}},
		{Value: "Piazza del Duomo", Weight: 1.1, Tags: []string{"common"}},
		{Value: "Via Garibaldi", Weight: 0.8, Tags: []string{"common"}},
	},
}

var dsIt_2 = &DataSet{
	Language: "it",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Ferrari", Weight: 2.0, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Fiat", Weight: 1.85, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Gucci", Weight: 1.7, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Prada", Weight: 1.55, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Versace", Weight: 1.4, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Benetton", Weight: 1.25, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Olivetti", Weight: 1.1, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Pirelli", Weight: 0.95, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Alitalia", Weight: 0.8, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Barilla", Weight: 0.6500000000000001, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
	},
}

var dsIt_3 = &DataSet{
	Language: "it",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "S.p.A.", Weight: 3.0, Tags: []string{"common"}},
		{Value: "S.r.l.", Weight: 2.7, Tags: []string{"common"}},
		{Value: "S.n.c.", Weight: 2.4, Tags: []string{"common"}},
		{Value: "S.a.s.", Weight: 2.1, Tags: []string{"formal"}},
		{Value: "Società", Weight: 1.8, Tags: []string{"formal"}},
		{Value: "Gruppo", Weight: 1.5, Tags: []string{"formal"}},
	},
}

var dsIt_4 = &DataSet{
	Language: "it",
	Country:  "IT",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Giulia", Weight: 2.0, Tags: []string{"classical", "youthful"}},
		{Value: "Francesca", Weight: 1.8, Tags: []string{"traditional", "french"}},
		{Value: "Anna", Weight: 1.6, Tags: []string{"biblical", "grace"}},
		{Value: "Sara", Weight: 1.4, Tags: []string{"biblical", "princess"}},
		{Value: "Chiara", Weight: 1.2, Tags: []string{"traditional", "bright"}},
		{Value: "Alessia", Weight: 1.1, Tags: []string{"modern", "defender"}},
		{Value: "Martina", Tags: []string{"traditional", "war"}},
		{Value: "Valentina", Weight: 0.9, Tags: []string{"traditional", "strong"}},
		{Value: "Elena", Weight: 0.8, Tags: []string{"classical", "light"}},
		{Value: "Giorgia", Weight: 0.7, Tags: []string{"traditional", "farmer"}},
	},
}

var dsIt_5 = &DataSet{
	Language: "it",
	Country:  "IT",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Marco", Weight: 2.0, Tags: []string{"traditional", "war"}},
		{Value: "Alessandro", Weight: 1.8, Tags: []string{"classical", "defender"}},
		{Value: "Luca", Weight: 1.6, Tags: []string{"traditional", "light"}},
		{Value: "Matteo", Weight: 1.4, Tags: []string{"biblical", "gift"}},
		{Value: "Andrea", Weight: 1.2, Tags: []string{"traditional", "brave"}},
		{Value: "Francesco", Weight: 1.1, Tags: []string{"religious", "french"}},
		{Value: "Giuseppe", Tags: []string{"biblical", "traditional"}},
		{Value: "Antonio", Weight: 0.9, Tags: []string{"traditional", "roman"}},
		{Value: "Lorenzo", Weight: 0.8, Tags: []string{"classical", "laurel"}},
		{Value: "Giovanni", Weight: 0.7, Tags: []string{"biblical", "traditional"}},
	},
}

var dsIt_6 = &DataSet{
	Language: "it",
	Country:  "IT",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Rossi", Weight: 2.0, Tags: []string{"descriptive"}},
		{Value: "Ferrari", Weight: 1.8, Tags: []string{"occupational"}},
		{Value: "Bianchi", Weight: 1.6, Tags: []string{"descriptive"}},
		{Value: "Romano", Weight: 1.4, Tags: []string{"geographical"}},
		{Value: "Galli", Weight: 1.2, Tags: []string{"geographical"}},
		{Value: "Conti", Weight: 1.1, Tags: []string{"noble"}},
		{Value: "De Luca", Tags: []string{"patronymic"}},
		{Value: "Mancini", Weight: 0.9, Tags: []string{"descriptive"}},
		{Value: "Costa", Weight: 0.8, Tags: []string{"geographical"}},
		{Value: "Giordano", Weight: 0.7, Tags: []string{"geographical"}},
	},
}

var dsIt_7 = &DataSet{
	Language: "it",
	Type:     "texts",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "testo"},
		{Value: "parola", Weight: 0.95},
		{Value: "frase", Weight: 0.9},
		{Value: "paragrafo", Weight: 0.85},
		{Value: "articolo", Weight: 0.8},
		{Value: "rapporto", Weight: 0.75},
		{Value: "studio", Weight: 0.7},
		{Value: "analisi", Weight: 0.6499999999999999},
		{Value: "contenuto", Weight: 0.6},
		{Value: "argomento", Weight: 0.55},
		{Value: "concetto", Weight: 0.5},
		{Value: "significato", Weight: 0.44999999999999996},
		{Value: "lingua", Weight: 0.3999999999999999},
		{Value: "letteratura", Weight: 0.35},
		{Value: "scrittura", Weight: 0.29999999999999993},
	},
}
