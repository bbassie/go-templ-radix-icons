// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/bbassie/go-templ-radix-icons/colors"

func SpaceEvenlyHorizontallyIcon(color ...string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"15\" height=\"15\" viewBox=\"0 0 15 15\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M14.4999 0.999992C14.2237 0.999992 13.9999 1.22385 13.9999 1.49999L13.9999 13.4999C13.9999 13.776 14.2237 13.9999 14.4999 13.9999C14.776 13.9999 14.9999 13.776 14.9999 13.4999L14.9999 1.49999C14.9999 1.22385 14.776 0.999992 14.4999 0.999992ZM0.499996 0.999992C0.223856 0.999992 -9.78509e-09 1.22385 -2.18556e-08 1.49999L4.07279e-07 13.4999C3.95208e-07 13.776 0.223855 13.9999 0.499996 13.9999C0.776136 13.9999 0.999992 13.776 0.999992 13.4999L0.999992 1.49999C0.999992 1.22385 0.776136 0.999992 0.499996 0.999992ZM1.99998 6.99994C1.99998 6.44766 2.44769 5.99995 2.99998 5.99995L5.99995 5.99995C6.55223 5.99995 6.99994 6.44766 6.99994 6.99994L6.99994 7.99993C6.99994 8.55221 6.55223 8.99992 5.99995 8.99992L2.99998 8.99992C2.4477 8.99992 1.99998 8.55221 1.99998 7.99993L1.99998 6.99994ZM8.99993 5.99995C8.44765 5.99995 7.99993 6.44766 7.99993 6.99994L7.99993 7.99993C7.99993 8.55221 8.44765 8.99992 8.99993 8.99992L11.9999 8.99992C12.5522 8.99992 12.9999 8.55221 12.9999 7.99993L12.9999 6.99994C12.9999 6.44766 12.5522 5.99995 11.9999 5.99995L8.99993 5.99995Z\" fill=\"")
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