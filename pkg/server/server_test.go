package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSolveRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	validPuzzleNumbers := []int{
		0, 0, 0, 5, 0, 0, 0, 0, 0,
		1, 0, 0, 8, 0, 2, 0, 9, 0,
		0, 0, 9, 0, 0, 0, 4, 0, 0,
		8, 0, 0, 9, 0, 1, 0, 4, 0,
		0, 0, 0, 0, 7, 0, 0, 0, 0,
		0, 6, 0, 0, 0, 0, 0, 0, 3,
		7, 0, 0, 0, 4, 0, 0, 0, 0,
		0, 8, 0, 2, 0, 7, 6, 0, 0,
		0, 0, 0, 0, 5, 0, 0, 2, 0,
	}

	solvedPuzzleNumbers := []int{
		6, 3, 8, 5, 9, 4, 2, 1, 7,
		1, 4, 7, 8, 6, 2, 3, 9, 5,
		2, 5, 9, 7, 1, 3, 4, 6, 8,
		8, 7, 3, 9, 2, 1, 5, 4, 6,
		5, 1, 4, 3, 7, 6, 9, 8, 2,
		9, 6, 2, 4, 8, 5, 1, 7, 3,
		7, 2, 5, 6, 4, 9, 8, 3, 1,
		4, 8, 1, 2, 3, 7, 6, 5, 9,
		3, 9, 6, 1, 5, 8, 7, 2, 4,
	}

	getSolveResponseForBody := func(t *testing.T, payload string) (*http.Response, vmSolvePuzzleResponse) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)

		req, err := http.NewRequest(http.MethodPost, "/solve", strings.NewReader(payload))
		if err != nil {
			t.Error(err)
		}
		req.Header.Add("Content-Type", "application/json")

		c.Request = req

		handleSolve(c)

		response := rr.Result()

		resVM := vmSolvePuzzleResponse{}
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&resVM); err != nil {
			t.Error(err)
		}

		return rr.Result(), resVM
	}

	t.Run("Malformed Request", func(t *testing.T) {
		bodyBytes, err := json.Marshal(gin.H{"numbers": "malformed"})
		if err != nil {
			t.Error(err)
		}

		response, resVM := getSolveResponseForBody(t, string(bodyBytes))

		if response.StatusCode != http.StatusBadRequest {
			t.Error("unnexpected response status", response.Status)
		}

		if resVM.Error != Unparsable {
			t.Error("unnexpected response error code", resVM.Error)
		}

		if resVM.Message != Unparsable.String() {
			t.Error("unnexpected response message", resVM.Message)
		}
	})

	t.Run("Puzzle Too Short", func(t *testing.T) {
		tooShort := validPuzzleNumbers[:80]

		bodyBytes, err := json.Marshal(vmSolvePuzzle{Numbers: tooShort})
		if err != nil {
			t.Error(err)
		}

		response, resVM := getSolveResponseForBody(t, string(bodyBytes))

		if response.StatusCode != http.StatusBadRequest {
			t.Error("unnexpected response status", response.Status)
		}
		if resVM.Error != Unparsable {
			t.Error("unnexpected response error code")
		}
		if resVM.Message != Unparsable.String() {
			t.Error("unnexpected response message")
		}
	})

	t.Run("Puzzle Too Long", func(t *testing.T) {
		tooLong := make([]int, 82)
		copy(tooLong, validPuzzleNumbers)

		bodyBytes, err := json.Marshal(vmSolvePuzzle{Numbers: tooLong})
		if err != nil {
			t.Error(err)
		}

		response, resVM := getSolveResponseForBody(t, string(bodyBytes))

		if response.StatusCode != http.StatusBadRequest {
			t.Error("unnexpected response status", response.Status)
		}
		if resVM.Error != Unparsable {
			t.Error("unnexpected response error code")
		}
		if resVM.Message != Unparsable.String() {
			t.Error("unnexpected response message")
		}
	})

	t.Run("Number Out of Range", func(t *testing.T) {
		oorPuzzle := make([]int, 81)
		copy(oorPuzzle, validPuzzleNumbers)

		oorPuzzle[40] = -1

		bodyBytes, err := json.Marshal(vmSolvePuzzle{Numbers: oorPuzzle})
		if err != nil {
			t.Error(err)
		}

		response, resVM := getSolveResponseForBody(t, string(bodyBytes))

		if response.StatusCode != http.StatusBadRequest {
			t.Error("unnexpected response status", response.Status)
		}
		if resVM.Error != Unparsable {
			t.Error("unnexpected response error code")
		}
		if resVM.Message != Unparsable.String() {
			t.Error("unnexpected response message")
		}

		oorPuzzle[40] = 10

		bodyBytes, err = json.Marshal(vmSolvePuzzle{Numbers: oorPuzzle})
		if err != nil {
			t.Error(err)
		}

		response, resVM = getSolveResponseForBody(t, string(bodyBytes))

		if response.StatusCode != http.StatusBadRequest {
			t.Error("unnexpected response status", response.Status)
		}
		if resVM.Error != Unparsable {
			t.Error("unnexpected response error code")
		}
		if resVM.Message != Unparsable.String() {
			t.Error("unnexpected response message")
		}
	})

	t.Run("Solve Puzzle", func(t *testing.T) {
		bodyBytes, err := json.Marshal(vmSolvePuzzle{Numbers: validPuzzleNumbers})
		if err != nil {
			t.Error(err)
		}

		response, resVM := getSolveResponseForBody(t, string(bodyBytes))

		if response.StatusCode != http.StatusOK {
			t.Error("unnexpected response status", response.Status)
		}
		if resVM.Error != NoError {
			t.Error("unnexpected response error code")
		}
		if resVM.Message != NoError.String() {
			t.Error("unnexpected response message")
		}

		if len(resVM.Numbers) != 81 {
			t.Error("unnexpected response number count", len(resVM.Numbers))
		}
		for i, v := range resVM.Numbers {
			if solvedPuzzleNumbers[i] != v {
				t.Error("unexpected value in solved puzzle", i)
			}
		}
	})

	t.Run("Unsolvable Puzzle", func(t *testing.T) {
		unsolvableNumbers := make([]int, 81)
		copy(unsolvableNumbers, validPuzzleNumbers)

		unsolvableNumbers[0] = 4

		bodyBytes, err := json.Marshal(vmSolvePuzzle{Numbers: unsolvableNumbers})
		if err != nil {
			t.Error(err)
		}

		response, resVM := getSolveResponseForBody(t, string(bodyBytes))

		if response.StatusCode != http.StatusConflict {
			t.Error("unnexpected response status", response.Status)
		}
		if resVM.Error != Unsolvable {
			t.Error("unnexpected response error code")
		}
		if resVM.Message != Unsolvable.String() {
			t.Error("unnexpected response message")
		}
	})
}
