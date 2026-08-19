package twilioInterface

import (
	"fmt"
	log "log/slog"
	"net/http"
	"path"

	"os"

	"github.com/spf13/viper"
	twilio "github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/twiml"
)

type Smscfg struct {
	SSID           string
	Api_SSID       string
	Api_Secret     string
	Phone_number   string
	Test_to_number string
}

var cfg Smscfg

func init() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Error("Failed to get user home directory.")
		os.Exit(1)
	}
	arg := os.Args[1]
	fmt.Println("args:", len(os.Args), os.Args)
	switch arg {
	case "dev":
		viper.AddConfigPath("../.settings")
	case "prod":
		viper.AddConfigPath(path.Join(homedir, ".settings"))
	}

	viper.SetConfigType("yaml")
	viper.SetConfigName("sms_config")

	err = viper.ReadInConfig()
	if err != nil {
		log.Error("failed to read sms config file")
		return
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Error("Failed to unmarshall sms config.")
	}
}
func CreateClient() *twilio.RestClient {
	//Creates and returns an authenticated client	
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username:   cfg.Api_SSID,
		Password:   cfg.Api_Secret,
		AccountSid: cfg.SSID,
	})
	return client
}
func smsHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	from := r.FormValue("From")
	to := r.FormValue("To")
	body := r.FormValue("Body")
	log.Info(" New SMS\nFrom: %s\nTo: %s\nBody: %s", from, to, body,"message")
	resp, err := twiml.Messages(
		[]twiml.Element{
			&twiml.MessagingMessage{
				Body: "Thanks for your message!",
			},
		},
	)
	if err != nil {
		http.Error(w, "Failed to build response", http.StatusInternalServerError)
		return
	}
	fmt.Println(resp)
	w.Header().Set("Content-Type", "application/xml")
	//fmt.Fprint(w, twimlStr)
}
