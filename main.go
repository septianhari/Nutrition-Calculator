package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type NutrisionalScore struct {
	Value    int
	Positive int
	Negative int
}

func calculateScore(energy, sugars, sfa, sodium, fruits, fibre, protein float64) NutrisionalScore {

	negative := 0
	if energy > 670 {
		negative++
	}
	if sugars > 10 {
		negative++
	}
	if sfa > 3 {
		negative++
	}
	if sodium > 300 {
		negative++
	}

	positive := 0
	if fruits > 60 {
		positive++
	}
	if fibre > 3 {
		positive++
	}
	if protein > 5 {
		positive++
	}

	value := negative - positive

	return NutrisionalScore{Value: value, Positive: positive, Negative: negative}
}

var tpl = template.Must(template.New("form").Parse(`
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>Nutriscore Calculator</title>
	<style>
		body { font-family: sans-serif; margin: 40px; }
		label { display: block; margin-top: 10px; }
	</style>
</head>
<body>
	<h2>Nutriscore Calculator</h2>
	<form method="POST" action="/">
		<label>Energy (kcal): <input type="number" name="energy" step="0.1" min="0"></label>
		<label>Sugars (g): <input type="number" name="sugars" step="0.1" min="0"></label>
		<label>Saturated Fat (g): <input type="number" name="sfa" step="0.1" min="0"></label>
		<label>Sodium (mg): <input type="number" name="sodium" step="0.1" min="0"></label>
		<label>Fruits (%): <input type="number" name="fruits" step="0.1" min="0"></label>
		<label>Fibre (g): <input type="number" name="fibre" step="0.1" min="0"></label>
		<label>Protein (g): <input type="number" name="protein" step="0.1" min="0"></label>
		<br>
		<button type="submit">Hitung</button>
	</form>
	{{if .}}
		<h3>Hasil:</h3>
		<p>Nilai Nutriscore: {{.Value}}</p>
		<p>Negative: {{.Negative}}</p>
		<p>Positive: {{.Positive}}</p>
		<p>Score: {{.Value}}</p>
	{{end}}
</body>
</html>
`))

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		energy, err := strconv.ParseFloat(r.FormValue("energy"), 64)
		if err != nil {
			http.Error(w, "Invalid energy value", http.StatusBadRequest)
			return
		}
		sugars, err := strconv.ParseFloat(r.FormValue("sugars"), 64)
		if err != nil {
			http.Error(w, "Invalid sugars value", http.StatusBadRequest)
			return
		}
		sfa, err := strconv.ParseFloat(r.FormValue("sfa"), 64)
		if err != nil {
			http.Error(w, "Invalid saturated fat value", http.StatusBadRequest)
			return
		}
		sodium, err := strconv.ParseFloat(r.FormValue("sodium"), 64)
		if err != nil {
			http.Error(w, "Invalid sodium value", http.StatusBadRequest)
			return
		}
		fruits, err := strconv.ParseFloat(r.FormValue("fruits"), 64)
		if err != nil {
			http.Error(w, "Invalid fruits value", http.StatusBadRequest)
			return
		}
		fibre, err := strconv.ParseFloat(r.FormValue("fibre"), 64)
		if err != nil {
			http.Error(w, "Invalid fibre value", http.StatusBadRequest)
			return
		}
		protein, err := strconv.ParseFloat(r.FormValue("protein"), 64)
		if err != nil {
			http.Error(w, "Invalid protein value", http.StatusBadRequest)
			return
		}

		// 🚨 Validation: reject negative values
		if energy < 0 || sugars < 0 || sfa < 0 || sodium < 0 || fruits < 0 || fibre < 0 || protein < 0 {
			http.Error(w, "Input tidak boleh negatif!", http.StatusBadRequest)
			return
		}

		score := calculateScore(energy, sugars, sfa, sodium, fruits, fibre, protein)
		tpl.Execute(w, score)
		return
	}
	tpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server jalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
