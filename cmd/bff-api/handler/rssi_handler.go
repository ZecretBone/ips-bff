package handler

import (
	"net/http"

	v1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/rssi/v1"
	rssiv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/shared/rssi/v1"
	rssiclient "github.com/ZecretBone/ips-bff/internal/repository/RSSIClient"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RSSIHandler interface {
	Get(ctx *gin.Context)
}

type rssiHandler struct {
	rssiClient rssiclient.Service
}

func ProvideRSSIHandler(rssiClient rssiclient.Service) RSSIHandler {
	return &rssiHandler{
		rssiClient: rssiClient,
	}
}

func (rs *rssiHandler) Get(ctx *gin.Context) {
	body := v1.CollectDataRequest{
		Signals: []*rssiv1.RSSI{
			{
				Ssid:        "newtONE",
				MacAddress:  "ab:cd:ef:gh:xy:xz",
				Strength:    60.3,
				PollingRate: 300,
				CreatedAt:   timestamppb.Now(),
			}, {
				Ssid:        "newtTWO",
				MacAddress:  "ab:cd:ef:gh:xy:x1",
				Strength:    50.3,
				PollingRate: 300,
				CreatedAt:   timestamppb.Now(),
			},
		},
		DeviceInfo: &rssiv1.DeviceInfo{
			DeviceId: "FXM1234",
			Models:   "androidOreo123",
		},
		Stage: rssiv1.StatCollectionStage_STAT_COLLECTION_STAGE_MULTIPLE,
		Position: &rssiv1.Position{
			X: 3,
			Y: 2,
			Z: 1,
		},
		Duration:  10,
		StartedAt: timestamppb.Now(),
		EndedAt:   timestamppb.Now(),
		CreatedAt: timestamppb.Now(),
	}

	if _, err := rs.rssiClient.CollectData(ctx, &body); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
}
