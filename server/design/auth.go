package design

import . "goa.design/goa/v3/dsl"

var _ = Service("auth", func() {
	Description("Service that handles Google OAuth2 login and callback")
	Method("googleLogin", func() {
		Description("Redirects the user to the Google OAuth2 consent screen")

		Result(func() {
			Attribute("redirect_url", String, "Google OAuth2 authorization URL")
			Required("redirect_url")
		})

		HTTP(func() {
			GET("/auth/google/login")
			Response(StatusOK)
		})
	})
 
	Method("googleCallback", func() {
		Description("Handles the callback from Google OAuth2 after user consent")

		Payload(func() {
			Attribute("code", String, "Authorization code returned by Google")
			Attribute("state", String, "Optional state parameter used for CSRF protection")
			Required("code")
		})

		Result(func() {
			Attribute("email", String, "User email obtained from Google")
			Attribute("name", String, "User's full name")
			Attribute("picture", String, "Profile picture URL")
			Attribute("token", String, "JWT issued by your backend")
			Required("email", "name", "token")
		})

		HTTP(func() {
			GET("/auth/google/callback")
			Param("code")
			Param("state")
			Response(StatusOK)
		})
	})
})
