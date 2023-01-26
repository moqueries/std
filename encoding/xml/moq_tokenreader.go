// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package xml

import (
	"encoding/xml"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that xml.TokenReader is mocked
// completely
var _ xml.TokenReader = (*MoqTokenReader_mock)(nil)

// MoqTokenReader holds the state of a moq of the TokenReader type
type MoqTokenReader struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqTokenReader_mock

	ResultsByParams_Token []MoqTokenReader_Token_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Token struct{}
		}
	}
}

// MoqTokenReader_mock isolates the mock interface of the TokenReader type
type MoqTokenReader_mock struct {
	Moq *MoqTokenReader
}

// MoqTokenReader_recorder isolates the recorder interface of the TokenReader
// type
type MoqTokenReader_recorder struct {
	Moq *MoqTokenReader
}

// MoqTokenReader_Token_params holds the params of the TokenReader type
type MoqTokenReader_Token_params struct{}

// MoqTokenReader_Token_paramsKey holds the map key params of the TokenReader
// type
type MoqTokenReader_Token_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqTokenReader_Token_resultsByParams contains the results for a given set of
// parameters for the TokenReader type
type MoqTokenReader_Token_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqTokenReader_Token_paramsKey]*MoqTokenReader_Token_results
}

// MoqTokenReader_Token_doFn defines the type of function needed when calling
// AndDo for the TokenReader type
type MoqTokenReader_Token_doFn func()

// MoqTokenReader_Token_doReturnFn defines the type of function needed when
// calling DoReturnResults for the TokenReader type
type MoqTokenReader_Token_doReturnFn func() (xml.Token, error)

// MoqTokenReader_Token_results holds the results of the TokenReader type
type MoqTokenReader_Token_results struct {
	Params  MoqTokenReader_Token_params
	Results []struct {
		Values *struct {
			Result1 xml.Token
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqTokenReader_Token_doFn
		DoReturnFn MoqTokenReader_Token_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqTokenReader_Token_fnRecorder routes recorded function calls to the
// MoqTokenReader moq
type MoqTokenReader_Token_fnRecorder struct {
	Params    MoqTokenReader_Token_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqTokenReader_Token_results
	Moq       *MoqTokenReader
}

// MoqTokenReader_Token_anyParams isolates the any params functions of the
// TokenReader type
type MoqTokenReader_Token_anyParams struct {
	Recorder *MoqTokenReader_Token_fnRecorder
}

// NewMoqTokenReader creates a new moq of the TokenReader type
func NewMoqTokenReader(scene *moq.Scene, config *moq.Config) *MoqTokenReader {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqTokenReader{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqTokenReader_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Token struct{}
			}
		}{ParameterIndexing: struct {
			Token struct{}
		}{
			Token: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the TokenReader type
func (m *MoqTokenReader) Mock() *MoqTokenReader_mock { return m.Moq }

func (m *MoqTokenReader_mock) Token() (result1 xml.Token, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqTokenReader_Token_params{}
	var results *MoqTokenReader_Token_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Token {
		paramsKey := m.Moq.ParamsKey_Token(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Token(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Token(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Token(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the TokenReader type
func (m *MoqTokenReader) OnCall() *MoqTokenReader_recorder {
	return &MoqTokenReader_recorder{
		Moq: m,
	}
}

func (m *MoqTokenReader_recorder) Token() *MoqTokenReader_Token_fnRecorder {
	return &MoqTokenReader_Token_fnRecorder{
		Params:   MoqTokenReader_Token_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqTokenReader_Token_fnRecorder) Any() *MoqTokenReader_Token_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Token(r.Params))
		return nil
	}
	return &MoqTokenReader_Token_anyParams{Recorder: r}
}

func (r *MoqTokenReader_Token_fnRecorder) Seq() *MoqTokenReader_Token_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Token(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqTokenReader_Token_fnRecorder) NoSeq() *MoqTokenReader_Token_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Token(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqTokenReader_Token_fnRecorder) ReturnResults(result1 xml.Token, result2 error) *MoqTokenReader_Token_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 xml.Token
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqTokenReader_Token_doFn
		DoReturnFn MoqTokenReader_Token_doReturnFn
	}{
		Values: &struct {
			Result1 xml.Token
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqTokenReader_Token_fnRecorder) AndDo(fn MoqTokenReader_Token_doFn) *MoqTokenReader_Token_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqTokenReader_Token_fnRecorder) DoReturnResults(fn MoqTokenReader_Token_doReturnFn) *MoqTokenReader_Token_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 xml.Token
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqTokenReader_Token_doFn
		DoReturnFn MoqTokenReader_Token_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqTokenReader_Token_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqTokenReader_Token_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Token {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqTokenReader_Token_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqTokenReader_Token_paramsKey]*MoqTokenReader_Token_results{},
		}
		r.Moq.ResultsByParams_Token = append(r.Moq.ResultsByParams_Token, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Token) {
			copy(r.Moq.ResultsByParams_Token[insertAt+1:], r.Moq.ResultsByParams_Token[insertAt:0])
			r.Moq.ResultsByParams_Token[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Token(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqTokenReader_Token_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqTokenReader_Token_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqTokenReader_Token_fnRecorder {
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
					Result1 xml.Token
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqTokenReader_Token_doFn
				DoReturnFn MoqTokenReader_Token_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqTokenReader) PrettyParams_Token(params MoqTokenReader_Token_params) string {
	return fmt.Sprintf("Token()")
}

func (m *MoqTokenReader) ParamsKey_Token(params MoqTokenReader_Token_params, anyParams uint64) MoqTokenReader_Token_paramsKey {
	m.Scene.T.Helper()
	return MoqTokenReader_Token_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqTokenReader) Reset() { m.ResultsByParams_Token = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqTokenReader) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Token {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Token(results.Params))
			}
		}
	}
}
