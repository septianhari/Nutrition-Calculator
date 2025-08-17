package main

import(
	"fmt"
)

func main() {
	
	ns :=GetNutrisionalSco(NutrisionalData{
		Energy: EnergyFromKcal(),
		Sugars: SugarsGram(),
		SaturatedFattyAcids: SaturatedFattyAcids(),
		Sodium: SodiumMilligram(),
		Fruits: FruitsPercent(), 
		Fibre: FibreGram(),
		Protein: ProteinGram(),
	}, Food)

		fmt.Printf("Nutritional Score: %d\n", ns.Value)
}