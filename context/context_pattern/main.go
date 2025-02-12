package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func logic(ctx context.Context, info string) (string, error){
	return "", nil
}

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		//wrap it with stuff
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte("error parsing form"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := r.FormValue("data")
	result, err := logic(ctx, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(result))
}

type ServiceCaller struct {
	client *http.Client
}

func (sc ServiceCaller) call(ctx context.Context, data string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"http://example.com?data="+data, nil)
	if err != nil {
		return "", err
	}
	resp, err := sc.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected error, status code %d", resp.StatusCode)
	}
	return processresp(resp.Body)
}

func processresp(body io.ReadCloser) (string, error) {
	return "", nil
}

func main() {
	ctx := context.Background()
	result, err := logic(ctx, "a string")
	fmt.Println(result, err)
}