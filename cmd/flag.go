package cmd

import (
	"fmt"
)

var (
// ad   = hero.NewHero(hero.AdCommand)
// sup  = hero.NewHero(hero.SupCommand)
// mid  = hero.NewHero(hero.MidCommand)
// top  = hero.NewHero(hero.TopCommand)
// jug  = hero.NewHero(hero.JugCommand)
// team = []*hero.Hero{ad, sup, mid, top, jug}
)

const (
	All  = "all"
	Help = "help"
)

func HeroSetting(command string) {

}

func HandleCommand(command string) {
	switch command {
	case Help:
		fmt.Println("help !")
	case All:
		// for _, hero := range team {
		// fmt.Println(hero.Status())
		// }
	case "af":
		// go ad.S1.Start()
		// fmt.Println(ad.Status())
		// utils.CopyToClipBoard(ad.Status())
	case "ac":
		// go ad.S2.Start()
	}
}
