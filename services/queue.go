package services

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

const (
	MinConfirm = 6
	MaxEvent   = 1000
)

type Queue struct {
	queue [][]interface{}

	sender   chan interface{}
	receiver chan<- []interface{}

	ctx    context.Context
	cancel context.CancelFunc
}

func NewQueue(ch chan<- []interface{}) *Queue {
	ctx, cancel := context.WithCancel(context.Background())
	return &Queue{
		sender:   make(chan interface{}, MaxEvent),
		receiver: ch,

		ctx:    ctx,
		cancel: cancel,
	}
}

func (q *Queue) Close() {
	q.cancel()
	close(q.sender)
}

func (q *Queue) Run(client *ethclient.Client) {
	go func() {
		blk := make(chan *types.Header)

		sub, err := client.SubscribeNewHead(context.Background(), blk)
		if err != nil {
			log.Error("Queue:", "error", err.Error())
			return
		}
		defer sub.Unsubscribe()

		for {
			select {
			case <-blk:
				var evts []interface{}
				for len(q.sender) != 0 {
					evts = append(evts, <-q.sender)
				}
				q.push(evts)
			case <-q.ctx.Done():
				return
			}
		}
	}()
}

func (q *Queue) Push(evt interface{}) {
	q.sender <- evt
}

func (q *Queue) push(evts []interface{}) {
	if len(q.queue) < MinConfirm {
		q.queue = append([][]interface{}{evts}, q.queue...)
	} else {
		q.receiver <- q.queue[MinConfirm-1]
		q.queue = append([][]interface{}{evts}, q.queue[:len(q.queue)-1]...)
	}
}
