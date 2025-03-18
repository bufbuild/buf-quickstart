package main

import (
	tagv1 "buf.build/gen/go/jrinehart/common/protocolbuffers/go/tag/v1"
	"buf.build/gen/go/jrinehart/invoice/connectrpc/go/invoice/v1/invoicev1connect"
	invoicev1 "buf.build/gen/go/jrinehart/invoice/protocolbuffers/go/invoice/v1"
	"connectrpc.com/connect"
	"context"
	"log"
	"net/http"
)

func main() {
	client := invoicev1connect.NewInvoiceServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)

	createValidInvoice(client)
	createInvalidInvoice(client)
}

func createValidInvoice(client invoicev1connect.InvoiceServiceClient) {
	_, err := client.CreateInvoice(
		context.Background(),
		connect.NewRequest(&invoicev1.CreateInvoiceRequest{
			Invoice: &invoicev1.Invoice{
				InvoiceId:  "invoice-one",
				CustomerId: "customer-one",
				LineItems: []*invoicev1.LineItem{
					{
						UnitPrice: 999,
						Quantity:  2,
					},
				},
			},
			Tags: &tagv1.Tags{
				Tag: []string{
					"bogo-campaign",
					"valued-customer",
				},
			},
		}),
	)
	if err != nil {
		log.Fatalf("error creating valid invoice: %v", err)
	}
	log.Println("Valid invoice created")
}

func createInvalidInvoice(client invoicev1connect.InvoiceServiceClient) {
	_, err := client.CreateInvoice(
		context.Background(),
		connect.NewRequest(&invoicev1.CreateInvoiceRequest{
			Invoice: &invoicev1.Invoice{
				InvoiceId:  "invoice-two",
				CustomerId: "customer-one",
			},
		}),
	)
	if err == nil {
		log.Fatalf("Oh no! An invalid invoice was created!")
	}
	log.Printf("correctly received error for invalid invoice: %v", err)
}
