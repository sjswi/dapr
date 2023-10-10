package components

import (
	"github.com/dapr/components-contrib/rdb/mysql"
	testLoader "github.com/dapr/dapr/pkg/components/rdb"
)

func init() {
	testLoader.DefaultRegistry.RegisterComponent(mysql.NewMySQL, "mysql")
}
