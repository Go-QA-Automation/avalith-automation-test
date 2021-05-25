package support

import (
	"log"

	"github.com/tebeka/selenium"
)

var (
	driver           selenium.WebDriver
	chromeDriverPath = `C:\Users\jonat\go\src\github.com\jonathanbs9\go-selenium-avalith-web-test\resources\chromedriver.exe`
	seleniumPath     = `C:\Users\jonat\go\src\github.com\jonathanbs9\go-selenium-avalith-web-test\resources\selenium-server-standalone-3.141.59.jar`
	port             = 4444
)

func WDInit() selenium.WebDriver {
	var err error

	ops := []selenium.ServiceOption{
		selenium.ChromeDriver(seleniumPath),
	}

	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, ops...)
	if err != nil {
		log.Println("Error al comenzar el servidor ChromeDriver ...", err.Error())
	}
	defer service.Stop()

	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Println("(support/base) | Error al instanciar driver Selenium -> ", err.Error())
	}
	driver.MaximizeWindow("")

	return driver
}
