package fake

func init() {
	registerDataSet("pt", "addresses", "cities", dsPt_0)
	registerDataSet("pt", "addresses", "streets", dsPt_1)
	registerDataSet("pt", "companies", "names", dsPt_2)
	registerDataSet("pt", "companies", "suffixes", dsPt_3)
	registerDataSet("pt", "names", "first_female", dsPt_4)
	registerDataSet("pt", "names", "first_male", dsPt_5)
	registerDataSet("pt", "names", "last", dsPt_6)
	registerDataSet("pt", "texts", "lorem", dsPt_7)
}

var dsPt_0 = &DataSet{
	Language: "pt",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "São Paulo", Weight: 2.0, Tags: []string{"major"}},
		{Value: "Rio de Janeiro", Weight: 1.8, Tags: []string{"major"}},
		{Value: "Brasília", Weight: 1.6, Tags: []string{"major"}},
		{Value: "Salvador", Weight: 1.4, Tags: []string{"medium"}},
		{Value: "Fortaleza", Weight: 1.2, Tags: []string{"medium"}},
		{Value: "Belo Horizonte", Tags: []string{"medium"}},
		{Value: "Manaus", Weight: 0.7999999999999998, Tags: []string{"medium"}},
		{Value: "Curitiba", Weight: 0.5999999999999999, Tags: []string{"medium"}},
		{Value: "Recife", Weight: 0.5, Tags: []string{"medium"}},
		{Value: "Goiânia", Weight: 0.5, Tags: []string{"medium"}},
	},
}

var dsPt_1 = &DataSet{
	Language: "pt",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Avenida Paulista", Weight: 2.0, Tags: []string{"common"}},
		{Value: "Rua Augusta", Weight: 1.7, Tags: []string{"common"}},
		{Value: "Avenida Brasil", Weight: 1.4, Tags: []string{"common"}},
		{Value: "Rua da Consolação", Weight: 1.1, Tags: []string{"common"}},
		{Value: "Avenida Atlântica", Weight: 0.8, Tags: []string{"common"}},
	},
}

var dsPt_2 = &DataSet{
	Language: "pt",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Petrobras", Weight: 2.0, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Vale", Weight: 1.85, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Itaú", Weight: 1.7, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Bradesco", Weight: 1.55, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "JBS", Weight: 1.4, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Embraer", Weight: 1.25, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Gerdau", Weight: 1.1, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Ambev", Weight: 0.95, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Natura", Weight: 0.8, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Magazine Luiza", Weight: 0.6500000000000001, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
	},
}

var dsPt_3 = &DataSet{
	Language: "pt",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "S.A.", Weight: 3.0, Tags: []string{"common"}},
		{Value: "Ltda.", Weight: 2.7, Tags: []string{"common"}},
		{Value: "EIRELI", Weight: 2.4, Tags: []string{"common"}},
		{Value: "Sociedade Anônima", Weight: 2.1, Tags: []string{"formal"}},
		{Value: "Limitada", Weight: 1.8, Tags: []string{"formal"}},
		{Value: "Empresa", Weight: 1.5, Tags: []string{"formal"}},
	},
}

var dsPt_4 = &DataSet{
	Language: "pt",
	Country:  "BR",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Maria", Weight: 2.0, Tags: []string{"biblical", "traditional"}},
		{Value: "Ana", Weight: 1.8, Tags: []string{"biblical", "grace"}},
		{Value: "Isabel", Weight: 1.6, Tags: []string{"biblical", "noble"}},
		{Value: "Francisca", Weight: 1.4, Tags: []string{"traditional", "french"}},
		{Value: "Teresa", Weight: 1.2, Tags: []string{"traditional", "harvest"}},
		{Value: "Catarina", Weight: 1.1, Tags: []string{"traditional", "pure"}},
		{Value: "Margarida", Tags: []string{"traditional", "pearl"}},
		{Value: "Joana", Weight: 0.9, Tags: []string{"biblical", "gracious"}},
		{Value: "Beatriz", Weight: 0.8, Tags: []string{"traditional", "blessed"}},
		{Value: "Carmen", Weight: 0.7, Tags: []string{"traditional", "song"}},
	},
}

var dsPt_5 = &DataSet{
	Language: "pt",
	Country:  "BR",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "João", Weight: 2.0, Tags: []string{"biblical", "traditional"}},
		{Value: "José", Weight: 1.8, Tags: []string{"biblical", "traditional"}},
		{Value: "António", Weight: 1.6, Tags: []string{"traditional", "roman"}},
		{Value: "Manuel", Weight: 1.4, Tags: []string{"biblical", "traditional"}},
		{Value: "Carlos", Weight: 1.2, Tags: []string{"traditional", "strong"}},
		{Value: "Francisco", Weight: 1.1, Tags: []string{"religious", "traditional"}},
		{Value: "Luis", Tags: []string{"traditional", "warrior"}},
		{Value: "Miguel", Weight: 0.9, Tags: []string{"biblical", "archangel"}},
		{Value: "Pedro", Weight: 0.8, Tags: []string{"biblical", "rock"}},
		{Value: "Ricardo", Weight: 0.7, Tags: []string{"traditional", "ruler"}},
	},
}

var dsPt_6 = &DataSet{
	Language: "pt",
	Country:  "BR",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Silva", Weight: 2.0, Tags: []string{"geographical"}},
		{Value: "Santos", Weight: 1.8, Tags: []string{"religious"}},
		{Value: "Ferreira", Weight: 1.6, Tags: []string{"occupational"}},
		{Value: "Pereira", Weight: 1.4, Tags: []string{"geographical"}},
		{Value: "Oliveira", Weight: 1.2, Tags: []string{"geographical"}},
		{Value: "Sousa", Weight: 1.1, Tags: []string{"geographical"}},
		{Value: "Costa", Tags: []string{"geographical"}},
		{Value: "Rodrigues", Weight: 0.9, Tags: []string{"patronymic"}},
		{Value: "Martins", Weight: 0.8, Tags: []string{"patronymic"}},
		{Value: "Jesus", Weight: 0.7, Tags: []string{"religious"}},
	},
}

var dsPt_7 = &DataSet{
	Language: "pt",
	Type:     "texts",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "texto"},
		{Value: "palavra", Weight: 0.95},
		{Value: "frase", Weight: 0.9},
		{Value: "parágrafo", Weight: 0.85},
		{Value: "artigo", Weight: 0.8},
		{Value: "relatório", Weight: 0.75},
		{Value: "estudo", Weight: 0.7},
		{Value: "análise", Weight: 0.6499999999999999},
		{Value: "conteúdo", Weight: 0.6},
		{Value: "assunto", Weight: 0.55},
		{Value: "conceito", Weight: 0.5},
		{Value: "significado", Weight: 0.44999999999999996},
		{Value: "idioma", Weight: 0.3999999999999999},
		{Value: "literatura", Weight: 0.35},
		{Value: "escrita", Weight: 0.29999999999999993},
	},
}
