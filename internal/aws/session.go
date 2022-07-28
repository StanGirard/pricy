package aws

// Import AWS session
import (
	"flag"

	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	sso = flag.Bool("sso", false, "Use AWS SSO")
)

func createSessionWithCredentials() *session.Session {
	// Create a new session that the SDK will use to load
	// credentials from credentials
	sess := session.Must(session.NewSession())

	return sess
}

func createSessionWithSSO() *session.Session {
	// Create a new session that the SDK will use to load
	// credentials from the shared credentials file.
	// Usefull for SSO
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return sess
}

func InitSession() *session.Session {
	// Create a new session that the SDK will use to load
	// credentials from. With either SSO or credentials
	flag.Parse()
	if *sso {
		return createSessionWithSSO()
	} else {
		return createSessionWithCredentials()
	}
}
