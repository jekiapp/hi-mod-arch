package user

import (
	"fmt"
	"net/http"

	"github.com/jekiapp/hi-mod/internal/config"
	"github.com/jekiapp/hi-mod/internal/model"
)

func GetUserInfo(cfg *config.Config, cli *http.Client, userID int64) (model.UserData, error) {
	// request to upstream to get user data
	url := fmt.Sprintf("https://%s/%s", cfg.User.Host, cfg.User.GetUserInfoPath)
	cli.Get(url)
	// check error
	// ...
	// Unmarshal the response into the object
	// ...
	return model.UserData{}, nil
}
