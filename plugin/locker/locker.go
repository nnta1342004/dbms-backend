package locker

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	goservice "github.com/leductoan3082004/go-sdk"
	logger "github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	"time"
)

type Locker interface {
	Lock(key string, expTime time.Duration) error
	Unlock(key string) error
	RLock(key string) error
}

type locker struct {
	prefix string
	logger logger.Logger

	rdbClient *redis.Client
	locker    *redsync.Redsync
}

func NewLocker(prefix string, gc goservice.ServiceContext) *locker {
	return &locker{
		prefix:    prefix,
		rdbClient: gc.MustGet(appCommon.PluginRedis).(*redis.Client),
	}
}

func (l *locker) GetPrefix() string {
	return l.prefix
}

func (l *locker) Get() interface{} {
	return l.locker
}

func (l *locker) Name() string {
	return "lockera"
}

func (l *locker) InitFlags() {
	prefix := l.prefix
	if l.prefix != "" {
		prefix += "-"
	}
}

func (l *locker) Configure() error {
	l.logger = logger.GetCurrent().GetLogger(l.prefix)
	l.logger.Info("Initialized locker plugin")

	client := l.rdbClient
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)
	l.locker = rs
	return nil
}

func (l *locker) Run() error {
	return l.Configure()
}

func (locker) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}

func (l *locker) Lock(key string, expTime time.Duration) error {
	return nil
}

func (l *locker) Unlock(key string) error {
	return nil
}

func (l *locker) RLock(key string) error {
	return nil
}
