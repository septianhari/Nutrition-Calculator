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
