// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package iotest

import (
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ErrReader_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type ErrReader_genType func(err error) io.Reader

// MoqErrReader_genType holds the state of a moq of the ErrReader_genType type
type MoqErrReader_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqErrReader_genType_mock

	ResultsByParams []MoqErrReader_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Err moq.ParamIndexing
		}
	}
}

// MoqErrReader_genType_mock isolates the mock interface of the
// ErrReader_genType type
type MoqErrReader_genType_mock struct {
	Moq *MoqErrReader_genType
}

// MoqErrReader_genType_params holds the params of the ErrReader_genType type
type MoqErrReader_genType_params struct{ Err error }

// MoqErrReader_genType_paramsKey holds the map key params of the
// ErrReader_genType type
type MoqErrReader_genType_paramsKey struct {
	Params struct{ Err error }
	Hashes struct{ Err hash.Hash }
}

// MoqErrReader_genType_resultsByParams contains the results for a given set of
// parameters for the ErrReader_genType type
type MoqErrReader_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqErrReader_genType_paramsKey]*MoqErrReader_genType_results
}

// MoqErrReader_genType_doFn defines the type of function needed when calling
// AndDo for the ErrReader_genType type
type MoqErrReader_genType_doFn func(err error)

// MoqErrReader_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ErrReader_genType type
type MoqErrReader_genType_doReturnFn func(err error) io.Reader

// MoqErrReader_genType_results holds the results of the ErrReader_genType type
type MoqErrReader_genType_results struct {
	Params  MoqErrReader_genType_params
	Results []struct {
		Values *struct {
			Result1 io.Reader
		}
		Sequence   uint32
		DoFn       MoqErrReader_genType_doFn
		DoReturnFn MoqErrReader_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqErrReader_genType_fnRecorder routes recorded function calls to the
// MoqErrReader_genType moq
type MoqErrReader_genType_fnRecorder struct {
	Params    MoqErrReader_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqErrReader_genType_results
	Moq       *MoqErrReader_genType
}

// MoqErrReader_genType_anyParams isolates the any params functions of the
// ErrReader_genType type
type MoqErrReader_genType_anyParams struct {
	Recorder *MoqErrReader_genType_fnRecorder
}

// NewMoqErrReader_genType creates a new moq of the ErrReader_genType type
func NewMoqErrReader_genType(scene *moq.Scene, config *moq.Config) *MoqErrReader_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqErrReader_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqErrReader_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Err moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Err moq.ParamIndexing
		}{
			Err: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ErrReader_genType type
func (m *MoqErrReader_genType) Mock() ErrReader_genType {
	return func(err error) io.Reader {
		m.Scene.T.Helper()
		moq := &MoqErrReader_genType_mock{Moq: m}
		return moq.Fn(err)
	}
}

func (m *MoqErrReader_genType_mock) Fn(err error) (result1 io.Reader) {
	m.Moq.Scene.T.Helper()
	params := MoqErrReader_genType_params{
		Err: err,
	}
	var results *MoqErrReader_genType_results
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
		result.DoFn(err)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(err)
	}
	return
}

func (m *MoqErrReader_genType) OnCall(err error) *MoqErrReader_genType_fnRecorder {
	return &MoqErrReader_genType_fnRecorder{
		Params: MoqErrReader_genType_params{
			Err: err,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqErrReader_genType_fnRecorder) Any() *MoqErrReader_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqErrReader_genType_anyParams{Recorder: r}
}

func (a *MoqErrReader_genType_anyParams) Err() *MoqErrReader_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqErrReader_genType_fnRecorder) Seq() *MoqErrReader_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqErrReader_genType_fnRecorder) NoSeq() *MoqErrReader_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqErrReader_genType_fnRecorder) ReturnResults(result1 io.Reader) *MoqErrReader_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 io.Reader
		}
		Sequence   uint32
		DoFn       MoqErrReader_genType_doFn
		DoReturnFn MoqErrReader_genType_doReturnFn
	}{
		Values: &struct {
			Result1 io.Reader
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqErrReader_genType_fnRecorder) AndDo(fn MoqErrReader_genType_doFn) *MoqErrReader_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqErrReader_genType_fnRecorder) DoReturnResults(fn MoqErrReader_genType_doReturnFn) *MoqErrReader_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 io.Reader
		}
		Sequence   uint32
		DoFn       MoqErrReader_genType_doFn
		DoReturnFn MoqErrReader_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqErrReader_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqErrReader_genType_resultsByParams
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
		results = &MoqErrReader_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqErrReader_genType_paramsKey]*MoqErrReader_genType_results{},
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
		r.Results = &MoqErrReader_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqErrReader_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqErrReader_genType_fnRecorder {
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
				Values *struct {
					Result1 io.Reader
				}
				Sequence   uint32
				DoFn       MoqErrReader_genType_doFn
				DoReturnFn MoqErrReader_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqErrReader_genType) PrettyParams(params MoqErrReader_genType_params) string {
	return fmt.Sprintf("ErrReader_genType(%#v)", params.Err)
}

func (m *MoqErrReader_genType) ParamsKey(params MoqErrReader_genType_params, anyParams uint64) MoqErrReader_genType_paramsKey {
	m.Scene.T.Helper()
	var errUsed error
	var errUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Err == moq.ParamIndexByValue {
			errUsed = params.Err
		} else {
			errUsedHash = hash.DeepHash(params.Err)
		}
	}
	return MoqErrReader_genType_paramsKey{
		Params: struct{ Err error }{
			Err: errUsed,
		},
		Hashes: struct{ Err hash.Hash }{
			Err: errUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqErrReader_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqErrReader_genType) AssertExpectationsMet() {
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
