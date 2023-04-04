package config

type Config struct {
	Mysql struct {
		DataSource string
	}

	Log struct {
		Mode     string
		Encoding string
		Path     string
		Level    string
	}

	//定时任务config
	CronConf struct {
		User string
	}
}
