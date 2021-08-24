package middlewares

import (
	"os"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pmh-only/gohttpfs/configloader"
)

func ListApi(config configloader.Configuration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() != "GET" {
			return c.Next()
		}

		stat, err := os.Stat(path.Join(*config.StoragePath, strings.Replace(c.Path(), *config.ServePrefix, "", 1)))
		if err != nil {
			return c.JSON(FileListResponse{Success: false, Message: strings.Replace(err.Error(), *config.StoragePath, "", 1)})
		}

		if !stat.IsDir() {
			return c.Next()
		}

		files, err := os.ReadDir(path.Join(*config.StoragePath, strings.Replace(c.Path(), *config.ServePrefix, "", 1)))
		if err != nil {
			return c.JSON(FileListResponse{Success: false, Message: strings.Replace(err.Error(), *config.StoragePath, "", 1)})
		}

		var filelist FileList
		for _, file := range files {
			if *config.HideDotfile && strings.HasPrefix(file.Name(), ".") {
				continue
			}

			fileInfo, err := file.Info()
			if err != nil {
				return c.JSON(FileListResponse{Success: false, Message: strings.Replace(err.Error(), *config.StoragePath, "", 1)})
			}

			filelist = append(filelist, FileListItem{
				FileName:    file.Name(),
				FileSize:    uint(fileInfo.Size()),
				ModifiedAt:  uint(fileInfo.ModTime().Unix()),
				IsDirectory: file.IsDir(),
			})
		}

		c.JSON(FileListResponse{Success: true, FileList: filelist})
		return nil
	}
}
