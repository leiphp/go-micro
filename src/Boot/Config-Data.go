package Boot

type ConfigInterface interface {
	Reload() error
	IsLoad() bool
}

type MySqlConfig struct {
	Dsn 	string
	Maxidle int
	Maxopen int
}

type RedisConfig struct {
	Ip string
	Port int
}

func (* MySqlConfig) Reload() error {
	return ReloadDB()
}

func (this *MySqlConfig) IsLoad() bool {
	if this.Dsn != "" && this.Maxopen > 0 && this.Maxidle > 0 {
		return true
	}
	return false
}

func (* RedisConfig) Reload() error {
	return nil
}

func (this *RedisConfig) IsLoad() bool {
	if this.Ip != "" && this.Port > 0 {
		return true
	}
	return false
}
