package main

import (
	"syscall/js"
)

func main() {
	// Set up the DOM on load
	document := js.Global().Get("document")
	setupHeader(document)
	setupForm(document)

	// Keep the WASM running
	select {}
}

func setupHeader(document js.Value) {
	// Add custom font
	head := document.Get("head")
	style := document.Call("createElement", "style")
	style.Set("innerHTML", `
		@import url('https://fonts.googleapis.com/css2?family=Playfair+Display:wght@400;700&family=Open+Sans:wght@300;400;600&display=swap');
		body {
			font-family: 'Open Sans', sans-serif;
			background-color: #f8f9fa;
			color: #333;
			margin: 0;
		}
		header {
			background-color: #002147;
			color: white;
			padding: 20px;
			text-align: center;
			font-family: 'Playfair Display', serif;
		}
		header h1 {
			margin: 0;
			font-size: 2.5rem;
		}
		form {
			background: white;
			padding: 20px;
			border-radius: 8px;
			box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
			max-width: 400px;
			margin: 20px auto;
		}
		label {
			display: block;
			margin-bottom: 15px;
			font-weight: 600;
		}
		input {
			width: 100%;
			padding: 10px;
			margin-top: 5px;
			border: 1px solid #ccc;
			border-radius: 4px;
		}
		button {
			background-color: #002147;
			color: white;
			border: none;
			padding: 10px 20px;
			border-radius: 4px;
			cursor: pointer;
			font-weight: 600;
		}
		button:hover {
			background-color: #01497c;
		}
	`)
	head.Call("appendChild", style)

	// Create the header
	header := document.Call("createElement", "header")
	h1 := document.Call("createElement", "h1")
	h1.Set("innerHTML", " Cruise Lines")
	header.Call("appendChild", h1)
	document.Get("body").Call("appendChild", header)
}

func setupForm(document js.Value) {
	// Create the form
	form := document.Call("createElement", "form")
	form.Set("className", "submission-form")

	// Add input fields
	addInputField(document, form, "Cabin Number", "cabinNumber", "text")
	addInputField(document, form, "First Name", "firstName", "text")
	addInputField(document, form, "Last Name", "lastName", "text")
	addInputField(document, form, "Date of Birth", "dateOfBirth", "date")

	// Add submit button
	button := document.Call("createElement", "button")
	button.Set("innerHTML", "Submit")
	button.Set("type", "submit")
	form.Call("appendChild", button)

	// Add form submit event listener
	form.Call("addEventListener", "submit", js.FuncOf(handleFormSubmit))

	// Append the form to the body
	document.Get("body").Call("appendChild", form)
}

func addInputField(document js.Value, form js.Value, labelText, name, inputType string) {
	label := document.Call("createElement", "label")
	label.Set("innerHTML", labelText)
	input := document.Call("createElement", "input")
	input.Set("name", name)
	input.Set("type", inputType)
	label.Call("appendChild", input)
	form.Call("appendChild", label)
}

func handleFormSubmit(this js.Value, args []js.Value) interface{} {
	event := args[0]
	event.Call("preventDefault") // Prevent default form submission

	// Retrieve form values
	form := event.Get("target")
	cabinNumber := form.Call("querySelector", "input[name='cabinNumber']").Get("value").String()
	firstName := form.Call("querySelector", "input[name='firstName']").Get("value").String()
	lastName := form.Call("querySelector", "input[name='lastName']").Get("value").String()
	dateOfBirth := form.Call("querySelector", "input[name='dateOfBirth']").Get("value").String()

	// Log the submitted values
	js.Global().Call("alert", "Form Submitted!\nCabin Number: "+cabinNumber+
		"\nFirst Name: "+firstName+
		"\nLast Name: "+lastName+
		"\nDate of Birth: "+dateOfBirth)

	// submit the form post request
	js.Global().Call("fetch", "http://localhost:8080", map[string]interface{}{
		"method": "POST",
		"body": js.Global().Get("JSON").Call("stringify", map[string]interface{}{
			"cabinNumber": cabinNumber,
			"firstName":   firstName,
			"lastName":    lastName,
			"dateOfBirth": dateOfBirth,
		}),
		"headers": map[string]interface{}{
			"Content-Type": "application/json",
		},
	}).Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		response := args[0]
		if !response.Get("ok").Bool() {
			js.Global().Call("alert", "Failed to submit form")
			return nil
		}
		response.Call("json").Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			js.Global().Call("alert", "Form submitted successfully!")
			return nil
		}))
		return nil
	})).Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		error := args[0]
		js.Global().Call("alert", "Failed to submit form: "+error.Get("message").String())
		return nil
	}))

	return nil
}
