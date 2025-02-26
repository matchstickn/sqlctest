package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	_ "embed"
	"log"
	"net"
	"os"

	"github.com/matchstickn/sqlctest/cmd"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

//go:embed certs/server.crt
var crt []byte

//go:embed certs/server.key
var key []byte

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		log.Println("detected env injection, if not:", err)
	}
	connstr := os.Getenv("NEONTECH_URL")
	https := os.Getenv("HTTPS")

	query, pq := cmd.SetUpDB(ctx, connstr)
	defer pq.Close(ctx)

	app := fiber.New(fiber.Config{
		AppName: "FlintCRUD",
	})

	cmd.SetUpRoutes(ctx, query, app)

	log.Println(query.GetUserTricks(ctx))

	if https == "true" {

		cert, err := tls.X509KeyPair(crt, key)
		if err != nil {
			log.Fatal(err)
		}

		caCertPool := x509.NewCertPool()
		if !caCertPool.AppendCertsFromPEM(cert.Certificate[0]) {
			log.Println("err: unable to append certs from pem to certpool")
			// log.Fatal(app.Listen(":4000"))
		}

		cfg := &tls.Config{
			Certificates:       []tls.Certificate{cert},
			RootCAs:            caCertPool,
			NextProtos:         []string{"http/1.1"},
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: true, // Set to true ONLY for development/testing
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			},
		}

		ln, err := net.Listen("tcp", ":4000")
		if err != nil {
			log.Fatal(err)
		}

		ln = tls.NewListener(ln, cfg)

		log.Fatal(app.Listener(ln))

		// log.Fatal(app.ListenTLSWithCertificate(":4000", cfg.Certificates[0]))
	} else {
		log.Fatal(app.Listen(":4000"))
	}
}
