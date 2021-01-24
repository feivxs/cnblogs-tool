package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/feivxs/cbtool/common"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)

var client = http.Client{}

func Upload(request common.UploadRequest) (string, error) {
	encodeRequest := buildRequestBody(request)
	//bodyBytes, _ := ioutil.ReadAll(encodeRequest)
	//fmt.Println("PostBody:\n" + string(bodyBytes))
	resp, _ := client.Post(GetConfig().MetaWebBlogUrl, "text/html", encodeRequest)
	respBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response: \n" + string(respBytes))
	return string(respBytes), nil
}

func buildRequestBody(request common.UploadRequest) *bytes.Buffer {
	buf := new(bytes.Buffer)

	buf.WriteString(`<?xml version="1.0"?><methodCall>`)
	buf.WriteString("<methodName>" + xmlEscape("metaWeblog.newMediaObject") + "</methodName>")
	buf.WriteString("<params>")

	buf.WriteString("<param><value><string>" + request.BlogId + "</string></value></param>")
	buf.WriteString("<param><value><string>" + request.Username + "</string></value></param>")
	buf.WriteString("<param><value><string>" + request.Password + "</string></value></param>")
	buf.WriteString("<param><value>" + toXml(request.File, true) + "</value></param>")

	buf.WriteString("</params></methodCall>")

	return buf
}

func xmlEscape(s string) string {
	var b bytes.Buffer
	for i := 0; i < len(s); i++ {
		c := s[i]
		if s, ok := xmlSpecial[c]; ok {
			b.WriteString(s)
		} else {
			b.WriteByte(c)
		}
	}
	return b.String()
}

var xmlSpecial = map[byte]string{
	'<':  "&lt;",
	'>':  "&gt;",
	'"':  "&quot;",
	'\'': "&apos;",
	'&':  "&amp;",
}

func toXml(v interface{}, typ bool) (s string) {
	if v == nil {
		return "<nil/>"
	}
	r := reflect.ValueOf(v)
	t := r.Type()
	k := t.Kind()

	if b, ok := v.([]byte); ok {
		return "<base64>" + base64.StdEncoding.EncodeToString(b) + "</base64>"
	}

	switch k {
	case reflect.Invalid:
		panic("unsupported type")
	case reflect.Bool:
		return fmt.Sprintf("<boolean>%v</boolean>", v)
	case reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if typ {
			return fmt.Sprintf("<int>%v</int>", v)
		}
		return fmt.Sprintf("%v", v)
	case reflect.Uintptr:
		panic("unsupported type")
	case reflect.Float32, reflect.Float64:
		if typ {
			return fmt.Sprintf("<double>%v</double>", v)
		}
		return fmt.Sprintf("%v", v)
	case reflect.Complex64, reflect.Complex128:
		panic("unsupported type")
	case reflect.Array:
		s = "<array><data>"
		for n := 0; n < r.Len(); n++ {
			s += "<value>"
			s += toXml(r.Index(n).Interface(), typ)
			s += "</value>"
		}
		s += "</data></array>"
		return s
	case reflect.Chan:
		panic("unsupported type")
	case reflect.Func:
		panic("unsupported type")
	case reflect.Interface:
		return toXml(r.Elem(), typ)
	case reflect.Map:
		s = "<struct>"
		for _, key := range r.MapKeys() {
			s += "<member>"
			s += "<name>" + xmlEscape(key.Interface().(string)) + "</name>"
			s += "<value>" + toXml(r.MapIndex(key).Interface(), typ) + "</value>"
			s += "</member>"
		}
		s += "</struct>"
		return s
	case reflect.Ptr:
		panic("unsupported type")
	case reflect.Slice:
		s = "<array><data>"
		for n := 0; n < r.Len(); n++ {
			s += "<value>"
			s += toXml(r.Index(n).Interface(), typ)
			s += "</value>"
		}
		s += "</data></array>"
		return s
	case reflect.String:
		if typ {
			return fmt.Sprintf("<string>%v</string>", xmlEscape(v.(string)))
		}
		return xmlEscape(v.(string))
	case reflect.Struct:
		if t.Name() == "Time" {
			return fmt.Sprintf("<dateTime.iso8601>%v</dateTime.iso8601>", v.(time.Time).String())
		}

		s = "<struct>"
		for n := 0; n < r.NumField(); n++ {
			rv := r.FieldByIndex([]int{n}).Interface()
			if rv != nil {
				s += "<member>"

				if t.Field(n).Tag.Get("xml") != "" {
					s += "<name>" + t.Field(n).Tag.Get("xml") + "</name>"
				} else {
					s += "<name>" + strings.ToLower(t.Field(n).Name) + "</name>"
				}
				s += "<value>" + toXml(rv, true) + "</value>"
				s += "</member>"
			}
		}
		s += "</struct>"
		return s
	case reflect.UnsafePointer:
		return toXml(r.Elem(), typ)
	}
	return
}
