package app

import (
	"time"
	"timestamps-assignment/api/domain"
)

const (
	layout = "20060102T150405Z"
)

type TimestampDataService struct {
}

func NewTimestampDataService() *TimestampDataService {
	return &TimestampDataService{}
}
func (s *TimestampDataService) TimestampsCalculation(interval string, timezone string, start string, end string) ([]string, *domain.ErrorResponse) {
	intervals := []string{"1h", "1d", "1mo", "1y"}
	timeZone, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, &domain.ErrorResponse{
			Status:      "Invalid location",
			Description: err.Error(),
		}
	}
	startTimestamp, err := time.Parse(layout, start)
	if err != nil {
		return nil, &domain.ErrorResponse{
			Status:      "Invalid start timestamp",
			Description: err.Error(),
		}
	}
	endTimestamp, err := time.Parse(layout, end)
	if err != nil {
		return nil, &domain.ErrorResponse{
			Status:      "Invalid end timestamp",
			Description: err.Error(),
		}
	}

	if endTimestamp.Before(startTimestamp) {
		return nil, &domain.ErrorResponse{
			Status:      "Invalid period",
			Description: "End date must be after start date",
		}
	}
	if !contains(intervals, interval) {
		return nil, &domain.ErrorResponse{
			Status:      "Unsupported period",
			Description: "Period selected must be on of {1h, 1d, 1mo, 1y}",
		}
	}

	t := startTimestamp.Add(time.Hour)
	t = t.In(timeZone)
	endTimestamp = endTimestamp.In(timeZone)
	t = truncateToHour(t)
	var intervalTimestamps []string
	for t.Before(endTimestamp) {
		if interval == "1mo" {
			if lastDateOfMonth(t).Before(endTimestamp) {
				intervalTimestamps = append(intervalTimestamps, lastDateOfMonth(t).In(timeZone).UTC().Format(layout))
			}
		} else if interval == "1y" {
			if lastDateOfYear(t).Before(endTimestamp) {
				intervalTimestamps = append(intervalTimestamps, lastDateOfYear(t).In(timeZone).UTC().Format(layout))
			}
		} else {

			intervalTimestamps = append(intervalTimestamps, t.In(timeZone).UTC().Format(layout))
		}
		t = addInterval(interval, t)

	}

	return intervalTimestamps, nil
}
