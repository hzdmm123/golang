package main

import (
	"os"
	"time"
	"errors"
	"os/signal"
)

type Runner struct {
	//interrupt发送的信号
	interrupt chan  os.Signal
	//complete报告通道处理任务完成
	complete chan error
	//报告处理任务已经超时
	timeout <-chan time.Time
	//持有一组以索引顺序依次执行的函数
	tasks []func(int)

}

var ErrTimeOut = errors.New("receive timeout")

var ErrInterrupt = errors.New("received interrupted")
//返回一个新的准备使用的runner
func New(d time.Duration)  *Runner{
	return &Runner{
		interrupt:make(chan  os.Signal,1),
		complete:make(chan error),
		timeout: time.After(d),

	}
}
//将一个任务附加到Runner上
func (r * Runner) Add(tasks ...func(int))  {
	r.tasks = append(r.tasks,tasks...)
}

func (r *Runner) start() error  {
	//接受所有的中断信号
	signal.Notify(r.interrupt,os.Interrupt)

	//用不同的goroutine执行不同的任务
	go func() {
		r.complete <-r.run()

	}()

	select {
		//当任务完成的时候发出信号
	case err:= <-r.complete:
			return err

	case <- r.timeout:
		return ErrTimeOut

	}
}

func (r *Runner) run() error  {
	for id, task := range r.tasks{
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		task(id)
	}

	return nil
}
func (r *Runner) gotInterrupt() bool{
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true

	default:
		return true
	}
}
