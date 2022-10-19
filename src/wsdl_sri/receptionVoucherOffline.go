package wsdl_sri

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"github.com/tiaguinho/gosoap"
)

type ValidarComprobante struct {
	XMLName xml.Name `xml:"http://ec.gob.sri.ws.recepcion validarComprobante"`
	Xml     []byte   `xml:"xml,omitempty" json:"xml,omitempty"`
}

type ValidarComprobanteResponse struct {
	XMLName                       xml.Name            `xml:"http://ec.gob.sri.ws.recepcion validarComprobanteResponse"`
	RespuestaRecepcionComprobante *RespuestaSolicitud `xml:"RespuestaRecepcionComprobante,omitempty" json:"RespuestaRecepcionComprobante,omitempty"`
}

type RespuestaSolicitud struct {
	Estado       string         `xml:"estado,omitempty" json:"estado,omitempty"`
	Comprobantes []*Comprobante `xml:"comprobantes>comprobante,omitempty" json:"comprobantes,omitempty"`
}

type Comprobante struct {
	XMLName     xml.Name            `xml:"comprobante"`
	ClaveAcceso string              `xml:"claveAcceso,omitempty" json:"claveAcceso,omitempty"`
	Mensajes    []*MensajeRecepcion `xml:"mensajes>mensaje,omitempty" json:"mensajes,omitempty"`
}

type MensajeRecepcion struct {
	XMLName              xml.Name `xml:"mensaje"`
	Identificador        string   `xml:"identificador,omitempty" json:"identificador,omitempty"`
	Mensaje              string   `xml:"mensaje,omitempty" json:"mensaje,omitempty"`
	InformacionAdicional string   `xml:"informacionAdicional,omitempty" json:"informacionAdicional,omitempty"`
	Tipo                 string   `xml:"tipo,omitempty" json:"tipo,omitempty"`
}

type RecepcionComprobantesOffline interface {
	ValidarComprobante(request *ValidarComprobante) (*ValidarComprobanteResponse, error)
}

type recepcionComprobantesOffline struct {
	client *gosoap.Client
}

func NewRecepcionComprobantesOffline(enviroment string) RecepcionComprobantesOffline {
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
		url = "https://cel.sri.gob.ec/comprobantes-electronicos-ws/RecepcionComprobantesOffline?wsdl"
	}
	url = "https://celcer.sri.gob.ec/comprobantes-electronicos-ws/RecepcionComprobantesOffline?wsdl"

	soap, err := gosoap.SoapClient(url, httpClient)
	if err != nil {
		fmt.Printf("Call error: %s", err)
		return nil
	}
	return &recepcionComprobantesOffline{
		client: soap,
	}
}

func (service *recepcionComprobantesOffline) ValidarComprobante(request *ValidarComprobante) (*ValidarComprobanteResponse, error) {
	res, err := service.client.Call("validarComprobante", request)
	if err != nil {
		fmt.Printf("Call error: %s", err)
		return nil, err
	}

	response := new(ValidarComprobanteResponse)
	fmt.Println(string(res.Body))

	xml.Unmarshal(res.Body, &response)
	return response, nil
}
