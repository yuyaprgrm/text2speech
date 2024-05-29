package worker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuyaprgrm/text2speech/pkg/worker"
)

func TestPoolAddWorker(t *testing.T) {
	var (
		worker01 = 0
		worker02 = 1
		worker03 = 2
	)
	pool := worker.NewPool[int]()
	pool.AddWorker(&worker01)
	pool.AddWorker(&worker02)
	pool.AddWorker(&worker03)
	assert.Equal(t, 3, pool.Size())
}

func TestPoolCheckout(t *testing.T) {
	var (
		worker01 = 0
	)
	pool := worker.NewPool[int]()
	pool.AddWorker(&worker01)
	w, ok := pool.Checkout()
	assert.True(t, ok)
	assert.Same(t, w.Upgrade(), &worker01)
}

func TestWorkerReturn(t *testing.T) {
	var (
		worker01 = 0
	)
	pool := worker.NewPool[int]()
	pool.AddWorker(&worker01)
	w, _ := pool.Checkout()
	w.Return()
	w, ok := pool.Checkout()
	assert.True(t, ok)
	assert.Same(t, w.Upgrade(), &worker01)

}

func TestWorkerDoubleReturn(t *testing.T) {
	var (
		worker01 = 0
	)
	pool := worker.NewPool[int]()
	pool.AddWorker(&worker01)
	w, _ := pool.Checkout()
	w.Return()
	assert.Panics(t, func() {
		w.Return()
	})
}

func TestWorkerUpgradeAfterRetur(t *testing.T) {
	var (
		worker01 = 0
	)
	pool := worker.NewPool[int]()
	pool.AddWorker(&worker01)
	w, _ := pool.Checkout()
	w.Return()
	assert.Panics(t, func() {
		w.Upgrade()
	})
}
