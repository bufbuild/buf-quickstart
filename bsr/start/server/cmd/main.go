package main

import (
	connect "connectrpc.com/connect"
	"connectrpc.com/validate"
	"context"
	"fmt"
	invoicev1 "github.com/bufbuild/buf-quickstarts/bsr/server/gen/invoice/v1"
	"github.com/bufbuild/buf-quickstarts/bsr/server/gen/invoice/v1/invoicev1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

const address = "localhost:8080"

func main() {
	mux := http.NewServeMux()

	// Create the validation interceptor provided by connectrpc.com/validate.
	validateInterceptor, err := validate.NewInterceptor()
	if err != nil {
		log.Fatalf("error creating interceptor: %v", err)
	}

	path, handler := invoicev1connect.NewInvoiceServiceHandler(&invoiceServiceServer{}, connect.WithInterceptors(validateInterceptor))
	mux.Handle(path, handler)
	fmt.Println("... Listening on", address)
	if err := http.ListenAndServe(
		address,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}

// invoiceServiceServer implements the PetStoreService API.
type invoiceServiceServer struct {
	invoicev1connect.UnimplementedInvoiceServiceHandler
}

// CreateInvoice creates an invoices with any associated tags.
func (s *invoiceServiceServer) CreateInvoice(
	_ context.Context,
	req *connect.Request[invoicev1.CreateInvoiceRequest],
) (*connect.Response[invoicev1.CreateInvoiceResponse], error) {
	invoice := req.Msg.GetInvoice()
	tags := req.Msg.GetTags()

	total := 0
	for _, li := range invoice.GetLineItems() {
		total += int(li.GetUnitPrice() * li.GetQuantity())
	}
	log.Printf("Creating invoice for customer %s for %v", invoice.GetCustomerId(), total)

	if len(tags.GetTag()) > 0 {
		for _, tag := range tags.GetTag() {
			log.Printf("  - %s", tag)
		}
	}

	return connect.NewResponse(&invoicev1.CreateInvoiceResponse{
		InvoiceId: invoice.GetInvoiceId(),
	}), nil
}
