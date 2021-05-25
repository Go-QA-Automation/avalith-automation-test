package support

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/tebeka/selenium"
)

var (
	driver           selenium.WebDriver
	chromeDriverPath string
	seleniumPath     string
	port             = 4444
)

func WDInit() selenium.WebDriver {
	var err error

	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("Error con el directorio ", err.Error())
	}

	if runtime.GOOS == "windows" {
		chromeDriverPath = path.Join(currentDir, "resources\\chromedriver.exe")
		seleniumPath = path.Join(currentDir, "resources\\selenium-server-standalone-3.141.59.jar")
	} else {
		chromeDriverPath = path.Join(currentDir, "resources/chromedriver")
		seleniumPath = path.Join(currentDir, "resources/selenium-server-standalone-3.141.59.jar")
	}

	ops := []selenium.ServiceOption{
		selenium.ChromeDriver(seleniumPath),
	}

	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, ops...)
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
