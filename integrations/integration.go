package integrations

// Integration defines what an integration is made of.
// Each integrations is responsible to register it self to the registry (see
// RegisterIntegration for details).
type Integration interface {
	// Identify is responsible of forwarding the identify call to the integration
	Identify(event Event) error

	// Enabled returns wether or not the integration is enabled/configured
	Enabled() bool
}

// Event defines the structure of the data we receive from the API
type Event struct {
	UserID     string            `json:"userID"`
	UserTraits map[string]string `json:"userTraits"`
	Timestamp  int64             `json:"timestamp"`
	ReceivedAt int64             `json:"receivedAt"`
}

// Validate the content of the event to be sure it has everything that's need
func (e Event) Validate() (missingParameters []string) {
	if e.UserID == "" {
		missingParameters = append(missingParameters, "userID")
	}
	if e.Timestamp == 0 {
		missingParameters = append(missingParameters, "timestamp")
	}
	if e.ReceivedAt == 0 {
		missingParameters = append(missingParameters, "receivedAt")
	}
	return
}
