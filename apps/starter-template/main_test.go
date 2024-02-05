package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MainTestSuite struct {
	suite.Suite
}

func (suite *MainTestSuite) TestHelloWorld() {
	expected := "Hello World"
	result := HelloWorld()
	assert.Equal(suite.T(), expected, result, "HelloWorld() returned unexpected result")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}
