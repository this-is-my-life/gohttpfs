package middlewares

import (
	"archive/zip"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pmh-only/gohttpfs/configloader"
)

func ArchiveApi(config configloader.Configuration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() != "GET" || len(c.Query("archive")) < 1 {
			return c.Next()
		}

		archive, err := ioutil.TempFile(os.TempDir(), "gohttpfs-")
		if err != nil {
			log.Println("[archive/tempfilegen]", err)
			return fiber.ErrInternalServerError
		}

		zipWriter := zip.NewWriter(archive)
		err = archiveFolder(zipWriter, path.Join(*config.StoragePath, strings.Replace(c.Path(), *config.ServePrefix, "", 1)), "", config)
		if err != nil {
			log.Println("[archive/archivefolder]", err)
			return fiber.ErrInternalServerError
		}

		zipWriter.Close()
		archive.Close()

		c.Download(archive.Name())

		return nil
	}
}

// from https://stackoverflow.com/a/49233329 - LeTigre, ahmelsayed
func archiveFolder(writer *zip.Writer, basePath, baseInZip string, config configloader.Configuration) error {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if *config.HideDotfile && strings.HasPrefix(file.Name(), ".") {
			continue
		}

		if !file.IsDir() {
			dat, err := ioutil.ReadFile(path.Join(basePath, file.Name()))
			if err != nil {
				return err
			}

			// Add some files to the archive.
			f, err := writer.Create(path.Join(basePath, file.Name()))
			if err != nil {
				return err
			}

			_, err = f.Write(dat)
			if err != nil {
				return err
			}
		} else if file.IsDir() {
			err := archiveFolder(writer, path.Join(basePath, file.Name()), path.Join(baseInZip, file.Name()), config)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
