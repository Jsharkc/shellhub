// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/shellhub-io/shellhub/pkg/models"
	mock "github.com/stretchr/testify/mock"

	paginator "github.com/shellhub-io/shellhub/pkg/api/paginator"

	time "time"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// BillingDeleteCustomer provides a mock function with given fields: ctx, namespace
func (_m *Store) BillingDeleteCustomer(ctx context.Context, namespace *models.Namespace) error {
	ret := _m.Called(ctx, namespace)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Namespace) error); ok {
		r0 = rf(ctx, namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BillingDeleteSubscription provides a mock function with given fields: ctx, tenantID
func (_m *Store) BillingDeleteSubscription(ctx context.Context, tenantID string) error {
	ret := _m.Called(ctx, tenantID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, tenantID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BillingRemoveInstance provides a mock function with given fields: ctx, subsID
func (_m *Store) BillingRemoveInstance(ctx context.Context, subsID string) error {
	ret := _m.Called(ctx, subsID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, subsID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BillingUpdateCustomer provides a mock function with given fields: ctx, namespace, custID
func (_m *Store) BillingUpdateCustomer(ctx context.Context, namespace *models.Namespace, custID string) error {
	ret := _m.Called(ctx, namespace, custID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Namespace, string) error); ok {
		r0 = rf(ctx, namespace, custID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BillingUpdateDeviceLimit provides a mock function with given fields: ctx, subscriptionID, newLimit
func (_m *Store) BillingUpdateDeviceLimit(ctx context.Context, subscriptionID string, newLimit int) (*models.Namespace, error) {
	ret := _m.Called(ctx, subscriptionID, newLimit)

	var r0 *models.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, string, int) *models.Namespace); ok {
		r0 = rf(ctx, subscriptionID, newLimit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, subscriptionID, newLimit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BillingUpdatePaymentFailed provides a mock function with given fields: ctx, subscriptionID, set, pf
func (_m *Store) BillingUpdatePaymentFailed(ctx context.Context, subscriptionID string, set bool, pf *models.PaymentFailed) (*models.Namespace, error) {
	ret := _m.Called(ctx, subscriptionID, set, pf)

	var r0 *models.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, string, bool, *models.PaymentFailed) *models.Namespace); ok {
		r0 = rf(ctx, subscriptionID, set, pf)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool, *models.PaymentFailed) error); ok {
		r1 = rf(ctx, subscriptionID, set, pf)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BillingUpdatePaymentID provides a mock function with given fields: ctx, namespace, paymentID
func (_m *Store) BillingUpdatePaymentID(ctx context.Context, namespace *models.Namespace, paymentID string) error {
	ret := _m.Called(ctx, namespace, paymentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Namespace, string) error); ok {
		r0 = rf(ctx, namespace, paymentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BillingUpdateSubscription provides a mock function with given fields: ctx, namespace, billing
func (_m *Store) BillingUpdateSubscription(ctx context.Context, namespace *models.Namespace, billing *models.Billing) error {
	ret := _m.Called(ctx, namespace, billing)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Namespace, *models.Billing) error); ok {
		r0 = rf(ctx, namespace, billing)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BillingUpdateSubscriptionPeriodEnd provides a mock function with given fields: ctx, subscriptionID, periodEnd
func (_m *Store) BillingUpdateSubscriptionPeriodEnd(ctx context.Context, subscriptionID string, periodEnd time.Time) error {
	ret := _m.Called(ctx, subscriptionID, periodEnd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time) error); ok {
		r0 = rf(ctx, subscriptionID, periodEnd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeviceCreate provides a mock function with given fields: ctx, d, hostname
func (_m *Store) DeviceCreate(ctx context.Context, d models.Device, hostname string) error {
	ret := _m.Called(ctx, d, hostname)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Device, string) error); ok {
		r0 = rf(ctx, d, hostname)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeviceDelete provides a mock function with given fields: ctx, uid
func (_m *Store) DeviceDelete(ctx context.Context, uid models.UID) error {
	ret := _m.Called(ctx, uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) error); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeviceGet provides a mock function with given fields: ctx, uid
func (_m *Store) DeviceGet(ctx context.Context, uid models.UID) (*models.Device, error) {
	ret := _m.Called(ctx, uid)

	var r0 *models.Device
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) *models.Device); ok {
		r0 = rf(ctx, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.UID) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceGetByMac provides a mock function with given fields: ctx, mac, tenant, status
func (_m *Store) DeviceGetByMac(ctx context.Context, mac string, tenant string, status string) (*models.Device, error) {
	ret := _m.Called(ctx, mac, tenant, status)

	var r0 *models.Device
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *models.Device); ok {
		r0 = rf(ctx, mac, tenant, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, mac, tenant, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceGetByName provides a mock function with given fields: ctx, name, tenant
func (_m *Store) DeviceGetByName(ctx context.Context, name string, tenant string) (*models.Device, error) {
	ret := _m.Called(ctx, name, tenant)

	var r0 *models.Device
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.Device); ok {
		r0 = rf(ctx, name, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceGetByUID provides a mock function with given fields: ctx, uid, tenant
func (_m *Store) DeviceGetByUID(ctx context.Context, uid models.UID, tenant string) (*models.Device, error) {
	ret := _m.Called(ctx, uid, tenant)

	var r0 *models.Device
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, string) *models.Device); ok {
		r0 = rf(ctx, uid, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.UID, string) error); ok {
		r1 = rf(ctx, uid, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceList provides a mock function with given fields: ctx, pagination, filters, status, sort, order
func (_m *Store) DeviceList(ctx context.Context, pagination paginator.Query, filters []models.Filter, status string, sort string, order string) ([]models.Device, int, error) {
	ret := _m.Called(ctx, pagination, filters, status, sort, order)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(context.Context, paginator.Query, []models.Filter, string, string, string) []models.Device); ok {
		r0 = rf(ctx, pagination, filters, status, sort, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, paginator.Query, []models.Filter, string, string, string) int); ok {
		r1 = rf(ctx, pagination, filters, status, sort, order)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, paginator.Query, []models.Filter, string, string, string) error); ok {
		r2 = rf(ctx, pagination, filters, status, sort, order)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeviceLookup provides a mock function with given fields: ctx, namespace, name
func (_m *Store) DeviceLookup(ctx context.Context, namespace string, name string) (*models.Device, error) {
	ret := _m.Called(ctx, namespace, name)

	var r0 *models.Device
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.Device); ok {
		r0 = rf(ctx, namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceRename provides a mock function with given fields: ctx, uid, name
func (_m *Store) DeviceRename(ctx context.Context, uid models.UID, name string) error {
	ret := _m.Called(ctx, uid, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, string) error); ok {
		r0 = rf(ctx, uid, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeviceSetOnline provides a mock function with given fields: ctx, uid, online
func (_m *Store) DeviceSetOnline(ctx context.Context, uid models.UID, online bool) error {
	ret := _m.Called(ctx, uid, online)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, bool) error); ok {
		r0 = rf(ctx, uid, online)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeviceUpdateStatus provides a mock function with given fields: ctx, uid, status
func (_m *Store) DeviceUpdateStatus(ctx context.Context, uid models.UID, status string) error {
	ret := _m.Called(ctx, uid, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, string) error); ok {
		r0 = rf(ctx, uid, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirewallRuleCreate provides a mock function with given fields: ctx, rule
func (_m *Store) FirewallRuleCreate(ctx context.Context, rule *models.FirewallRule) error {
	ret := _m.Called(ctx, rule)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.FirewallRule) error); ok {
		r0 = rf(ctx, rule)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirewallRuleDelete provides a mock function with given fields: ctx, id
func (_m *Store) FirewallRuleDelete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirewallRuleGet provides a mock function with given fields: ctx, id
func (_m *Store) FirewallRuleGet(ctx context.Context, id string) (*models.FirewallRule, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.FirewallRule
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.FirewallRule); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.FirewallRule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FirewallRuleList provides a mock function with given fields: ctx, pagination
func (_m *Store) FirewallRuleList(ctx context.Context, pagination paginator.Query) ([]models.FirewallRule, int, error) {
	ret := _m.Called(ctx, pagination)

	var r0 []models.FirewallRule
	if rf, ok := ret.Get(0).(func(context.Context, paginator.Query) []models.FirewallRule); ok {
		r0 = rf(ctx, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.FirewallRule)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, paginator.Query) int); ok {
		r1 = rf(ctx, pagination)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, paginator.Query) error); ok {
		r2 = rf(ctx, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FirewallRuleUpdate provides a mock function with given fields: ctx, id, rule
func (_m *Store) FirewallRuleUpdate(ctx context.Context, id string, rule models.FirewallRuleUpdate) (*models.FirewallRule, error) {
	ret := _m.Called(ctx, id, rule)

	var r0 *models.FirewallRule
	if rf, ok := ret.Get(0).(func(context.Context, string, models.FirewallRuleUpdate) *models.FirewallRule); ok {
		r0 = rf(ctx, id, rule)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.FirewallRule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, models.FirewallRuleUpdate) error); ok {
		r1 = rf(ctx, id, rule)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStats provides a mock function with given fields: ctx
func (_m *Store) GetStats(ctx context.Context) (*models.Stats, error) {
	ret := _m.Called(ctx)

	var r0 *models.Stats
	if rf, ok := ret.Get(0).(func(context.Context) *models.Stats); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Stats)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LicenseLoad provides a mock function with given fields: ctx
func (_m *Store) LicenseLoad(ctx context.Context) (*models.License, error) {
	ret := _m.Called(ctx)

	var r0 *models.License
	if rf, ok := ret.Get(0).(func(context.Context) *models.License); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.License)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LicenseSave provides a mock function with given fields: ctx, license
func (_m *Store) LicenseSave(ctx context.Context, license *models.License) error {
	ret := _m.Called(ctx, license)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.License) error); ok {
		r0 = rf(ctx, license)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NamespaceAddMember provides a mock function with given fields: ctx, tenantID, ID
func (_m *Store) NamespaceAddMember(ctx context.Context, tenantID string, ID string) (*models.Namespace, error) {
	ret := _m.Called(ctx, tenantID, ID)

	var r0 *models.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.Namespace); ok {
		r0 = rf(ctx, tenantID, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tenantID, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceCreate provides a mock function with given fields: ctx, tenantID
func (_m *Store) NamespaceCreate(ctx context.Context, tenantID *models.Namespace) (*models.Namespace, error) {
	ret := _m.Called(ctx, tenantID)

	var r0 *models.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, *models.Namespace) *models.Namespace); ok {
		r0 = rf(ctx, tenantID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Namespace) error); ok {
		r1 = rf(ctx, tenantID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceDelete provides a mock function with given fields: ctx, tenantID
func (_m *Store) NamespaceDelete(ctx context.Context, tenantID string) error {
	ret := _m.Called(ctx, tenantID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, tenantID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NamespaceGet provides a mock function with given fields: ctx, tenantID
func (_m *Store) NamespaceGet(ctx context.Context, tenantID string) (*models.Namespace, error) {
	ret := _m.Called(ctx, tenantID)

	var r0 *models.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Namespace); ok {
		r0 = rf(ctx, tenantID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tenantID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceGetByName provides a mock function with given fields: ctx, tenantID
func (_m *Store) NamespaceGetByName(ctx context.Context, tenantID string) (*models.Namespace, error) {
	ret := _m.Called(ctx, tenantID)

	var r0 *models.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Namespace); ok {
		r0 = rf(ctx, tenantID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tenantID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceGetFirst provides a mock function with given fields: ctx, ID
func (_m *Store) NamespaceGetFirst(ctx context.Context, ID string) (*models.Namespace, error) {
	ret := _m.Called(ctx, ID)

	var r0 *models.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Namespace); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceGetSessionRecord provides a mock function with given fields: ctx, tenantID
func (_m *Store) NamespaceGetSessionRecord(ctx context.Context, tenantID string) (bool, error) {
	ret := _m.Called(ctx, tenantID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, tenantID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tenantID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceList provides a mock function with given fields: ctx, pagination, filters, export
func (_m *Store) NamespaceList(ctx context.Context, pagination paginator.Query, filters []models.Filter, export bool) ([]models.Namespace, int, error) {
	ret := _m.Called(ctx, pagination, filters, export)

	var r0 []models.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, paginator.Query, []models.Filter, bool) []models.Namespace); ok {
		r0 = rf(ctx, pagination, filters, export)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Namespace)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, paginator.Query, []models.Filter, bool) int); ok {
		r1 = rf(ctx, pagination, filters, export)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, paginator.Query, []models.Filter, bool) error); ok {
		r2 = rf(ctx, pagination, filters, export)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NamespaceRemoveMember provides a mock function with given fields: ctx, tenantID, ID
func (_m *Store) NamespaceRemoveMember(ctx context.Context, tenantID string, ID string) (*models.Namespace, error) {
	ret := _m.Called(ctx, tenantID, ID)

	var r0 *models.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.Namespace); ok {
		r0 = rf(ctx, tenantID, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tenantID, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceRename provides a mock function with given fields: ctx, tenantID, name
func (_m *Store) NamespaceRename(ctx context.Context, tenantID string, name string) (*models.Namespace, error) {
	ret := _m.Called(ctx, tenantID, name)

	var r0 *models.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.Namespace); ok {
		r0 = rf(ctx, tenantID, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tenantID, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceSetSessionRecord provides a mock function with given fields: ctx, sessionRecord, tenantID
func (_m *Store) NamespaceSetSessionRecord(ctx context.Context, sessionRecord bool, tenantID string) error {
	ret := _m.Called(ctx, sessionRecord, tenantID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, bool, string) error); ok {
		r0 = rf(ctx, sessionRecord, tenantID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NamespaceUpdate provides a mock function with given fields: ctx, tenantID, namespace
func (_m *Store) NamespaceUpdate(ctx context.Context, tenantID string, namespace *models.Namespace) error {
	ret := _m.Called(ctx, tenantID, namespace)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *models.Namespace) error); ok {
		r0 = rf(ctx, tenantID, namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PrivateKeyCreate provides a mock function with given fields: ctx, key
func (_m *Store) PrivateKeyCreate(ctx context.Context, key *models.PrivateKey) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.PrivateKey) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PrivateKeyGet provides a mock function with given fields: ctx, fingerprint
func (_m *Store) PrivateKeyGet(ctx context.Context, fingerprint string) (*models.PrivateKey, error) {
	ret := _m.Called(ctx, fingerprint)

	var r0 *models.PrivateKey
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.PrivateKey); ok {
		r0 = rf(ctx, fingerprint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PrivateKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, fingerprint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublicKeyCreate provides a mock function with given fields: ctx, key
func (_m *Store) PublicKeyCreate(ctx context.Context, key *models.PublicKey) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.PublicKey) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublicKeyDelete provides a mock function with given fields: ctx, fingerprint, tenant
func (_m *Store) PublicKeyDelete(ctx context.Context, fingerprint string, tenant string) error {
	ret := _m.Called(ctx, fingerprint, tenant)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, fingerprint, tenant)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublicKeyGet provides a mock function with given fields: ctx, fingerprint, tenant
func (_m *Store) PublicKeyGet(ctx context.Context, fingerprint string, tenant string) (*models.PublicKey, error) {
	ret := _m.Called(ctx, fingerprint, tenant)

	var r0 *models.PublicKey
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.PublicKey); ok {
		r0 = rf(ctx, fingerprint, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PublicKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, fingerprint, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublicKeyList provides a mock function with given fields: ctx, pagination
func (_m *Store) PublicKeyList(ctx context.Context, pagination paginator.Query) ([]models.PublicKey, int, error) {
	ret := _m.Called(ctx, pagination)

	var r0 []models.PublicKey
	if rf, ok := ret.Get(0).(func(context.Context, paginator.Query) []models.PublicKey); ok {
		r0 = rf(ctx, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PublicKey)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, paginator.Query) int); ok {
		r1 = rf(ctx, pagination)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, paginator.Query) error); ok {
		r2 = rf(ctx, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PublicKeyUpdate provides a mock function with given fields: ctx, fingerprint, tenant, key
func (_m *Store) PublicKeyUpdate(ctx context.Context, fingerprint string, tenant string, key *models.PublicKeyUpdate) (*models.PublicKey, error) {
	ret := _m.Called(ctx, fingerprint, tenant, key)

	var r0 *models.PublicKey
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *models.PublicKeyUpdate) *models.PublicKey); ok {
		r0 = rf(ctx, fingerprint, tenant, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PublicKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *models.PublicKeyUpdate) error); ok {
		r1 = rf(ctx, fingerprint, tenant, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SessionCreate provides a mock function with given fields: ctx, session
func (_m *Store) SessionCreate(ctx context.Context, session models.Session) (*models.Session, error) {
	ret := _m.Called(ctx, session)

	var r0 *models.Session
	if rf, ok := ret.Get(0).(func(context.Context, models.Session) *models.Session); ok {
		r0 = rf(ctx, session)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Session) error); ok {
		r1 = rf(ctx, session)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SessionCreateRecordFrame provides a mock function with given fields: ctx, uid, recordSession
func (_m *Store) SessionCreateRecordFrame(ctx context.Context, uid models.UID, recordSession *models.RecordedSession) error {
	ret := _m.Called(ctx, uid, recordSession)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, *models.RecordedSession) error); ok {
		r0 = rf(ctx, uid, recordSession)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionDeleteActives provides a mock function with given fields: ctx, uid
func (_m *Store) SessionDeleteActives(ctx context.Context, uid models.UID) error {
	ret := _m.Called(ctx, uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) error); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionDeleteRecordFrame provides a mock function with given fields: ctx, uid
func (_m *Store) SessionDeleteRecordFrame(ctx context.Context, uid models.UID) error {
	ret := _m.Called(ctx, uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) error); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionGet provides a mock function with given fields: ctx, uid
func (_m *Store) SessionGet(ctx context.Context, uid models.UID) (*models.Session, error) {
	ret := _m.Called(ctx, uid)

	var r0 *models.Session
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) *models.Session); ok {
		r0 = rf(ctx, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.UID) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SessionGetRecordFrame provides a mock function with given fields: ctx, uid
func (_m *Store) SessionGetRecordFrame(ctx context.Context, uid models.UID) ([]models.RecordedSession, int, error) {
	ret := _m.Called(ctx, uid)

	var r0 []models.RecordedSession
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) []models.RecordedSession); ok {
		r0 = rf(ctx, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.RecordedSession)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, models.UID) int); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, models.UID) error); ok {
		r2 = rf(ctx, uid)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SessionList provides a mock function with given fields: ctx, pagination
func (_m *Store) SessionList(ctx context.Context, pagination paginator.Query) ([]models.Session, int, error) {
	ret := _m.Called(ctx, pagination)

	var r0 []models.Session
	if rf, ok := ret.Get(0).(func(context.Context, paginator.Query) []models.Session); ok {
		r0 = rf(ctx, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Session)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, paginator.Query) int); ok {
		r1 = rf(ctx, pagination)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, paginator.Query) error); ok {
		r2 = rf(ctx, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SessionSetAuthenticated provides a mock function with given fields: ctx, uid, authenticated
func (_m *Store) SessionSetAuthenticated(ctx context.Context, uid models.UID, authenticated bool) error {
	ret := _m.Called(ctx, uid, authenticated)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, bool) error); ok {
		r0 = rf(ctx, uid, authenticated)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionSetLastSeen provides a mock function with given fields: ctx, uid
func (_m *Store) SessionSetLastSeen(ctx context.Context, uid models.UID) error {
	ret := _m.Called(ctx, uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) error); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionSetRecorded provides a mock function with given fields: ctx, uid, recorded
func (_m *Store) SessionSetRecorded(ctx context.Context, uid models.UID, recorded bool) error {
	ret := _m.Called(ctx, uid, recorded)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, bool) error); ok {
		r0 = rf(ctx, uid, recorded)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionUpdateDeviceUID provides a mock function with given fields: ctx, oldUID, newUID
func (_m *Store) SessionUpdateDeviceUID(ctx context.Context, oldUID models.UID, newUID models.UID) error {
	ret := _m.Called(ctx, oldUID, newUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, models.UID) error); ok {
		r0 = rf(ctx, oldUID, newUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserCreate provides a mock function with given fields: ctx, user
func (_m *Store) UserCreate(ctx context.Context, user *models.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserCreateToken provides a mock function with given fields: ctx, token
func (_m *Store) UserCreateToken(ctx context.Context, token *models.UserTokenRecover) error {
	ret := _m.Called(ctx, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.UserTokenRecover) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserDelete provides a mock function with given fields: ctx, ID
func (_m *Store) UserDelete(ctx context.Context, ID string) error {
	ret := _m.Called(ctx, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserDeleteTokens provides a mock function with given fields: ctx, ID
func (_m *Store) UserDeleteTokens(ctx context.Context, ID string) error {
	ret := _m.Called(ctx, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserGetByEmail provides a mock function with given fields: ctx, email
func (_m *Store) UserGetByEmail(ctx context.Context, email string) (*models.User, error) {
	ret := _m.Called(ctx, email)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserGetByID provides a mock function with given fields: ctx, ID, ns
func (_m *Store) UserGetByID(ctx context.Context, ID string, ns bool) (*models.User, int, error) {
	ret := _m.Called(ctx, ID, ns)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) *models.User); ok {
		r0 = rf(ctx, ID, ns)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) int); ok {
		r1 = rf(ctx, ID, ns)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, bool) error); ok {
		r2 = rf(ctx, ID, ns)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UserGetByTenant provides a mock function with given fields: ctx, tenant
func (_m *Store) UserGetByTenant(ctx context.Context, tenant string) (*models.User, error) {
	ret := _m.Called(ctx, tenant)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(ctx, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserGetByUsername provides a mock function with given fields: ctx, username
func (_m *Store) UserGetByUsername(ctx context.Context, username string) (*models.User, error) {
	ret := _m.Called(ctx, username)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserGetToken provides a mock function with given fields: ctx, ID
func (_m *Store) UserGetToken(ctx context.Context, ID string) (*models.UserTokenRecover, error) {
	ret := _m.Called(ctx, ID)

	var r0 *models.UserTokenRecover
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.UserTokenRecover); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserTokenRecover)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserList provides a mock function with given fields: ctx, pagination, filters
func (_m *Store) UserList(ctx context.Context, pagination paginator.Query, filters []models.Filter) ([]models.User, int, error) {
	ret := _m.Called(ctx, pagination, filters)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(context.Context, paginator.Query, []models.Filter) []models.User); ok {
		r0 = rf(ctx, pagination, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, paginator.Query, []models.Filter) int); ok {
		r1 = rf(ctx, pagination, filters)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, paginator.Query, []models.Filter) error); ok {
		r2 = rf(ctx, pagination, filters)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UserUpdateAccountStatus provides a mock function with given fields: ctx, ID
func (_m *Store) UserUpdateAccountStatus(ctx context.Context, ID string) error {
	ret := _m.Called(ctx, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserUpdateData provides a mock function with given fields: ctx, data, ID
func (_m *Store) UserUpdateData(ctx context.Context, data *models.User, ID string) error {
	ret := _m.Called(ctx, data, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User, string) error); ok {
		r0 = rf(ctx, data, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserUpdateFromAdmin provides a mock function with given fields: ctx, name, username, email, password, ID
func (_m *Store) UserUpdateFromAdmin(ctx context.Context, name string, username string, email string, password string, ID string) error {
	ret := _m.Called(ctx, name, username, email, password, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string) error); ok {
		r0 = rf(ctx, name, username, email, password, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserUpdatePassword provides a mock function with given fields: ctx, newPassword, ID
func (_m *Store) UserUpdatePassword(ctx context.Context, newPassword string, ID string) error {
	ret := _m.Called(ctx, newPassword, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, newPassword, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
