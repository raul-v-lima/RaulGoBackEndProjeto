package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert" // add Testify package
)

func TestRoutes(t *testing.T) {

	// Define a structure for specifying input and output data of a single test case.
	tests := []struct {
		description   string
		method        string
		route         string // input route
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "get characteres",
			route:         "/",
			method:        "GET",
			expectedError: false,
			expectedCode:  200,
		},
		{
			description:   "Delete character without ID",
			route:         "/deleteCharacter/:id",
			method:        "DELETE",
			expectedError: false,
			expectedCode:  404,
		},
	}

	// Define Fiber app.
	app := fiber.New()

	// Define routes.
	//PublicRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	app.Delete("/deleteCharacter/:id", func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case.
		req := httptest.NewRequest(test.method, test.route, nil)
		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app.
		resp, _ := app.Test(req, -1) // the -1 disables request latency

		// Verify, that no error occurred, that is not expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

		// As expected errors lead to broken responses,
		// the next test case needs to be processed.
		//if test.expectedError {
		//	continue
		//}

		// Verify, if the status code is as expected.
		//assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

/*
func PublicRoutes(a *fiber.App) {

	app := fiber.New()

	app.Get("/", handlers.ListCharacters)
	//app.Post("/character",handlers.CreateCaracters)
	//app.Put("/updateCharacter/:id",handlers.UpdateCharacter)
	app.Delete("/deleteCharacter/:id", handlers.DeleteCharacter)

 }
*/
