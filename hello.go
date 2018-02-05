package main

import (
	"fmt"
	"strconv"
	"os"
	"sort"
	"log"
)

var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

type Vecteur3D [3]float32
type Intervalle struct {
	debut int
	fin   int
}

type Sequence []int

func (s Sequence) Len() int {
	// longueur de s
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	// retourne vrai si l'élément d'indice i est inférieur à l'élément d'indice j
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	// échange deux éléments
	s[i], s[j] = s[j], s[i]
}

/* Un commentaire sur
   plusieurs lignes */
func main_hello() {
	// Un commentaire sur une ligne
	fmt.Printf("Bonjour, monde!\n")

	var james = "James"
	var nom = "Bond"
	var nomComplet = james + " " + nom
	fmt.Println(nomComplet)

	var tab = [...] float32{0.1, 0.2, 0.3, 0.5, 0.8}
	var size = len(tab)
	fmt.Println("Size of tab : " + strconv.Itoa(size))
	//x:= Somme(tab)
	//fmt.Printf(x)

	//var tranche = tab[0:3]
	//var autre = tab[:]
	var fibonacci = []int{1, 1, 2, 3, 5, 8}
	fmt.Println(strconv.Itoa(Somme(fibonacci)))

	x := []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	fmt.Println(x)

	//var tranche = make([]int,10,100)
	const Pi = 3.14159
	fmt.Println(Pi)
	fmt.Println(HOME)

	var vec Vecteur3D
	vec[0] = 0
	vec[1] = 0
	vec[2] = 0
	fmt.Println(vec)

	intervalle := Intervalle{0, 3}
	inter2 := Intervalle{fin: 5}
	fmt.Println(inter2)
	fmt.Println(intervalle)

	inter3 := new(Intervalle)
	inter4 := &Intervalle{0, 5}
	inter3.fin += inter3.debut
	fmt.Println(inter3)
	fmt.Println(inter4)

	var Constantes = map[string]float32{
		"Pi": 3.14159,
		"e":  2.71,
		"g":  9.81,
	}
	fmt.Println(Constantes)
	valeur, ok := Constantes["g"]
	fmt.Println(valeur)
	fmt.Println(ok)
	//Constantes["kB"] = 0, false
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(Min(1, 3, 2, 0))

	tab5 := []int{1, 3, 2, 5}
	fmt.Println(Min(tab5...))

	somme, moyenne := SommeEtMoyenne([] float32{0.1, 0.2, 0.3, 0.5, 0.8})
	fmt.Printf("%g, %g\n", somme, moyenne)

	var seq Sequence
	sort.Sort(seq)

	defer func() {
		if err := recover(); err != nil {
			log.Println("tache echouee:", err)
		}
	}()
	panic("pas de valeur pour $USER")

}

func Somme(a []int) (somme int) {
	for _, v := range a {
		somme += v
	}
	return
}
func Min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func SommeEtMoyenne(a []float32) (float32, float32) {
	var somme float32 = 0.0
	for _, x := range a {
		somme += x
	}
	return somme, somme / float32(len(a))
}
