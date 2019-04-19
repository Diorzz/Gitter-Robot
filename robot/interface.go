package robot

type Robot interface {
	Ask(msg string) (string, error)
}

func NewRobot(rtype string) Robot {
	switch rtype {
	case "qingyun":
		return NewQingYunRobot()
	default:
		return nil
	}
}
