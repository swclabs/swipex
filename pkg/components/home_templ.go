// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import (
	"github.com/a-h/templ"
	templruntime "github.com/a-h/templ/runtime"
)

func HomeIndex() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html lang=\"en\"><style>\r\n      * {\r\n        box-sizing: border-box;\r\n        margin: 0;\r\n        padding: 0;\r\n      }\r\n\r\n      body,\r\n      html {\r\n        font-family: Arial, sans-serif;\r\n        background-color: #f0f0f0;\r\n        position: relative;\r\n      }\r\n\r\n      .login-container {\r\n        position: absolute;\r\n        top: 50%;\r\n        left: 50%;\r\n        display: flex;\r\n        transform: translate(-50%, -50%);\r\n        background: #fff;\r\n        justify-content: center;\r\n        align-items: center;\r\n        border-radius: 8px;\r\n      }\r\n\r\n      .login-form {\r\n        background: #fff;\r\n        padding: 20px;\r\n        border-radius: 8px;\r\n        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);\r\n        width: 300px;\r\n      }\r\n\r\n      .login-form h2 {\r\n        margin-bottom: 20px;\r\n        text-align: center;\r\n      }\r\n\r\n      .form-group {\r\n        margin-bottom: 15px;\r\n      }\r\n\r\n      .form-group label {\r\n        display: block;\r\n        margin-bottom: 5px;\r\n      }\r\n\r\n      .form-group input {\r\n        width: 100%;\r\n        padding: 10px;\r\n        border: 1px solid hsl(0, 0%, 80%);\r\n        border-radius: 4px;\r\n      }\r\n\r\n      .logo {\r\n        display: flex;\r\n        gap: 10px;\r\n        justify-content: center;\r\n        width: 300px;\r\n      }\r\n\r\n      button {\r\n        width: 100%;\r\n        padding: 10px;\r\n        background-color: #007bff;\r\n        border: none;\r\n        border-radius: 4px;\r\n        color: #fff;\r\n        font-size: 16px;\r\n        cursor: pointer;\r\n      }\r\n\r\n      button:hover {\r\n        background-color: #0056b3;\r\n      }\r\n\r\n      .text {\r\n        padding: 10px;\r\n        text-align: center;\r\n        font-size: small;\r\n      }\r\n    </style><body><div class=\"login-container\"><div><div class=\"logo\"><img src=\"https://raw.githubusercontent.com/swclabs/swipex/main/logo/hcmut.png\" height=\"40px\"> <img src=\"https://raw.githubusercontent.com/swclabs/swipex/79ca340c27b51ad91e5b8d23d910b9cd5b9c905d/logo/logo.svg\" height=\"40px\"></div><p class=\"text\">Designed for the final <br>thesis at HCMUT-VNUHCM</p></div><form class=\"login-form\" method=\"POST\" action=\"/auth\"><div class=\"form-group\"><label for=\"email\">Email</label> <input type=\"text\" name=\"email\" required></div><div class=\"form-group\"><label for=\"password\">Password</label> <input type=\"password\" name=\"password\" required></div><button type=\"submit\">Login</button></form></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
