// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package template

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Index() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>HTMX Example</title><script src=\"https://cdn.tailwindcss.com\"></script><script defer src=\"https://cdn.jsdelivr.net/npm/alpinejs@3.14.3/dist/cdn.min.js\"></script><script src=\"https://unpkg.com/htmx.org\"></script></head><body><nav class=\"max-w-7xl mx-auto p-2\"><div class=\"flex justify-between item-center\"><div><h2 class=\"text-2xl font-bold\">Logo</h2></div><div><ul class=\"flex justify-between item-center gap-3 text-lg font-semibold\"><li><a>Home</a></li><li><a href=\"/about\">About</a></li><li><a>Contact</a></li></ul></div><div class=\"flex gap-4 justify-between item-center\"><div><button class=\"bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded\">Sign In </button></div><div><button class=\"bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded\">Sign Up</button></div></div></div></nav><header><div><h2 class=\"text-4xl text-center font-bold\">This is banner section</h2><button x-data=\"{ label: &#39;Click Here&#39; }\" x-text=\"label\"></button> <button x-data @click=\"alert(&#39;I\\&#39;ve been clicked!&#39;)\">Click Me</button></div></header></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
