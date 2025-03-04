package theory

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserViewHandler(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		user        HttpTestUser
	}
	tests := []struct {
		name    string
		request string
		users   map[string]HttpTestUser
		want    want
	}{
		{
			name: "simple test #1",
			users: map[string]HttpTestUser{
				"id1": {
					ID:        "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			want: want{
				contentType: "application/json",
				statusCode:  200,
				user: HttpTestUser{ID: "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			request: "/users?user_id=id1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, tt.request, nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(UserViewHandler(tt.users))
			h(w, request)

			result := w.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

			userResult, err := io.ReadAll(result.Body)
			require.NoError(t, err)
			err = result.Body.Close()
			require.NoError(t, err)

			var user HttpTestUser
			err = json.Unmarshal(userResult, &user)
			require.NoError(t, err)

			assert.Equal(t, tt.want.user, user)
		})
	}
}
