package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutrisionalScore struct {
	Value     int
	Positive  int
	Negative  int
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
	Energy              EnergyKJ
	Sugars              SugarsGram
	SaturatedFattyAcids SaturatedFattyAcidsGram
	Sodium              SodiumMilligram
	Fruits              FruitsPercent
	Fibre               FibreGram
	Protein             ProteinGram
	IsWater             bool
}

// ---------- Helpers Input ----------
var in = bufio.NewReader(os.Stdin)

func inputFloat(prompt string) float64 {
	for {
		fmt.Print(prompt)
		line, _ := in.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			return 0
		}
		v, err := strconv.ParseFloat(strings.ReplaceAll(line, ",", "."), 64)
		if err == nil {
			return v
		}
		fmt.Println("  -> Angka tidak valid, coba lagi.")
	}
}

func inputInt(prompt string, min, max int) int {
	for {
		fmt.Print(prompt)
		line, _ := in.ReadString('\n')
		line = strings.TrimSpace(line)
		n, err := strconv.Atoi(line)
		if err == nil && n >= min && n <= max {
			return n
		}
		fmt.Printf("  -> Masukkan angka %d..%d\n", min, max)
	}
}

// “Constructor” input (hindari bentrok nama dengan type)
func EnergyFromKcalInput() EnergyKJ {
	kcal := inputFloat("Energi per 100g/ml (kcal): ")
	return EnergyKJ(kcal * 4.184) // 1 kcal = 4.184 kJ
}
func InputSugarsGram() SugarsGram {
	return SugarsGram(inputFloat("Gula per 100g/ml (gram): "))
}
func InputSaturatedFattyAcidsGram() SaturatedFattyAcidsGram {
	return SaturatedFattyAcidsGram(inputFloat("Lemak jenuh per 100g/ml (gram): "))
}
func InputSodiumMilligram() SodiumMilligram {
	return SodiumMilligram(inputFloat("Natrium per 100g/ml (mg): "))
}
func InputFruitsPercent() FruitsPercent {
	v := inputFloat("Persen buah/kacang/sayur/legum (%) (0-100): ")
	if v < 0 {
		v = 0
	}
	if v > 100 {
		v = 100
	}
	return FruitsPercent(v)
}
func InputFibreGram() FibreGram {
	return FibreGram(inputFloat("Serat per 100g/ml (gram): "))
}
func InputProteinGram() ProteinGram {
	return ProteinGram(inputFloat("Protein per 100g/ml (gram): "))
}

// ---------- Skor (contoh tabel mirip Nutri-Score klasik) ----------
func (e EnergyKJ) GetPoints(st ScoreType) int {
	if st == Water {
		return 0
	}
	switch {
	case e <= 335:
		return 0
	case e <= 670:
		return 1
	case e <= 1005:
		return 2
	case e <= 1340:
		return 3
	case e <= 1675:
		return 4
	case e <= 2010:
		return 5
	case e <= 2345:
		return 6
	case e <= 2680:
		return 7
	case e <= 3015:
		return 8
	case e <= 3350:
		return 9
	default:
		return 10
	}
}

func (s SugarsGram) GetPoints(st ScoreType) int {
	if st == Water {
		return 0
	}
	switch {
	case s <= 4.5:
		return 0
	case s <= 9:
		return 1
	case s <= 13.5:
		return 2
	case s <= 18:
		return 3
	case s <= 22.5:
		return 4
	case s <= 27:
		return 5
	case s <= 31:
		return 6
	case s <= 36:
		return 7
	case s <= 40:
		return 8
	case s <= 45:
		return 9
	default:
		return 10
	}
}

func (sfa SaturatedFattyAcidsGram) GetPoints(st ScoreType) int {
	switch {
	case sfa <= 1:
		return 0
	case sfa <= 2:
		return 1
	case sfa <= 3:
		return 2
	case sfa <= 4:
		return 3
	case sfa <= 5:
		return 4
	case sfa <= 6:
		return 5
	case sfa <= 7:
		return 6
	case sfa <= 8:
		return 7
	case sfa <= 9:
		return 8
	case sfa <= 10:
		return 9
	default:
		return 10
	}
}

func (s SodiumMilligram) GetPoints(st ScoreType) int {
	switch {
	case s <= 90:
		return 0
	case s <= 180:
		return 1
	case s <= 270:
		return 2
	case s <= 360:
		return 3
	case s <= 450:
		return 4
	case s <= 540:
		return 5
	case s <= 630:
		return 6
	case s <= 720:
		return 7
	case s <= 810:
		return 8
	case s <= 900:
		return 9
	default:
		return 10
	}
}

func (f FruitsPercent) GetPoints(st ScoreType) int {
	switch {
	case f < 40:
		return 0
	case f < 60:
		return 1
	case f < 80:
		return 2
	default: // >= 80%
		return 5
	}
}

func (f FibreGram) GetPoints(st ScoreType) int {
	switch {
	case f < 0.9:
		return 0
	case f < 1.9:
		return 1
	case f < 2.8:
		return 2
	case f < 3.7:
		return 3
	case f < 4.7:
		return 4
	default:
		return 5
	}
}

func (p ProteinGram) GetPoints(st ScoreType) int {
	switch {
	case p < 1.6:
		return 0
	case p < 3.2:
		return 1
	case p < 4.8:
		return 2
	case p < 6.4:
		return 3
	default:
		return 5
	}
}

// ---------- Hitung Nilai ----------
func GetNutrisionalSco(n NutrisionalData, st ScoreType) NutrisionalScore {
	if st == Water {
		return NutrisionalScore{
			Value:     0,
			Positive:  0,
			Negative:  0,
			ScoreType: st,
		}
	}

	fruitPoints := n.Fruits.GetPoints(st)
	fibrePoints := n.Fibre.GetPoints(st)

	negative := n.Energy.GetPoints(st) +
		n.Sugars.GetPoints(st) +
		n.SaturatedFattyAcids.GetPoints(st) +
		n.Sodium.GetPoints(st)

	positive := fruitPoints + fibrePoints + n.Protein.GetPoints(st)

	value := negative - positive

	return NutrisionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: st,
	}
}

func chooseScoreType() ScoreType {
	fmt.Println("Pilih tipe produk:")
	fmt.Println("  1) Food")
	fmt.Println("  2) Beverage")
	fmt.Println("  3) Water")
	fmt.Println("  4) Cheese")
	choice := inputInt("Masukkan pilihan [1-4]: ", 1, 4)
	switch choice {
	case 1:
		return Food
	case 2:
		return Beverage
	case 3:
		return Water
	default:
		return Cheese
	}
}

func main() {
	st := chooseScoreType()

	// Kumpulkan input
	data := NutrisionalData{
		Energy:              EnergyFromKcalInput(),
		Sugars:              InputSugarsGram(),
		SaturatedFattyAcids: InputSaturatedFattyAcidsGram(),
		Sodium:              InputSodiumMilligram(),
		Fruits:              InputFruitsPercent(),
		Fibre:               InputFibreGram(),
		Protein:             InputProteinGram(),
	}

	ns := GetNutrisionalSco(data, st)

	fmt.Printf("\n--- Hasil ---\n")
	fmt.Printf("Negative points: %d\n", ns.Negative)
	fmt.Printf("Positive points: %d\n", ns.Positive)
	fmt.Printf("Nutritional Score (neg - pos): %d\n", ns.Value)
}
