package propertyupdated

import (
	"fmt"
	datetime "github.com/go-fed/activity/streams/values/dateTime"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
	"time"
)

// ActivityStreamsUpdatedProperty is the functional property "updated". It is
// permitted to be a single default-valued value type.
type ActivityStreamsUpdatedProperty struct {
	xmlschemaDateTimeMember time.Time
	hasDateTimeMember       bool
	unknown                 interface{}
	iri                     *url.URL
	alias                   string
}

// DeserializeUpdatedProperty creates a "updated" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeUpdatedProperty(m map[string]interface{}, aliasMap map[string]string) (*ActivityStreamsUpdatedProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/ns/activitystreams"]; ok {
		alias = a
	}
	propName := "updated"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "updated")
	}
	i, ok := m[propName]

	if ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
			if err == nil && len(u.Scheme) > 0 {
				this := &ActivityStreamsUpdatedProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if v, err := datetime.DeserializeDateTime(i); err == nil {
			this := &ActivityStreamsUpdatedProperty{
				alias:                   alias,
				hasDateTimeMember:       true,
				xmlschemaDateTimeMember: v,
			}
			return this, nil
		}
		this := &ActivityStreamsUpdatedProperty{
			alias:   alias,
			unknown: i,
		}
		return this, nil
	}
	return nil, nil
}

// NewActivityStreamsUpdatedProperty creates a new updated property.
func NewActivityStreamsUpdatedProperty() *ActivityStreamsUpdatedProperty {
	return &ActivityStreamsUpdatedProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling IsXMLSchemaDateTime
// afterwards will return false.
func (this *ActivityStreamsUpdatedProperty) Clear() {
	this.unknown = nil
	this.iri = nil
	this.hasDateTimeMember = false
}

// Get returns the value of this property. When IsXMLSchemaDateTime returns false,
// Get will return any arbitrary value.
func (this ActivityStreamsUpdatedProperty) Get() time.Time {
	return this.xmlschemaDateTimeMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this ActivityStreamsUpdatedProperty) GetIRI() *url.URL {
	return this.iri
}

// HasAny returns true if the value or IRI is set.
func (this ActivityStreamsUpdatedProperty) HasAny() bool {
	return this.IsXMLSchemaDateTime() || this.iri != nil
}

// IsIRI returns true if this property is an IRI.
func (this ActivityStreamsUpdatedProperty) IsIRI() bool {
	return this.iri != nil
}

// IsXMLSchemaDateTime returns true if this property is set and not an IRI.
func (this ActivityStreamsUpdatedProperty) IsXMLSchemaDateTime() bool {
	return this.hasDateTimeMember
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsUpdatedProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/ns/activitystreams": this.alias}
	var child map[string]string

	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this ActivityStreamsUpdatedProperty) KindIndex() int {
	if this.IsXMLSchemaDateTime() {
		return 0
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this ActivityStreamsUpdatedProperty) LessThan(o vocab.ActivityStreamsUpdatedProperty) bool {
	// LessThan comparison for if either or both are IRIs.
	if this.IsIRI() && o.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	} else if this.IsIRI() {
		// IRIs are always less than other values, none, or unknowns
		return true
	} else if o.IsIRI() {
		// This other, none, or unknown value is always greater than IRIs
		return false
	}
	// LessThan comparison for the single value or unknown value.
	if !this.IsXMLSchemaDateTime() && !o.IsXMLSchemaDateTime() {
		// Both are unknowns.
		return false
	} else if this.IsXMLSchemaDateTime() && !o.IsXMLSchemaDateTime() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsXMLSchemaDateTime() && o.IsXMLSchemaDateTime() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return datetime.LessDateTime(this.Get(), o.Get())
	}
}

// Name returns the name of this property: "updated".
func (this ActivityStreamsUpdatedProperty) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "updated"
	} else {
		return "updated"
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsUpdatedProperty) Serialize() (interface{}, error) {
	if this.IsXMLSchemaDateTime() {
		return datetime.SerializeDateTime(this.Get())
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// Set sets the value of this property. Calling IsXMLSchemaDateTime afterwards
// will return true.
func (this *ActivityStreamsUpdatedProperty) Set(v time.Time) {
	this.Clear()
	this.xmlschemaDateTimeMember = v
	this.hasDateTimeMember = true
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *ActivityStreamsUpdatedProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}
