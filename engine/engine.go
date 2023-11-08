package engine

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/CPU-commits/VeleroMC/db"
	"github.com/CPU-commits/VeleroMC/utils"
)

type engine struct {
	engine *db.Engine
}

func (engine *engine) Download(version string) ([]byte, error) {
	// Exists
	existsVersion, err := engine.engine.ExistsVersion(version)
	if err != nil {
		return nil, err
	}
	if !existsVersion {
		return nil, errVersionNotExists
	}
	// Download version
	url := engine.engine.BuildDownloadVersionURL(version)
	fmt.Printf("Downloading server...\n")
	res, err := utils.Fetch(url)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Success download!\n")
	// Read file in memory
	file, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (*engine) executeJavaCmd(fileName string) error {
	return utils.DoCmd(
		"java",
		"-Xmx1024M",
		"-Xms1024M",
		"-jar",
		fileName,
		"nogui",
	)
}

func (engine *engine) Build(javaFile []byte) error {
	fmt.Println("Start building server")
	// Save file locally
	folder := "minecraft"
	fileName := "server.jar"

	err := utils.SaveFile(path.Join(folder, fileName), javaFile)
	if err != nil {
		return err
	}
	// Execute java file
	err = os.Chdir(folder)
	if err != nil {
		return err
	}
	if err := engine.executeJavaCmd(fileName); err != nil {
		return err
	}
	// Agree to the EULA
	eula, err := utils.ReadFile("eula.txt")
	if err != nil {
		return err
	}
	agreeEula := strings.ReplaceAll(eula, "eula=false", "eula=true")
	err = utils.SaveFile("eula.txt", []byte(agreeEula))
	if err != nil {
		return err
	}
	// Execute java file (again)
	if err := engine.executeJavaCmd(fileName); err != nil {
		return err
	}

	return nil
}

func (engine *engine) Run(path string) error {
	fileName := filepath.Base(path)
	folder := filepath.Dir(path)

	if err := os.Chdir(folder); err != nil {
		return err
	}

	return engine.executeJavaCmd(fileName)
}

func (engine *engine) List() ([]string, error) {
	return engine.engine.Versions, nil
}

func newEngine(engineName string) iEngine {
	engineDb, err := engineRepository.FindOneByEngine(engineName)
	if err != nil {
		panic(err)
	}

	return &engine{
		engine: engineDb,
	}
}
