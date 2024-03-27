package main

import (
	"log"
	"testing"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func TestSignInSelenium(t *testing.T) {
	// initialize a Chrome browser instance on port 4444
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)

	if err != nil {
		log.Fatal("Error:", err)
	}

	defer service.Stop()

	// configure the browser options

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		//"--headless", // comment out this line for testing
	}})

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// visit the target page
	err = driver.Get("https://bas-meshit-astana-e930fe4b68dc.herokuapp.com/login")
	if err != nil {
		log.Fatal("Error:", err)
	}

	formElement, err := driver.FindElement(selenium.ByID, "form")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// fill in the login form fields
	emailElement, err := formElement.FindElement(selenium.ByName, "email")
	if err != nil {
		// Handle error
	}

	emailElement.SendKeys("fafsdfas04@gmail.com")

	passwdElement, err := formElement.FindElement(selenium.ByName, "passwd")
	if err != nil {
		// Handle error
	}

	passwdElement.SendKeys("qwerty_1")

	// submit the form
	formElement.Submit()

	_, err = driver.GetCookies()
	if err != nil {
		t.Errorf("Expected status che tam, %v", err)
		return
	}
}
