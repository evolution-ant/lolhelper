package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lolhelper/skills"
	"github.com/lolhelper/utils"

	"github.com/lolhelper/hero"
	"github.com/peterh/liner"
)

var (
	history_fn = filepath.Join(os.TempDir(), ".liner_example_history")
	names      = []string{"john", "james", "mary", "nancy"}
)

const (
	maren1 = "我***你个****你****全家*******"
)

func main() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	line.SetCompleter(func(line string) (c []string) {
		for _, n := range names {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})

	if f, err := os.Open(history_fn); err == nil {
		line.ReadHistory(f)
		f.Close()
	}
	heros := []*hero.Hero{}
	for i, heroCommand := range hero.HeroCommands {
		// isOk := false
		// for {
		if command, err := line.Prompt("输入" + hero.HeroNames[i] + "技能首字母 比如闪现和治疗：sxzl  "); err == nil {
			// isOk = isValid(command)
			// if isOk {
			// 	break
			// } else {
			// 	fmt.Println("输入错误,重新输入")
			// 	continue
			// }
			h := hero.NewHero(heroCommand + command)
			heros = append(heros, h)
			fmt.Println(h.S1.OutName, h.S2.OutName)
		} else if err == liner.ErrPromptAborted {
			os.Exit(0)
		} else {
			log.Print("Error reading line: ", err)
		}
		// }
	}
	for _, v := range heros {
		fmt.Println(v.String())
	}

	// go Looper(heros)
	for {
		if command, err := line.Prompt(""); err == nil {
			if command == "sb" {
				fmt.Println(maren1)
				utils.CopyToClipBoard(maren1)
				continue
			}
			if command == "all" {
				allString := ""
				for _, v := range heros {
					allString += v.String() + "\r\n"
				}
				fmt.Println(allString)
				utils.CopyToClipBoard(allString)
				continue
			}
			if len(command) > 3 || len(command) == 0 {
				fmt.Println("输入错误")
				continue
			}
			switch string(command[0]) {
			case hero.AdCommand:
				handleCmd(heros[0], command)
			case hero.SupCommand:
				handleCmd(heros[1], command)
			case hero.MidCommand:
				handleCmd(heros[2], command)
			case hero.JugCommand:
				handleCmd(heros[3], command)
			case hero.TopCommand:
				handleCmd(heros[4], command)
			}
		} else if err == liner.ErrPromptAborted {
			os.Exit(0)
		} else {
			log.Print("Error reading line: ", err)
		}
	}
	if f, err := os.Create(history_fn); err != nil {
		log.Print("Error writing history file: ", err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}
}

func Looper(heros []*hero.Hero) {
	timer1 := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-timer1.C:
			allString := ""
			for _, v := range heros {
				allString += v.String() + "\r\n"
			}
			fmt.Println(allString)
			utils.CopyToClipBoard(allString)
		}
	}
}

func isValid(command string) bool {
	isHero := false
	isSkill := false
	if len(command) == 2 {
		for _, v := range hero.HeroCommands {
			if string(command[0]) == v {
				isHero = true
			}
		}
		for _, v := range skills.SkillCommands {
			if string(command[1]) == v {
				isSkill = true
			}
		}
	}
	return isHero && isSkill
}
func handleCmd(hero *hero.Hero, command string) {
	if len(command) == 1 {
		fmt.Println(hero.String())
		utils.CopyToClipBoard(hero.String())
	} else {
		hero.StartCD(string(command[1:3]))
	}
}
