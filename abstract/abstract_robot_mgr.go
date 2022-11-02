package abstract

type IAbstractRobotMgr interface {
	InitAllRobot()
	AddRobot(robot IRobot)
	GetRobotLs() []IRobot
	InitRobotMgr(robotNum int, tranTotal int64, startIndex int)
	InitBase()
}

type AbstractRobotMgr struct {
	robotLs    []IRobot
	RobotNum   int
	TranTotal  int64
	StartIndex int
}

func NewAbstractRobotMgr() *AbstractRobotMgr {
	var abstractRobotMgr = AbstractRobotMgr{
		robotLs: make([]IRobot, 0),
	}
	return &abstractRobotMgr
}

func (a *AbstractRobotMgr) InitRobotMgr(robotNum int, tranTotal int64, startIndex int) {
	a.RobotNum = robotNum
	a.TranTotal = tranTotal
	a.StartIndex = startIndex
}

func (a *AbstractRobotMgr) AddRobot(robot IRobot) {
	a.robotLs = append(a.robotLs, robot)
}

func (a *AbstractRobotMgr) SetRobotLs(num []IRobot) {
	a.robotLs = num
}

func (a *AbstractRobotMgr) GetRobotLs() []IRobot {
	return a.robotLs
}

func (a *AbstractRobotMgr) InitAllRobot() {
}

func (a *AbstractRobotMgr) GetRobotNum() int {
	return a.RobotNum
}

func (a *AbstractRobotMgr) SetRobotNum(num int) {
	a.RobotNum = num
}

func (a *AbstractRobotMgr) SetTranTotal(num int64) {
	a.TranTotal = num
}

func (a *AbstractRobotMgr) GetTranTotal() int64 {
	return a.TranTotal
}

func (a *AbstractRobotMgr) SetIndexStart(startIndex int) {
	a.StartIndex = startIndex
}
func (a *AbstractRobotMgr) GetIndexStart() int {
	return a.StartIndex
}
func (a *AbstractRobotMgr) InitBase() {
}
