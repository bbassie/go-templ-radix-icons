// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/bbassie/go-templ-radix-icons/colors"

func DragHandleVerticalIcon(color ...string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"15\" height=\"15\" viewBox=\"0 0 15 15\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M4.09998 12.5C4.09998 12.7209 4.27906 12.9 4.49998 12.9C4.72089 12.9 4.89998 12.7209 4.89998 12.5L4.89998 2.50002C4.89998 2.27911 4.72089 2.10003 4.49998 2.10003C4.27906 2.10003 4.09998 2.27911 4.09998 2.50002L4.09998 12.5ZM6.09998 12.5C6.09998 12.7209 6.27906 12.9 6.49998 12.9C6.72089 12.9 6.89998 12.7209 6.89998 12.5L6.89998 2.50002C6.89998 2.27911 6.72089 2.10003 6.49998 2.10003C6.27906 2.10003 6.09998 2.27911 6.09998 2.50002L6.09998 12.5ZM8.49998 12.9C8.27906 12.9 8.09998 12.7209 8.09998 12.5L8.09998 2.50002C8.09998 2.27911 8.27906 2.10002 8.49998 2.10002C8.72089 2.10002 8.89998 2.27911 8.89998 2.50002L8.89998 12.5C8.89998 12.7209 8.72089 12.9 8.49998 12.9ZM10.1 12.5C10.1 12.7209 10.2791 12.9 10.5 12.9C10.7209 12.9 10.9 12.7209 10.9 12.5L10.9 2.50002C10.9 2.27911 10.7209 2.10002 10.5 2.10002C10.2791 2.10002 10.1 2.27911 10.1 2.50002L10.1 12.5Z\" fill=\"")
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