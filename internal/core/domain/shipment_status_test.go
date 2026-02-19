package domain

import (
	"testing"
)

func TestIsValidTransition(t *testing.T) {
	tests := []struct {
		name          string
		currentStatus ShipmentStatusList
		nextStatus    ShipmentStatusList
		expectedValue bool
	}{
		{
			name:          "created to inwarehouse invalid transition",
			currentStatus: ShipmentCreatedStatus,
			nextStatus:    ShipmentInWarehouseStatus,
			expectedValue: false,
		},
		{
			name:          "created to intransit invalid transition",
			currentStatus: ShipmentCreatedStatus,
			nextStatus:    ShipmentInTransitStatus,
			expectedValue: false,
		},
		{
			name:          "created to delivered invalid transition",
			currentStatus: ShipmentCreatedStatus,
			nextStatus:    ShipmentDeliveredStatus,
			expectedValue: false,
		},
		{
			name:          "pickedup to intransit invalid transition",
			currentStatus: ShipmentPickedUpStatus,
			nextStatus:    ShipmentInTransitStatus,
			expectedValue: false,
		},
		{
			name:          "created to pickedup valid transition",
			currentStatus: ShipmentCreatedStatus,
			nextStatus:    ShipmentPickedUpStatus,
			expectedValue: true,
		},
		{
			name:          "pickedup to pickedup valid transition",
			currentStatus: ShipmentPickedUpStatus,
			nextStatus:    ShipmentInWarehouseStatus,
			expectedValue: true,
		},
		{
			name:          "inwarehouse to intransit valid transition",
			currentStatus: ShipmentInWarehouseStatus,
			nextStatus:    ShipmentInTransitStatus,
			expectedValue: true,
		},
		{
			name:          "intransit to delivered valid transition",
			currentStatus: ShipmentInTransitStatus,
			nextStatus:    ShipmentDeliveredStatus,
			expectedValue: true,
		},
	}

	for _, tt := range tests {
		// Run each case as a sub-test
		t.Run(tt.name, func(t *testing.T) {
			// Execute
			isValidTransition := IsValidTransition(tt.currentStatus, tt.nextStatus)

			// Assertions
			if isValidTransition != tt.expectedValue {
			}
		})
	}
}
