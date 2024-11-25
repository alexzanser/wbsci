package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloWorldServer(t *testing.T) {
	// Создаем тестовый HTTP-сервер
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Hello, World!")
		if err != nil {
			t.Errorf("unexpected error writing response: %v", err)
		}
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	// Выполняем GET-запрос к тестовому серверу
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("failed to send GET request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Проверяем тело ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	expected := "Hello, World!\n"
	if strings.TrimSpace(string(body)) != strings.TrimSpace(expected) {
		t.Errorf("expected response body %q, got %q", expected, string(body))
	}
}
