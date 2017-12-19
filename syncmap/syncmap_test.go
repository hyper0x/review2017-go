package syncmap

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"math/rand"
	"testing"
)

const (
	// number 代表用于测试的键值对的一般数量。
	number = 1000
)

// -- Store new key -- //

func benchmarkSyncMapNewKey(
	b *testing.B,
	syncMap SyncMap) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			syncMap.Store(rand.Int63n(number), randString())
		}
	})
}

func BenchmarkStoreNewKey(b *testing.B) {
	b.Run("type=Official", func(b *testing.B) {
		benchmarkSyncMapNewKey(b, NewOfficialSyncMap())
	})
	b.Run("type=Simple", func(b *testing.B) {
		benchmarkSyncMapNewKey(b, NewSimpleSyncMap())
	})
}

// -- Store existing key -- //

func benchmarkStoreExistingKey(
	b *testing.B,
	data map[interface{}]interface{},
	syncMap SyncMap) {
	for key, value := range data {
		syncMap.Store(key, value)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			syncMap.Store(rand.Int63n(number), randString())
		}
	})
}

func BenchmarkStoreExistingKey(b *testing.B) {
	b.Run("type=Official", func(b *testing.B) {
		data := genTestingData(number)
		syncMap := NewOfficialSyncMap()
		benchmarkStoreExistingKey(b, data, syncMap)
	})
	b.Run("type=Simple", func(b *testing.B) {
		data := genTestingData(number)
		syncMap := NewSimpleSyncMap()
		benchmarkStoreExistingKey(b, data, syncMap)
	})
}

// -- LoadOrStore -- //

func benchmarkLoadOrStore(
	b *testing.B,
	data map[interface{}]interface{},
	syncMap SyncMap) {
	count := number / 2
	for key, value := range data {
		if count == 0 {
			break
		}
		syncMap.Store(key, value)
		count--
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			syncMap.Store(rand.Int63n(number), randString())
		}
	})
}

func BenchmarkLoadOrStore(b *testing.B) {
	b.Run("type=Official", func(b *testing.B) {
		data := genTestingData(number)
		syncMap := NewOfficialSyncMap()
		benchmarkLoadOrStore(b, data, syncMap)
	})
	b.Run("type=Simple", func(b *testing.B) {
		data := genTestingData(number)
		syncMap := NewSimpleSyncMap()
		benchmarkLoadOrStore(b, data, syncMap)
	})
}

// -- Load -- //

func benchmarkLoad(
	b *testing.B,
	data map[interface{}]interface{},
	syncMap SyncMap) {
	for key, value := range data {
		syncMap.Store(key, value)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			syncMap.Load(rand.Int63n(number))
		}
	})
}

func BenchmarkLoad(b *testing.B) {
	b.Run("type=Official", func(b *testing.B) {
		data := genTestingData(number)
		syncMap := NewOfficialSyncMap()
		benchmarkLoad(b, data, syncMap)
	})
	b.Run("type=Simple", func(b *testing.B) {
		data := genTestingData(number)
		syncMap := NewSimpleSyncMap()
		benchmarkLoad(b, data, syncMap)
	})
}

// -- Delete -- //

func benchmarkDelete(
	b *testing.B,
	data map[interface{}]interface{},
	syncMap SyncMap) {
	for key, value := range data {
		syncMap.Store(key, value)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			syncMap.Delete(rand.Int63n(number))
		}
	})
}

func BenchmarkDelete(b *testing.B) {
	b.Run("type=Official", func(b *testing.B) {
		data := genTestingData(number)
		syncMap := NewOfficialSyncMap()
		benchmarkDelete(b, data, syncMap)
	})
	b.Run("type=Simple", func(b *testing.B) {
		data := genTestingData(number)
		syncMap := NewSimpleSyncMap()
		benchmarkDelete(b, data, syncMap)
	})
}

// -- Helper -- //

// genTestingData 用于生成测试用数据。
func genTestingData(number int) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	if number <= 0 {
		return m
	}
	for i := 0; i < number; i++ {
		m[i] = randString()
	}
	return m
}

// randString 会生成并返回一个伪随机字符串。
func randString() string {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, rand.Int31())
	return hex.EncodeToString(buf.Bytes())
}
