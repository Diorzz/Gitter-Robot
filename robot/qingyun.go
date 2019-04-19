package robot

import (
	"encoding/json"
	"fmt"

	"github.com/Gitter-Robot/request"
)

const (
	qingYunBaseUrl = "http://api.qingyunke.com/api.php?key=free&appid=0&msg="
)

type QingYunRobot struct {
	ApiUrl string
}

func NewQingYunRobot() *QingYunRobot {
	return &QingYunRobot{
		ApiUrl: qingYunBaseUrl,
	}
}

// Ask returns the result of robot.
func (robot *QingYunRobot) Ask(msg string) (string, error) {
	// QingYun robot api use http get method to get result.
	httpProxy := &request.HttpRequest{}
	res, err := httpProxy.Get(robot.ApiUrl + msg)
	if err != nil {
		return "", err
	}

	var result QingYunResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}
	if result.Result != 0 {
		return "", fmt.Errorf("QingYun return error")
	}
	return result.Content, nil

}

// The result struct of Qingyun robot.
type QingYunResult struct {
	Result  int    `json:"result"`
	Content string `json:"content"`
}
