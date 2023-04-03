package app

import (
	"errors"
	"reflect"
	"testing"
)

func Test_TimestampsMatching_1h_OK(t *testing.T) {
	period := "1h"
	timezone := "Europe/Athens"
	t1 := "20210714T204603Z"
	t2 := "20210715T123456Z"

	service := NewTimestampDataService()

	result, err := service.TimestampsCalculation(period, timezone, t1, t2)
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if len(result) != len(Test_TimestampsMatching_1h_OK_Data) {
		t.Errorf("Result items %d not equal to expected items %d", len(result), len(Test_TimestampsMatching_1h_OK_Data))
	}
	if !reflect.DeepEqual(result, Test_TimestampsMatching_1h_OK_Data) {
		t.Errorf("Result data not matching expected data")
	}
}

func Test_TimestampsMatching_1mo_OK(t *testing.T) {
	period := "1mo"
	timezone := "Europe/Athens"
	t1 := "20210214T204603Z"
	t2 := "20211115T123456Z"

	service := NewTimestampDataService()

	result, err := service.TimestampsCalculation(period, timezone, t1, t2)
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if len(result) != len(Test_TimestampsMatching_1mo_OK_Data) {
		t.Errorf("Result items %d not equal to expected items %d", len(result), len(Test_TimestampsMatching_1mo_OK_Data))
	}
	if !reflect.DeepEqual(result, Test_TimestampsMatching_1mo_OK_Data) {
		t.Errorf("Result data not matching expected data")
	}
}

func Test_TimestampsMatching_UnsupportedPeriod(t *testing.T) {
	period := "1w"
	timezone := "Europe/Athens"
	t1 := "20210714T204603Z"
	t2 := "20210715T123456Z"

	service := NewTimestampDataService()
	expectedError := errors.New("Unsupported period")
	_, err := service.TimestampsCalculation(period, timezone, t1, t2)
	if err == nil {
		t.Error("Expected error: ")
	}
	if err.Status != expectedError.Error() {
		t.Errorf("Error %s not equal to expected error %s", err.Status, expectedError.Error())
	}
}

func Test_TimestampsMatching_InvalidPeriod(t *testing.T) {
	period := "1w"
	timezone := "Europe/Athens"
	t1 := "20210715T123456Z"
	t2 := "20210714T204603Z"

	service := NewTimestampDataService()
	expectedError := errors.New("Invalid period")
	_, err := service.TimestampsCalculation(period, timezone, t1, t2)
	if err == nil {
		t.Error("Expected error: ")
	}
	if err.Status != expectedError.Error() {
		t.Errorf("Error %s not equal to expected error %s", err.Status, expectedError.Error())
	}
}

func Test_TimestampsMatching_InvalidTimezone(t *testing.T) {
	period := "1d"
	timezone := "Asia/Athens"
	t1 := "20210714T204603Z"
	t2 := "20210715T123456Z"

	service := NewTimestampDataService()
	expectedError := errors.New("Invalid location")
	_, err := service.TimestampsCalculation(period, timezone, t1, t2)
	if err == nil {
		t.Error("Expected error: ")
	}
	if err.Status != expectedError.Error() {
		t.Errorf("Error %s not equal to expected error %s", err.Status, expectedError.Error())
	}
}
