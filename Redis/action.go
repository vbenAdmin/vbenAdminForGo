package Redis

import (
	"strconv"
)

var userTokenKey = "UT"

func joinKey(userId uint) string {
	return userTokenKey + strconv.Itoa(int(userId))
}
func SetUserToken(userId uint, value int64) error {
	return Set(joinKey(userId), value, 0)
}

func GetUserToken(userId uint) (int64, error) {
	value, err := Get(joinKey(userId))
	valueInt, _ := strconv.ParseInt(value, 10, 64)
	return valueInt, err
}
