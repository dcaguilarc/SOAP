package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// Definir la estructura de la respuesta SOAP
type Envelope struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	XMLNs   string   `xml:"xmlns:soapenv,attr"`
	Body    Body     `xml:"soapenv:Body"`
}

type Body struct {
	HolaMundoResponse HolaMundoResponse `xml:"HollaMundoResponse"`
}

type HolaMundoResponse struct {
	Message string `xml:"mensaje"`
}

func main() {
	// Manejar la ruta para el WSDL
	http.HandleFunc("/wsdl", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		wsdl := `
            <definitions xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
                xmlns:web="http://www.example.com/webservice">
                <message name="HollaMundoRequest">
                    <part name="message" type="xsd:string"/>
                </message>
                <message name="HollaMundoResponse">
                    <part name="mensaje" type="xsd:string"/>
                </message>
                <portType name="HollaMundoPortType">
                    <operation name="HollaMundo">
                        <input message="web:HollaMundoRequest"/>
                        <output message="web:HollaMundoResponse"/>
                    </operation>
                </portType>
                <binding name="HollaMundoBinding" type="web:HollaMundoPortType">
                    <soap:binding transport="http://schemas.xmlsoap.org/soap/http"/>
                    <operation name="HollaMundo">
                        <soap:operation soapAction=""/>
                        <input>
                            <soap:body use="literal"/>
                        </input>
                        <output>
                            <soap:body use="literal"/>
                        </output>
                    </operation>
                </binding>
                <service name="HollaMundoService">
                    <port name="HollaMundoPort" binding="web:HollaMundoBinding">
                        <soap:address location="http://localhost:8080/"/>
                    </port>
                </service>
            </definitions>
        `
		fmt.Fprintln(w, wsdl)
	})

	// Manejar la ruta SOAP para el servicio "HollaMundo"
	http.HandleFunc("/HollaMundo", func(w http.ResponseWriter, r *http.Request) {
		// Crear la respuesta con el mensaje "Hola Mundo"
		holaMundoResponse := HolaMundoResponse{
			Message: "Hola Mundo",
		}

		// Crear el sobre SOAP (SOAP envelope)
		envelope := Envelope{
			XMLNs: "http://schemas.xmlsoap.org/soap/envelope/",
			Body: Body{
				HolaMundoResponse: holaMundoResponse,
			},
		}

		// Establecer el tipo de contenido como XML
		w.Header().Set("Content-Type", "application/xml")

		// Codificar la respuesta a XML y escribirla en la respuesta HTTP
		xml.NewEncoder(w).Encode(envelope)
	})

	// Iniciar el servidor SOAP en el puerto 8080
	fmt.Println("Servidor SOAP corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
