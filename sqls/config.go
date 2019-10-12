package sqls


type Configs struct {
	db_host string
	db_user string
	db_password string
	db_name string
}

func setConfig()*Configs{
	return &Configs{"47.101.67.178","root","wys1993%$#@!","shopx"}
}

func GetConfig()*Configs{
	return setConfig()
}
