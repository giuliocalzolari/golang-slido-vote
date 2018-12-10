package main

import (
	"reflect"
	"testing"
)

func TestVoteUp(t *testing.T) {
	type args struct {
		auth        AuthData
		question_id string
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
			if got := VoteUp(tt.args.auth, tt.args.question_id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VoteUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthEvent(t *testing.T) {
	type args struct {
		event_uuid string
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
			if got := AuthEvent(tt.args.event_uuid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEventData(t *testing.T) {
	type args struct {
		event_code string
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
			if got := GetEventData(tt.args.event_code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEventData() = %v, want %v", got, tt.want)
			}
		})
	}
}
