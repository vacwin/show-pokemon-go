# show-pokemon-go

A CLI tool to print pokemon sprites in your terminal. Fork of
[pokemon-colorscripts-go](https://github.com/ollyjarvis/pokemon-colorscripts-go),
rewritten to use only Go standard library (no external dependencies).

## Requirements

A terminal with true color support. Most modern terminals have this.
More info in [this gist](https://gist.github.com/XVilka/8346728).

## Installation

Requires Go installed and in PATH.

```
git clone https://github.com/vacwin/show-pokemon-go.git
cd show-pokemon-go
make install
```

This will:
- Build the `showpkm` binary
- Install it to `~/.local/bin/`
- Copy pokemon data to `~/.local/share/showpkm/`

Make sure `~/.local/bin` is in your PATH:

```
export PATH="$HOME/.local/bin:$PATH"
```

### Uninstall

```
make uninstall
```

## Usage

```
Usage: showpkm [options]

Options:
  -a          Print list of all pokemon
  -n <name>   Select pokemon by name
              Exceptions: nidoran-f, nidoran-m, mr-mime, farfetchd, type-null
  -f <form>   Show an alternative form of a pokemon
  -s          Show the shiny version
  -l          Show a larger sprite
  -r <gen>    Generation(s) to randomly select from (default: 1-8)
              Range: 2-6 | List: 1,3,6 | Number: 5
```

### Examples

Random pokemon:
```
showpkm
```

Specific pokemon:
```
showpkm -n charizard
```

Shiny version:
```
showpkm -n spheal -s
```

Large sprite:
```
showpkm -n spheal -l
```

Random from generation 1:
```
showpkm -r 1
```

Random from generations 1-3:
```
showpkm -r 1-3
```

Random from generations 1, 3 and 6:
```
showpkm -r 1,3,6
```

Alternative form:
```
showpkm -n deoxys -f defense
```

### Running on terminal startup

Add `showpkm` to your `.bashrc` or `.zshrc`. For fish, override `fish_greeting`:

```fish
function fish_greeting
    showpkm
end
```

## How it works

The binary prints text files containing unicode characters with ANSI color codes.
Sprites are stored in `~/.local/share/showpkm/colorscripts/` with variants for
regular/shiny and small/large sizes. Pokemon metadata lives in `pokemon.json`.

Sprites were generated from [PokéSprite](https://msikma.github.io/pokesprite/)
using [these scripts](https://gitlab.com/phoneybadger/pokemon-generator-scripts).

## Similar projects

- [pokemon-colorscripts](https://gitlab.com/phoneybadger/pokemon-colorscripts) (original, Python)
- [pokemon-colorscripts-go](https://github.com/ollyjarvis/pokemon-colorscripts-go) (upstream of this fork)
- [pokeget](https://github.com/talwat/pokeget)
- [pokeshell](https://github.com/acxz/pokeshell)
- [krabby](https://github.com/yannjor/krabby)

## Credits

- Pokemon designs, names, branding: [The Pokémon Company](https://pokemon.com)
- Box art sprites: [PokéSprite](https://msikma.github.io/pokesprite/)
- Original Go rewrite: [ollyjarvis](https://github.com/ollyjarvis)
- Original project: [phoneybadger](https://gitlab.com/phoneybadger)

## License

MIT
