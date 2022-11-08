package abstract

import "github.com/tester-rep/spr-go/channel"

type IRobot interface {
	Init() bool
	RunAction() int
	Close()
	Run(robot IRobot)
	IsEnableSampler() bool
	SetStopRobot(sign bool)
	IsActive() bool
	SetActive(isActive bool)
}

type AbstractRobot struct {
	Index         int
	TranTotal     int64
	initSuccess   bool
	StopRobot     bool
	enableSampler bool
	isActive      bool
}

func NewAbstractRobot(index int, tranTotal int64) *AbstractRobot {
	var robot = new(AbstractRobot)
	robot.Index = index
	robot.TranTotal = tranTotal
	robot.StopRobot = false
	robot.enableSampler = true
	robot.isActive = false
	return robot
}

func (a *AbstractRobot) Init() bool {
	return true
}

func (a *AbstractRobot) RunAction() int {
	return 0
}

func (a *AbstractRobot) Close() {

}

func (a *AbstractRobot) Run(robot IRobot) {
	//机器人初始化
	a.initSuccess = robot.Init()
	if !a.initSuccess {
	}
	var i = int64(1)
	for a.initSuccess { //初始化成功才能跑
		if a.StopRobot { //如果有退出信号 退出
			break
		}
		if a.TranTotal != -1 && i > a.TranTotal {
			break
		}
		//非0，表示有错误
		if 0 != robot.RunAction() {
			break
		}
		i++
	} //end for
	robot.Close()
	channel.StopNum <- 1
}

func (a *AbstractRobot) IsEnableSampler() bool {
	return a.enableSampler
}

func (a *AbstractRobot) SetStopRobot(sign bool) {
	a.StopRobot = sign
}

func (a *AbstractRobot) IsActive() bool {
	return a.isActive
}
func (a *AbstractRobot) SetActive(isActive bool) {
	a.isActive = isActive
}
