package wsdl_sri

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"github.com/tiaguinho/gosoap"
)

type RespuestaAutorizacion AnyType

type AutorizacionComprobante struct {
	XMLName                xml.Name `xml:"http://ec.gob.sri.ws.autorizacion autorizacionComprobante"`
	ClaveAccesoComprobante string   `xml:"claveAccesoComprobante,omitempty" json:"claveAccesoComprobante,omitempty"`
}

type AutorizacionComprobanteResponse struct {
	XMLName                          xml.Name              `xml:"http://ec.gob.sri.ws.autorizacion autorizacionComprobanteResponse"`
	RespuestaAutorizacionComprobante *RespuestaComprobante `xml:"RespuestaAutorizacionComprobante,omitempty" json:"RespuestaAutorizacionComprobante,omitempty"`
}

type RespuestaComprobante struct {
	ClaveAccesoConsultada string          `xml:"claveAccesoConsultada,omitempty" json:"claveAccesoConsultada,omitempty"`
	NumeroComprobantes    string          `xml:"numeroComprobantes,omitempty" json:"numeroComprobantes,omitempty"`
	Autorizaciones        []*Autorizacion `xml:"autorizaciones>autorizacion,omitempty" json:",omitempty"`
}

type Autorizacion struct {
	XMLName            xml.Name               `xml:"autorizacion"`
	Estado             string                 `xml:"estado,omitempty" json:"estado,omitempty"`
	NumeroAutorizacion string                 `xml:"numeroAutorizacion,omitempty" json:"numeroAutorizacion,omitempty"`
	FechaAutorizacion  string                 `xml:"fechaAutorizacion,omitempty" json:"fechaAutorizacion,omitempty"`
	Ambiente           string                 `xml:"ambiente,omitempty" json:"ambiente,omitempty"`
	Comprobante        string                 `xml:"comprobante,omitempty" json:"comprobante,omitempty"`
	Mensajes           []*MensajeAutorizacion `xml:"mensajes>mensaje,omitempty" json:",omitempty"`
}

type MensajeAutorizacion struct {
	XMLName              xml.Name `xml:"mensaje"`
	Identificador        string   `xml:"identificador,omitempty" json:"identificador,omitempty"`
	Mensaje              string   `xml:"mensaje,omitempty" json:"mensaje,omitempty"`
	InformacionAdicional string   `xml:"informacionAdicional,omitempty" json:"informacionAdicional,omitempty"`
	Tipo                 string   `xml:"tipo,omitempty" json:"tipo,omitempty"`
}

type AutorizacionComprobanteLote struct {
	XMLName         xml.Name `xml:"http://ec.gob.sri.ws.autorizacion autorizacionComprobanteLote"`
	ClaveAccesoLote string   `xml:"claveAccesoLote,omitempty" json:"claveAccesoLote,omitempty"`
}

type AutorizacionComprobanteLoteResponse struct {
	XMLName                   xml.Name       `xml:"http://ec.gob.sri.ws.autorizacion autorizacionComprobanteLoteResponse"`
	RespuestaAutorizacionLote *RespuestaLote `xml:"RespuestaAutorizacionLote,omitempty" json:"RespuestaAutorizacionLote,omitempty"`
}

type RespuestaLote struct {
	ClaveAccesoLoteConsultada string          `xml:"claveAccesoLoteConsultada,omitempty" json:"claveAccesoLoteConsultada,omitempty"`
	NumeroComprobantesLote    string          `xml:"numeroComprobantesLote,omitempty" json:"numeroComprobantesLote,omitempty"`
	Autorizaciones            []*Autorizacion `xml:"autorizaciones>autorizacion,omitempty" json:",omitempty"`
}

type AutorizacionComprobantesOffline interface {
	AutorizacionComprobante(request *AutorizacionComprobante) (*AutorizacionComprobanteResponse, error)
	AutorizacionComprobanteLote(request *AutorizacionComprobanteLote) (*AutorizacionComprobanteLoteResponse, error)
}

type autorizacionComprobantesOffline struct {
	client *gosoap.Client
}

func NewAutorizacionComprobantesOffline(enviroment string) AutorizacionComprobantesOffline {
	t := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{
		Timeout:   1200 * time.Millisecond,
		Transport: t,
	}

	var url string
	if enviroment == "2" {
		//Only Production
		url = "https://cel.sri.gob.ec/comprobantes-electronicos-ws/AutorizacionComprobantesOffline?wsdl"
	}
	url = "https://celcer.sri.gob.ec/comprobantes-electronicos-ws/AutorizacionComprobantesOffline?wsdl"

	soap, err := gosoap.SoapClient(url, httpClient)
	if err != nil {
		fmt.Printf("Call error: %s", err)
		return nil
	}

	return &autorizacionComprobantesOffline{
		client: soap,
	}
}

func (service *autorizacionComprobantesOffline) AutorizacionComprobante(request *AutorizacionComprobante) (*AutorizacionComprobanteResponse, error) {
	res, err := service.client.Call("autorizacionComprobante", request)
	if err != nil {
		fmt.Printf("Call error: %s", err)
		return nil, err
	}

	response := new(AutorizacionComprobanteResponse)
	res.Unmarshal(&response)
	return response, nil
}

func (service *autorizacionComprobantesOffline) AutorizacionComprobanteLote(request *AutorizacionComprobanteLote) (*AutorizacionComprobanteLoteResponse, error) {
	res, err := service.client.Call("autorizacionComprobanteLote", request)
	if err != nil {
		fmt.Printf("Call error: %s", err)
		return nil, err
	}

	response := new(AutorizacionComprobanteLoteResponse)
	res.Unmarshal(&response)
	return response, nil
}
