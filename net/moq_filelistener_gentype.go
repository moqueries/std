// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package net

import (
	"fmt"
	"math/bits"
	"net"
	"os"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// FileListener_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type FileListener_genType func(f *os.File) (ln net.Listener, err error)

// MoqFileListener_genType holds the state of a moq of the FileListener_genType
// type
type MoqFileListener_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFileListener_genType_mock

	ResultsByParams []MoqFileListener_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			F moq.ParamIndexing
		}
	}
}

// MoqFileListener_genType_mock isolates the mock interface of the
// FileListener_genType type
type MoqFileListener_genType_mock struct {
	Moq *MoqFileListener_genType
}

// MoqFileListener_genType_params holds the params of the FileListener_genType
// type
type MoqFileListener_genType_params struct{ F *os.File }

// MoqFileListener_genType_paramsKey holds the map key params of the
// FileListener_genType type
type MoqFileListener_genType_paramsKey struct {
	Params struct{ F *os.File }
	Hashes struct{ F hash.Hash }
}

// MoqFileListener_genType_resultsByParams contains the results for a given set
// of parameters for the FileListener_genType type
type MoqFileListener_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFileListener_genType_paramsKey]*MoqFileListener_genType_results
}

// MoqFileListener_genType_doFn defines the type of function needed when
// calling AndDo for the FileListener_genType type
type MoqFileListener_genType_doFn func(f *os.File)

// MoqFileListener_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the FileListener_genType type
type MoqFileListener_genType_doReturnFn func(f *os.File) (ln net.Listener, err error)

// MoqFileListener_genType_results holds the results of the
// FileListener_genType type
type MoqFileListener_genType_results struct {
	Params  MoqFileListener_genType_params
	Results []struct {
		Values *struct {
			Ln  net.Listener
			Err error
		}
		Sequence   uint32
		DoFn       MoqFileListener_genType_doFn
		DoReturnFn MoqFileListener_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFileListener_genType_fnRecorder routes recorded function calls to the
// MoqFileListener_genType moq
type MoqFileListener_genType_fnRecorder struct {
	Params    MoqFileListener_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFileListener_genType_results
	Moq       *MoqFileListener_genType
}

// MoqFileListener_genType_anyParams isolates the any params functions of the
// FileListener_genType type
type MoqFileListener_genType_anyParams struct {
	Recorder *MoqFileListener_genType_fnRecorder
}

// NewMoqFileListener_genType creates a new moq of the FileListener_genType
// type
func NewMoqFileListener_genType(scene *moq.Scene, config *moq.Config) *MoqFileListener_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqFileListener_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqFileListener_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				F moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			F moq.ParamIndexing
		}{
			F: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the FileListener_genType type
func (m *MoqFileListener_genType) Mock() FileListener_genType {
	return func(f *os.File) (_ net.Listener, _ error) {
		m.Scene.T.Helper()
		moq := &MoqFileListener_genType_mock{Moq: m}
		return moq.Fn(f)
	}
}

func (m *MoqFileListener_genType_mock) Fn(f *os.File) (ln net.Listener, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqFileListener_genType_params{
		F: f,
	}
	var results *MoqFileListener_genType_results
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
		result.DoFn(f)
	}

	if result.Values != nil {
		ln = result.Values.Ln
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		ln, err = result.DoReturnFn(f)
	}
	return
}

func (m *MoqFileListener_genType) OnCall(f *os.File) *MoqFileListener_genType_fnRecorder {
	return &MoqFileListener_genType_fnRecorder{
		Params: MoqFileListener_genType_params{
			F: f,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqFileListener_genType_fnRecorder) Any() *MoqFileListener_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqFileListener_genType_anyParams{Recorder: r}
}

func (a *MoqFileListener_genType_anyParams) F() *MoqFileListener_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqFileListener_genType_fnRecorder) Seq() *MoqFileListener_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFileListener_genType_fnRecorder) NoSeq() *MoqFileListener_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFileListener_genType_fnRecorder) ReturnResults(ln net.Listener, err error) *MoqFileListener_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Ln  net.Listener
			Err error
		}
		Sequence   uint32
		DoFn       MoqFileListener_genType_doFn
		DoReturnFn MoqFileListener_genType_doReturnFn
	}{
		Values: &struct {
			Ln  net.Listener
			Err error
		}{
			Ln:  ln,
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFileListener_genType_fnRecorder) AndDo(fn MoqFileListener_genType_doFn) *MoqFileListener_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFileListener_genType_fnRecorder) DoReturnResults(fn MoqFileListener_genType_doReturnFn) *MoqFileListener_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Ln  net.Listener
			Err error
		}
		Sequence   uint32
		DoFn       MoqFileListener_genType_doFn
		DoReturnFn MoqFileListener_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFileListener_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFileListener_genType_resultsByParams
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
		results = &MoqFileListener_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFileListener_genType_paramsKey]*MoqFileListener_genType_results{},
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
		r.Results = &MoqFileListener_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFileListener_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFileListener_genType_fnRecorder {
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
					Ln  net.Listener
					Err error
				}
				Sequence   uint32
				DoFn       MoqFileListener_genType_doFn
				DoReturnFn MoqFileListener_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFileListener_genType) PrettyParams(params MoqFileListener_genType_params) string {
	return fmt.Sprintf("FileListener_genType(%#v)", params.F)
}

func (m *MoqFileListener_genType) ParamsKey(params MoqFileListener_genType_params, anyParams uint64) MoqFileListener_genType_paramsKey {
	m.Scene.T.Helper()
	var fUsed *os.File
	var fUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.F == moq.ParamIndexByValue {
			fUsed = params.F
		} else {
			fUsedHash = hash.DeepHash(params.F)
		}
	}
	return MoqFileListener_genType_paramsKey{
		Params: struct{ F *os.File }{
			F: fUsed,
		},
		Hashes: struct{ F hash.Hash }{
			F: fUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqFileListener_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFileListener_genType) AssertExpectationsMet() {
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