package index

import (
	"encoding/json"
	"go.uber.org/zap"
	"gotest/common/module/logs"
	"os"
	"strings"
)

const (
	storagePath = "./data/message.json"
)

// ManageMessage 数据处理方法
type ManageMessage interface {
	DealWithMessage(msgType int, data []byte)        // 处理消息
	Persistence(msg ...Massage)                      // 持久化方法
	GetPersistence(id string, msgType int) []Massage // 获取持久化数据
}

type ManageInstance struct {
	data []Massage
}

// DealWithMessage 处理消息方法
func (m *ManageInstance) DealWithMessage(msgType int, data []byte) {
	logs.Logger.Info("websocket", zap.Int("type", msgType), zap.String("data", string(data)))
}

// Persistence 数据持久化
func (m *ManageInstance) Persistence(msg ...Massage) {
	logs.Logger.Info("Persistence run")
	// 判断路径是否存在,不存在则创建
	if isPathExist(storagePath) {
		// 获取全部的数据
		if m.data == nil || len(m.data) == 0 {
			m.data = m.GetPersistence("", 0)
		}
		for _, v := range msg {
			// 当类型为订阅类型时才进行存储
			if v.Type == WsMessageTypeSub {
				m.data = append(m.data, msg...)
			}
		}
		logs.Logger.Info("Persistence", zap.Reflect("msg", m.data))
		byteData, err := json.Marshal(m.data)
		if err != nil {
			logs.Logger.Error("Marshal error", zap.Error(err))
			return
		}

		logs.Logger.Info("Persistence", zap.ByteString("msg", byteData))
		// 将数据写入json 文件
		if err = os.WriteFile(storagePath, byteData, 0664); err != nil {
			logs.Logger.Error("WriteFile error", zap.Error(err))
			return
		}
	}
}

// GetPersistence 获取持久化数据
func (m *ManageInstance) GetPersistence(id string, msgType int) []Massage {
	logs.Logger.Info("GetPersistence run")
	if !isPathExist(storagePath) {
		return nil
	}
	data := make([]Massage, 0)
	if m.data == nil || len(m.data) == 0 {
		// 读取消息
		logs.Logger.Info("GetPersistence", zap.String("path", storagePath))
		storageData, err := os.ReadFile(storagePath)
		if err != nil || storageData == nil || len(storageData) == 0 {
			logs.Logger.Error("read file error", zap.Error(err))
			return nil
		}
		if err = json.Unmarshal(storageData, &data); err != nil {
			logs.Logger.Error("unmarshal persistence data error", zap.Error(err))
			return nil
		}
	}

	dataTemp := make([]Massage, 0)
	for _, v := range data {
		logs.Logger.Info("GetPersistence", zap.String("for", string(v.Data)))
		switch {
		case id == v.Id && msgType == 0:
			// 获取指定实例全部数据
			dataTemp = append(dataTemp, v)
		case id == v.Id && msgType == v.Type:
			// 获取指定实例，指定类型数据
			dataTemp = append(dataTemp, v)
		case id == "" && msgType == 0:
			// 获取全部数据
			return data
		}
		return dataTemp
	}
	return nil
}

func isPathExist(path string) bool {
	index := strings.LastIndex(path, "/")
	path = path[:index]
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			logs.Logger.Error("Error creating directory:" + err.Error())
			return false
		}
		logs.Logger.Info("Directory created successfully:" + path)
		return true
	} else if err != nil {
		logs.Logger.Error("Error creating directory:" + err.Error())
		return false
	}
	return true
}
