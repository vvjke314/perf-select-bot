package chooser

import (
	"sort"

	"github.com/vvjke314/MPPR/lab1/internal/model"
)

func Choose(progress model.Progress) string {
	var results = map[string]int{
		"Dior Sauvage от Dior":              0,
		"Bleu de Chanel от Chanel":          0,
		"Acqua Di Gio от Giorgio Armani":    0,
		"Terre d'Hermes от Hermes":          0,
		"Le Male от Jean Paul Gaultier":     0,
		"Chanel No. 5 от Chanel":            0,
		"La Vie Est Belle от Lancôme":       0,
		"Black Opium от Yves Saint Laurent": 0,
		"Daisy от Marc Jacobs":              0,
		"Alien от Thierry Mugler":           0,
	}

	if progress.Q1Answer > 10000 {
		results["Dior Sauvage от Dior"] += 1
		results["Acqua Di Gio от Giorgio Armani"] += 1
		results["Terre d'Hermes от Hermes"] += 1
		results["Black Opium от Yves Saint Laurent"] += 1
		results["Daisy от Marc Jacobs"] += 1
	}

	if progress.Q2Answer == 2 {
		results["Dior Sauvage от Dior"] += 1
		results["Acqua Di Gio от Giorgio Armani"] += 1
		results["Terre d'Hermes от Hermes"] += 1
		results["Chanel No. 5 от Chanel"] += 1
		results["Alien от Thierry Mugler"] += 1
	}

	if progress.Q5Answer == 1 || progress.Q7Answer == 2 {
		results["Chanel No. 5 от Chanel"] += 1
		results["La Vie Est Belle от Lancôme"] += 1
		results["Black Opium от Yves Saint Laurent"] += 1
		results["Daisy от Marc Jacobs"] += 1
		results["Alien от Thierry Mugler"] += 1
	} else {
		results["Dior Sauvage от Dior"] += 1
		results["Bleu de Chanel от Chanel"] += 1
		results["Acqua Di Gio от Giorgio Armani"] += 1
		results["Terre d'Hermes от Hermes"] += 1
		results["Le Male от Jean Paul Gaultier"] += 1
	}

	if progress.Q8Answer == 3 {
		results["Dior Sauvage от Dior"] += 1
		results["Bleu de Chanel от Chanel"] += 1
		results["Acqua Di Gio от Giorgio Armani"] += 1
		results["Black Opium от Yves Saint Laurent"] += 1
		results["Daisy от Marc Jacobs"] += 1
	} else if progress.Q8Answer == 2 {
		results["Terre d'Hermes от Hermes"] += 1
		results["Le Male от Jean Paul Gaultier"] += 1
	} else {
		results["Chanel No. 5 от Chanel"] += 1
		results["La Vie Est Belle от Lancôme"] += 1
	}

	if progress.Q19Answer == 1 {
		results["Dior Sauvage от Dior"] += 1
		results["Chanel No. 5 от Chanel"] += 1
		results["La Vie Est Belle от Lancôme"] += 1
		results["Daisy от Marc Jacobs"] += 1
		results["Terre d'Hermes от Hermes"] += 1
	}

	srtResults := sortMap(results)

	return srtResults[0].Key
}

func sortMap(results map[string]int) PairList {
	pl := make(PairList, len(results))
	i := 0
	for k, v := range results {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
