package logic

import "github.com/jekiapp/hi-mod-arch/internal/config"

func Init(cfg *config.Config) error {
	// you can create various initialization in logic layer as needed
	// for example: callwrapper, featureflag, default value, etc.
	//
	// the pattern should follow as in domain init.go
	return nil
}
