package queue

import (
	"context"
	"sync"
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/nika-gromova/o-architecture-patterns/game/errors"
	"github.com/nika-gromova/o-architecture-patterns/game/tests/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_StartThread(t *testing.T) {
	var (
		ctrl              = gomock.NewController(t)
		cmdMock           = mocks.NewMockCommand(ctrl)
		errorsHandler     = errors.NewHandler()
		ctx               = context.Background()
		cmds              = make(chan base.Command, 100)
		q                 = New(ctx, errorsHandler, cmds)
		wg                = sync.WaitGroup{} // синхронизация основного потока и потока обработки очереди
		started, finished = false, false
		testThread        = &Thread{
			Processor: q,
			OnStart: func() {
				started = true
			},
			OnStop: func() {
				defer wg.Done()
				finished = true
			},
		}
		startCmd = &StartThreadCommand{
			thread: testThread,
		}
	)
	q.Add(cmdMock) // добавление тестовой команды в очередь

	cmdMock.EXPECT().Execute().DoAndReturn(func() error {
		testThread.Stop() // остановка обработки очереди "вручную"
		return nil
	})

	wg.Add(1)
	err := startCmd.Execute()

	require.NoError(t, err)
	wg.Wait() // ожидание завершения потока обработки очереди
	require.True(t, started)
	require.True(t, finished)
}

func Test_HardStop(t *testing.T) {
	var (
		ctrl          = gomock.NewController(t)
		cmdMock       = mocks.NewMockCommand(ctrl)
		errorsHandler = errors.NewHandler()
		ctx           = context.Background()
		cmds          = make(chan base.Command, 100)
		q             = New(ctx, errorsHandler, cmds)
		wg            = sync.WaitGroup{}
		cond          = sync.NewCond(&sync.Mutex{}) // условная переменная, чтобы дождаться выполнения первой команды в очереди
		testThread    = &Thread{
			Processor: q,
			OnStart:   func() {},
			OnStop: func() {
				defer wg.Done()
			},
		}
		startCmd = &StartThreadCommand{
			thread: testThread,
		}
		hardStopCmd = &HardStopCommand{
			thread: testThread,
		}
		cmd1Processed = false
	)
	q.Add(cmdMock) // добавление тестовой команды в очередь

	cmdMock.EXPECT().Execute().DoAndReturn(func() error {
		cond.L.Lock()
		cmd1Processed = true
		cond.L.Unlock()
		cond.Signal()
		return nil
	}).Times(1)

	// ожидание выполнения первой команды, после этого запуск команды на остановку обработки очереди
	go func() {
		cond.L.Lock()
		defer cond.L.Unlock()
		for !cmd1Processed {
			cond.Wait()
		}
		_ = hardStopCmd.Execute()
	}()

	wg.Add(1)
	err := startCmd.Execute()

	require.NoError(t, err)
	wg.Wait() // ожидание завершения потока обработки очереди
}

func Test_SoftStop(t *testing.T) {
	var (
		ctrl          = gomock.NewController(t)
		cmdMock       = mocks.NewMockCommand(ctrl)
		errorsHandler = errors.NewHandler()
		ctx           = context.Background()
		cmds          = make(chan base.Command, 100)
		q             = New(ctx, errorsHandler, cmds)
		wg            = sync.WaitGroup{}
		testThread    = &Thread{
			Processor: q,
			OnStart:   func() {},
			OnStop: func() {
				defer wg.Done()
			},
		}
		startCmd = &StartThreadCommand{
			thread: testThread,
		}
		softStopCmd = &SoftStopCommand{
			thread: testThread,
		}
	)
	q.Add(cmdMock) // добавление тестовых команд в очередь
	q.Add(cmdMock)

	cmdMock.EXPECT().Execute().Return(nil).Times(2)

	wg.Add(1)
	err := startCmd.Execute()

	go func() {
		_ = softStopCmd.Execute()
	}()

	require.NoError(t, err)
	wg.Wait() // ожидание завершения потока обработки очереди
}
