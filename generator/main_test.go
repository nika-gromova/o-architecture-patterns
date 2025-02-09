package main

import (
	"os"
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/stretchr/testify/require"
)

type TestInterface interface {
	GetVelocityVector() (base.Vector, error)
	SetVelocityVector(vector base.Vector) error
	Test(args []any) error
}

func TestGenerateAdapter(t *testing.T) {
	var (
		file          = "main_test.go"
		testInterface = "TestInterface"
		outputFile    = "result.go"
		expected      = []byte(`package main


import (
	"context"
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/nika-gromova/o-architecture-patterns/game/ioc"
)


type TestInterfaceAdapter struct {
	obj any
	ctx context.Context
}

func (a *TestInterfaceAdapter) SetObject (obj any) {
    a.obj = obj
}
func (a *TestInterfaceAdapter) SetContext (ctx context.Context) {
    a.ctx = ctx
}


func (a *TestInterfaceAdapter) GetVelocityVector() (base.Vector, error) {
    result, err := ioc.Resolve(a.ctx, "", a.obj, )
    if err != nil {
        var resultTyped base.Vector
        return resultTyped, err
    }
	resultTyped, ok := result.(base.Vector)
	if !ok {
		return resultTyped, fmt.Errorf("type conversion error")
	}
	return resultTyped, nil
}

func (a *TestInterfaceAdapter) SetVelocityVector(vector base.Vector) (error) {
    result, err := ioc.Resolve(a.ctx, "TestInterfaceAdapter:setvelocityvector", a.obj, vector)
    if err != nil {
        return err
    }
	resultCommand, ok := result.(base.Command)
	if !ok {
		return fmt.Errorf("failed to convert command")
	}
	return resultCommand.Execute()
}

func (a *TestInterfaceAdapter) Test(args []any) (error) {
    result, err := ioc.Resolve(a.ctx, "TestInterfaceAdapter:test", a.obj, args)
    if err != nil {
        return err
    }
	resultCommand, ok := result.(base.Command)
	if !ok {
		return fmt.Errorf("failed to convert command")
	}
	return resultCommand.Execute()
}
`)
	)

	process(file, testInterface, outputFile)

	actual, err := os.ReadFile(outputFile)
	require.NoError(t, err)

	require.Equal(t, expected, actual)

	err = os.Remove(outputFile)
	require.NoError(t, err)
}
