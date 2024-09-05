// Code generated by assertions generator; DO NOT EDIT.

package resourceassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
)

type WarehouseResourceAssert struct {
	*assert.ResourceAssert
}

func WarehouseResource(t *testing.T, name string) *WarehouseResourceAssert {
	t.Helper()

	return &WarehouseResourceAssert{
		ResourceAssert: assert.NewResourceAssert(name, "resource"),
	}
}

func ImportedWarehouseResource(t *testing.T, id string) *WarehouseResourceAssert {
	t.Helper()

	return &WarehouseResourceAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "imported resource"),
	}
}

///////////////////////////////////
// Attribute value string checks //
///////////////////////////////////

func (w *WarehouseResourceAssert) HasAutoResumeString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("auto_resume", expected))
	return w
}

func (w *WarehouseResourceAssert) HasAutoSuspendString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("auto_suspend", expected))
	return w
}

func (w *WarehouseResourceAssert) HasCommentString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("comment", expected))
	return w
}

func (w *WarehouseResourceAssert) HasEnableQueryAccelerationString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("enable_query_acceleration", expected))
	return w
}

func (w *WarehouseResourceAssert) HasFullyQualifiedNameString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("fully_qualified_name", expected))
	return w
}

func (w *WarehouseResourceAssert) HasInitiallySuspendedString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("initially_suspended", expected))
	return w
}

func (w *WarehouseResourceAssert) HasMaxClusterCountString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("max_cluster_count", expected))
	return w
}

func (w *WarehouseResourceAssert) HasMaxConcurrencyLevelString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("max_concurrency_level", expected))
	return w
}

func (w *WarehouseResourceAssert) HasMinClusterCountString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("min_cluster_count", expected))
	return w
}

func (w *WarehouseResourceAssert) HasNameString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("name", expected))
	return w
}

func (w *WarehouseResourceAssert) HasQueryAccelerationMaxScaleFactorString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("query_acceleration_max_scale_factor", expected))
	return w
}

func (w *WarehouseResourceAssert) HasResourceMonitorString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("resource_monitor", expected))
	return w
}

func (w *WarehouseResourceAssert) HasScalingPolicyString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("scaling_policy", expected))
	return w
}

func (w *WarehouseResourceAssert) HasStatementQueuedTimeoutInSecondsString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("statement_queued_timeout_in_seconds", expected))
	return w
}

func (w *WarehouseResourceAssert) HasStatementTimeoutInSecondsString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("statement_timeout_in_seconds", expected))
	return w
}

func (w *WarehouseResourceAssert) HasWarehouseSizeString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("warehouse_size", expected))
	return w
}

func (w *WarehouseResourceAssert) HasWarehouseTypeString(expected string) *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueSet("warehouse_type", expected))
	return w
}

////////////////////////////
// Attribute empty checks //
////////////////////////////

func (w *WarehouseResourceAssert) HasNoAutoResume() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("auto_resume"))
	return w
}

func (w *WarehouseResourceAssert) HasNoAutoSuspend() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("auto_suspend"))
	return w
}

func (w *WarehouseResourceAssert) HasNoComment() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("comment"))
	return w
}

func (w *WarehouseResourceAssert) HasNoEnableQueryAcceleration() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("enable_query_acceleration"))
	return w
}

func (w *WarehouseResourceAssert) HasNoFullyQualifiedName() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("fully_qualified_name"))
	return w
}

func (w *WarehouseResourceAssert) HasNoInitiallySuspended() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("initially_suspended"))
	return w
}

func (w *WarehouseResourceAssert) HasNoMaxClusterCount() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("max_cluster_count"))
	return w
}

func (w *WarehouseResourceAssert) HasNoMaxConcurrencyLevel() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("max_concurrency_level"))
	return w
}

func (w *WarehouseResourceAssert) HasNoMinClusterCount() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("min_cluster_count"))
	return w
}

func (w *WarehouseResourceAssert) HasNoName() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("name"))
	return w
}

func (w *WarehouseResourceAssert) HasNoQueryAccelerationMaxScaleFactor() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("query_acceleration_max_scale_factor"))
	return w
}

func (w *WarehouseResourceAssert) HasNoResourceMonitor() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("resource_monitor"))
	return w
}

func (w *WarehouseResourceAssert) HasNoScalingPolicy() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("scaling_policy"))
	return w
}

func (w *WarehouseResourceAssert) HasNoStatementQueuedTimeoutInSeconds() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("statement_queued_timeout_in_seconds"))
	return w
}

func (w *WarehouseResourceAssert) HasNoStatementTimeoutInSeconds() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("statement_timeout_in_seconds"))
	return w
}

func (w *WarehouseResourceAssert) HasNoWarehouseSize() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("warehouse_size"))
	return w
}

func (w *WarehouseResourceAssert) HasNoWarehouseType() *WarehouseResourceAssert {
	w.AddAssertion(assert.ValueNotSet("warehouse_type"))
	return w
}