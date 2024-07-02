package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_asciiArtLiveHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantBody       string
	}{
		{
			name: "Valid request",
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					form := url.Values{}
					form.Add("text", "Hello, World!")
					form.Add("banner", "standard")
					req := httptest.NewRequest(http.MethodPost, "/ascii-art-live", bytes.NewBufferString(form.Encode()))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					return req
				}(),
			},
			wantStatusCode: http.StatusOK,
			wantBody:       "  _   _      _ _        __        __         _     _ \n | | | | ___| | | ___   __\\ \\      / /__  _ __| | __| |\n | |_| |/ _ \\ | |/ _ \\ / _` \\ \\ /\\ / / _ \\| '__| |/ _` |\n |  _  |  __/ | | (_) | (_| |\\ V  V / (_) | |  | | (_| |\n |_| |_|\\___|_|_|\\___/ \\__,_| \\_/\\_/ \\___/|_|  |_|\\__,_|\n",
		},
		{
			name: "Missing text parameter",
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					form := url.Values{}
					form.Add("banner", "standard")
					req := httptest.NewRequest(http.MethodPost, "/ascii-art-live", bytes.NewBufferString(form.Encode()))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					return req
				}(),
			},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "Bad Request\n",
		},
		{
			name: "Missing banner parameter",
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					form := url.Values{}
					form.Add("text", "Hello, World!")
					req := httptest.NewRequest(http.MethodPost, "/ascii-art-live", bytes.NewBufferString(form.Encode()))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					return req
				}(),
			},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "Bad Request\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asciiArtLiveHandler(tt.args.w, tt.args.r)
			rec := tt.args.w.(*httptest.ResponseRecorder)
			if rec.Code != tt.wantStatusCode {
				t.Errorf("asciiArtLiveHandler() status code = %v, want %v", rec.Code, tt.wantStatusCode)
			}
			if rec.Body.String() != tt.wantBody {
				t.Errorf("asciiArtLiveHandler() body = %v, want %v", rec.Body.String(), tt.wantBody)
			}
		})
	}
}

func Test_asciiArtHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantBody       string
	}{
		{
			name: "Valid request",
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					form := url.Values{}
					form.Add("text", "Hello, World!")
					form.Add("banner", "standard")
					req := httptest.NewRequest(http.MethodPost, "/ascii", bytes.NewBufferString(form.Encode()))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					return req
				}(),
			},
			wantStatusCode: http.StatusOK,
			wantBody:       "  _   _      _ _        __        __         _     _ \n | | | | ___| | | ___   __\\ \\      / /__  _ __| | __| |\n | |_| |/ _ \\ | |/ _ \\ / _` \\ \\ /\\ / / _ \\| '__| |/ _` |\n |  _  |  __/ | | (_) | (_| |\\ V  V / (_) | |  | | (_| |\n |_| |_|\\___|_|_|\\___/ \\__,_| \\_/\\_/ \\___/|_|  |_|\\__,_|\n",
		},
		{
			name: "Missing text parameter",
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					form := url.Values{}
					form.Add("banner", "standard")
					req := httptest.NewRequest(http.MethodPost, "/ascii", bytes.NewBufferString(form.Encode()))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					return req
				}(),
			},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "Bad Request\n",
		},
		{
			name: "Missing banner parameter",
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					form := url.Values{}
					form.Add("text", "Hello, World!")
					req := httptest.NewRequest(http.MethodPost, "/ascii", bytes.NewBufferString(form.Encode()))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					return req
				}(),
			},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "Bad Request\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asciiArtHandler(tt.args.w, tt.args.r)
			rec := tt.args.w.(*httptest.ResponseRecorder)
			if rec.Code != tt.wantStatusCode {
				t.Errorf("asciiArtHandler() status code = %v, want %v", rec.Code, tt.wantStatusCode)
			}
			if rec.Body.String() != tt.wantBody {
				t.Errorf("asciiArtHandler() body = %v, want %v", rec.Body.String(), tt.wantBody)
			}
		})
	}
}

func Test_indexHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
	}{
		{
			name: "Valid GET request",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/", nil),
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "Invalid POST request",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/", nil),
			},
			wantStatusCode: http.StatusMethodNotAllowed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indexHandler(tt.args.w, tt.args.r)
			rec := tt.args.w.(*httptest.ResponseRecorder)
			if rec.Code != tt.wantStatusCode {
				t.Errorf("indexHandler() status code = %v, want %v", rec.Code, tt.wantStatusCode)
			}
		})
	}
}
