// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/bbassie/go-templ-radix-icons/colors"

func HobbyKnifeIcon(color ...string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"15\" height=\"15\" viewBox=\"0 0 15 15\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M12.3536 13.3536C12.1583 13.5488 11.8417 13.5488 11.6465 13.3536L6.39645 8.10355C6.36478 8.07188 6.33824 8.03702 6.31685 8H5.00002C4.78719 8 4.59769 7.86528 4.52777 7.66426L2.12777 0.764277C2.05268 0.548387 2.13355 0.309061 2.3242 0.182972C2.51486 0.0568819 2.76674 0.0761337 2.93602 0.229734L8.336 5.12972C8.44044 5.22449 8.50001 5.35897 8.50001 5.5V5.81684C8.53702 5.83824 8.57189 5.86478 8.60356 5.89645L13.8536 11.1464C14.0488 11.3417 14.0488 11.6583 13.8536 11.8536L12.3536 13.3536ZM8.25 6.95711L7.45711 7.75L12 12.2929L12.7929 11.5L8.25 6.95711ZM3.71669 2.28845L5.35549 7H6.2929L7.50001 5.79289V5.72146L3.71669 2.28845Z\" fill=\"")
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