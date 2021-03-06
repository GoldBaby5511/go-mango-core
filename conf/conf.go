package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoldBaby5511/go-mango-core/log"
	n "github.com/GoldBaby5511/go-mango-core/network"
	"github.com/GoldBaby5511/go-mango-core/util"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	ArgAppName         string = "-Name"
	ArgAppType         string = "-Type"
	ArgAppId           string = "-Id"
	ArgCenterAddr      string = "-CenterAddr"
	ArgListenOnAddr    string = "-ListenOnAddr"
	ArgDockerRun       string = "-DockerRun"
	ArgDefaultBasePort string = "-DefaultBasePort"

	//ζε‘ηΆζ
	AppStateNone              = 0
	AppStateStarting          = 1
	AppStateRunning           = 2
	AppStateMaintenance       = 4
	AppStateMaintenanceFinish = 8
	AppStateOffline           = 16
)

var (
	LenStackBuf               = 4096
	GoLen                     = 10000
	TimerDispatcherLen        = 10000
	AsynCallLen               = 10000
	ChanRPCLen                = 10000
	DefaultBasePort    uint32 = 10000
	AppInfo            BaseInfo
	ApplogDir          string
)

type BaseInfo struct {
	Name         string
	Type         uint32
	Id           uint32
	ListenOnAddr string
	CenterAddr   string
}

func LoadBaseConfig(name string) {
	AppInfo.Name = name
	if AppInfo.Name != "" {
		data, err := ioutil.ReadFile(fmt.Sprintf("configs/%s/%s.json", AppInfo.Name, AppInfo.Name))
		if err == nil {
			err = json.Unmarshal(data, &AppInfo)
		}
	}
	args := os.Args
	if v, ok := util.ParseArgsString(ArgAppName, args); ok {
		AppInfo.Name = v
	}
	if v, ok := util.ParseArgsUint32(ArgAppType, args); ok {
		AppInfo.Type = v
	}
	if v, ok := util.ParseArgsUint32(ArgAppId, args); ok {
		AppInfo.Id = v
	}
	if v, ok := util.ParseArgsString(ArgCenterAddr, args); ok {
		AppInfo.CenterAddr = v
	}
	if v, ok := util.ParseArgsString(ArgListenOnAddr, args); ok {
		AppInfo.ListenOnAddr = v
	}
	if v, ok := util.ParseArgsUint32(ArgDefaultBasePort, args); ok {
		DefaultBasePort = v
	}
	if AppInfo.ListenOnAddr == "" {
		AppInfo.ListenOnAddr = fmt.Sprintf("0.0.0.0:%d", DefaultBasePort+AppInfo.Id)
	}
	if AppInfo.CenterAddr == "" && AppInfo.Type != n.AppCenter && AppInfo.Type != n.AppLogger {
		AppInfo.CenterAddr = fmt.Sprintf("127.0.0.1:%v", DefaultBasePort+50)
		log.Debug("", "δ½Ώη¨ι»θ?€ε°ε,CenterAddr=%v", AppInfo.CenterAddr)
	}
	if RunInLocalDocker() {
		AppInfo.CenterAddr = "center:" + strconv.Itoa(util.GetPortFromIPAddress(AppInfo.CenterAddr))
	}

	if util.PortInUse(util.GetPortFromIPAddress(AppInfo.ListenOnAddr)) {
		log.Fatal("εε§ε", "η«―ε£[%v]ε·²θ’«ε η¨,θ―·ζ£ζ₯θΏθ‘η―ε’", util.GetPortFromIPAddress(AppInfo.ListenOnAddr))
		return
	}

	if AppInfo.Name == "" || AppInfo.Type == 0 || AppInfo.Id == 0 || AppInfo.ListenOnAddr == "" ||
		(AppInfo.CenterAddr == "" && AppInfo.Type != n.AppCenter && AppInfo.Type != n.AppLogger) {
		log.Fatal("εε§ε", "εε§εζ°εΌεΈΈ,θ―·ζ£ζ₯,AppInfo=%v", AppInfo)
		return
	}

	//εε»Ίζ₯εΏη?ε½
	if err := makeAppLogDir(); err != nil {
		log.Fatal("εε§ε", "εε»Ίζ₯εΏη?ε½ε€±θ΄₯,err=%v", err)
		return
	}

	log.Debug("", "εΊη‘ε±ζ§,%v,logη?ε½=%v", AppInfo, ApplogDir)
}

func RunInLocalDocker() bool {
	args := os.Args
	if v, ok := util.ParseArgsUint32(ArgDockerRun, args); ok && v == 1 {
		return true
	}
	return false
}

func makeAppLogDir() error {
	curPath, err := util.GetCurrentPath()
	if err != nil {
		return errors.New("θ·εε½εθ·―εΎε€±θ΄₯")
	}
	pathName := fmt.Sprintf("%slog/%s%d/", curPath, AppInfo.Name, AppInfo.Id)
	err = os.MkdirAll(pathName, os.ModePerm)
	if err != nil {
		return errors.New("ζδ»Άθ·―εΎεε»Ίε€±θ΄₯")
	}
	ApplogDir = pathName
	return nil
}
