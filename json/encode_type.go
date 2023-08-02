package json

import (
	"errors"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func HandlerType(ctx *Context, kind protoreflect.Kind, gf *protogen.GeneratedFile, mapKey bool, name string) error {
	switch kind {
	case protoreflect.BoolKind:
		Bool(gf, mapKey, name)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Uint64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Fixed32Kind,
		protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
		Integer(gf, mapKey, name)
	case protoreflect.DoubleKind:
		Float(gf, name, 64)
	case protoreflect.FloatKind:
		Float(gf, name, 32)
	case protoreflect.StringKind:
		String(gf, name)
	case protoreflect.BytesKind:
		Bytes(gf, name)
	case protoreflect.EnumKind:
		Enum(gf, name)
	case protoreflect.MessageKind:
		MessageWriteType(ctx, gf, name)
	default:
		return errors.New("not support type " + kind.String())
	}
	return nil
}

func Integer(gf *protogen.GeneratedFile, mapKey bool, name string) {
	if mapKey {
		gf.P(Buf, WriteByte, "('\"')")
	}
	protoimplPackage := protogen.GoImportPath("strconv")
	gf.P(Buf, WriteString, "(", protoimplPackage.Ident("FormatUint"), "(uint64(", name, "),10))")
	if mapKey {
		gf.P(Buf, WriteByte, "('\"')")
	}
}

func Float(gf *protogen.GeneratedFile, name string, bitSize int) {
	protoimplPackage := protogen.GoImportPath("strconv")
	gf.P(Buf, WriteString, "(", protoimplPackage.Ident("FormatFloat"),
		"(float64(", name, "),'f', -1,", bitSize, "))")
}

func String(gf *protogen.GeneratedFile, name string) {
	gf.P(Buf, WriteByte, "('\"')")
	gf.P(Buf, WriteString, "(", name, ")")
	gf.P(Buf, WriteByte, "('\"')")
}

func Bytes(gf *protogen.GeneratedFile, name string) {
	protoimplPackage := protogen.GoImportPath("encoding/base64")
	gf.P(Buf, WriteByte, "('\"')")
	gf.P(Buf, WriteString, "(", protoimplPackage.Ident("StdEncoding"),
		".EncodeToString", "(", name, "))")
	gf.P(Buf, WriteByte, "('\"')")
}

func Enum(gf *protogen.GeneratedFile, name string) {
	gf.P(Buf, WriteByte, "('\"')")
	gf.P(Buf, WriteString, "(", name, ".String())")
	gf.P(Buf, WriteByte, "('\"')")
}

func Bool(gf *protogen.GeneratedFile, mapKey bool, name string) {
	gf.P("if ", name, " {")
	if mapKey {
		gf.P(Buf, WriteString, `("\"true\"")`)
	} else {
		gf.P(Buf, WriteString, `("true")`)
	}
	gf.P("} else {")
	if mapKey {
		gf.P(Buf, WriteString, `("\"false\"")`)
	} else {
		gf.P(Buf, WriteString, `("false")`)
	}
	gf.P("}")
}

func MessageWriteType(ctx *Context, gf *protogen.GeneratedFile, name string) {
	gf.P("if data, err := ", name, ".", ctx.EncodeMethodName, "(); err != nil {")
	gf.P("return nil,err")
	gf.P("} else {")
	gf.P(Buf, WriteBytes, "(data)")
	gf.P("}")
}
