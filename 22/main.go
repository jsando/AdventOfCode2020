package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	players := parseInput(os.Args[1])
	winner := playCombatGame(players)
	fmt.Printf("Part 1: %d\n", winner.score())

	players = parseInput(os.Args[1])
	winner = playRecursiveCombatGame(players)
	fmt.Printf("Part 2: %d\n", winner.score())
}

func playCombatGame(players []*Player) *Player {
	round := 1
	for {
		winner := hasWinner(players)
		if winner != nil {
			return winner
		}
		fmt.Printf("-- Round %d --\n", round)
		playCombatRound(players)
		round++
	}
}

func playCombatRound(players []*Player) {
	for _, player := range players {
		fmt.Printf("%s's deck: %v\n", player.name, player.deck)
	}
	plays := []*Play{}
	for _, player := range players {
		play := player.getPlay()
		if play != nil {
			plays = append(plays, play)
			fmt.Printf("%s plays: %d\n", play.player.name, play.card)
		}
	}
	sort.Slice(plays, func(i, j int) bool {
		return plays[i].card > plays[j].card
	})
	winner := plays[0].player
	fmt.Printf("%s wins the round!\n\n", winner.name)
	for i := 0; i < len(plays); i++ {
		winner.addCard(plays[i].card)
	}
}

func playRecursiveCombatGame(players []*Player) *Player {
	round := 1
	cardsInRound := map[string]bool{}
	for {
		winner := hasWinner(players)
		if winner != nil {
			return winner
		}
		fmt.Printf("-- Round %d --\n", round)
		cardKey := ""
		for _, player := range players {
			cardKey += fmt.Sprintf("%s's deck: %v", player.name, player.deck)
		}
		if cardsInRound[cardKey] {
			// automatic win for player 1
			return players[0]
		}
		playRecursiveCombatRound(players)
		round++
		cardsInRound[cardKey] = true
	}
}

func playRecursiveCombatRound(players []*Player) {
	for _, player := range players {
		fmt.Printf("%s's deck: %v\n", player.name, player.deck)
	}
	recurse := true
	plays := []*Play{}
	for _, player := range players {
		play := player.getPlay()
		if play != nil {
			plays = append(plays, play)
			fmt.Printf("%s plays: %d\n", play.player.name, play.card)
			if play.card > len(play.player.deck) {
				recurse = false
			}
		}
	}
	recurse = recurse && len(plays) > 1
	if recurse {
		subPlayers := []*Player{}
		for _, play := range plays {
			subDeck := make([]int, play.card)
			copy(subDeck, play.player.deck[0:play.card])
			subPlayer := &Player{name: play.player.name, deck: subDeck}
			subPlayers = append(subPlayers, subPlayer)
		}
		subWinner := playRecursiveCombatGame(subPlayers)
		// haha, I assumed part2 would extend to more than 2 plays so I used a list ... but
		// now the order just depends on which one of the two players won.  Hence this awkwardness.
		var winner *Player
		for _, play := range plays {
			if play.player.name == subWinner.name {
				winner = play.player
				winner.addCard(play.card)
				break
			}
		}
		for _, play := range plays {
			if play.player.name != subWinner.name {
				winner.addCard(play.card)
			}
		}
	} else {
		sort.Slice(plays, func(i, j int) bool {
			return plays[i].card > plays[j].card
		})
		winner := plays[0].player
		fmt.Printf("%s wins the round!\n\n", winner.name)
		for i := 0; i < len(plays); i++ {
			winner.addCard(plays[i].card)
		}
	}
}

func hasWinner(players []*Player) *Player {
	zeroCount := 0
	var winner *Player
	for _, player := range players {
		if len(player.deck) == 0 {
			zeroCount++
		} else {
			winner = player
		}
	}
	if zeroCount == len(players)-1 {
		return winner
	}
	return nil
}

type Play struct {
	player *Player
	card   int
}

type Player struct {
	name string
	deck []int
}

func (p *Player) getPlay() *Play {
	if len(p.deck) == 0 {
		return nil
	}
	play := &Play{player: p, card: p.deck[0]}
	p.deck = p.deck[1:]
	return play
}

func (p *Player) addCard(card int) {
	p.deck = append(p.deck, card)
}

func (p *Player) score() int {
	score := 0
	for i, card := range p.deck {
		score += card * (len(p.deck) - i)
	}
	return score
}

func parseInput(path string) []*Player {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	sections := strings.Split(string(bytes), "\n\n")
	players := []*Player{}
	for _, section := range sections {
		players = append(players, parsePlayer(section))
	}
	return players
}

func parsePlayer(section string) *Player {
	lines := strings.Split(section, "\n")
	name := ""
	deck := []int{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "Player ") {
			name = line
		} else {
			val, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			deck = append(deck, val)
		}
	}
	return &Player{name: name, deck: deck}
}
