package hero

import (
	"errors"
	"strconv"

	"github.com/lolhelper/skills"
)

const (
	AdCommand  = "a"
	SupCommand = "f"
	MidCommand = "z"
	JugCommand = "d"
	TopCommand = "s"
)

var (
	HeroCommands = []string{
		AdCommand,
		SupCommand,
		MidCommand,
		JugCommand,
		TopCommand,
	}
	HeroNames = []string{
		adName,
		supName,
		midName,
		jugName,
		topName,
	}
)

const (
	adName  = "AD "
	supName = "SUP"
	midName = "MID"
	jugName = "JUG"
	topName = "TOP"
)

type Hero struct {
	S1      *skills.Skill
	S2      *skills.Skill
	OutName string
}

func (hero *Hero) StartCD(skCommand string) error {
	if skCommand == hero.S1.Command {
		go hero.S1.Start()
		return nil
	} else if skCommand == hero.S2.Command {
		go hero.S2.Start()
		return nil
	}
	return errors.New("there is no this skill!")
}

func (hero *Hero) String() (status string) {
	status = "[" + hero.OutName + "]" + " "
	if hero.S1.On == true {
		status += hero.S1.OutName + " " + "√    "
	} else {
		status += hero.S1.OutName + " " + strconv.FormatFloat(hero.S1.CD, 'f', 0, 64) + "秒"
	}
	if hero.S2.On == true {
		status += " " + hero.S2.OutName + " " + "√    "
	} else {
		status += " " + hero.S2.OutName + " " + strconv.FormatFloat(hero.S2.CD, 'f', 0, 64) + "秒"
	}

	return
}

func NewHero(command string) (hero *Hero) {
	hero = InitHero(command)
	switch string(command[0]) {
	case AdCommand:
		hero.OutName = adName
	case SupCommand:
		hero.OutName = supName
	case JugCommand:
		hero.OutName = jugName
	case MidCommand:
		hero.OutName = midName
	case TopCommand:
		hero.OutName = topName
	}
	return
}

func InitHero(command string) (hero *Hero) {
	hero = &Hero{}
	hero.S1 = skills.NewSkill(string(command[1:3]))
	hero.S2 = skills.NewSkill(string(command[3:5]))
	return
}
