package mapas

import "fmt"

func ShowMaps() {
	countries := make(map[string]string)
	countries["México:D.F"] = "D.F"
	countries["Argentina"] = "Buenos Aires"
	countries["Venezuela"] = "Caracas"
	countries["Argentina"] = "Buenos Días"

	fmt.Println(countries)
	fmt.Println(countries["México"])

	campeonatos := map[string]int{
		"Barcelona":    39,
		"Chivas":       32,
		"Real Madrid":  36,
		"Boca Juniors": 37,
	}

	for key, score := range campeonatos {
		fmt.Printf("Equipo %s tiene un puntaje de %d \n", key, score)
	}

	fmt.Println(":::::::Delete:::::")
	delete(campeonatos, "Chivas")
	fmt.Println(campeonatos)

	puntajeChiva, existChiva := campeonatos["Chivas"]
	fmt.Printf("Puntaje %d, existe %t \n", puntajeChiva, existChiva)

	puntajeBarcelona, existBarcelona := campeonatos["Barcelona"]
	fmt.Printf("Puntaje %d, existe %t \n", puntajeBarcelona, existBarcelona)

}
