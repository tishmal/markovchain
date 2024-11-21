package logic

import "fmt"

func Help() {
	fmt.Println(`Markov Chain text generator.

Usage:
  markovchain [-w <N>] [-p <S>] [-l <N>]
  markovchain --help

Options:
  --help  Show this screen.
  -w N    Number of maximum words (default: 100, max: 10000)
  -p S    Starting prefix
  -l N    Prefix length (default: 2, max: 5)`)
}
