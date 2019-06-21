package skills

import (
	"time"
)

const (
	burnCd       = 5 * time.Minute
	cureCd       = 5 * time.Minute
	disCd        = 5 * time.Minute
	flashCd      = 5 * time.Minute
	rCd          = 5 * time.Minute
	runCd        = 5 * time.Minute
	tpCd         = 5 * time.Minute
	weakCd       = 5 * time.Minute
	refineCd     = 5 * time.Minute
	screenCd     = 5 * time.Minute
	disciplineCd = 5 * time.Minute
	on           = true
)
const (
	burnCommand       = "dr"
	cureCommand       = "zl"
	flashCommand      = "sx"
	rCommand          = "r"
	runCommand        = "jp"
	tpCommand         = "cs"
	weakCommand       = "xr"
	refineCommand     = "jh"
	screenCommand     = "pz"
	disciplineCommand = "cj"
)

var (
	SkillCommands = []string{
		burnCommand,
		cureCommand,
		flashCommand,
		rCommand,
		runCommand,
		tpCommand,
		weakCommand,
		refineCommand,
		screenCommand,
		disciplineCommand,
	}
)

const (
	burnName       = "点燃"
	cureName       = "治疗"
	flashName      = "闪现"
	rName          = "大招"
	runName        = "疾跑"
	tpName         = "传送"
	weakName       = "虚弱"
	refineName     = "净化"
	screenName     = "屏障"
	disciplineName = "惩戒"
)

type Skill struct {
	CD      float64
	On      bool
	OutName string
	Command string
}

func (skill *Skill) String() {
	return
}

func NewSkill(command string) (skill *Skill) {
	skill = &Skill{}
	skill.On = on
	switch command {
	case burnCommand:
		skill.CD = burnCd.Seconds()
		skill.OutName = burnName
		skill.Command = burnCommand
	case cureCommand:
		skill.CD = cureCd.Seconds()
		skill.OutName = cureName
		skill.Command = cureCommand
	case flashCommand:
		skill.CD = flashCd.Seconds()
		skill.OutName = flashName
		skill.Command = flashCommand
	case rCommand:
		skill.CD = rCd.Seconds()
		skill.OutName = rName
		skill.Command = rCommand
	case runCommand:
		skill.CD = runCd.Seconds()
		skill.OutName = runName
		skill.Command = runCommand
	case tpCommand:
		skill.CD = tpCd.Seconds()
		skill.OutName = tpName
		skill.Command = tpCommand
	case weakCommand:
		skill.CD = weakCd.Seconds()
		skill.OutName = weakName
		skill.Command = weakCommand

	case refineCommand:
		skill.CD = refineCd.Seconds()
		skill.OutName = refineName
		skill.Command = refineCommand

	case screenCommand:
		skill.CD = screenCd.Seconds()
		skill.OutName = screenName
		skill.Command = screenCommand

	case disciplineCommand:
		skill.CD = disciplineCd.Seconds()
		skill.OutName = disciplineName
		skill.Command = disciplineCommand
	}
	return
}

func (skill *Skill) Start() {
	if skill.On == false {
		return
	} else {
		skill.On = false
	}
	timer1 := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer1.C:
			skill.CD--
			if skill.CD == 0 {
				skill.On = true
				return
			}
		}
	}
}
