// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// BindToDevice_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type BindToDevice_genType func(fd int, device string) (err error)

// MoqBindToDevice_genType holds the state of a moq of the BindToDevice_genType
// type
type MoqBindToDevice_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqBindToDevice_genType_mock

	ResultsByParams []MoqBindToDevice_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Fd     moq.ParamIndexing
			Device moq.ParamIndexing
		}
	}
}

// MoqBindToDevice_genType_mock isolates the mock interface of the
// BindToDevice_genType type
type MoqBindToDevice_genType_mock struct {
	Moq *MoqBindToDevice_genType
}

// MoqBindToDevice_genType_params holds the params of the BindToDevice_genType
// type
type MoqBindToDevice_genType_params struct {
	Fd     int
	Device string
}

// MoqBindToDevice_genType_paramsKey holds the map key params of the
// BindToDevice_genType type
type MoqBindToDevice_genType_paramsKey struct {
	Params struct {
		Fd     int
		Device string
	}
	Hashes struct {
		Fd     hash.Hash
		Device hash.Hash
	}
}

// MoqBindToDevice_genType_resultsByParams contains the results for a given set
// of parameters for the BindToDevice_genType type
type MoqBindToDevice_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqBindToDevice_genType_paramsKey]*MoqBindToDevice_genType_results
}

// MoqBindToDevice_genType_doFn defines the type of function needed when
// calling AndDo for the BindToDevice_genType type
type MoqBindToDevice_genType_doFn func(fd int, device string)

// MoqBindToDevice_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the BindToDevice_genType type
type MoqBindToDevice_genType_doReturnFn func(fd int, device string) (err error)

// MoqBindToDevice_genType_results holds the results of the
// BindToDevice_genType type
type MoqBindToDevice_genType_results struct {
	Params  MoqBindToDevice_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqBindToDevice_genType_doFn
		DoReturnFn MoqBindToDevice_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqBindToDevice_genType_fnRecorder routes recorded function calls to the
// MoqBindToDevice_genType moq
type MoqBindToDevice_genType_fnRecorder struct {
	Params    MoqBindToDevice_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqBindToDevice_genType_results
	Moq       *MoqBindToDevice_genType
}

// MoqBindToDevice_genType_anyParams isolates the any params functions of the
// BindToDevice_genType type
type MoqBindToDevice_genType_anyParams struct {
	Recorder *MoqBindToDevice_genType_fnRecorder
}

// NewMoqBindToDevice_genType creates a new moq of the BindToDevice_genType
// type
func NewMoqBindToDevice_genType(scene *moq.Scene, config *moq.Config) *MoqBindToDevice_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqBindToDevice_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqBindToDevice_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Fd     moq.ParamIndexing
				Device moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Fd     moq.ParamIndexing
			Device moq.ParamIndexing
		}{
			Fd:     moq.ParamIndexByValue,
			Device: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the BindToDevice_genType type
func (m *MoqBindToDevice_genType) Mock() BindToDevice_genType {
	return func(fd int, device string) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqBindToDevice_genType_mock{Moq: m}
		return moq.Fn(fd, device)
	}
}

func (m *MoqBindToDevice_genType_mock) Fn(fd int, device string) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqBindToDevice_genType_params{
		Fd:     fd,
		Device: device,
	}
	var results *MoqBindToDevice_genType_results
	for _, resultsByParams := range m.Moq.ResultsByParams {
		paramsKey := m.Moq.ParamsKey(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(fd, device)
	}

	if result.Values != nil {
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(fd, device)
	}
	return
}

func (m *MoqBindToDevice_genType) OnCall(fd int, device string) *MoqBindToDevice_genType_fnRecorder {
	return &MoqBindToDevice_genType_fnRecorder{
		Params: MoqBindToDevice_genType_params{
			Fd:     fd,
			Device: device,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqBindToDevice_genType_fnRecorder) Any() *MoqBindToDevice_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqBindToDevice_genType_anyParams{Recorder: r}
}

func (a *MoqBindToDevice_genType_anyParams) Fd() *MoqBindToDevice_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqBindToDevice_genType_anyParams) Device() *MoqBindToDevice_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqBindToDevice_genType_fnRecorder) Seq() *MoqBindToDevice_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqBindToDevice_genType_fnRecorder) NoSeq() *MoqBindToDevice_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqBindToDevice_genType_fnRecorder) ReturnResults(err error) *MoqBindToDevice_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqBindToDevice_genType_doFn
		DoReturnFn MoqBindToDevice_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqBindToDevice_genType_fnRecorder) AndDo(fn MoqBindToDevice_genType_doFn) *MoqBindToDevice_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqBindToDevice_genType_fnRecorder) DoReturnResults(fn MoqBindToDevice_genType_doReturnFn) *MoqBindToDevice_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqBindToDevice_genType_doFn
		DoReturnFn MoqBindToDevice_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqBindToDevice_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqBindToDevice_genType_resultsByParams
	for n, res := range r.Moq.ResultsByParams {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqBindToDevice_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqBindToDevice_genType_paramsKey]*MoqBindToDevice_genType_results{},
		}
		r.Moq.ResultsByParams = append(r.Moq.ResultsByParams, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams) {
			copy(r.Moq.ResultsByParams[insertAt+1:], r.Moq.ResultsByParams[insertAt:0])
			r.Moq.ResultsByParams[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqBindToDevice_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqBindToDevice_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqBindToDevice_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults or DoReturnResults must be called before calling Repeat")
		return nil
	}
	r.Results.Repeat.Repeat(r.Moq.Scene.T, repeaters)
	last := r.Results.Results[len(r.Results.Results)-1]
	for n := 0; n < r.Results.Repeat.ResultCount-1; n++ {
		if r.Sequence {
			last = struct {
				Values     *struct{ Err error }
				Sequence   uint32
				DoFn       MoqBindToDevice_genType_doFn
				DoReturnFn MoqBindToDevice_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqBindToDevice_genType) PrettyParams(params MoqBindToDevice_genType_params) string {
	return fmt.Sprintf("BindToDevice_genType(%#v, %#v)", params.Fd, params.Device)
}

func (m *MoqBindToDevice_genType) ParamsKey(params MoqBindToDevice_genType_params, anyParams uint64) MoqBindToDevice_genType_paramsKey {
	m.Scene.T.Helper()
	var fdUsed int
	var fdUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Fd == moq.ParamIndexByValue {
			fdUsed = params.Fd
		} else {
			fdUsedHash = hash.DeepHash(params.Fd)
		}
	}
	var deviceUsed string
	var deviceUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Device == moq.ParamIndexByValue {
			deviceUsed = params.Device
		} else {
			deviceUsedHash = hash.DeepHash(params.Device)
		}
	}
	return MoqBindToDevice_genType_paramsKey{
		Params: struct {
			Fd     int
			Device string
		}{
			Fd:     fdUsed,
			Device: deviceUsed,
		},
		Hashes: struct {
			Fd     hash.Hash
			Device hash.Hash
		}{
			Fd:     fdUsedHash,
			Device: deviceUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqBindToDevice_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqBindToDevice_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams(results.Params))
			}
		}
	}
}