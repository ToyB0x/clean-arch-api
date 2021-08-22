package service

// キャッシュサービス
type MemStoreService interface {
	Get(key string) ([]byte, error)
	Add(key string, value []byte, sec int) error
}
