package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

type Pokemon struct {
	Idx int `json:"idx"`
	Slug string `json:"slug"`
	Gen int `json:"gen"`
	Name map[string]string `json:"name"`
	Forms []string `json:"forms"`
}

var (
	PROGRAM_DIR string
	GENERATIONS = [][]int{{1, 151}, {152, 251}, {252, 386}, {387, 493}, {494, 649}, {650, 721}, {722, 809}, {810, 898}}
)

const (
	SHINY_RATE = 1.0 / 128.0
)

func print_file(filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	fmt.Print(string(data))

	return nil
}

func loadPokemons() ([]Pokemon, error) {
	data, err := os.ReadFile(fmt.Sprintf("%s/pokemon.json", PROGRAM_DIR))
	if err != nil {
		return nil, err
	}
	var pokemons []Pokemon
	err = json.Unmarshal(data, &pokemons)
	return pokemons, err
}

func list_pokemon_names() error {
	pokemons, err := loadPokemons()
	if err != nil {
		return err
	}
	for _, p := range pokemons {
		fmt.Println(p.Name["en"])
	}
	return nil
}

func show_pokemon_by_name(name string, shiny bool, is_large bool, form string) error {
	dir := fmt.Sprintf("%s/colorscripts", PROGRAM_DIR)

	if is_large {
		dir = fmt.Sprintf("%s/large", dir)
	} else {
		dir = fmt.Sprintf("%s/small", dir)
	}

	if shiny {
		dir = fmt.Sprintf("%s/shiny", dir)
	} else {
		dir = fmt.Sprintf("%s/regular", dir)
	}

	if form != "" {
		pokemons, err := loadPokemons()
		if err != nil {
			return err
		}

		var alt_forms []string
		found := false

		for _, p := range pokemons {
			if strings.EqualFold(p.Slug, name) {
				for _, f := range p.Forms {
					if f != "regular" {
						if f == form {
							found = true
						}
						alt_forms = append(alt_forms, f)
					}
				}
				break
			}
		}

		if found {
			name = fmt.Sprintf("%s-%s", name, form)
		} else {
			fmt.Printf("invalid form '%s' for pokemon %s\n", form, name)
			if len(alt_forms) == 0 {
				fmt.Printf("no alternative forms for %s\n", name)
			} else {
				fmt.Printf("available alternative forms are")
				for _, f := range alt_forms {
					fmt.Printf("- %s\n", f)
				}
			}
			os.Exit(0)
		}
	}

	file := fmt.Sprintf("%s/%s", dir, name)
	return print_file(file)
}

func show_random_pokemon(generations string, shiny bool, is_large bool) error {
    var gen int

    if strings.Contains(generations, ",") {
        parts := strings.Split(generations, ",")
        picked := parts[rand.IntN(len(parts))]
        g, err := strconv.Atoi(strings.TrimSpace(picked))
        if err != nil {
            return fmt.Errorf("invalid generation '%s'", picked)
        }
        gen = g
    } else if strings.Contains(generations, "-") {
        parts := strings.Split(generations, "-")
        lo, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
        hi, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
        if err1 != nil || err2 != nil || lo > hi {
            return fmt.Errorf("invalid generation range '%s'", generations)
        }
        gen = rand.IntN(hi-lo+1) + lo
    } else if generations != "" {
        g, err := strconv.Atoi(generations)
        if err != nil {
            return fmt.Errorf("invalid generation '%s'", generations)
        }
        gen = g
    } else {
        gen = rand.IntN(8) + 1
    }

    if gen < 1 || gen > 8 {
        return fmt.Errorf("generation %d out of range (1-8)", gen)
    }

    pokemons, err := loadPokemons()
    if err != nil {
        return err
    }

    lo := GENERATIONS[gen-1][0]
    hi := GENERATIONS[gen-1][1]
    idx := rand.IntN(hi-lo) + lo
    name := pokemons[idx].Slug

    if !shiny {
        shiny = rand.Float32() <= SHINY_RATE
    }

    return show_pokemon_by_name(name, shiny, is_large, "")
}

func main() {
	var (
		all    = flag.Bool("a", false, "Print list of all pokemon")
		name   = flag.String("n", "", "Select pokemon by name")
		form   = flag.String("f", "", "Show an alternative form of a pokemon")
		shiny  = flag.Bool("s", false, "Show the shiny version of the pokemon")
		large  = flag.Bool("l", false, "Show a larger version of the sprite")
		random = flag.String("r", "1-8", "Generation(s) to randomly select from (e.g. 2-6, 1,3,6, 5)")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: pokemon-colorscripts-go [options]\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, "  -a          Print list of all pokemon\n")
		fmt.Fprintf(os.Stderr, "  -n <name>   Select pokemon by name\n")
		fmt.Fprintf(os.Stderr, "              Exceptions: nidoran-f, nidoran-m, mr-mime, farfetchd, type-null\n")
		fmt.Fprintf(os.Stderr, "  -f <form>   Show an alternative form of a pokemon\n")
		fmt.Fprintf(os.Stderr, "  -s          Show the shiny version\n")
		fmt.Fprintf(os.Stderr, "  -l          Show a larger sprite\n")
		fmt.Fprintf(os.Stderr, "  -r <gen>    Generation(s) to randomly select from (default: 1-8)\n")
		fmt.Fprintf(os.Stderr, "              Range: 2-6 | List: 1,3,6 | Number: 5\n")
	}

	flag.Parse()

	var err error
	if *all {
		err = list_pokemon_names()
	} else if *name != "" {
		err = show_pokemon_by_name(*name, *shiny, *large, *form)
	} else {
		err = show_random_pokemon(*random, *shiny, *large)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
