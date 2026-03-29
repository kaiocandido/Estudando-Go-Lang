//Goblinocus é um país que leva muito a sério a previsão do tempo. Como você é um desenvolvedor renomado, responsável e competente, eles pediram que você escrevesse um programa que pudesse prever as condições climáticas atuais de várias cidades de Goblinocus. Você estava ocupado na época e pediu a um de seus amigos para fazer o trabalho. Depois de um tempo, o presidente da Goblinocus entrou em contato com você e disse que não entendia o código do seu amigo. Ao verificar o código, você descobre que seu amigo não agiu como um programador responsável e não há comentários no código. Você se sente obrigado a esclarecer o programa para que os goblins também possam entendê-lo.

// Package weather provides functionality to describe current weather conditions for a given location.
package weather

var (
	// CurrentCondition stores the current weather condition (e.g., sunny, rainy).
	CurrentCondition string

	// CurrentLocation stores the name of the location for which the weather is reported.
	CurrentLocation string
)

// Forecast returns a formatted string describing the current weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
