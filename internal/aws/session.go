package aws

// Import AWS session
import (
	"github.com/aws/aws-sdk-go/aws/session"
)

func CreateSessionWithCredentials() *session.Session {
	// Create a new session that the SDK will use to load
	// credentials from the shared credentials file.
	sess := session.Must(session.NewSession())

	return sess
}

func CreateSessionWithSSO() *session.Session {
	// Create a new session that the SDK will use to load
	// credentials from the shared credentials file.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return sess
}
