package notify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type EmailDTO struct {
	Email          string
	MessageContent string
}

type EmailSender interface {
	Send(token EmailDTO) error
}

type email struct {
}

func (e email) Send(email EmailDTO) error {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		fmt.Printf("Unable to read client secret file: %v", err)
		return err
	}
	config, err := google.ConfigFromJSON(b, gmail.MailGoogleComScope)
	if err != nil {
		fmt.Printf("Unable to parse client secret file to config: %v", err)
		return err
	}
	client, err := getClient(config)
	if err != nil {
		fmt.Printf("Unable to retrieve Gmail client: %v", err)
		return err
	}

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		fmt.Printf("Unable to retrieve Gmail client: %v", err)
		return err
	}
	//追記
	msgStr := "From: 'me'\r\n" +
		fmt.Sprintf("To: %v\r\n", email.Email) + //送信先
		"Subject: [PearOpenData] 認証トークンの送信\r\n" +
		"\r\n" + email.MessageContent

	fmt.Println("msgStr: ", msgStr)
	reader := strings.NewReader(msgStr)
	transformer := japanese.ISO2022JP.NewEncoder()
	msgISO2022JP, err := ioutil.ReadAll(transform.NewReader(reader, transformer))
	if err != nil {
		fmt.Printf("Unable to convert to ISO2022JP: %v", err)
		return err
	}

	msg := msgISO2022JP
	message := gmail.Message{}
	message.Raw = base64.StdEncoding.EncodeToString(msg)
	if _, err = srv.Users.Messages.Send("me", &message).Do(); err != nil {
		fmt.Printf("%v", err)
	}
	return nil
}

func NewEmailSender() EmailSender {
	return &email{}
}

func getClient(config *oauth2.Config) (*http.Client, error) {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		return nil, err
	}
	return config.Client(context.Background(), tok), nil
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}
