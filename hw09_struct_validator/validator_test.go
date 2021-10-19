package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{App{Version: "xxxx"}, ValidationErrors{{Field: "Version", Err: ErrInvalidLen}}},
		{App{Version: "xxxxx"}, nil},
		{App{Version: "xxxxxx"}, ValidationErrors{{Field: "Version", Err: ErrInvalidLen}}},
		{User{ID: "AAAAAA-BBBBBB-CCCCCC-DDDDDD-EEEEEE", Age: 65, Email: "aaa bbb", Role: "guest", Phones: []string{"123456789013"}, meta: []byte("xxx")}, ValidationErrors{{Field: "ID", Err: ErrInvalidLen}, {Field: "Age", Err: ErrValueGreatMax}, {Field: "Email", Err: ErrNoMatchRegexp}, {Field: "Role", Err: ErrValueNotInSet}, {Field: "Phones", Err: ErrInvalidLen}}},
		{User{ID: "AAAAAA-BBBBBB-CCCCCC-DDDDDD-EEEEEE-F", Age: 30, Email: "mybox@gmail.com", Role: "admin", Phones: []string{"12345678901"}}, nil},
		{Response{Code: 200, Body: ""}, nil},
		{Response{Code: 888, Body: ""}, ValidationErrors{{Field: "Code", Err: ErrValueNotInSet}}},
		{Token{[]byte("fff"), []byte("ddd"), []byte("xxx")}, nil},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.in)
			if tt.expectedErr == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}
