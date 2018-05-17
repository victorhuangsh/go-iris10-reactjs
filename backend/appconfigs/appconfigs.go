package appconfigs

import (
	"sync"
	"github.com/kataras/iris"
)

var instance *AppConfigs
var once sync.Once
//var mu sync.Mutex

func GetInst( ) *AppConfigs {
	return GetInstance(nil)
}

func GetInstance( iris_app *iris.Application ) *AppConfigs {
	once.Do(func() {
		instance = &AppConfigs{}
		instance.Init(iris_app)
	})

	/*
	// The lock mode
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		instance = &AppConfigs{}     // unnecessary locking if instance already created
		instance.Init(iris_app)
	}
	*/

	return instance
}

type AppConfigs struct {
	G_H5Host string
	G_GOHost string
	G_LogPath string
	RuntimeMode int
}

func (p *AppConfigs) Init(iris_app *iris.Application) {

	var m_app *iris.Application
	if (iris_app != nil){
		m_app = iris_app
	}else {
		m_app = iris.New()
	}

	m_app.Configure(iris.WithConfiguration(iris.YAML("./backend/configs/cfg_run_app.yml")))
	cfg_all := m_app.ConfigurationReadOnly().GetOther()
	if _, ok := cfg_all["H5Host"]; ok {
		p.G_H5Host = cfg_all["H5Host"].(string)
	}
	if _, ok := cfg_all["GOHost"]; ok {
		p.G_GOHost = cfg_all["GOHost"].(string)
	}

	if _, ok := cfg_all["LogPath"]; ok {
		p.G_LogPath = cfg_all["LogPath"].(string)
	}

	if _, ok := cfg_all["RuntimeMode"]; ok {
		m_int_runtime := cfg_all["RuntimeMode"]
		p.RuntimeMode = m_int_runtime.(int)
		/*
		m_str_runtime := m_int_runtime.(string)
		if runtimeMode, err := strconv.Atoi(m_str_runtime); err == nil {
			fmt.Printf("i=%d, type: %T\n", runtimeMode, runtimeMode)
		}else{
			p.RuntimeMode = runtimeMode
		}
		*/
	}
}

func (p *AppConfigs) SetH5Host( h5Host string ){
	p.G_H5Host = h5Host
}

func (p *AppConfigs) SetGOHost( goHost string ){
	p.G_GOHost = goHost
}

func (p *AppConfigs) SetLogPath( input string ){
	p.G_LogPath = input
}

func (p *AppConfigs) SetRuntimeMode( input int ){
	p.RuntimeMode = input
}