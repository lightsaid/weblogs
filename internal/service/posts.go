package service

import (
	"fmt"
	"os"
)

func GetDedaultPostThumb() string {
	return fmt.Sprintf("%s", os.Getenv("ASSETS_PREFIX"))
}
