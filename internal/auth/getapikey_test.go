package auth

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func Test_GetAPIKey_AuthHeaderEmpty(t *testing.T) {
	//Arrange
	var headerMock = http.Header{}
	headerMock.Add("Authorization", "")
	//Act
	got, _ := GetAPIKey(headerMock)
	want := ""
	//Assert
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func Test_GetAPIKey_AuthHeaderNotLongEnough(t *testing.T) {
	//Arrange
	var headerMock = http.Header{}
	headerMock.Add("Authorization", "A")
	//Act
	got, _ := GetAPIKey(headerMock)
	want := ""
	//Assert
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func Test_GetAPIKey_AuthHeaderIsCorrectButEmpty(t *testing.T) {
	//Arrange
	var headerMock = http.Header{}
	headerMock.Add("Authorization", "ApiKey")
	//Act
	got, _ := GetAPIKey(headerMock)
	want := ""
	//Assert
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func Test_GetAPIKey_AuthHeaderIsCorrect(t *testing.T) {
	//Arrange
	var headerMock = http.Header{}
	headerMock.Add("Authorization", fmt.Sprintf("ApiKey %v", TestHeaderValue))
	//Act
	got, _ := GetAPIKey(headerMock)
	want := "AnythingElse"
	//Assert
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
