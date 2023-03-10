// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package http

import (
	"fmt"
	"math/bits"
	"net/http"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// FileServer_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type FileServer_genType func(root http.FileSystem) http.Handler

// MoqFileServer_genType holds the state of a moq of the FileServer_genType
// type
type MoqFileServer_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFileServer_genType_mock

	ResultsByParams []MoqFileServer_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Root moq.ParamIndexing
		}
	}
}

// MoqFileServer_genType_mock isolates the mock interface of the
// FileServer_genType type
type MoqFileServer_genType_mock struct {
	Moq *MoqFileServer_genType
}

// MoqFileServer_genType_params holds the params of the FileServer_genType type
type MoqFileServer_genType_params struct{ Root http.FileSystem }

// MoqFileServer_genType_paramsKey holds the map key params of the
// FileServer_genType type
type MoqFileServer_genType_paramsKey struct {
	Params struct{ Root http.FileSystem }
	Hashes struct{ Root hash.Hash }
}

// MoqFileServer_genType_resultsByParams contains the results for a given set
// of parameters for the FileServer_genType type
type MoqFileServer_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFileServer_genType_paramsKey]*MoqFileServer_genType_results
}

// MoqFileServer_genType_doFn defines the type of function needed when calling
// AndDo for the FileServer_genType type
type MoqFileServer_genType_doFn func(root http.FileSystem)

// MoqFileServer_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the FileServer_genType type
type MoqFileServer_genType_doReturnFn func(root http.FileSystem) http.Handler

// MoqFileServer_genType_results holds the results of the FileServer_genType
// type
type MoqFileServer_genType_results struct {
	Params  MoqFileServer_genType_params
	Results []struct {
		Values *struct {
			Result1 http.Handler
		}
		Sequence   uint32
		DoFn       MoqFileServer_genType_doFn
		DoReturnFn MoqFileServer_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFileServer_genType_fnRecorder routes recorded function calls to the
// MoqFileServer_genType moq
type MoqFileServer_genType_fnRecorder struct {
	Params    MoqFileServer_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFileServer_genType_results
	Moq       *MoqFileServer_genType
}

// MoqFileServer_genType_anyParams isolates the any params functions of the
// FileServer_genType type
type MoqFileServer_genType_anyParams struct {
	Recorder *MoqFileServer_genType_fnRecorder
}

// NewMoqFileServer_genType creates a new moq of the FileServer_genType type
func NewMoqFileServer_genType(scene *moq.Scene, config *moq.Config) *MoqFileServer_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqFileServer_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqFileServer_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Root moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Root moq.ParamIndexing
		}{
			Root: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the FileServer_genType type
func (m *MoqFileServer_genType) Mock() FileServer_genType {
	return func(root http.FileSystem) http.Handler {
		m.Scene.T.Helper()
		moq := &MoqFileServer_genType_mock{Moq: m}
		return moq.Fn(root)
	}
}

func (m *MoqFileServer_genType_mock) Fn(root http.FileSystem) (result1 http.Handler) {
	m.Moq.Scene.T.Helper()
	params := MoqFileServer_genType_params{
		Root: root,
	}
	var results *MoqFileServer_genType_results
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
		result.DoFn(root)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(root)
	}
	return
}

func (m *MoqFileServer_genType) OnCall(root http.FileSystem) *MoqFileServer_genType_fnRecorder {
	return &MoqFileServer_genType_fnRecorder{
		Params: MoqFileServer_genType_params{
			Root: root,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqFileServer_genType_fnRecorder) Any() *MoqFileServer_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqFileServer_genType_anyParams{Recorder: r}
}

func (a *MoqFileServer_genType_anyParams) Root() *MoqFileServer_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqFileServer_genType_fnRecorder) Seq() *MoqFileServer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFileServer_genType_fnRecorder) NoSeq() *MoqFileServer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFileServer_genType_fnRecorder) ReturnResults(result1 http.Handler) *MoqFileServer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 http.Handler
		}
		Sequence   uint32
		DoFn       MoqFileServer_genType_doFn
		DoReturnFn MoqFileServer_genType_doReturnFn
	}{
		Values: &struct {
			Result1 http.Handler
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFileServer_genType_fnRecorder) AndDo(fn MoqFileServer_genType_doFn) *MoqFileServer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFileServer_genType_fnRecorder) DoReturnResults(fn MoqFileServer_genType_doReturnFn) *MoqFileServer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 http.Handler
		}
		Sequence   uint32
		DoFn       MoqFileServer_genType_doFn
		DoReturnFn MoqFileServer_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFileServer_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFileServer_genType_resultsByParams
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
		results = &MoqFileServer_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFileServer_genType_paramsKey]*MoqFileServer_genType_results{},
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
		r.Results = &MoqFileServer_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFileServer_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFileServer_genType_fnRecorder {
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
					Result1 http.Handler
				}
				Sequence   uint32
				DoFn       MoqFileServer_genType_doFn
				DoReturnFn MoqFileServer_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFileServer_genType) PrettyParams(params MoqFileServer_genType_params) string {
	return fmt.Sprintf("FileServer_genType(%#v)", params.Root)
}

func (m *MoqFileServer_genType) ParamsKey(params MoqFileServer_genType_params, anyParams uint64) MoqFileServer_genType_paramsKey {
	m.Scene.T.Helper()
	var rootUsed http.FileSystem
	var rootUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Root == moq.ParamIndexByValue {
			rootUsed = params.Root
		} else {
			rootUsedHash = hash.DeepHash(params.Root)
		}
	}
	return MoqFileServer_genType_paramsKey{
		Params: struct{ Root http.FileSystem }{
			Root: rootUsed,
		},
		Hashes: struct{ Root hash.Hash }{
			Root: rootUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqFileServer_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFileServer_genType) AssertExpectationsMet() {
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
