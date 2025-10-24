package design

import . "goa.design/goa/v3/dsl"

var JWTAuth = JWTSecurity("jwt", func() {
    Description("Autenticación JWT")
    Scope("api:read", "Read access")
    Scope("api:write", "Write access")
})

var RegisterPayload = Type("RegisterPayload", func() {
    Description("Datos necesarios para registrar un usuario")
    Attribute("email", String, "Correo electrónico del usuario", func() {
        Format(FormatEmail) 
        Example("usuario@ejemplo.com")
    })
    Attribute("password", String, "Contraseña del usuario", func() {
        MinLength(8)
        Example("MiPassword123!")
    })
    Attribute("name", String, "Nombre del usuario", func() {
        MinLength(2)
        Example("Juan Pérez")
    })
    Required("email", "password", "name")
})

var LoginPayload = Type("LoginPayload", func() {
    Description("Datos necesarios para iniciar sesión")
    Attribute("email", String, "Correo electrónico del usuario", func() {
        Format(FormatEmail)
        Example("usuario@ejemplo.com")
    })
    Attribute("password", String, "Contraseña del usuario", func() {
        Example("MiPassword123!")
    })
    Required("email", "password")
})

var TokenResponse = Type("TokenResponse", func() {
    Description("Respuesta de autenticación JWT")
    Attribute("access_token", String, "Token de acceso JWT", func() {
        Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...")
    })
    Attribute("token_type", String, "Tipo de token", func() {
        Default("Bearer")
        Example("Bearer")
    })
    Attribute("expires_in", Int, "Segundos hasta que expire el token", func() {
        Example(3600)
    })
    Required("access_token", "token_type")
})

var _ = Service("users", func() {
    Description("Servicio para registro, login y gestión de usuarios")
    
    Method("register", func() {
        Description("Registrar un nuevo usuario")
        Payload(RegisterPayload)
        Result(TokenResponse)
        
        HTTP(func() {
            POST("/register")
            Response(StatusCreated)
            Response(StatusBadRequest, func() {
                Description("Datos inválidos")
            })
            Response(StatusConflict, func() {
                Description("El usuario ya existe")
            })
            Response(StatusInternalServerError)
        })
    })
    
    Method("login", func() {
        Description("Iniciar sesión y obtener JWT")
        Payload(LoginPayload)
        Result(TokenResponse)
        
        HTTP(func() {
            POST("/login")
            Response(StatusOK)
            Response(StatusBadRequest, func() {
                Description("Datos inválidos")
            })
            Response(StatusUnauthorized, func() {
                Description("Credenciales incorrectas")
            })
            Response(StatusInternalServerError)
        })
    })
    
})