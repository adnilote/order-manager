package redis

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/adnilote/order-manager/app/business/entities"
	pb "github.com/adnilote/order-manager/app/business/entities/proto"
	"github.com/adnilote/order-manager/app/config"
	"github.com/stretchr/testify/require"
)

var store *Client

func init() {
	configFileName := os.Getenv("TEST_CONFIG")
	if configFileName == "" {
		configFileName = "config-test.yaml"
	}

	err := config.LoadConfig(configFileName)
	if err != nil {
		log.Fatal("Load config failed: ", err)
	}

	store, err = NewClient(context.TODO(), config.Config.Redis)
	if err != nil {
		panic(err)
	}
}

func TestTotalPositions(t *testing.T) {

	position := *pb.NewPopulatedTotalPosition(randy{}, true)
	err := store.SaveTotalPosition(position)
	require.Nil(t, err)

	position2, err := store.GetTotalPosition(entities.GetTotalPositionKey(position))
	require.Nil(t, err)
	require.True(t, position.Equal(position2))

	err = store.DeleteTotalPosition(position)
	require.Nil(t, err)

	pos, err := store.GetTotalPosition(entities.GetTotalPositionKey(position))
	require.True(t, pos.Equal(pb.TotalPosition{}))
	require.Equal(t, entities.ErrNotFound, err)
}

type randy struct {
}

func (r randy) Float32() float32 {
	return float32(5)
}
func (r randy) Float64() float64 {
	return float64(5.5)
}

func (r randy) Int31() int32 {
	return int32(5)
}
func (r randy) Int63() int64 {
	return int64(5)
}
func (r randy) Uint32() uint32 {
	return uint32(5456)
}
func (r randy) Intn(_ int) int {
	return int(5)
}
