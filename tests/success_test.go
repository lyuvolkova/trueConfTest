package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"refactoring/internal"
	"refactoring/internal/router"
	"refactoring/internal/storage"
)

func TestUsers(t *testing.T) {
	store := getNewFile(t)
	defer os.Remove(store)

	repo := storage.NewRepository(store)
	err := repo.Load()
	if err != nil {
		t.Fatal(err)
	}

	r := router.Router(repo)

	ts := httptest.NewServer(r)
	defer ts.Close()

	url := ts.URL + "/api/v1/users"
	t.Run("test1: create user", func(t *testing.T) {
		data := []byte(
			`{
					"display_name": "myTEST1",
					"email": "test1"
					}`)
		b := bytes.NewReader(data)
		res, err := http.Post(url, "application/json", b)
		if err != nil {
			t.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		res.Body.Close()
		s := map[string]string{}
		err = json.Unmarshal(body, &s)
		if err != nil {
			t.Fatal("don`t unmarshal file", err)
		}

		url += "/" + s["user_id"]
		user := findUser(t, url)
		if user.Email != "test1" && user.DisplayName != "myTEST1" {
			log.Fatal("user is not equal expected ")
		}
	})

	url = ts.URL + "/api/v1/users"
	t.Run("test2: delete user", func(t *testing.T) {
		url += "/4"
		data := []byte(``)
		b := bytes.NewReader(data)
		req, err := http.NewRequest("DELETE", url, b)
		if err != nil {
			t.Fatal(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != 200 {
			t.Fatal("user don`t delete")
		}

		res, err = http.Get(url)
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != 404 {
			t.Fatal("user don`t delete")
		}
		res.Body.Close()
	})

	url = ts.URL + "/api/v1/users"
	t.Run("test3: update user", func(t *testing.T) {
		url += "/1"
		data := []byte(`{"display_name": "UPDATE_USER"}`)
		b := bytes.NewReader(data)

		req, err := http.NewRequest("PATCH", url, b)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		_, err = http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		user := findUser(t, url)
		if user.DisplayName != "UPDATE_USER" {
			t.Fatal("user is not update ", err)
		}
	})
}

func getNewFile(t *testing.T) string {
	src, err := os.Open("../users.json")
	if err != nil {
		t.Fatal(err)
	}
	defer src.Close()

	file, err := os.CreateTemp(t.TempDir(), "users.json")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, src)
	if err != nil {
		t.Fatal(err)
	}
	return file.Name()
}

func findUser(t *testing.T, url string) internal.User {
	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()
	user := internal.User{}
	err = json.Unmarshal(body, &user)

	return user
}
