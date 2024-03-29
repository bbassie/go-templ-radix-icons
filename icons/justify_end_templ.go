// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/bbassie/go-templ-radix-icons/colors"

func JustifyEndIcon(color ...string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"15\" height=\"15\" viewBox=\"0 0 15 15\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M13.95 1.49953C13.95 1.251 13.7485 1.04953 13.5 1.04953C13.2514 1.04953 13.05 1.251 13.05 1.49953L13.05 13.4995C13.05 13.7481 13.2514 13.9495 13.5 13.9495C13.7485 13.9495 13.95 13.7481 13.95 13.4995L13.95 1.49953ZM3.99997 5.99997L11 5.99997L11 8.99997L3.99997 8.99997L3.99997 5.99997ZM11.25 4.99997C11.6642 4.99997 12 5.33576 12 5.74997L12 9.24997C12 9.66418 11.6642 9.99997 11.25 9.99997L3.74997 9.99997C3.33576 9.99997 2.99997 9.66418 2.99997 9.24997L2.99997 5.74997C2.99998 5.33576 3.33576 4.99997 3.74998 4.99997L11.25 4.99997Z\" fill=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(colors.ColorsOrFallback(color...)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
