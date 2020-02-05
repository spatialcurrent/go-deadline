// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package deadline

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	Exit      = func(ctx context.Context) { os.Exit(0) }
	ExitError = func(ctx context.Context) { os.Exit(1) }
)

type Deadline struct {
	duration  time.Duration
	function  func(ctx context.Context)
	started   bool
	cancelled bool
	finished  bool
	mutex     *sync.Mutex
}

func (d *Deadline) Start(ctx context.Context) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if d.started {
		return errors.New("cannot start deadline: deadline was already started")
	}
	d.started = true
	go func() {
		time.Sleep(d.duration)
		d.mutex.Lock()
		defer d.mutex.Unlock()
		if !d.cancelled {
			d.function(ctx)
		}
		d.finished = true
	}()
	return nil
}

func (d *Deadline) Cancel() error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if d.cancelled {
		return errors.New("cannot cancel deadline: deadline was already cancelled")
	}
	if d.finished {
		return errors.New("cannot cancel deadline: deadline is already finished")
	}
	d.cancelled = true
	return nil
}

func (d *Deadline) Started() bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.started
}

func (d *Deadline) Cancelled() bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.cancelled
}

func (d *Deadline) Finished() bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.finished
}

func (d *Deadline) Reset() {
	d.started = false
	d.cancelled = false
	d.finished = false
}

func New(duration time.Duration, f func(ctx context.Context)) (*Deadline, error) {
	if duration <= 0 {
		return nil, fmt.Errorf("invalid duration, expecting value greater than zero, but found %s", duration)
	}
	d := &Deadline{
		duration: duration,
		function: f,
		mutex:    &sync.Mutex{},
	}
	return d, nil
}
