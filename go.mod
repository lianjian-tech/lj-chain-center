module lj-chain-center

go 1.17

require (
	git.huawei.com/poissonsearch/wienerchain/contract/docker-container/contract-go/contractapi v0.0.0-00010101000000-0000000000
	github.com/ethereum/go-ethereum v1.10.7
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-gonic/gin v1.7.4
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.3.0
	github.com/ipfs/go-ipfs-api v0.2.0
	github.com/jinzhu/copier v0.3.2
	github.com/jmoiron/sqlx v1.3.4
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron v1.2.0
	github.com/shirou/gopsutil v3.21.7+incompatible
	github.com/shopspring/decimal v1.2.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/valyala/fasthttp v1.29.0
	github.com/willf/pad v0.0.0-20200313202418-172aa767f2a4
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	gopkg.in/yaml.v2 v2.4.0
)

replace git.huawei.com/poissonsearch/wienerchain/contract/docker-container/contract-go/contractapi => ./client/huawei/internal
