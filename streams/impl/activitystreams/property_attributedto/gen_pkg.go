package propertyattributedto

import vocab "github.com/go-fed/activity/streams/vocab"

var mgr privateManager

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializeLinkActivityStreams returns the deserialization method for
	// the "LinkInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeLinkActivityStreams() func(map[string]interface{}, map[string]string) (vocab.LinkInterface, error)
	// DeserializeObjectActivityStreams returns the deserialization method for
	// the "ObjectInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeObjectActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ObjectInterface, error)
}

// SetManager sets the manager package-global variable. For internal use only, do
// not use as part of Application behavior. Must be called at golang init time.
func SetManager(m privateManager) {
	mgr = m
}