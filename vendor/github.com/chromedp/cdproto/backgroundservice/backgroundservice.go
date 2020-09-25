// Package backgroundservice provides the Chrome DevTools Protocol
// commands, types, and events for the BackgroundService domain.
//
// Defines events for background web platform features.
//
// Generated by the cdproto-gen command.
package backgroundservice

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// StartObservingParams enables event updates for the service.
type StartObservingParams struct {
	Service ServiceName `json:"service"`
}

// StartObserving enables event updates for the service.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService#method-startObserving
//
// parameters:
//   service
func StartObserving(service ServiceName) *StartObservingParams {
	return &StartObservingParams{
		Service: service,
	}
}

// Do executes BackgroundService.startObserving against the provided context.
func (p *StartObservingParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandStartObserving, p, nil)
}

// StopObservingParams disables event updates for the service.
type StopObservingParams struct {
	Service ServiceName `json:"service"`
}

// StopObserving disables event updates for the service.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService#method-stopObserving
//
// parameters:
//   service
func StopObserving(service ServiceName) *StopObservingParams {
	return &StopObservingParams{
		Service: service,
	}
}

// Do executes BackgroundService.stopObserving against the provided context.
func (p *StopObservingParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandStopObserving, p, nil)
}

// SetRecordingParams set the recording state for the service.
type SetRecordingParams struct {
	ShouldRecord bool        `json:"shouldRecord"`
	Service      ServiceName `json:"service"`
}

// SetRecording set the recording state for the service.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService#method-setRecording
//
// parameters:
//   shouldRecord
//   service
func SetRecording(shouldRecord bool, service ServiceName) *SetRecordingParams {
	return &SetRecordingParams{
		ShouldRecord: shouldRecord,
		Service:      service,
	}
}

// Do executes BackgroundService.setRecording against the provided context.
func (p *SetRecordingParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetRecording, p, nil)
}

// ClearEventsParams clears all stored data for the service.
type ClearEventsParams struct {
	Service ServiceName `json:"service"`
}

// ClearEvents clears all stored data for the service.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService#method-clearEvents
//
// parameters:
//   service
func ClearEvents(service ServiceName) *ClearEventsParams {
	return &ClearEventsParams{
		Service: service,
	}
}

// Do executes BackgroundService.clearEvents against the provided context.
func (p *ClearEventsParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandClearEvents, p, nil)
}

// Command names.
const (
	CommandStartObserving = "BackgroundService.startObserving"
	CommandStopObserving  = "BackgroundService.stopObserving"
	CommandSetRecording   = "BackgroundService.setRecording"
	CommandClearEvents    = "BackgroundService.clearEvents"
)
