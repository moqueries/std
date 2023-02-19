// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package bufio

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ScanWords_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type ScanWords_genType func(data []byte, atEOF bool) (advance int, token []byte, err error)

// MoqScanWords_genType holds the state of a moq of the ScanWords_genType type
type MoqScanWords_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqScanWords_genType_mock

	ResultsByParams []MoqScanWords_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Data  moq.ParamIndexing
			AtEOF moq.ParamIndexing
		}
	}
}

// MoqScanWords_genType_mock isolates the mock interface of the
// ScanWords_genType type
type MoqScanWords_genType_mock struct {
	Moq *MoqScanWords_genType
}

// MoqScanWords_genType_params holds the params of the ScanWords_genType type
type MoqScanWords_genType_params struct {
	Data  []byte
	AtEOF bool
}

// MoqScanWords_genType_paramsKey holds the map key params of the
// ScanWords_genType type
type MoqScanWords_genType_paramsKey struct {
	Params struct{ AtEOF bool }
	Hashes struct {
		Data  hash.Hash
		AtEOF hash.Hash
	}
}

// MoqScanWords_genType_resultsByParams contains the results for a given set of
// parameters for the ScanWords_genType type
type MoqScanWords_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqScanWords_genType_paramsKey]*MoqScanWords_genType_results
}

// MoqScanWords_genType_doFn defines the type of function needed when calling
// AndDo for the ScanWords_genType type
type MoqScanWords_genType_doFn func(data []byte, atEOF bool)

// MoqScanWords_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ScanWords_genType type
type MoqScanWords_genType_doReturnFn func(data []byte, atEOF bool) (advance int, token []byte, err error)

// MoqScanWords_genType_results holds the results of the ScanWords_genType type
type MoqScanWords_genType_results struct {
	Params  MoqScanWords_genType_params
	Results []struct {
		Values *struct {
			Advance int
			Token   []byte
			Err     error
		}
		Sequence   uint32
		DoFn       MoqScanWords_genType_doFn
		DoReturnFn MoqScanWords_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqScanWords_genType_fnRecorder routes recorded function calls to the
// MoqScanWords_genType moq
type MoqScanWords_genType_fnRecorder struct {
	Params    MoqScanWords_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqScanWords_genType_results
	Moq       *MoqScanWords_genType
}

// MoqScanWords_genType_anyParams isolates the any params functions of the
// ScanWords_genType type
type MoqScanWords_genType_anyParams struct {
	Recorder *MoqScanWords_genType_fnRecorder
}

// NewMoqScanWords_genType creates a new moq of the ScanWords_genType type
func NewMoqScanWords_genType(scene *moq.Scene, config *moq.Config) *MoqScanWords_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqScanWords_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqScanWords_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Data  moq.ParamIndexing
				AtEOF moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Data  moq.ParamIndexing
			AtEOF moq.ParamIndexing
		}{
			Data:  moq.ParamIndexByHash,
			AtEOF: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ScanWords_genType type
func (m *MoqScanWords_genType) Mock() ScanWords_genType {
	return func(data []byte, atEOF bool) (_ int, _ []byte, _ error) {
		m.Scene.T.Helper()
		moq := &MoqScanWords_genType_mock{Moq: m}
		return moq.Fn(data, atEOF)
	}
}

func (m *MoqScanWords_genType_mock) Fn(data []byte, atEOF bool) (advance int, token []byte, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqScanWords_genType_params{
		Data:  data,
		AtEOF: atEOF,
	}
	var results *MoqScanWords_genType_results
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
		result.DoFn(data, atEOF)
	}

	if result.Values != nil {
		advance = result.Values.Advance
		token = result.Values.Token
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		advance, token, err = result.DoReturnFn(data, atEOF)
	}
	return
}

func (m *MoqScanWords_genType) OnCall(data []byte, atEOF bool) *MoqScanWords_genType_fnRecorder {
	return &MoqScanWords_genType_fnRecorder{
		Params: MoqScanWords_genType_params{
			Data:  data,
			AtEOF: atEOF,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqScanWords_genType_fnRecorder) Any() *MoqScanWords_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqScanWords_genType_anyParams{Recorder: r}
}

func (a *MoqScanWords_genType_anyParams) Data() *MoqScanWords_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqScanWords_genType_anyParams) AtEOF() *MoqScanWords_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqScanWords_genType_fnRecorder) Seq() *MoqScanWords_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqScanWords_genType_fnRecorder) NoSeq() *MoqScanWords_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqScanWords_genType_fnRecorder) ReturnResults(advance int, token []byte, err error) *MoqScanWords_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Advance int
			Token   []byte
			Err     error
		}
		Sequence   uint32
		DoFn       MoqScanWords_genType_doFn
		DoReturnFn MoqScanWords_genType_doReturnFn
	}{
		Values: &struct {
			Advance int
			Token   []byte
			Err     error
		}{
			Advance: advance,
			Token:   token,
			Err:     err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqScanWords_genType_fnRecorder) AndDo(fn MoqScanWords_genType_doFn) *MoqScanWords_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqScanWords_genType_fnRecorder) DoReturnResults(fn MoqScanWords_genType_doReturnFn) *MoqScanWords_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Advance int
			Token   []byte
			Err     error
		}
		Sequence   uint32
		DoFn       MoqScanWords_genType_doFn
		DoReturnFn MoqScanWords_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqScanWords_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqScanWords_genType_resultsByParams
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
		results = &MoqScanWords_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqScanWords_genType_paramsKey]*MoqScanWords_genType_results{},
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
		r.Results = &MoqScanWords_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqScanWords_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqScanWords_genType_fnRecorder {
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
					Advance int
					Token   []byte
					Err     error
				}
				Sequence   uint32
				DoFn       MoqScanWords_genType_doFn
				DoReturnFn MoqScanWords_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqScanWords_genType) PrettyParams(params MoqScanWords_genType_params) string {
	return fmt.Sprintf("ScanWords_genType(%#v, %#v)", params.Data, params.AtEOF)
}

func (m *MoqScanWords_genType) ParamsKey(params MoqScanWords_genType_params, anyParams uint64) MoqScanWords_genType_paramsKey {
	m.Scene.T.Helper()
	var dataUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Data == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The data parameter can't be indexed by value")
		}
		dataUsedHash = hash.DeepHash(params.Data)
	}
	var atEOFUsed bool
	var atEOFUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.AtEOF == moq.ParamIndexByValue {
			atEOFUsed = params.AtEOF
		} else {
			atEOFUsedHash = hash.DeepHash(params.AtEOF)
		}
	}
	return MoqScanWords_genType_paramsKey{
		Params: struct{ AtEOF bool }{
			AtEOF: atEOFUsed,
		},
		Hashes: struct {
			Data  hash.Hash
			AtEOF hash.Hash
		}{
			Data:  dataUsedHash,
			AtEOF: atEOFUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqScanWords_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqScanWords_genType) AssertExpectationsMet() {
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
