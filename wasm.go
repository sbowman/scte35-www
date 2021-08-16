package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"syscall/js"

	"github.com/Comcast/scte35-go/pkg/scte35"
)

const (
	OutputString = "string"
	OutputJSON   = "json"
	OutputXML    = "xml"
)

var (
	ErrInvalidType = errors.New("expected a Base64 value")
)

func decoderRing() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}

		if args[0].Type() != js.TypeString {
			return ErrInvalidType
		}

		base64 := args[0].String()
		sis, err := scte35.DecodeBase64(base64)
		if err != nil {
			return js.Error{Value: js.ValueOf(err.Error())}
		}

		outputType := OutputString
		if len(args) > 1 && args[1].Type() == js.TypeString {
			outputType = args[1].String()
		}

		switch outputType {
		case OutputString:
			return sis.Table("", "  ")

		case OutputJSON:
			doc, err := json.MarshalIndent(sis, "", "  ")
			if err != nil {
				return js.Error{Value: js.ValueOf(err.Error())}
			}

			return string(doc)

		case OutputXML:
			doc, err := xml.MarshalIndent(sis, "", "  ")
			if err != nil {
				return js.Error{Value: js.ValueOf(err.Error())}
			}

			return string(doc)
		}

		return nil
	})
}

func main() {
	fmt.Println("Initialized SCTE 35 library...")
	js.Global().Set("decode", decoderRing())
	<-make(chan struct{})
}
