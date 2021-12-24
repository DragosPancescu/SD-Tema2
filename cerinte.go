package main

import (
	"fmt"
	"sync"
)

func print_nice_input(input [][]string) {
	fmt.Println("\nInput: ")
	for _, arr := range input {
		fmt.Println(arr)
	}
}

/*
Cerinta 1.
Se dă un vector de vectori, care conțin mai multe cuvinte ce pot avea in componența lor doar litere.
Să se afle numărul mediu de cuvinte care conțin numar par de vocale si un numar divizibil cu 3 de consoane folosind tehnica map-reduce.
*/

// Returneaza numarul de vocale si consoane dintr-un string.
func get_nr_vocale_consoane(input string) (int, int) {
	nr_vocale := 0
	nr_consoane := 0

	for _, litera := range input {
		switch litera {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			{
				nr_vocale += 1
			}
		default:
			{
				nr_consoane += 1
			}
		}
	}

	return nr_vocale, nr_consoane
}

func mapper1(intrare <-chan string, iesire chan<- map[string]int) {
	contorizare := map[string]int{}

	for cuvant := range intrare {

		// extragem numarul de vocale si consoane
		nr_vocale, nr_consoane := get_nr_vocale_consoane(cuvant)

		if nr_vocale%2 == 0 && nr_consoane%3 == 0 {
			contorizare["valid"] = contorizare["valid"] + 1
		}
	}

	iesire <- contorizare
	close(iesire)
}

func reducer1(intrare <-chan int, iesire chan<- float32) {
	suma, contor := 0, 0

	for n := range intrare {
		suma += n
		contor++
	}

	iesire <- float32(suma) / float32(contor)
	close(iesire)
}

func citire_input1(iesire [3]chan<- string) {
	intrare := [][]string{
		{"aabbb", "ebep", "blablablaa", "hijk", "wsww"},            //1
		{"abba", "eeeppp", "cocor", "ppppppaa", "qwerty", "acasq"}, //2
		{"lalala", "lalal", "papapa", "papap"},                     //3
	}

	print_nice_input(intrare)

	for i := range iesire {
		go func(caracter chan<- string, cuvinte []string) {
			for _, c := range cuvinte {
				caracter <- c
			}
			close(caracter)
		}(iesire[i], intrare[i])
	}
}

func genereaza_amestecare1(intrare []<-chan map[string]int, iesire [1]chan<- int) {
	//procesul de sincronizare
	var sincronizare sync.WaitGroup

	sincronizare.Add(len(intrare))

	for _, caracter := range intrare {
		go func(c <-chan map[string]int) {
			for i := range c {
				contor_valid, ok := i["valid"]
				if ok {
					iesire[0] <- contor_valid
				}
			}
			sincronizare.Done()
		}(caracter)
	}
	go func() {
		sincronizare.Wait()
		close(iesire[0])
	}()
}

func scrie_medie1(intrare []<-chan float32) {
	var sincronizare sync.WaitGroup
	sincronizare.Add(len(intrare))

	for i := 0; i < len(intrare); i++ {
		go func(pozitieComponent int, caracter <-chan float32) {
			for medie := range caracter {
				fmt.Printf("Numărul mediu de cuvinte care conțin numar par de vocale si un numar divizibil cu 3 de consoane este: %f\n", medie)
			}
			sincronizare.Done()
		}(i, intrare[i])
	}
	sincronizare.Wait()
}

func cerinta1() {
	dimensiune := 12

	componentaText1 := make(chan string, dimensiune)
	componentaText2 := make(chan string, dimensiune)
	componentaText3 := make(chan string, dimensiune)

	componentaMap1 := make(chan map[string]int, dimensiune)
	componentaMap2 := make(chan map[string]int, dimensiune)
	componentaMap3 := make(chan map[string]int, dimensiune)

	componentReduce1 := make(chan int, dimensiune)

	componentaMedie1 := make(chan float32, dimensiune)

	go citire_input1([3]chan<- string{componentaText1, componentaText2, componentaText3})

	go mapper1(componentaText1, componentaMap1)
	go mapper1(componentaText2, componentaMap2)
	go mapper1(componentaText3, componentaMap3)

	go genereaza_amestecare1([]<-chan map[string]int{componentaMap1, componentaMap2, componentaMap3}, [1]chan<- int{componentReduce1})

	go reducer1(componentReduce1, componentaMedie1)

	scrie_medie1([]<-chan float32{componentaMedie1})
}

/*
Cerinta 2.
Se dă un vector de vectori, care conțin mai multe cuvinte ce pot avea in componența lor atȃt litere cȃt și cifre.
Să se afle numărul mediu de cuvinte care sunt palindrom folosind tehnica map-reduce.
*/

func is_palindrom(input string) bool {

	for_size := 0
	if len(input)%2 == 0 {
		for_size = len(input) / 2
	} else {
		for_size = (len(input) - 1) / 2
	}

	for i := 0; i < for_size; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}
	return true
}

func mapper2(intrare <-chan string, iesire chan<- map[string]int) {
	contorizare := map[string]int{}

	for cuvant := range intrare {

		is_palindrom := is_palindrom(cuvant)

		if is_palindrom {
			contorizare["valid"] = contorizare["valid"] + 1
		}
	}

	iesire <- contorizare
	close(iesire)
}

func citire_input2(iesire [3]chan<- string) {
	intrare := [][]string{
		{"a1551a", "parc", "ana", "minim", "1pcl3"},                // 1
		{"calabalac", "tivit", "leu", "zece10", "ploaie", "9ana9"}, // 2
		{"lalalal", "tema", "papa", "ger"},                         // 3
	}

	print_nice_input(intrare)

	for i := range iesire {
		go func(caracter chan<- string, cuvinte []string) {
			for _, c := range cuvinte {
				caracter <- c
			}
			close(caracter)
		}(iesire[i], intrare[i])
	}
}

func scrie_medie2(intrare []<-chan float32) {
	var sincronizare sync.WaitGroup
	sincronizare.Add(len(intrare))

	for i := 0; i < len(intrare); i++ {
		go func(pozitieComponent int, caracter <-chan float32) {
			for medie := range caracter {
				fmt.Printf("Numărul mediu de palindroame este: %f\n", medie)
			}
			sincronizare.Done()
		}(i, intrare[i])
	}
	sincronizare.Wait()
}

func cerinta2() {
	dimensiune := 12

	componentaText1 := make(chan string, dimensiune)
	componentaText2 := make(chan string, dimensiune)
	componentaText3 := make(chan string, dimensiune)

	componentaMap1 := make(chan map[string]int, dimensiune)
	componentaMap2 := make(chan map[string]int, dimensiune)
	componentaMap3 := make(chan map[string]int, dimensiune)

	componentReduce1 := make(chan int, dimensiune)

	componentaMedie1 := make(chan float32, dimensiune)

	go citire_input2([3]chan<- string{componentaText1, componentaText2, componentaText3})

	go mapper2(componentaText1, componentaMap1)
	go mapper2(componentaText2, componentaMap2)
	go mapper2(componentaText3, componentaMap3)

	go genereaza_amestecare1([]<-chan map[string]int{componentaMap1, componentaMap2, componentaMap3}, [1]chan<- int{componentReduce1})

	go reducer1(componentReduce1, componentaMedie1)

	scrie_medie2([]<-chan float32{componentaMedie1})
}
