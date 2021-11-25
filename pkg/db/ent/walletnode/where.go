// Code generated by entc, DO NOT EDIT.

package walletnode

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocation), v))
	})
}

// HostVendor applies equality check predicate on the "host_vendor" field. It's identical to HostVendorEQ.
func HostVendor(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHostVendor), v))
	})
}

// PublicIP applies equality check predicate on the "public_ip" field. It's identical to PublicIPEQ.
func PublicIP(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublicIP), v))
	})
}

// LocalIP applies equality check predicate on the "local_ip" field. It's identical to LocalIPEQ.
func LocalIP(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocalIP), v))
	})
}

// CreatetimeUtc applies equality check predicate on the "createtime_utc" field. It's identical to CreatetimeUtcEQ.
func CreatetimeUtc(v int64) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatetimeUtc), v))
	})
}

// LastOnlineTimeUtc applies equality check predicate on the "last_online_time_utc" field. It's identical to LastOnlineTimeUtcEQ.
func LastOnlineTimeUtc(v int64) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastOnlineTimeUtc), v))
	})
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUUID), v))
	})
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.WalletNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WalletNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUUID), v...))
	})
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.WalletNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WalletNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUUID), v...))
	})
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUUID), v))
	})
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUUID), v))
	})
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUUID), v))
	})
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUUID), v))
	})
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUUID), v))
	})
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUUID), v))
	})
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUUID), v))
	})
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUUID), v))
	})
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUUID), v))
	})
}

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocation), v))
	})
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocation), v))
	})
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...string) predicate.WalletNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WalletNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLocation), v...))
	})
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...string) predicate.WalletNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WalletNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLocation), v...))
	})
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocation), v))
	})
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocation), v))
	})
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocation), v))
	})
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocation), v))
	})
}

// LocationContains applies the Contains predicate on the "location" field.
func LocationContains(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLocation), v))
	})
}

// LocationHasPrefix applies the HasPrefix predicate on the "location" field.
func LocationHasPrefix(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLocation), v))
	})
}

// LocationHasSuffix applies the HasSuffix predicate on the "location" field.
func LocationHasSuffix(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLocation), v))
	})
}

// LocationEqualFold applies the EqualFold predicate on the "location" field.
func LocationEqualFold(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLocation), v))
	})
}

// LocationContainsFold applies the ContainsFold predicate on the "location" field.
func LocationContainsFold(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLocation), v))
	})
}

// HostVendorEQ applies the EQ predicate on the "host_vendor" field.
func HostVendorEQ(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHostVendor), v))
	})
}

// HostVendorNEQ applies the NEQ predicate on the "host_vendor" field.
func HostVendorNEQ(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHostVendor), v))
	})
}

// HostVendorIn applies the In predicate on the "host_vendor" field.
func HostVendorIn(vs ...string) predicate.WalletNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WalletNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHostVendor), v...))
	})
}

// HostVendorNotIn applies the NotIn predicate on the "host_vendor" field.
func HostVendorNotIn(vs ...string) predicate.WalletNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WalletNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHostVendor), v...))
	})
}

// HostVendorGT applies the GT predicate on the "host_vendor" field.
func HostVendorGT(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHostVendor), v))
	})
}

// HostVendorGTE applies the GTE predicate on the "host_vendor" field.
func HostVendorGTE(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHostVendor), v))
	})
}

// HostVendorLT applies the LT predicate on the "host_vendor" field.
func HostVendorLT(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHostVendor), v))
	})
}

// HostVendorLTE applies the LTE predicate on the "host_vendor" field.
func HostVendorLTE(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHostVendor), v))
	})
}

// HostVendorContains applies the Contains predicate on the "host_vendor" field.
func HostVendorContains(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHostVendor), v))
	})
}

// HostVendorHasPrefix applies the HasPrefix predicate on the "host_vendor" field.
func HostVendorHasPrefix(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHostVendor), v))
	})
}

// HostVendorHasSuffix applies the HasSuffix predicate on the "host_vendor" field.
func HostVendorHasSuffix(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHostVendor), v))
	})
}

// HostVendorEqualFold applies the EqualFold predicate on the "host_vendor" field.
func HostVendorEqualFold(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHostVendor), v))
	})
}

// HostVendorContainsFold applies the ContainsFold predicate on the "host_vendor" field.
func HostVendorContainsFold(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHostVendor), v))
	})
}

// PublicIPEQ applies the EQ predicate on the "public_ip" field.
func PublicIPEQ(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublicIP), v))
	})
}

// PublicIPNEQ applies the NEQ predicate on the "public_ip" field.
func PublicIPNEQ(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPublicIP), v))
	})
}

// PublicIPIn applies the In predicate on the "public_ip" field.
func PublicIPIn(vs ...string) predicate.WalletNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WalletNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPublicIP), v...))
	})
}

// PublicIPNotIn applies the NotIn predicate on the "public_ip" field.
func PublicIPNotIn(vs ...string) predicate.WalletNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WalletNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPublicIP), v...))
	})
}

// PublicIPGT applies the GT predicate on the "public_ip" field.
func PublicIPGT(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPublicIP), v))
	})
}

// PublicIPGTE applies the GTE predicate on the "public_ip" field.
func PublicIPGTE(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPublicIP), v))
	})
}

// PublicIPLT applies the LT predicate on the "public_ip" field.
func PublicIPLT(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPublicIP), v))
	})
}

// PublicIPLTE applies the LTE predicate on the "public_ip" field.
func PublicIPLTE(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPublicIP), v))
	})
}

// PublicIPContains applies the Contains predicate on the "public_ip" field.
func PublicIPContains(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPublicIP), v))
	})
}

// PublicIPHasPrefix applies the HasPrefix predicate on the "public_ip" field.
func PublicIPHasPrefix(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPublicIP), v))
	})
}

// PublicIPHasSuffix applies the HasSuffix predicate on the "public_ip" field.
func PublicIPHasSuffix(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPublicIP), v))
	})
}

// PublicIPEqualFold applies the EqualFold predicate on the "public_ip" field.
func PublicIPEqualFold(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPublicIP), v))
	})
}

// PublicIPContainsFold applies the ContainsFold predicate on the "public_ip" field.
func PublicIPContainsFold(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPublicIP), v))
	})
}

// LocalIPEQ applies the EQ predicate on the "local_ip" field.
func LocalIPEQ(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocalIP), v))
	})
}

// LocalIPNEQ applies the NEQ predicate on the "local_ip" field.
func LocalIPNEQ(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocalIP), v))
	})
}

// LocalIPIn applies the In predicate on the "local_ip" field.
func LocalIPIn(vs ...string) predicate.WalletNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WalletNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLocalIP), v...))
	})
}

// LocalIPNotIn applies the NotIn predicate on the "local_ip" field.
func LocalIPNotIn(vs ...string) predicate.WalletNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WalletNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLocalIP), v...))
	})
}

// LocalIPGT applies the GT predicate on the "local_ip" field.
func LocalIPGT(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocalIP), v))
	})
}

// LocalIPGTE applies the GTE predicate on the "local_ip" field.
func LocalIPGTE(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocalIP), v))
	})
}

// LocalIPLT applies the LT predicate on the "local_ip" field.
func LocalIPLT(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocalIP), v))
	})
}

// LocalIPLTE applies the LTE predicate on the "local_ip" field.
func LocalIPLTE(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocalIP), v))
	})
}

// LocalIPContains applies the Contains predicate on the "local_ip" field.
func LocalIPContains(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLocalIP), v))
	})
}

// LocalIPHasPrefix applies the HasPrefix predicate on the "local_ip" field.
func LocalIPHasPrefix(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLocalIP), v))
	})
}

// LocalIPHasSuffix applies the HasSuffix predicate on the "local_ip" field.
func LocalIPHasSuffix(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLocalIP), v))
	})
}

// LocalIPEqualFold applies the EqualFold predicate on the "local_ip" field.
func LocalIPEqualFold(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLocalIP), v))
	})
}

// LocalIPContainsFold applies the ContainsFold predicate on the "local_ip" field.
func LocalIPContainsFold(v string) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLocalIP), v))
	})
}

// CreatetimeUtcEQ applies the EQ predicate on the "createtime_utc" field.
func CreatetimeUtcEQ(v int64) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatetimeUtc), v))
	})
}

// CreatetimeUtcNEQ applies the NEQ predicate on the "createtime_utc" field.
func CreatetimeUtcNEQ(v int64) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatetimeUtc), v))
	})
}

// CreatetimeUtcIn applies the In predicate on the "createtime_utc" field.
func CreatetimeUtcIn(vs ...int64) predicate.WalletNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WalletNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatetimeUtc), v...))
	})
}

// CreatetimeUtcNotIn applies the NotIn predicate on the "createtime_utc" field.
func CreatetimeUtcNotIn(vs ...int64) predicate.WalletNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WalletNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatetimeUtc), v...))
	})
}

// CreatetimeUtcGT applies the GT predicate on the "createtime_utc" field.
func CreatetimeUtcGT(v int64) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatetimeUtc), v))
	})
}

// CreatetimeUtcGTE applies the GTE predicate on the "createtime_utc" field.
func CreatetimeUtcGTE(v int64) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatetimeUtc), v))
	})
}

// CreatetimeUtcLT applies the LT predicate on the "createtime_utc" field.
func CreatetimeUtcLT(v int64) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatetimeUtc), v))
	})
}

// CreatetimeUtcLTE applies the LTE predicate on the "createtime_utc" field.
func CreatetimeUtcLTE(v int64) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatetimeUtc), v))
	})
}

// LastOnlineTimeUtcEQ applies the EQ predicate on the "last_online_time_utc" field.
func LastOnlineTimeUtcEQ(v int64) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastOnlineTimeUtc), v))
	})
}

// LastOnlineTimeUtcNEQ applies the NEQ predicate on the "last_online_time_utc" field.
func LastOnlineTimeUtcNEQ(v int64) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastOnlineTimeUtc), v))
	})
}

// LastOnlineTimeUtcIn applies the In predicate on the "last_online_time_utc" field.
func LastOnlineTimeUtcIn(vs ...int64) predicate.WalletNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WalletNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastOnlineTimeUtc), v...))
	})
}

// LastOnlineTimeUtcNotIn applies the NotIn predicate on the "last_online_time_utc" field.
func LastOnlineTimeUtcNotIn(vs ...int64) predicate.WalletNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WalletNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastOnlineTimeUtc), v...))
	})
}

// LastOnlineTimeUtcGT applies the GT predicate on the "last_online_time_utc" field.
func LastOnlineTimeUtcGT(v int64) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastOnlineTimeUtc), v))
	})
}

// LastOnlineTimeUtcGTE applies the GTE predicate on the "last_online_time_utc" field.
func LastOnlineTimeUtcGTE(v int64) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastOnlineTimeUtc), v))
	})
}

// LastOnlineTimeUtcLT applies the LT predicate on the "last_online_time_utc" field.
func LastOnlineTimeUtcLT(v int64) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastOnlineTimeUtc), v))
	})
}

// LastOnlineTimeUtcLTE applies the LTE predicate on the "last_online_time_utc" field.
func LastOnlineTimeUtcLTE(v int64) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastOnlineTimeUtc), v))
	})
}

// HasCoin applies the HasEdge predicate on the "coin" edge.
func HasCoin() predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CoinTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CoinTable, CoinPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCoinWith applies the HasEdge predicate on the "coin" edge with a given conditions (other predicates).
func HasCoinWith(preds ...predicate.CoinInfo) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CoinInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CoinTable, CoinPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WalletNode) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WalletNode) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WalletNode) predicate.WalletNode {
	return predicate.WalletNode(func(s *sql.Selector) {
		p(s.Not())
	})
}
