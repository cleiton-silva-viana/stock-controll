package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Entity_GenerateUUIDv7(t *testing.T) {
	// Act
	var uuid = GenerateUUID()

	// Assert
	assert.Len(t, uuid, 36)
}

func Test_Entity_ConvertToUUIDv7_NoError(t *testing.T) {
	// Arrange
	uuid := "01924a4e-98f5-7abe-840c-04ee74dde766"

	// Act

	var flag = IsValidUUUID(uuid)

	// Assert
	assert.True(t, flag)
}

type testField struct {
	testDescription string
	uuid            string
}

func Test_Entity_ConvertToUUIDv7_WithError(t *testing.T) {
	// Arrange
	testsCases := []testField{
		{
			testDescription: "uuid is empyt",
			uuid:            "",
		},
		{
			testDescription: "uuid is filled empty characters",
			uuid:            "                                    ",
		},
		{
			testDescription: "uuid has invalid format",
			uuid:            "019252a1kd83a-7700-be9e-e7d3c8c89b68",
		},
	}

	for _, tt := range testsCases {

		// Act
		var flag = IsValidUUUID(tt.uuid)


		// Assert
		assert.False(t, flag)
	}
}
