package server

type responseError int64

const NoError responseError = 0
const Unparsable responseError = 1
const Unsolvable responseError = 2

func (e responseError) String() string {
	switch e {
	case NoError:
		return "okay"
	case Unparsable:
		return "failed to parse puzzle"
	case Unsolvable:
		return "failed to solve puzzle"
	default:
		panic("response error not supported")
	}
}

type vmSolvePuzzle struct {
	Numbers []int `json:"numbers"`
}

type vmSolvePuzzleResponse struct {
	Numbers []int         `json:"numbers"`
	Error   responseError `json:"error"`
	Message string        `json:"message"`
}

func newVMSolvePuzzleResponse(errCode responseError, numbers []int) vmSolvePuzzleResponse {
	return vmSolvePuzzleResponse{
		Numbers: numbers,
		Error:   errCode,
		Message: errCode.String(),
	}
}
