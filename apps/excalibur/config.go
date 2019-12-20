package excalibur

type Config struct {
	AccountOption
	GateHttpOption
	MongodbOption
}

type GateHttpOption struct {
	ListenAddress string
	UseTLS        bool
	DevModeEnable bool
}

type AccountOption struct {
}

type GameServiceOption struct {
	DevModeEnable bool
}

type MongodbOption struct {
	MongoAddress string
}

var devModeConfig = Config{

	GateHttpOption: GateHttpOption{
		ListenAddress: ":11002",
		UseTLS:        true,
		DevModeEnable: true,
	},

	MongodbOption: MongodbOption{
		MongoAddress: "mongodb://127.0.0.1:27017/excalibur_dev",
	},
}

var betaModeConfig = Config{
	GateHttpOption: GateHttpOption{
		ListenAddress: ":11002",
		UseTLS:        true,
		DevModeEnable: true,
	},
	MongodbOption: MongodbOption{
		MongoAddress: "mongodb://127.0.0.1:27017/excalibur_beta",
	},
}

var approvalModeConfig = Config{

	GateHttpOption: GateHttpOption{
		ListenAddress: ":11003",
		UseTLS:        true,
		DevModeEnable: true,
	},
	MongodbOption: MongodbOption{
		MongoAddress: "mongodb://127.0.0.1:27017/excalibur_approval",
	},
}

var prodModeConfig = Config{

	GateHttpOption: GateHttpOption{
		ListenAddress: "127.0.0.1:20001",
		UseTLS:        true,
		DevModeEnable: false,
	},
	MongodbOption: MongodbOption{
		MongoAddress: "mongodb://127.0.0.1:27017/excalibur_prod",
	},
}

var taptapModeConfig = Config{

	GateHttpOption: GateHttpOption{
		ListenAddress: "127.0.0.1:20002",
		UseTLS:        true,
		DevModeEnable: false,
	},
	MongodbOption: MongodbOption{
		MongoAddress: "mongodb://127.0.0.1:27017/excalibur_taptap",
	},
}
