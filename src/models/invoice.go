package models

import "encoding/xml"

type Factura struct {
	XMLName                     xml.Name                     `xml:"factura"`
	Id                          string                       `xml:"id,attr"`
	Version                     string                       `xml:"version,attr"`
	InfoTributaria              *InfoTributaria              `xml:"infoTributaria"`
	InfoFactura                 *InfoFactura                 `xml:"infoFactura"`
	Detalles                    []*Detalle                   `xml:"detalles>detalle"`
	Reembolsos                  []*ReembolsoDetalle          `xml:"reembolsos>reembolsoDetalle"`
	Retenciones                 []*Retencion                 `xml:"retenciones>retencion"`
	InfoSustitutivaGuiaRemision *InfoSustitutivaGuiaRemision `xml:"infoSustitutivaGuiaRemision"`
	OtrosRubrosTerceros         []*Rubro                     `xml:"otrosRubrosTerceros>rubro"`
	TipoNegociable              *TipoNegociable              `xml:"tipoNegociable"`
	MaquinaFiscal               *MaquinaFiscal               `xml:"maquinaFiscal"`
	InfoAdicional               []*CampoAdicional            `xml:"infoAdicional>campoAdicional"`
}

type InfoTributaria struct {
	XMLName            xml.Name `xml:"infoTributaria"`
	Ambiente           string   `xml:"ambiente"`
	TipoEmision        string   `xml:"tipoEmision"`
	RazonSocial        string   `xml:"razonSocial"`
	NombreComercial    string   `xml:"nombreComercial"`
	Ruc                string   `xml:"ruc"`
	ClaveAcceso        string   `xml:"claveAcceso"`
	CodDoc             string   `xml:"codDoc"`
	Estab              string   `xml:"estab"`
	PtoEmi             string   `xml:"ptoEmi"`
	Secuencial         string   `xml:"secuencial"`
	DirMatriz          string   `xml:"dirMatriz"`
	AgenteRetencion    string   `xml:"agenteRetencion"`
	ContribuyenteRimpe string   `xml:"contribuyenteRimpe"`
}

type InfoFactura struct {
	XMLName                     xml.Name         `xml:"infoFactura"`
	FechaEmision                string           `xml:"fechaEmision"`
	DirEstablecimiento          string           `xml:"dirEstablecimiento"`
	ContribuyenteEspecial       string           `xml:"contribuyenteEspecial"`
	ObligadoContabilidad        string           `xml:"obligadoContabilidad"`
	ComercioExterior            string           `xml:"comercioExterior"`
	IncoTermFactura             string           `xml:"incoTermFactura"`
	LugarIncoTerm               string           `xml:"lugarIncoTerm"`
	PaisOrigen                  string           `xml:"paisOrigen"`
	PuertoEmbarque              string           `xml:"puertoEmbarque"`
	PuertoDestino               string           `xml:"puertoDestino"`
	PaisDestino                 string           `xml:"paisDestino"`
	PaisAdquisicion             string           `xml:"paisAdquisicion"`
	TipoIdentificacionComprador string           `xml:"tipoIdentificacionComprador"`
	GuiaRemision                string           `xml:"guiaRemision"`
	RazonSocialComprador        string           `xml:"razonSocialComprador"`
	IdentificacionComprador     string           `xml:"identificacionComprador"`
	DireccionComprador          string           `xml:"direccionComprador"`
	TotalSinImpuestos           string           `xml:"totalSinImpuestos"`
	TotalSubsidio               string           `xml:"totalSubsidio"`
	IncoTermTotalSinImpuestos   string           `xml:"incoTermTotalSinImpuestos"`
	TotalDescuento              string           `xml:"totalDescuento"`
	CodDocReembolso             string           `xml:"codDocReembolso"`
	TotalComprobantesReembolso  string           `xml:"totalComprobantesReembolso"`
	TotalBaseImponibleReembolso string           `xml:"totalBaseImponibleReembolso"`
	TotalImpuestoReembolso      string           `xml:"totalImpuestoReembolso"`
	TotalConImpuestos           []*TotalImpuesto `xml:"totalConImpuestos>totalImpuesto"`
	Compensaciones              []*Compensacion  `xml:"compensaciones>compensacion"`
	Propina                     string           `xml:"propina"`
	FleteInternacional          string           `xml:"fleteInternacional"`
	SeguroInternacional         string           `xml:"seguroInternacional"`
	GastosAduaneros             string           `xml:"gastosAduaneros"`
	GastosTransporteOtros       string           `xml:"gastosTransporteOtros"`
	ImporteTotal                string           `xml:"importeTotal"`
	Moneda                      string           `xml:"moneda"`
	Placa                       string           `xml:"placa"`
	Pagos                       []*Pago          `xml:"pagos>pago"`
	ValorRetIva                 string           `xml:"valorRetIva,omitempty"`
	ValorRetRenta               string           `xml:"valorRetRenta,omitempty"`
}

type TotalImpuesto struct {
	XMLName            xml.Name `xml:"totalImpuesto"`
	Codigo             string   `xml:"codigo"`
	CodigoPorcentaje   string   `xml:"codigoPorcentaje"`
	DescuentoAdicional string   `xml:"descuentoAdicional"`
	BaseImponible      string   `xml:"baseImponible"`
	Tarifa             string   `xml:"tarifa"`
	Valor              string   `xml:"valor"`
	ValorDevolucionIva string   `xml:"valorDevolucionIva"`
}

type Compensacion struct {
	XMLName xml.Name `xml:"compensacion"`
	Codigo  string   `xml:"codigo"`
	Tarifa  string   `xml:"tarifa"`
	Valor   string   `xml:"valor"`
}

type Pago struct {
	XMLName      xml.Name `xml:"pago"`
	FormaPago    string   `xml:"formaPago"`
	Total        string   `xml:"total"`
	Plazo        string   `xml:"plazo"`
	UnidadTiempo string   `xml:"unidadTiempo,omitempty"`
}

type Detalle struct {
	XMLName                xml.Name        `xml:"detalle"`
	CodigoPrincipal        string          `xml:"codigoPrincipal"`
	CodigoAuxiliar         string          `xml:"codigoAuxiliar"`
	Descripcion            string          `xml:"descripcion"`
	UnidadMedida           string          `xml:"unidadMedida"`
	Cantidad               string          `xml:"cantidad"`
	PrecioUnitario         string          `xml:"precioUnitario"`
	PrecioSinSubsidio      string          `xml:"precioSinSubsidio"`
	Descuento              string          `xml:"descuento"`
	PrecioTotalSinImpuesto string          `xml:"precioTotalSinImpuesto"`
	DetallesAdicionales    []*DetAdicional `xml:"detallesAdicionales>detAdicional"`
	Impuestos              []*Impuesto     `xml:"impuestos>impuesto"`
}
type DetAdicional struct {
	XMLName xml.Name `xml:"detAdicional"`
	Nombre  string   `xml:"nombre,attr"`
	Valor   string   `xml:"valor,attr"`
}

type Impuesto []struct {
	XMLName          xml.Name `xml:"impuesto"`
	Codigo           string   `xml:"codigo"`
	CodigoPorcentaje string   `xml:"codigoPorcentaje"`
	Tarifa           string   `xml:"tarifa"`
	BaseImponible    string   `xml:"baseImponible"`
	Valor            string   `xml:"valor"`
}

type ReembolsoDetalle []struct {
	XMLName                              xml.Name                 `xml:"reembolsoDetalle"`
	TipoIdentificacionProveedorReembolso string                   `xml:"tipoIdentificacionProveedorReembolso"`
	IdentificacionProveedorReembolso     string                   `xml:"identificacionProveedorReembolso"`
	CodPaisPagoProveedorReembolso        string                   `xml:"codPaisPagoProveedorReembolso"`
	TipoProveedorReembolso               string                   `xml:"tipoProveedorReembolso"`
	CodDocReembolso                      string                   `xml:"codDocReembolso"`
	EstabDocReembolso                    string                   `xml:"estabDocReembolso"`
	PtoEmiDocReembolso                   string                   `xml:"ptoEmiDocReembolso"`
	SecuencialDocReembolso               string                   `xml:"secuencialDocReembolso"`
	FechaEmisionDocReembolso             string                   `xml:"fechaEmisionDocReembolso"`
	NumeroautorizacionDocReemb           string                   `xml:"numeroautorizacionDocReemb"`
	DetalleImpuestos                     []*DetalleImpuesto       `xml:"detalleImpuestos>detalleImpuesto"`
	CompensacionesReembolso              []*CompensacionReembolso `xml:"compensacionesReembolso>compensacionReembolso"`
}

type DetalleImpuesto struct {
	XMLName                xml.Name `xml:"detalleImpuesto"`
	Codigo                 string   `xml:"codigo"`
	CodigoPorcentaje       string   `xml:"codigoPorcentaje"`
	Tarifa                 string   `xml:"tarifa"`
	BaseImponibleReembolso string   `xml:"baseImponibleReembolso"`
	ImpuestoReembolso      string   `xml:"impuestoReembolso"`
}

type CompensacionReembolso struct {
	XMLName xml.Name `xml:"compensacionReembolso"`
	Codigo  string   `xml:"codigo"`
	Tarifa  string   `xml:"tarifa"`
	Valor   string   `xml:"valor"`
}

type Retencion struct {
	XMLName          xml.Name `xml:"retencion"`
	Codigo           string   `xml:"codigo"`
	CodigoPorcentaje string   `xml:"codigoPorcentaje"`
	Tarifa           string   `xml:"tarifa"`
	Valor            string   `xml:"valor"`
}

type InfoSustitutivaGuiaRemision struct {
	XMLName                         xml.Name   `xml:"infoSustitutivaGuiaRemision"`
	DirPartida                      string     `xml:"dirPartida"`
	DirDestinatario                 string     `xml:"dirDestinatario"`
	FechaIniTransporte              string     `xml:"fechaIniTransporte"`
	FechaFinTransporte              string     `xml:"fechaFinTransporte"`
	RazonSocialTransportista        string     `xml:"razonSocialTransportista"`
	TipoIdentificacionTransportista string     `xml:"tipoIdentificacionTransportista"`
	RucTransportista                string     `xml:"rucTransportista"`
	Placa                           string     `xml:"placa"`
	Destinos                        []*Destino `xml:"destinos>destino"`
}
type Destino struct {
	XMLName          xml.Name `xml:"destino"`
	MotivoTraslado   string   `xml:"motivoTraslado"`
	DocAduaneroUnico string   `xml:"docAduaneroUnico"`
	CodEstabDestino  string   `xml:"codEstabDestino"`
	Ruta             string   `xml:"ruta"`
}

type Rubro struct {
	XMLName  xml.Name `xml:"rubro"`
	Concepto string   `xml:"concepto"`
	Total    string   `xml:"total"`
}

type TipoNegociable struct {
	XMLName xml.Name `xml:"tipoNegociable"`
	Correo  string   `xml:"correo"`
}

type MaquinaFiscal struct {
	XMLName xml.Name `xml:"maquinaFiscal"`
	Marca   string   `xml:"marca"`
	Modelo  string   `xml:"modelo"`
	Serie   string   `xml:"serie"`
}

type CampoAdicional struct {
	XMLName xml.Name `xml:"campoAdicional"`
	Nombre  string   `xml:"nombre,attr"`
}
