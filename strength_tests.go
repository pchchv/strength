package strength

import "testing"

func TestValidate(t *testing.T) {
	err := Validate("mypass", 50)
	expectedError := "insecure password, try including more special characters, using uppercase letters, using numbers or using a longer password"
	if err.Error() != expectedError {
		t.Errorf("Wanted %v, got %v", expectedError, err)
	}

	err = Validate("MYPASS", 50)
	expectedError = "insecure password, try including more special characters, using lowercase letters, using numbers or using a longer password"
	if err.Error() != expectedError {
		t.Errorf("Wanted %v, got %v", expectedError, err)
	}

	err = Validate("mypassword", 4)
	if err != nil {
		t.Errorf("Err should be nil")
	}

	err = Validate("aGoo0dMi#oFChaR2", 80)
	if err != nil {
		t.Errorf("Err should be nil")
	}

	expectedError = "insecure password, try including more special characters, using lowercase letters, using uppercase letters or using a longer password"
	err = Validate("123", 60)
	if err.Error() != expectedError {
		t.Errorf("Wanted %v, got %v", expectedError, err)
	}
}
