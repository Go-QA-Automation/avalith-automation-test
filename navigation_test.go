package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Go-QA-Automation/avalith-automation-test/support"
	"github.com/cucumber/godog"
	"github.com/tebeka/selenium"
)

var (
	Driver selenium.WebDriver
)

// Main page
func accedoALaPaginaPrincipal() error {
	log.Println("Accedo a página  de Avalith.net ")
	Driver.Get("https://avalith-iota.vercel.app/")

	return nil
}

// Click right menu
func hagoClickEnElMenu() error {
	divMenu, err := Driver.FindElement(selenium.ByCSSSelector, "#__next > div > div > header > div > button")
	if err != nil {
		log.Println("Error al traer el menu -> ", err.Error())
	}
	divMenu.Click()
	Driver.SetImplicitWaitTimeout(time.Second * 6)
	return nil
}

// Click on 'About Us'
func hagoClickEnElEnlaceAboutUs() error {
	aboutUs5, err := Driver.FindElement(selenium.ByCSSSelector, "#__next > div > div > div > div > ul > li:nth-child(1) > button")

	if err != nil {
		log.Println("Error al hacer click en About Us", err.Error())
	}
	Driver.SetImplicitWaitTimeout(time.Second * 10)
	aboutUs5.Click()

	return nil
}

// Click on 'Services'
func hagoClickEnElEnlaceServices() error {
	//serviceXPath := `//*[@id="gatsby-focus-wrapper"]/div/div/div/div/ul/li[2]/a`
	service, err := Driver.FindElement(selenium.ByCSSSelector, "#__next > div > div > div > div > ul > li:nth-child(2) > button")

	if err != nil {
		log.Println("Error al hacer click en Service", err.Error())
	}
	Driver.SetImplicitWaitTimeout(time.Second * 10)
	service.Click()

	return nil
}

// Click ok 'Careers'
func hagoClickEnElEnlaceCareers() error {
	//careerXPATH := `//*[@id="gatsby-focus-wrapper"]/div/div/div/div/ul/li[3]/a`
	career, err := Driver.FindElement(selenium.ByCSSSelector, "#__next > div > div > div > div > ul > li:nth-child(3) > button")
	if err != nil {
		log.Println("Error al hacer click en Career", err.Error())
	}
	Driver.SetImplicitWaitTimeout(time.Second * 10)
	career.Click()
	return nil
}

// Click on 'Our Partners'
func hagoClickEnElEnlaceOurPartners() error {
	//ourPartnerXPATH := `//*[@id="gatsby-focus-wrapper"]/div/div/div/div/ul/li[4]/a`
	ourPartner, err := Driver.FindElement(selenium.ByCSSSelector, "#__next > div > div > div > div > ul > li:nth-child(4) > button")
	if err != nil {
		log.Println("Error al hacer click en OurPartners", err.Error())
	}
	Driver.SetImplicitWaitTimeout(time.Second * 10)
	ourPartner.Click()
	Driver.SetImplicitWaitTimeout(time.Second * 12)
	return nil
}

// Click on 'Contact Us'
func hagoClickEnElEnlaceContactUs() error {
	//contactUsPath := `//*[@id="gatsby-focus-wrapper"]/div/div/div/div/ul/li[5]/a`
	contactUs, err := Driver.FindElement(selenium.ByCSSSelector, "#__next > div > div > div > div > ul > li.pt-12.pb-8 > button")
	if err != nil {
		log.Println("Error al obtener enlance a Contact Us | ", err.Error())
	}
	contactUs.Click()
	Driver.SetImplicitWaitTimeout(time.Second * 10)
	return nil
}

// Click on 'Privacy Policy'
func hagoClickEnPrivacyPolicy() error {
	privacyLink, err := Driver.FindElement(selenium.ByCSSSelector, `#__next > div > footer > div.grid.grid-rows-1.grid-cols-1.md\:grid-cols-5.pb-12.gap-8.text-center > button:nth-child(4)`)
	if err != nil {
		return fmt.Errorf("Error al obtener enlance a Privacy Policy | %s", err.Error())
	}

	privacyLink.Click()
	//	Driver.SetImplicitWaitTimeout(time.Second * 100)
	return nil
}

// Click ok 'Terms of Use'
func hagoClickEnTermsOfUse() error {
	termsLink, err := Driver.FindElement(selenium.ByCSSSelector, `#__next > div > footer > div.grid.grid-rows-1.grid-cols-1.md\:grid-cols-5.pb-12.gap-8.text-center > button:nth-child(5)`)
	if err != nil {
		return fmt.Errorf("Error al obtener enlance a Terms of Use | %s", err.Error())
	}
	Driver.SetImplicitWaitTimeout(time.Second * 5)
	termsLink.Click()
	Driver.SetImplicitWaitTimeout(time.Second * 10)
	return nil
}

// Positioned in 'About us' page
func estoyEnLaPginaAboutUsConTexto(text string) error {
	Driver.SetImplicitWaitTimeout(time.Second * 10)
	aboutUsdiv, err := Driver.FindElement(selenium.ByCSSSelector, `#__next > div > main > div.container.mx-auto.md\:mt-20.px-6 > section.flex.flex-col-reverse.lg\:flex-row.items-center > div:nth-child(1) > h3`)
	if err != nil {
		log.Println(err)
	}
	aboutUsText, _ := aboutUsdiv.Text()
	if aboutUsText != text {
		log.Fatal("No son iguales los textos")
		log.Fatalf("mensaje obtenido : %s  | Mensaje esperado: %s", aboutUsText, text)
	}

	return nil
}

// Positioned in 'Services' page
func estoyEnLaPginaServicesConTexto(textServicePage string) error {
	servicesdiv, err := Driver.FindElement(selenium.ByCSSSelector, `#__next > div > main > section.container.mx-auto.px-6.py-8 > div.flex.flex-col.md\:flex-row.items-center.pb-8 > div.w-full.md\:w-1\/2.md\:px-8 > p`)
	if err != nil {
		log.Println(err)
	}
	servicesText, _ := servicesdiv.Text()
	if servicesText != textServicePage {
		log.Fatalf("mensaje obtenido : %s  | Mensaje esperado: %s", textServicePage, servicesText)
	}

	return nil
}

// Positioned in 'Careers' page
func estoyEnLaPginaCareersConTexto(careerText string) error {
	careerElement, err := Driver.FindElement(selenium.ByCSSSelector, `#__next > div > main > section.container.mx-auto.md\:mt-20.px-6.pb-20 > div.flex.flex-col.md\:flex-row.items-center.pb-20 > div.w-full.md\:w-1\/2.md\:px-8 > h5`)
	if err != nil {
		log.Println(err)
	}
	careerElementText, _ := careerElement.Text()
	if careerElementText != careerText {
		log.Fatalf("mensaje obtenido : %s  | Mensaje esperado: %s", careerElementText, careerText)
	}

	return nil
}

// Positioned in 'Our Partners' page
func estoyEnLaPginaOurPartnersConTexto(ourPartnerText string) error {
	Driver.SetImplicitWaitTimeout(time.Second * 10)
	ourPartnerElement, err := Driver.FindElement(selenium.ByCSSSelector, `#__next > div > main > section.container.mx-auto.px-6.py-12 > p`)
	if err != nil {
		log.Println("Error al obtener elemento dentro de la página Our Partners | ", err.Error())
	}
	ourPartnerElementText, _ := ourPartnerElement.Text()
	if ourPartnerElementText != ourPartnerText {
		log.Fatalf("mensaje obtenido : %v  | Mensaje esperado: %v", ourPartnerElementText, ourPartnerText)
	}
	return nil
}

// Positioned in 'Contact us' page
func estoyEnLaPginaContactUsConTexto(contactUsText string) error {
	Driver.SetImplicitWaitTimeout(time.Second * 10)
	contactUsElement, err := Driver.FindElement(selenium.ByCSSSelector, `#__next > div > main > section > div.flex.flex-col.md\:flex-row.items-center.pb-12 > div > p`)
	if err != nil {
		log.Println("Error al obtener elemento Contact Us | ", err.Error())
	}
	contactUsElementText, _ := contactUsElement.Text()

	if contactUsElementText != contactUsText {
		return fmt.Errorf("mensaje obtenido : %s  | Mensaje esperado: %s", contactUsElementText, contactUsText)
	}
	return nil
}

// Positioned in 'Privacy Policy' page
func estoyEnLaPaginaDePrivacyPolicyConTitulo(title string) error {
	Driver.SetImplicitWaitTimeout(time.Second * 100)
	privacyTitle, err := Driver.FindElement(selenium.ByCSSSelector, `#__next > div > main > section > p:nth-child(7)`)
	if err != nil {
		return fmt.Errorf("Error al obtener elemento Privacy Policy | %s", err.Error())
	}
	log.Println(privacyTitle.Text())
	privacyTitleText, _ := privacyTitle.Text()
	if privacyTitleText != title {
		return fmt.Errorf("Titulo obtenido : %s  | Título esperado: %s", privacyTitleText, title)
	}
	return nil
}

// Positioned in 'Terms of use' page
func estoyEnLaPaginaDeTermsOfUseConTitulo(text string) error {
	termsOfUse, err := Driver.FindElement(selenium.ByCSSSelector, `#__next > div > main > section > h4:nth-child(6)`)
	if err != nil {
		return fmt.Errorf("Error al obtener el elemento Terms of Use | %s", err.Error())
	}
	log.Println(termsOfUse.Text())
	termsTitle, _ := termsOfUse.Text()
	if termsTitle != text {
		return fmt.Errorf("Titulo obtenido : %s | Titulo esperado: %s", termsTitle, text)
	}
	return nil
}

// Enter and invalid email
func ingresoUnEmailConFormatoIncorrecto(invalidEmail string) error {
	campoEmail, err := Driver.FindElement(selenium.ByCSSSelector, `#__next > div > footer > div.grid.grid-rows-1.grid-cols-1.md\:grid-cols-1.lg\:grid-cols-3.pb-12.gap-12.sm\:gap-8.text-center > div:nth-child(2) > form > input`)
	if err != nil {
		log.Println("Error al encontrar input ")
	}
	Driver.SetImplicitWaitTimeout(time.Second * 5)
	campoEmail.SendKeys(invalidEmail)

	return nil
}

// Click on 'Subscribe' button
func alHacerClickEnElBotnSuscribe() error {
	buttonSubscribe, err := Driver.FindElement(selenium.ByCSSSelector, `#__next > div > footer > div.grid.grid-rows-1.grid-cols-1.md\:grid-cols-1.lg\:grid-cols-3.pb-12.gap-12.sm\:gap-8.text-center > div:nth-child(2) > form > button`)
	if err != nil {
		log.Println("Error al buscar botón de suscripción")
	}
	buttonSubscribe.Click()

	return nil
}

// A message pops up
func apareceUnMensaje(msg string) error {
	messageAlert, err := Driver.FindElement(selenium.ByCSSSelector, `#__next > div > main > section > div.flex.flex-col.md\:flex-row.items-center.pb-12 > div > div`)
	if err != nil {
		log.Println("Error al obtener mensaje de alerta")
	}

	alerta, err := messageAlert.Text()
	if err != nil {
		log.Println("Error al recibir el texto del alerta")
	}

	if alerta != msg {
		return fmt.Errorf("El mensaje obtenido no es el mensaje esperado")
	}

	return nil
}

func ingresoNombreEmailYMensaje(fullName, email, msg string) error {
	nameDiv, err := Driver.FindElement(selenium.ByXPATH, `/html/body/div[1]/div/main/section/div[1]/div/form/input[1]`)
	if err != nil {
		log.Println("Error al buscar campo 'Full name'")
	}
	emailDiv, err := Driver.FindElement(selenium.ByXPATH, `//*[@id="email"]`)
	if err != nil {
		log.Println("Error al buscar campo 'email'")
	}
	msgDiv, err := Driver.FindElement(selenium.ByXPATH, `//*[@id="mensaje"]`)
	if err != nil {
		log.Println("Error al buscar el campo 'your message'")
	}
	/*captchaDiv, err := Driver.FindElement(selenium.ByCSSSelector, `#recaptcha-anchor > div.recaptcha-checkbox-checkmark`)
	if err != nil {
		log.Println("Error al buscar el captcha")
	}*/

	nameDiv.SendKeys(fullName)
	emailDiv.SendKeys(email)
	msgDiv.SendKeys(msg)
	//captchaDiv.Click()

	return nil
}

func alHacerClickEnSendIt() error {
	SendItDiv, err := Driver.FindElement(selenium.ByCSSSelector, `#__next > div > main > section > div.flex.flex-col.md\:flex-row.items-center.pb-12 > div > form > button`)
	if err != nil {
		log.Println("Error al buscar boton 'Send it' ")
	}
	SendItDiv.Click()
	return nil
}

func apareceUnMensajeDeError(errorMessage string) error {
	messageDiv, err := Driver.FindElement(selenium.ByCSSSelector, `#__next > div > main > section > div.flex.flex-col.md\:flex-row.items-center.pb-12 > div > div > div > p`)
	if err != nil {
		log.Println("Error al buscar el mensaje de error")
	}
	message, err := messageDiv.Text()
	if err != nil {
		log.Print("No existe cuerpo de mensaje")
	}
	if message != errorMessage {
		log.Println("Respuesa esperada distinta de respuesta obtenida")
	}
	return nil
}

// Init Scenario
func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
		Driver = support.WDInit()
		Driver.SetImplicitWaitTimeout(8 * time.Second)
	})

	ctx.Step(`^accedo a la pagina principal$`, accedoALaPaginaPrincipal)
	ctx.Step(`^hago click en el menu$`, hagoClickEnElMenu)

	// About Us - Services - Careers - Our Patners - Contact Us - Privacy Policy - Terms of Use

	ctx.Step(`^hago click en el enlace About Us$`, hagoClickEnElEnlaceAboutUs)
	ctx.Step(`^hago click en el enlace Services$`, hagoClickEnElEnlaceServices)
	ctx.Step(`^hago click en el enlace careers$`, hagoClickEnElEnlaceCareers)
	ctx.Step(`^hago click en el enlace our partners$`, hagoClickEnElEnlaceOurPartners)
	ctx.Step(`^hago click en el enlace Contact Us$`, hagoClickEnElEnlaceContactUs)
	ctx.Step(`^hago click en privacy policy$`, hagoClickEnPrivacyPolicy)
	ctx.Step(`^hago click en terms of use$`, hagoClickEnTermsOfUse)

	ctx.Step(`^estoy en la página About Us con texto "([^"]*)"$`, estoyEnLaPginaAboutUsConTexto)
	ctx.Step(`^estoy en la página Services con texto "([^"]*)"$`, estoyEnLaPginaServicesConTexto)
	ctx.Step(`^estoy en la página careers con texto "([^"]*)"$`, estoyEnLaPginaCareersConTexto)
	ctx.Step(`^estoy en la página Our Partners con texto "([^"]*)"$`, estoyEnLaPginaOurPartnersConTexto)
	ctx.Step(`^estoy en la página Contact Us con texto "([^"]*)"$`, estoyEnLaPginaContactUsConTexto)
	ctx.Step(`^estoy en la pagina de privacy policy con titulo "([^"]*)"$`, estoyEnLaPaginaDePrivacyPolicyConTitulo)
	ctx.Step(`^estoy en la pagina de terms of use con titulo "([^"]*)"$`, estoyEnLaPaginaDeTermsOfUseConTitulo)

	// Invalid Email
	ctx.Step(`^ingreso un email con formato incorrecto "([^"]*)"$`, ingresoUnEmailConFormatoIncorrecto)
	ctx.Step(`^al hacer click en el botón \'Suscribe\'$`, alHacerClickEnElBotnSuscribe)
	ctx.Step(`^aparece un mensaje "([^"]*)"$`, apareceUnMensaje)
	ctx.Step(`^aparece un mensaje de error "([^"]*)"$`, apareceUnMensajeDeError)

	//ctx.Step(`^accedo a la pagina Contact us$`, accedoALaPaginaContactUs)
	ctx.Step(`^al hacer click en Send It$`, alHacerClickEnSendIt)
	ctx.Step(`^ingreso nombre "([^"]*)" e-mail "([^"]*)" y mensaje "([^"]*)"$`, ingresoNombreEmailYMensaje)

	ctx.AfterScenario(func(sc *godog.Scenario, err error) {
		log.Println("Closing Chrome!")
		Driver.SetImplicitWaitTimeout(20 * time.Second)
		Driver.Quit()
	})
}
