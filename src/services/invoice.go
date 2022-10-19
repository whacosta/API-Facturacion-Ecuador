package services

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/whacosta/API-Facturacion-Ecuador/src/models"
	"github.com/whacosta/API-Facturacion-Ecuador/src/request"
	"github.com/whacosta/API-Facturacion-Ecuador/src/utils"
	"github.com/whacosta/API-Facturacion-Ecuador/src/wsdl_sri"
	"net/http"
	"strconv"
)

func GenerateInvoice(w http.ResponseWriter, r *http.Request) {
	body := &request.InvoiceRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Fprint(w, err)
		return
	}

	newAccessKey := utils.AccessKey{
		EmissionDate: body.InvoiceInfo.EmissionDate,
		VoucherType:  "01",
		RUC:          body.TaxInfo.Ruc,
		Environment:  body.TaxInfo.Environment,
		Serie:        "001001",
		Sequencial:   body.TaxInfo.Sequencial,
		NumericCode:  "12345678",
		EmissionType: "1",
	}

	taxInfo := &models.InfoTributaria{
		Ambiente:    body.TaxInfo.Environment,
		TipoEmision: "1",
		RazonSocial: body.TaxInfo.BusinessName,
		Ruc:         body.TaxInfo.Ruc,
		Secuencial:  body.TaxInfo.Sequencial,
		ClaveAcceso: newAccessKey.GenerateAccessKey(),
		CodDoc:      "01",
		Estab:       "001",
		PtoEmi:      "001",
		DirMatriz:   "Guayaquil",
	}

	var paymests []*models.Pago
	paymest := &models.Pago{
		FormaPago: "01",
		Total:     "4.20",
		Plazo:     "0",
	}
	paymests = append(paymests, paymest)

	invoiceInfo := &models.InfoFactura{
		FechaEmision:                body.InvoiceInfo.EmissionDate,
		TipoIdentificacionComprador: body.InvoiceInfo.IdentificationType,
		IdentificacionComprador:     body.InvoiceInfo.Identification,
		DireccionComprador:          body.InvoiceInfo.Address,
		Pagos:                       paymests,
	}

	var details []*models.Detalle
	for _, detail := range body.Details {
		unitPrice := strconv.FormatFloat(detail.UnitPrice, 'f', -1, 32)

		details = append(details, &models.Detalle{
			CodigoPrincipal: detail.ProductCode,
			Descripcion:     detail.Description,
			PrecioUnitario:  unitPrice,
		})
	}

	invoice := &models.Factura{
		Id:             "comprobante",
		Version:        "2.1.0",
		InfoTributaria: taxInfo,
		InfoFactura:    invoiceInfo,
		Detalles:       details,
	}

	out, _ := xml.MarshalIndent(invoice, " ", "  ")

	validateComprobante := wsdl_sri.ValidarComprobante{
		Xml: out,
	}

	response, err := SendInvoice(validateComprobante)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	/*if response.RespuestaRecepcionComprobante.Estado == "RECIBIDA" {
		autorizacionComprobante := wsdl_sri.AutorizacionComprobante{
			ClaveAccesoComprobante: taxInfo.ClaveAcceso,
		}

		auth, err := AuthInvoice(autorizacionComprobante)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": auth})
		return
	}*/
	fmt.Fprint(w, response)

}

func SendInvoice(request wsdl_sri.ValidarComprobante) (*wsdl_sri.ValidarComprobanteResponse, error) {
	receptor := wsdl_sri.NewRecepcionComprobantesOffline("1")
	return receptor.ValidarComprobante(&request)
}

func AuthInvoice(request wsdl_sri.AutorizacionComprobante) (*wsdl_sri.AutorizacionComprobanteResponse, error) {
	receptor := wsdl_sri.NewAutorizacionComprobantesOffline("1")
	return receptor.AutorizacionComprobante(&request)
}
