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

// FileConn_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type FileConn_genType func(f *os.File) (c net.Conn, err error)

// MoqFileConn_genType holds the state of a moq of the FileConn_genType type
type MoqFileConn_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFileConn_genType_mock

	ResultsByParams []MoqFileConn_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			F moq.ParamIndexing
		}
	}
}

// MoqFileConn_genType_mock isolates the mock interface of the FileConn_genType
// type
type MoqFileConn_genType_mock struct {
	Moq *MoqFileConn_genType
}

// MoqFileConn_genType_params holds the params of the FileConn_genType type
type MoqFileConn_genType_params struct{ F *os.File }

// MoqFileConn_genType_paramsKey holds the map key params of the
// FileConn_genType type
type MoqFileConn_genType_paramsKey struct {
	Params struct{ F *os.File }
	Hashes struct{ F hash.Hash }
}

// MoqFileConn_genType_resultsByParams contains the results for a given set of
// parameters for the FileConn_genType type
type MoqFileConn_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFileConn_genType_paramsKey]*MoqFileConn_genType_results
}

// MoqFileConn_genType_doFn defines the type of function needed when calling
// AndDo for the FileConn_genType type
type MoqFileConn_genType_doFn func(f *os.File)

// MoqFileConn_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the FileConn_genType type
type MoqFileConn_genType_doReturnFn func(f *os.File) (c net.Conn, err error)

// MoqFileConn_genType_results holds the results of the FileConn_genType type
type MoqFileConn_genType_results struct {
	Params  MoqFileConn_genType_params
	Results []struct {
		Values *struct {
			C   net.Conn
			Err error
		}
		Sequence   uint32
		DoFn       MoqFileConn_genType_doFn
		DoReturnFn MoqFileConn_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFileConn_genType_fnRecorder routes recorded function calls to the
// MoqFileConn_genType moq
type MoqFileConn_genType_fnRecorder struct {
	Params    MoqFileConn_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFileConn_genType_results
	Moq       *MoqFileConn_genType
}

// MoqFileConn_genType_anyParams isolates the any params functions of the
// FileConn_genType type
type MoqFileConn_genType_anyParams struct {
	Recorder *MoqFileConn_genType_fnRecorder
}

// NewMoqFileConn_genType creates a new moq of the FileConn_genType type
func NewMoqFileConn_genType(scene *moq.Scene, config *moq.Config) *MoqFileConn_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqFileConn_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqFileConn_genType_mock{},

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

// Mock returns the moq implementation of the FileConn_genType type
func (m *MoqFileConn_genType) Mock() FileConn_genType {
	return func(f *os.File) (_ net.Conn, _ error) {
		m.Scene.T.Helper()
		moq := &MoqFileConn_genType_mock{Moq: m}
		return moq.Fn(f)
	}
}

func (m *MoqFileConn_genType_mock) Fn(f *os.File) (c net.Conn, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqFileConn_genType_params{
		F: f,
	}
	var results *MoqFileConn_genType_results
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
		c = result.Values.C
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		c, err = result.DoReturnFn(f)
	}
	return
}

func (m *MoqFileConn_genType) OnCall(f *os.File) *MoqFileConn_genType_fnRecorder {
	return &MoqFileConn_genType_fnRecorder{
		Params: MoqFileConn_genType_params{
			F: f,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqFileConn_genType_fnRecorder) Any() *MoqFileConn_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqFileConn_genType_anyParams{Recorder: r}
}

func (a *MoqFileConn_genType_anyParams) F() *MoqFileConn_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqFileConn_genType_fnRecorder) Seq() *MoqFileConn_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFileConn_genType_fnRecorder) NoSeq() *MoqFileConn_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFileConn_genType_fnRecorder) ReturnResults(c net.Conn, err error) *MoqFileConn_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			C   net.Conn
			Err error
		}
		Sequence   uint32
		DoFn       MoqFileConn_genType_doFn
		DoReturnFn MoqFileConn_genType_doReturnFn
	}{
		Values: &struct {
			C   net.Conn
			Err error
		}{
			C:   c,
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFileConn_genType_fnRecorder) AndDo(fn MoqFileConn_genType_doFn) *MoqFileConn_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFileConn_genType_fnRecorder) DoReturnResults(fn MoqFileConn_genType_doReturnFn) *MoqFileConn_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			C   net.Conn
			Err error
		}
		Sequence   uint32
		DoFn       MoqFileConn_genType_doFn
		DoReturnFn MoqFileConn_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFileConn_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFileConn_genType_resultsByParams
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
		results = &MoqFileConn_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFileConn_genType_paramsKey]*MoqFileConn_genType_results{},
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
		r.Results = &MoqFileConn_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFileConn_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFileConn_genType_fnRecorder {
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
					C   net.Conn
					Err error
				}
				Sequence   uint32
				DoFn       MoqFileConn_genType_doFn
				DoReturnFn MoqFileConn_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFileConn_genType) PrettyParams(params MoqFileConn_genType_params) string {
	return fmt.Sprintf("FileConn_genType(%#v)", params.F)
}

func (m *MoqFileConn_genType) ParamsKey(params MoqFileConn_genType_params, anyParams uint64) MoqFileConn_genType_paramsKey {
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
	return MoqFileConn_genType_paramsKey{
		Params: struct{ F *os.File }{
			F: fUsed,
		},
		Hashes: struct{ F hash.Hash }{
			F: fUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqFileConn_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFileConn_genType) AssertExpectationsMet() {
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
