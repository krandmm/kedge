// This file was automatically generated by informer-gen

package v1

import (
	internalinterfaces "github.com/openshift/origin/pkg/user/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Groups returns a GroupInformer.
	Groups() GroupInformer
	// Identities returns a IdentityInformer.
	Identities() IdentityInformer
	// Users returns a UserInformer.
	Users() UserInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// Groups returns a GroupInformer.
func (v *version) Groups() GroupInformer {
	return &groupInformer{factory: v.SharedInformerFactory}
}

// Identities returns a IdentityInformer.
func (v *version) Identities() IdentityInformer {
	return &identityInformer{factory: v.SharedInformerFactory}
}

// Users returns a UserInformer.
func (v *version) Users() UserInformer {
	return &userInformer{factory: v.SharedInformerFactory}
}
