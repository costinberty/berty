package mdns

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"berty.tech/weshnet/pkg/netmanager"
	"berty.tech/weshnet/pkg/testutil"
)

type fakeService struct {
	start_count int
	close_count int
	mu          sync.Mutex
	wgStart     sync.WaitGroup
	wgClose     sync.WaitGroup
}

func (s *fakeService) getStartCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.start_count
}

func (s *fakeService) getCloseCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.close_count
}

func (s *fakeService) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.start_count++
	s.wgStart.Done()
	return nil
}

func (s *fakeService) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.close_count++
	s.wgClose.Done()
	return nil
}

func TestMdns(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, cleanup := testutil.Logger(t)
	defer cleanup()

	net := netmanager.NewNetManager(netmanager.ConnectivityInfo{})

	require.NotNil(t, net)

	svc := &fakeService{0, 0, sync.Mutex{}, sync.WaitGroup{}, sync.WaitGroup{}}

	go func() {
		mdnNetworkManagerConfig := NetworkManagerConfig{
			Logger:              logger,
			NetManager:          net,
			Service:             svc,
			InitialConnectivity: netmanager.ConnectivityInfo{},
		}
		NetworkManagerHandler(ctx, mdnNetworkManagerConfig)
	}()

	require.Equal(t, 0, svc.getCloseCount())
	require.Equal(t, 0, svc.getStartCount())

	svc.wgClose.Add(1)
	svc.wgStart.Add(1)
	net.UpdateState(netmanager.ConnectivityInfo{
		State:   netmanager.ConnectivityStateOn,
		NetType: netmanager.ConnectivityNetWifi,
	})
	svc.wgClose.Wait()
	svc.wgStart.Wait()

	require.Equal(t, 1, svc.getCloseCount())
	require.Equal(t, 1, svc.getStartCount())

	svc.wgClose.Add(1)
	net.UpdateState(netmanager.ConnectivityInfo{
		State: netmanager.ConnectivityStateOff,
	})
	svc.wgClose.Wait()

	require.Equal(t, 2, svc.getCloseCount())
	require.Equal(t, 1, svc.getStartCount())

	svc.wgClose.Add(1)
	svc.wgStart.Add(1)
	net.UpdateState(netmanager.ConnectivityInfo{
		State:   netmanager.ConnectivityStateOn,
		NetType: netmanager.ConnectivityNetWifi,
	})
	svc.wgClose.Wait()
	svc.wgStart.Wait()

	require.Equal(t, 3, svc.getCloseCount())
	require.Equal(t, 2, svc.getStartCount())
}
