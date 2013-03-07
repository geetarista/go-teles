package teles

type TelesError struct {
	ErrorString string
}

func (err *TelesError) Error() string {
	return err.ErrorString
}

func errInvalidResponse(resp string) error {
	return &TelesError{ErrorString: "Got response: " + resp}
}

func errCommandFailed(cmd, attempt string) error {
	return &TelesError{
		ErrorString: "Failed to send command to teles server: " + cmd + ". Attempt: " + attempt,
	}
}

func errSendFailed(cmd, attempts string) error {
	return &TelesError{
		ErrorString: "Failed to send command '" + cmd + "' after " + attempts + " attempts!",
	}
}
