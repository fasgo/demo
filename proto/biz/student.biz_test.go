package biz

import (
	"encoding/xml"
	"fmt"
	"github.com/fasgo/base/kits"
	"testing"
)

const data = `
<?xml version="1.0" encoding="utf-8" ?>
<!DOCTYPE resources [
   <!ELEMENT resource (code, message, status-code)>
   <!ATTLIST accept-language CDATA "">
   <!ELEMENT code (#PCDATA)>
   <!ELEMENT message (#PCDATA)>
   <!ELEMENT status-code (#PCDATA)>
]>
<!-- accept-language使用 iso_language_code或iso_language_code-ISO_COUNTRY_CODE, 多值用逗号分割 -->
<resources accept-language="en,en-US,en-UK">
    <resource>
        <!-- 必需: 错误代码 -->
        <code>1001</code>
        <!-- 必需: 错误消息 -->
        <message>测试%v</message>
        <!-- 可选: 状态码, 默认210 -->
        <status-code>403</status-code>
    </resource>
</resources>
`

type resource struct {
	Code       int    `xml:"code"`
	Message    string `xml:"message"`
	StatusCode int    `xml:"status-code"`
}

type resources struct {
	XMLName        xml.Name    `xml:"resources"`
	AcceptLanguage string      `xml:"accept-language,attr"`
	Resources      []*resource `xml:"resource"`
}

func TestTagServiceService_All(t *testing.T) {
	var rs *resources
	err := xml.Unmarshal([]byte(data), &rs)
	if err != nil {
		panic(err)
	}

	fmt.Println("accept-language: ", rs.AcceptLanguage)
	for _, r := range rs.Resources {
		fmt.Println("result: ", kits.ToJson(r))
	}
}
