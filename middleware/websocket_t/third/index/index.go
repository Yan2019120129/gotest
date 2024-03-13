package index

import (
	"context"
	"encoding/json"
	"github.com/fasthttp/websocket"
	"go.uber.org/zap"
	"gotest/common/module/logs"
	"os"
	"strings"
	"sync"
	"time"
)

// Instance 全局唯一实例
var Instance = &WsManage{
	WsConnMap: make(map[string]*Ws),
	statusMap: make(map[string]*WsManageConfig),
	lock:      sync.Mutex{},
	once:      sync.Once{},
	Interval:  2,
}

// WsManageConfig websocket 实例状态处理
type WsManageConfig struct {
	status        int           // 状态：停止，开启，
	IsPersistence bool          // 是否持久化
	IsCleanUp     bool          // 是否过期清理
	Expiration    time.Duration // 过期时间
	Persistence
	Config
}

// WsManage websocket 实例
type WsManage struct {
	once     sync.Once     // 保证状态检测只运行一次
	lock     sync.Mutex    // websocket 实例读写锁
	Interval time.Duration // 间隔多少秒检查websocket 状态
	Persistence
	statusMap map[string]*WsManageConfig // 实例状态：启动，关闭
	WsConnMap map[string]*Ws             // 存储 websocket 实例
}

// NewWsManage 新建websocket 管理实例
func NewWsManage(Interval time.Duration) *WsManage {
	return &WsManage{
		once:      sync.Once{},
		lock:      sync.Mutex{},
		Interval:  Interval,
		statusMap: make(map[string]*WsManageConfig),
		WsConnMap: make(map[string]*Ws),
	}
}

// NewWs 新建websocket 实例
func (i *WsManage) NewWs(cfg ...WsManageConfig) *WsManage {
	for _, v := range cfg {
		if _, ok := i.WsConnMap[v.Id]; !ok {
			logs.Logger.Info("NewWs run")
			// 创建实例
			i.WsConnMap[v.Id] = NewWs(&Config{
				Addr:          v.Addr,
				Id:            v.Id,
				Pulse:         v.Pulse,
				Nor:           v.Nor,
				ManageMessage: v.ManageMessage,
			})

			// 设置状态
			i.statusMap[v.Id] = &WsManageConfig{
				status:        WsStatusStop,
				IsPersistence: v.IsPersistence,
				IsCleanUp:     v.IsCleanUp,
				Expiration:    v.Expiration,
			}
		}
	}

	return i
}

// Run 运行websocket
func (i *WsManage) Run(uuids ...string) *WsManage {
	// 不传参数则启动全部实例
	if len(uuids) < 1 {
		i.Get()
		for k, v := range i.WsConnMap {
			if i.statusMap[k].status == WsStatusStart {
				continue
			}
			v.Run()
			i.statusMap[k].status = WsStatusStart
		}
		return i
	}

	// 启动指定实例
	for _, v := range uuids {
		if p, ok := i.WsConnMap[v]; ok {
			if i.statusMap[v].status == WsStatusStart {
				continue
			}
			p.Run()
		} else {
			logs.Logger.Error("could not find " + v)
		}
		logs.Logger.Info("websocket", zap.String(v, "run"))
	}

	i.once.Do(func() {
		i.StatusMonitor()
	})
	return i
}

func (i *WsManage) StatusMonitor() {
	logs.Logger.Info("StatusMonitor run")
	go func() {
		ch := time.NewTicker(i.Interval * time.Second)
		for {
			for k, v := range i.WsConnMap {
				if err := v.ctx.instance.Err(); err != nil {
					logs.Logger.Error("websocket", zap.String(k, "close"))
				}
				// 状态修改为停止
				i.statusMap[k].status = WsStatusStop
			}
			<-ch.C
		}
	}()
}

// SetWs 设置 websocket。
func (i *WsManage) SetWs(uuid string, ws *Ws) *WsManage {
	// 添加协程使用WaitGroup管理线程状态
	if _, ok := i.WsConnMap[uuid]; ok {
		i.WsConnMap[uuid] = ws
		i.statusMap[uuid] = &WsManageConfig{
			status:        WsStatusStop,
			IsPersistence: false,
			IsCleanUp:     false,
			Expiration:    0,
		}
	}
	return i
}

// GetUUIDAll 获取全部实例的UUID
func (i *WsManage) GetUUIDAll() []string {
	i.lock.Lock()
	defer i.lock.Unlock()
	uuids := make([]string, 0)
	for k, _ := range i.WsConnMap {
		uuids = append(uuids, k)
	}
	return uuids
}

// ConnectWS 连接okx websocket。
func (i *WsManage) reconnection(uuids ...string) (err error) {
	for _, v := range uuids {
		if p, ok := i.statusMap[v]; ok {
			switch p.status {
			case WsStatusStart:
			case WsStatusStop:
				i.setWsConn(v, getWebSocketInstance(v))
				if err != nil {
					logs.Logger.Error(err.Error())
					return err
				}
			default:
				panic("unhandled default case")
			}

		}
	}
	return nil
}

// SendMessage 发送数据
func (i *WsManage) SendMessage(msg ...*Massage) *WsManage {
	logs.Logger.Info("SendMessage run")
	for _, v := range msg {
		if p, ok := i.WsConnMap[v.Id]; ok {
			p.SendMessage()
		}
	}
	return i
}

// SendMessageJson 发送数据
func (i *WsManage) SendMessageJson(msg ...*Massage) *WsManage {
	logs.Logger.Info("SendMessage run")
	for _, v := range msg {
		if p, ok := i.WsConnMap[v.Id]; ok {
			p.SendMessageJson(*v)
		}
	}
	return i
}

// GetWs 获取ws实例
func (i *WsManage) GetWs(uuid string) *Ws {
	if _, ok := i.WsConnMap[uuid]; !ok {
		return i.WsConnMap[uuid]
	}
	return nil
}

// setWs 设置ws实例
func (i *WsManage) setWs(uuid string, ws *Ws) *WsManage {
	if _, ok := i.WsConnMap[uuid]; !ok {
		i.WsConnMap[uuid] = ws
	}
	return i
}

// GetServerAddr 获取服务地址
func (i *WsManage) GetServerAddr(uuid string) string {
	if p, ok := i.WsConnMap[uuid]; !ok {
		return p.serverAdder
	}
	return ""
}

// SetServerAddr 设置服务器地址
func (i *WsManage) SetServerAddr(uuid, addr string) *WsManage {
	if _, ok := i.WsConnMap[uuid]; !ok {
		i.WsConnMap[uuid].serverAdder = addr
	}
	return i
}

// GetWsConn 获取websocket实例
func (i *WsManage) GetWsConn(uuid string) *websocket.Conn {
	if p, ok := i.WsConnMap[uuid]; !ok {
		return p.instance
	}
	return nil
}

// SetWsConn 设置服务器地址
func (i *WsManage) setWsConn(uuid string, conn *websocket.Conn) {
	if _, ok := i.WsConnMap[uuid]; !ok {
		i.WsConnMap[uuid].instance = conn
		i.statusMap[uuid] = &WsManageConfig{
			status:        -1,
			IsPersistence: false,
			IsCleanUp:     false,
			Expiration:    0,
		}
	}
}

// GetContext 获取上下文
func (i *WsManage) GetContext(uuid string) *Ctx {
	if _, ok := i.WsConnMap[uuid]; ok {
		return i.WsConnMap[uuid].ctx
	}
	return nil
}

// setContext 设置上下文
func (i *WsManage) setContext(uuid string) {
	if _, ok := i.WsConnMap[uuid]; ok {
		instance, c := context.WithCancel(context.Background())
		i.WsConnMap[uuid].ctx.instance = instance
		i.WsConnMap[uuid].ctx.close = c
	}
}

// Set 数据持久化
func (i *WsManage) Set(configs ...Config) {
	logs.Logger.Info("Persistence run")
	// 判断路径是否存在,不存在则创建
	if isPathExist(storagePath) {
		oldCfgs := i.Get()
		configMap := map[string]*Config{}
		for _, v := range oldCfgs {
			configMap[v.Id] = v
		}
		for _, v := range configs {
			configMap[v.Id] = &v
		}
		logs.Logger.Info("Persistence", zap.Reflect("msg", configMap))
		byteData, err := json.Marshal(configMap)
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

// Get 获取持久化数据
func (i *WsManage) Get(ids ...string) []*Config {
	logs.Logger.Info("GetPersistence run")
	if !isPathExist(storagePath) {
		return nil
	}

	storageData, err := os.ReadFile(storagePath)
	if err != nil || storageData == nil || len(storageData) == 0 {
		logs.Logger.Error("read file error", zap.Error(err))
		return nil
	}

	configMap := map[string]*Config{}
	if err = json.Unmarshal(storageData, &configMap); err != nil {
		logs.Logger.Error("unmarshal persistence data error", zap.Error(err))
		return nil
	}

	cfgs := []*Config{}
	for _, v := range ids {
		logs.Logger.Info("GetPersistence", zap.String("path", storagePath))
		if _, ok := configMap[v]; ok {
			cfgs = append(cfgs, configMap[v])
		}
	}

	for _, v := range configMap {
		cfgs = append(cfgs, v)
	}

	// 不穿参数为获取全部实例化数据
	if len(ids) == 0 {
		return cfgs
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
