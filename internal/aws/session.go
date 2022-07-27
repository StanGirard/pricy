package aws

// Import AWS session
import (
	"github.com/aws/aws-sdk-go/aws/session"
)

func createSession() *session.Session {
	// Create a new session that the SDK will use to load
	// credentials from the shared credentials file.
	sess := session.Must(session.NewSession())

	return sess
}
