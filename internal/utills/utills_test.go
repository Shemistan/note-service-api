package utills

import (
	"testing"

	"github.com/Shemistan/note-service-api/internal/app/api"
	"github.com/stretchr/testify/require"
)

func TestSwapKeyAndValue(t *testing.T) {
	t.Run("input value equal nil", func(t *testing.T) {
		expectedRes := make(map[string]int64)

		res, err := SwapKeyAndValue(nil)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("len input value equal zero", func(t *testing.T) {
		req := make(map[int64]string)
		expectedRes := map[string]int64{}

		res, err := SwapKeyAndValue(req)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("success case", func(t *testing.T) {
		req := map[int64]string{1: "one", 2: "two"}
		expectedRes := map[string]int64{"one": 1, "two": 2}

		res, err := SwapKeyAndValue(req)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("input identical keys", func(t *testing.T) {
		req := map[int64]string{1: "one", 2: "two", 3: "one", 4: "two"}
		expectedRes := map[string]int64{"one": 1, "two": 2}

		res, err := SwapKeyAndValue(req)
		actual := equalMap(res, expectedRes)
		require.Nil(t, err)
		require.True(t, actual)
	})
}

func equalMap(output, expected map[string]int64) bool {
	if len(expected) != len(output) {
		return false
	}

	for key := range output {
		if _, found := expected[key]; !found {
			return false
		}
	}

	return true
}

func TestFilterSlice(t *testing.T) {
	t.Run("input value equal nil", func(t *testing.T) {
		var expectedRes []string

		res, err := FilterSlice(nil)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("len input values equal zero", func(t *testing.T) {
		req := []string{""}
		var expectedRes []string

		res, err := FilterSlice(req)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("not success case", func(t *testing.T) {
		req := []string{"r", "t", "z", "1"}
		var expectedRes []string

		res, err := FilterSlice(req)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("nothing passed the filter", func(t *testing.T) {
		req := []string{"a", "t", "d", "1"}
		expectedRes := []string{"a", "d"}

		res, err := FilterSlice(req)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("everything passed the filter", func(t *testing.T) {
		req := []string{"a", "b", "c", "d"}
		expectedRes := []string{"a", "b", "c", "d"}

		res, err := FilterSlice(req)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})
}

func TestConvertSliceToMap(t *testing.T) {
	t.Run("input value equal nil", func(t *testing.T) {
		expectedRes := make(map[int64]api.Note)

		res, err := ConvertSliceToMap(nil)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("len input values equal zero", func(t *testing.T) {
		req := make([]api.Note, 0)
		expectedRes := make(map[int64]api.Note)

		res, err := ConvertSliceToMap(req)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("success case", func(t *testing.T) {
		req := []api.Note{
			{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
			{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		}
		expectedRes := map[int64]api.Note{
			1: {Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
			2: {Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		}

		res, err := ConvertSliceToMap(req)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})
}

func TestSplitSlice(t *testing.T) {
	var (
		req = []api.Note{
			{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
			{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
			{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
			{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
		}
	)

	t.Run("input value equal nil", func(t *testing.T) {
		expectedError := "error input values"

		_, err := SplitSlice(nil, 1)
		require.Error(t, err)
		require.Equal(t, expectedError, err.Error())

	})

	t.Run("len input values equal zero ", func(t *testing.T) {
		reqZV := make([]api.Note, 0)
		expectedError := "error input values"

		_, err := SplitSlice(reqZV, 0)
		require.Error(t, err)
		require.Equal(t, expectedError, err.Error())

	})

	t.Run("batch size equal 1", func(t *testing.T) {
		expectedRes := [][]api.Note{
			{{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6}},
			{{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7}},
			{{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6}},
			{{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7}},
		}

		res, err := SplitSlice(req, 1)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("the number value is a multiple of batch size", func(t *testing.T) {
		expectedRes := [][]api.Note{
			{
				{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
				{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
			},
			{
				{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
				{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
			},
		}

		res, err := SplitSlice(req, 2)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("the number value is not a multiple of batch size", func(t *testing.T) {
		expectedRes := [][]api.Note{
			{
				{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
				{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
				{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
			},
			{
				{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
			},
		}

		res, err := SplitSlice(req, 3)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("numeric value is larger than the batch size", func(t *testing.T) {
		expectedRes := [][]api.Note{
			{
				{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
				{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
				{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
				{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
			},
		}

		res, err := SplitSlice(req, 6)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("batch size less or equal zero", func(t *testing.T) {
		expectedError := "error input values"

		_, err := SplitSlice(req, 0)
		require.NotNil(t, err)
		require.Equal(t, expectedError, err.Error())

	})
}
