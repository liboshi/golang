package soap

import (
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *Header  `xml:",omitempty"`
	Body    interface{}
}

type Header struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
}

type Fault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	Code    string   `xml:"faultcode"`
	String  string   `xml:"faultstring"`
	Detail  struct {
		Fault types.AnyType `xml:",any,typeattr"`
	} `xml:"detail"`
}

func (f *Fault) VimFault() types.AnyType {
	return f.Detail.Fault
}
