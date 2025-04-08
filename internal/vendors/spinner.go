package spinner

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

type Spinner struct {
	spinner *spinner.Spinner
	active  bool
	startTime  time.Time
	minShowTime time.Duration
}


type SpinnerOptions struct {
	Message   string
	SpinSpeed time.Duration
	SpinType  int
	MinShowTime time.Duration
}


func DefaultOptions() SpinnerOptions {
	return SpinnerOptions{
		Message:   "Working...",
		SpinSpeed: 100 * time.Millisecond,
		SpinType:  9, 
		MinShowTime: 500 * time.Millisecond,
		
	}
}

func New(options SpinnerOptions) *Spinner {
	s := spinner.New(spinner.CharSets[options.SpinType], options.SpinSpeed)
	s.Suffix = " " + options.Message
	return &Spinner{
		spinner: s,
		active:  false,
		minShowTime: options.MinShowTime,
	}
}

func (s *Spinner) Start() {
	if !s.active {
		s.spinner.Start()
		s.startTime = time.Now()
		s.active = true
	}
}

func (s *Spinner) Stop() {
	if s.active {

		elapsed := time.Since(s.startTime)
		if elapsed < s.minShowTime {
			time.Sleep(s.minShowTime - elapsed)
		}

		s.spinner.Stop()
		s.active = false
	}
}


func (s *Spinner) UpdateMessage(message string) {
	s.spinner.Suffix = " " + message
}


func (s *Spinner) UpdateMessagef(format string, args ...interface{}) {
	s.spinner.Suffix = " " + fmt.Sprintf(format, args...)
}


func StartWithMessage(message string) *Spinner {
	options := DefaultOptions()
	options.Message = message
	s := New(options)
	s.Start()
	return s
}

func WithContext(message string, fn func() error) error {
	s := StartWithMessage(message)
	defer s.Stop()
	
	return fn()
}

type Step struct {
	Name     string
	Action   func() error
	Progress int
}

func RunSteps(steps []Step, totalMessage string) error {
	s := StartWithMessage(totalMessage)
	defer s.Stop()
	
	total := len(steps)
	for i, step := range steps {
		s.UpdateMessagef("[%d/%d] %s", i+1, total, step.Name)
		if err := step.Action(); err != nil {
			return err
		}
	}
	
	return nil
}