# SOAP
* instalar GO
In the VC terminal add these commands

* go mod init hello-soap
* go get github.com/tiaguinho/gosoap

and with that we can run the program
* go run main.go

* here show de port 
 http://localhost:8080

in an other terminal

* curl -X POST http://localhost:8080/HollaMundo -H "Content-Type: text/xml" -d @soap_request.xml
