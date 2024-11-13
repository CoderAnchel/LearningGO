package models

type Eurofighter struct {
	General          GeneralInfo
	Especificaciones Especificaciones
	Motorizacion     Motorizacion
	Armamento        Armamento
	Usuarios         []string
}

type GeneralInfo struct {
	Fabricantes         []string
	PaisesInvolucrados  []string
	PrimeraIntroduccion int
	Uso                 string
}

type Especificaciones struct {
	Tripulacion        int
	Longitud           float32 // en metros
	Envergadura        float32 // en metros
	Altura             float32 // en metros
	PesoEnVacio        int     // en kg
	PesoMaximoDespegue int     // en kg
}

type Motorizacion struct {
	NumeroDeMotores int
	TipoDeMotores   string
	EmpujeMaximo    int     // en libras
	VelocidadMaxima float32 // en Mach
	TechoDeServicio int     // en metros
	Alcance         int     // en km
}

type Armamento struct {
	AireAire   []string
	AireTierra []string
	Cañon      Cañon
}

type Cañon struct {
	Modelo          string
	Calibre         float32 // en mm
	CapacidadCargas int
}

func EuroData() Eurofighter {
	euroFighter := Eurofighter{
		General: GeneralInfo{
			Fabricantes:         []string{"Airbus", "BAE Systems", "Leonardo"},
			PaisesInvolucrados:  []string{"Reino Unido", "Alemania", "Italia", "España"},
			PrimeraIntroduccion: 2003,
			Uso:                 "Multirrol",
		},
		Especificaciones: Especificaciones{
			Tripulacion:        1,
			Longitud:           15.96,
			Envergadura:        10.95,
			Altura:             5.28,
			PesoEnVacio:        11000,
			PesoMaximoDespegue: 23500,
		},
		Motorizacion: Motorizacion{
			NumeroDeMotores: 2,
			TipoDeMotores:   "Eurojet EJ200 turbofan",
			EmpujeMaximo:    20000,
			VelocidadMaxima: 2.0,
			TechoDeServicio: 16800,
			Alcance:         2900,
		},
		Armamento: Armamento{
			AireAire:   []string{"AIM-120 AMRAAM", "Meteor", "AIM-9 Sidewinder", "IRIS-T"},
			AireTierra: []string{"Paveway", "JDAM", "Brimstone", "Storm Shadow"},
			Cañon: Cañon{
				Modelo:          "Mauser BK-27",
				Calibre:         27.0,
				CapacidadCargas: 150,
			},
		},
		Usuarios: []string{"Reino Unido", "Alemania", "Italia", "España", "Austria", "Arabia Saudita", "Qatar", "Kuwait"},
	}
	return euroFighter
}
