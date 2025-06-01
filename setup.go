package main

import (
	"time"
)

type Setup struct {
	ID             uint64                  `json:"id"`
	SiteID         string                  `json:"site_id"`
	Origin         string                  `json:"origin"`
	EntityType     string                  `json:"entity_type"`
	EntityTypeID   string                  `json:"entity_type_id"`
	CollectorID    uint64                  `json:"collector_id"`
	CreatedBy      string                  `json:"created_by"`
	ExpiredAt      *time.Time              `json:"expired_at"`
	InternalDetail InternalDetail          `json:"internal_detail"`
	ExternalDetail *map[string]interface{} `json:"external_detail"`
	CreatedAt      time.Time               `json:"created_at"`
}

type InternalDetail struct {
	UniversalFlow string `json:"universal_flow"`
	Coupon        string `json:"coupon"`
	DeviceID      uint64 `json:"device_id"`
	Source        string `json:"source"`
	Strategy      string `json:"strategy"`
	Campaign      string `json:"campaign"`
	Segment       string `json:"segment"`
	UserID        string `json:"user_id"`
	FakePrice     bool   `json:"fake_price"`
	Plan          string `json:"plan"`
	AudienceType  string `json:"audience_type"`
	Label         string `json:"label"`
}
