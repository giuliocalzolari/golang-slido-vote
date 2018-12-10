package main

import (
	"reflect"
	"testing"
)

func TestVoteUp(t *testing.T) {
	type args struct {
		auth       AuthData
		questionID string
	}
	tests := []struct {
		name string
		args args
		want VoteData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VoteUp(tt.args.auth, tt.args.questionID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VoteUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthEvent(t *testing.T) {
	type args struct {
		eventUUID string
	}
	tests := []struct {
		name string
		args args
		want AuthData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AuthEvent(tt.args.eventUUID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEventData(t *testing.T) {
	type args struct {
		eventCode string
	}
	tests := []struct {
		name string
		args args
		want EventData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEventData(tt.args.eventCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEventData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printUsageErrorAndExit(t *testing.T) {
	type args struct {
		format string
		values []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printUsageErrorAndExit(tt.args.format, tt.args.values...)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
