package core

import (
	"fmt"
	"operate/conf"
	"operate/utils"
)

var (
	HistoryTypeUpload     = "Upload"
	HistoryTypeJumpServer = "JumpServer"
)

// History 历史记录
func History(n string) error {
	c := fmt.Sprintf("grep %s %s | tail -n %s", HistoryTypeUpload, conf.Conf.Log.Dir, n)
	valr, err := utils.ExecCommand(c)
	if err != nil {
		return err
	}
	fmt.Println(valr)
	return nil
}
