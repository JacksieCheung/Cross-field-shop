module Data-acquisition-subsystem

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.7
	github.com/jinzhu/gorm v1.9.16
	github.com/satori/go.uuid v1.2.0
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.8.1
	github.com/teris-io/shortid v0.0.0-20201117134242-e59966efd125
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/willf/pad v0.0.0-20200313202418-172aa767f2a4
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.uber.org/zap v1.21.0
	golang.org/x/crypto v0.0.0-20220518034528-6f7dac969898 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/spf13/afero => github.com/spf13/afero v1.5.1
