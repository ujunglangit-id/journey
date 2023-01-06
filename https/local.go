package https

import (
	"log"

	"github.com/ujunglangit-id/httpscerts"
	"github.com/ujunglangit-id/journey/configuration"
	"github.com/ujunglangit-id/journey/filenames"
)

func checkCertificates() {
	// Check https certificates. If they are not available generate temporary ones for testing.
	if err := httpscerts.Check(filenames.HttpsCertFilename, filenames.HttpsKeyFilename); err != nil {
		log.Println("Warning: couldn't load https certs. Generating new ones. Replace " + filenames.HttpsCertFilename + " and " + filenames.HttpsKeyFilename + " with your own certificates as soon as possible!")
		if err := httpscerts.Generate(filenames.HttpsCertFilename, filenames.HttpsKeyFilename, configuration.Config.HttpsUrl); err != nil {
			log.Fatal("Error: Couldn't create https certificates.")
			return
		}
	}
}
