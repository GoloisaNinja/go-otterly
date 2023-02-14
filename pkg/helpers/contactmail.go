package helpers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const CHARSET = "UTF-8"

func SendEmail(sess *session.Session, toAddresses []*string, emailText, sender, subject string) error {
	msgwsender := `Otterly Contact Form Submission` + "\n\n" + "Message Originator: " + "\n" + sender + "\n\n" + "Message Content: " + "\n" + emailText
	sesClient := ses.New(sess)
	_, err := sesClient.SendEmail(&ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: toAddresses,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String(CHARSET),
					Data:    aws.String(msgwsender),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CHARSET),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String("admin@otterly.games"),
	})
	return err
}
