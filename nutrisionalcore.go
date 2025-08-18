package main

type ScoreType int

const (
	FoodScore ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutrisionalScore struct {
	Value int
	Positive int
	Negative int
	ScoreType ScoreType
}

type EnergyKJ float64

type SugarsGram float64

type SaturatedFattyAcidsGram float64

type SodiumMilligram float64

type FruitsPercent float64

type FibreGram float64

type ProteinGram float64

type NutrisionalData struct {
	Energy EnergyKJ
	Sugars SugarsGram
	SaturatedFattyAcids SaturatedFattyAcidsGram
	Sodium SodiumMilligram
	Fruits FruitsPercent
	Fibre FibreGram
	Protein ProteinGram
	IsWater bool
}

func GetPoints(e EnergyKJ) GetPoints(st ScoreType) int{

}

func GetPoints(s SugarsGram) GetPoints(st ScoreType) int{

}

func GetPoints(sfa SaturatedFattyAcidsGram) GetPoints(st ScoreType) int{
}

func GetNutrisionalSco(n NutrisionalData, st ScoreType) NutrisionalScore{

	value := 0
	positive := 0
	negative := 0

	if st != Water {
		fruitPoints := n.Fruits.GetPoints(st)
		fibrePoints := n.Fibre.GetPoints(st)

		negative = n.Energy.GetPoints(st) + n.Sugars.GetPoints(st) + n.SaturatedFattyAcids.GetPoints(st) + n.Sodium.GetPoints(st)
		positive = fruitPoints + fibrePoints + n.Protein.GetPoints(st)
	}

	return NutrisionalScore{
		Value: value,
		Positive: positive,
		Negative: negative,
		ScoreType: st,
	}
}
