package pinger_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/echovl/pinger"
	"github.com/echovl/pinger/mocks"
	"github.com/golang/mock/gomock"
)

func TestPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)
	core := pinger.NewCore(mockDB, 60*time.Second)
	hosts := []*pinger.Host{
		{ID: 1, Name: "Arch Linux", URL: "https://archlinux.org/"},
		{ID: 2, Name: "Github", URL: "https://github.com/"},
	}

	mockDB.EXPECT().
		GetHost(context.Background(), 1).Return(nil, nil)
	mockDB.EXPECT().
		GetHost(context.Background(), 2).Return(nil, nil)

	mockDB.EXPECT().
		UpsertHost(context.Background(), hosts[0]).Return(nil)
	mockDB.EXPECT().
		UpsertHost(context.Background(), hosts[1]).Return(nil)

	if err := core.Ping(hosts...); err != nil {
		t.Error(err)
	}

	notFoundErr := errors.New("host not found")
	mockDB.EXPECT().
		GetHost(context.Background(), 1).Return(nil, notFoundErr)
	mockDB.EXPECT().
		GetHost(context.Background(), 2).Return(nil, nil)
	mockDB.EXPECT().
		UpsertHost(context.Background(), hosts[1]).Return(nil)

	if err := core.Ping(hosts...); err != nil {
		t.Error(err)
	}

	unexpectedErr := errors.New("upsert failed for some reason")
	mockDB.EXPECT().
		GetHost(context.Background(), 1).Return(nil, nil)
	mockDB.EXPECT().
		UpsertHost(context.Background(), hosts[0]).Return(unexpectedErr)

	if err := core.Ping(hosts...); err == nil {
		t.Errorf("expected err: %s", unexpectedErr)
	}

}
