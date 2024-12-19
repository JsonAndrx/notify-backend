package config

import(
	"os"
	"github.com/tidwall/gjson"
	"notify-backend/api/utils/debug"
)


func GetCodeError(key string) string {
	jsonConstants, errLoadConstants := os.ReadFile("./api/utils/response/codes.json")
	if errLoadConstants != nil {
		debug.LogError(errLoadConstants)
	}
	return gjson.Get(string(jsonConstants), key).String()
}