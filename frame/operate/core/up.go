package core

import (
	"encoding/json"
	"fmt"
	"operate/conf"
	"operate/utils"
)

// Up 上传文件
func Up(path string) error {
	if !utils.IsExistFile(path) {
		return fmt.Errorf("this %s is exist", path)
	}

	instance := utils.NewHttp()
	v, err := instance.PostFile(conf.Conf.Base.TargetServer+"/v1/files", path)
	if err != nil {
		return err
	}

	resp := struct {
		Expire int      `json:"expire"`
		Files  []string `json:"files"`
	}{}

	_ = json.Unmarshal(v, &resp)

	for _, fileName := range resp.Files {
		fmt.Println(fileName)
	}

	return nil
}
