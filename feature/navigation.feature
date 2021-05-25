Feature: Navegación
    Para que pueda acceder a la pagína
    Siendo un usuario cualquiera

    @AboutUs
    Scenario: Navegacion a About Us
        Given accedo a la pagina principal
        When hago click en el menu
        And hago click en el enlace About Us
        Then estoy en la página About Us con texto "Two brothers with the same passion, the same dream: building their careers and enhancing the IT industry."
    
#    @Services
    #Scenario: Navegacion a Services
     #   Given accedo a la pagina principal
      #  When hago click en el menu
       # And hago click en el enlace Services
        #Then estoy en la página Services con texto "From Avalith we offer you support at times where you need it most, with complete support for all your projects."

   #@Careers
   # Scenario: Navegacion a Careers
    #    Given accedo a la pagina principal
    #    When hago click en el menu
     #   And hago click en el enlace careers
      #  Then estoy en la página careers con texto "Improve your skills and boost your genius."

    #@OurPartners
    #Scenario: Navegacion a Our Partners
     #   Given accedo a la pagina principal
        #When hago click en el menu
        #And hago click en el enlace our partners
        #Then estoy en la página Our Partners con texto "Renowned companies around the world are part of our ally network. We have more than 50 Partners who rely on us for software development."

    #@ContactUs
    #Scenario: Navegacion a Contact Us
        #Given accedo a la pagina principal
        #When hago click en el menu
        #And hago click en el enlace Contact Us 
        #Then estoy en la página Contact Us con texto "We want to know more about you!. Write to us at hello@avalith.net or leave us your information and we will get in touch!"

    #@PrivacyPolicy
    #Scenario: Privacy Policy
        #Given accedo a la pagina principal
        #When hago click en privacy policy
        #Then estoy en la pagina de privacy policy con titulo "We collect different types of information for various purposes, including but not limited to providing and improving our Site and services."
    
#    @TermsOfUse
    #Scenario: Terms of use
        #Given accedo a la pagina principal
        #When hago click en terms of use
        #Then estoy en la pagina de terms of use con titulo "Site Rules"

    #@InvalidEmailSubscription
    #Scenario: Email inválido
        #Given accedo a la pagina principal
        #When ingreso un email con formato incorrecto "invalidemail@avalith"
        #And al hacer click en el botón 'Suscribe'
        #Then aparece un mensaje "The email you entered is not valid."
    
    #@ContactUsAlternative1
    #Scenario: Contact Us -  Empty full name
        #Given accedo a la pagina principal
        #When hago click en el menu
        #And hago click en el enlace Contact Us 
        #And ingreso nombre "" e-mail "testing@avalith.net" y mensaje "Este es un mensaje de prueba de al menos 20 caracteres"
        #And al hacer click en Send It
        #Then aparece un mensaje de error "fill in all fields"