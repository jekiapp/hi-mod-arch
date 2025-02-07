package user

import (
	"fmt"
	"net/http"

	"github.com/jekiapp/hi-mod-arch/internal/config"
	"github.com/jekiapp/hi-mod-arch/internal/model"
)

func Init(cfg *config.Config) error {
	userSvcURL = fmt.Sprintf("https://%s/%s", cfg.User.Host, cfg.User.GetUserInfoPath)
	return nil
}

var userSvcURL string

func GetUserInfo(cfg *config.Config, cli *http.Client, userID int64) (model.UserData, error) {
	// request to upstream to get user data
	cli.Get(userSvcURL)
	// check error
	// ...
	// Unmarshal the response into the object
	// ...
	return model.UserData{}, nil
}
