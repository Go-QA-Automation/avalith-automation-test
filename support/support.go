package support

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/tebeka/selenium"
)

type WDriver struct {
	chromeDriverPath string
	seleniumPath     string
	port             int
}

type IDriver interface {
	WDInit() selenium.WebDriver
}

func getConfig() (chromeDriverPath, seleniumPath string, err error) {
	currDir, err := os.Getwd()
	if err != nil {
		return chromeDriverPath, seleniumPath, err
	}
	if runtime.GOOS == "windows" {
		chromeDriverPath = path.Join(currDir, "resources\\chromedriver.exe")
		seleniumPath = path.Join(currDir, "resources\\selenium-server-standalone-3.141.59.jar")
	} else {
		chromeDriverPath = path.Join(currDir, "resources/chromedriver")
		seleniumPath = path.Join(currDir, "resources/selenium-server-standalone-3.141.59.jar")
	}
	return chromeDriverPath, seleniumPath, nil
}

func New() IDriver {
	chPath, sPath, _ := getConfig()
	return &WDriver{
		chromeDriverPath: chPath,
		seleniumPath:     sPath,
		port:             4444,
	}
}

func (d *WDriver) WDInit() (driver selenium.WebDriver) {
	ops := []selenium.ServiceOption{
		selenium.ChromeDriver(d.seleniumPath),
	}
	service, err := selenium.NewChromeDriverService(d.chromeDriverPath, d.port, ops...)
	if err != nil {
		log.Println("Error al comenzar el servidor ChromeDriver ...", err.Error())
	}
	defer service.Stop()

	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	driver, err = selenium.NewRemote(caps, "")
	if err != nil {
		log.Println("(support/base) | Error al instanciar driver Selenium -> ", err.Error())
	}
	driver.MaximizeWindow("")

	return driver
}
