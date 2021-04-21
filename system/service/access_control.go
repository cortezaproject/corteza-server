package service

import (
	"github.com/cortezaproject/corteza-server/pkg/messagebus"
)

// Addition to list of defined resources
// we need that because messagebus resource is not compliant
func (svc accessControl) list() []map[string]string {
	return []map[string]string{
		{"type": messagebus.QueueResourceType, "any": messagebus.QueueRbacResource(0), "op": "read"},
		{"type": messagebus.QueueResourceType, "any": messagebus.QueueRbacResource(0), "op": "update"},
		{"type": messagebus.QueueResourceType, "any": messagebus.QueueRbacResource(0), "op": "delete"},
		{"type": messagebus.QueueResourceType, "any": messagebus.QueueRbacResource(0), "op": "queue.read"},
		{"type": messagebus.QueueResourceType, "any": messagebus.QueueRbacResource(0), "op": "queue.write"},
	}
}
