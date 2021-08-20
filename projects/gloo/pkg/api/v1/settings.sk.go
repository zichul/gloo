// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"log"
	"os"
	"sort"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewSettings(namespace, name string) *Settings {
	settings := &Settings{}
	settings.SetMetadata(&core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return settings
}

func (r *Settings) SetMetadata(meta *core.Metadata) {
	r.Metadata = meta
}

// Deprecated
func (r *Settings) SetStatus(status *core.Status) {
	r.SetStatusForNamespace(status)
}

// Deprecated
func (r *Settings) GetStatus() *core.Status {
	if r != nil {
		s, _ := r.GetStatusForNamespace()
		return s
	}
	return nil
}

func (r *Settings) SetNamespacedStatuses(statuses *core.NamespacedStatuses) {
	r.NamespacedStatuses = statuses
}

// SetStatusForNamespace inserts the specified status into the NamespacedStatuses.Statuses map for
// the current namespace (as specified by POD_NAMESPACE env var).  If the resource does not yet
// have a NamespacedStatuses, one will be created.
// Note: POD_NAMESPACE environment variable must be set for this function to behave as expected.
// If unset, a podNamespaceErr is returned.
func (r *Settings) SetStatusForNamespace(status *core.Status) error {
	podNamespace := os.Getenv("POD_NAMESPACE")
	if podNamespace == "" {
		return errors.NewPodNamespaceErr()
	}
	if r.GetNamespacedStatuses() == nil {
		r.SetNamespacedStatuses(&core.NamespacedStatuses{})
	}
	if r.GetNamespacedStatuses().GetStatuses() == nil {
		r.GetNamespacedStatuses().Statuses = make(map[string]*core.Status)
	}
	r.GetNamespacedStatuses().GetStatuses()[podNamespace] = status
	return nil
}

// GetStatusForNamespace returns the status stored in the NamespacedStatuses.Statuses map for the
// controller specified by the POD_NAMESPACE env var, or nil if no status exists for that
// controller.
// Note: POD_NAMESPACE environment variable must be set for this function to behave as expected.
// If unset, a podNamespaceErr is returned.
func (r *Settings) GetStatusForNamespace() (*core.Status, error) {
	podNamespace := os.Getenv("POD_NAMESPACE")
	if podNamespace == "" {
		return nil, errors.NewPodNamespaceErr()
	}
	if r.GetNamespacedStatuses() == nil {
		return nil, nil
	}
	if r.GetNamespacedStatuses().GetStatuses() == nil {
		return nil, nil
	}
	return r.GetNamespacedStatuses().GetStatuses()[podNamespace], nil
}

func (r *Settings) MustHash() uint64 {
	hashVal, err := r.Hash(nil)
	if err != nil {
		log.Panicf("error while hashing: (%s) this should never happen", err)
	}
	return hashVal
}

func (r *Settings) GroupVersionKind() schema.GroupVersionKind {
	return SettingsGVK
}

type SettingsList []*Settings

func (list SettingsList) Find(namespace, name string) (*Settings, error) {
	for _, settings := range list {
		if settings.GetMetadata().Name == name && settings.GetMetadata().Namespace == namespace {
			return settings, nil
		}
	}
	return nil, errors.Errorf("list did not find settings %v.%v", namespace, name)
}

func (list SettingsList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, settings := range list {
		ress = append(ress, settings)
	}
	return ress
}

func (list SettingsList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, settings := range list {
		ress = append(ress, settings)
	}
	return ress
}

func (list SettingsList) Names() []string {
	var names []string
	for _, settings := range list {
		names = append(names, settings.GetMetadata().Name)
	}
	return names
}

func (list SettingsList) NamespacesDotNames() []string {
	var names []string
	for _, settings := range list {
		names = append(names, settings.GetMetadata().Namespace+"."+settings.GetMetadata().Name)
	}
	return names
}

func (list SettingsList) Sort() SettingsList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list SettingsList) Clone() SettingsList {
	var settingsList SettingsList
	for _, settings := range list {
		settingsList = append(settingsList, resources.Clone(settings).(*Settings))
	}
	return settingsList
}

func (list SettingsList) Each(f func(element *Settings)) {
	for _, settings := range list {
		f(settings)
	}
}

func (list SettingsList) EachResource(f func(element resources.Resource)) {
	for _, settings := range list {
		f(settings)
	}
}

func (list SettingsList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Settings) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

// Kubernetes Adapter for Settings

func (o *Settings) GetObjectKind() schema.ObjectKind {
	t := SettingsCrd.TypeMeta()
	return &t
}

func (o *Settings) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Settings)
}

func (o *Settings) DeepCopyInto(out *Settings) {
	clone := resources.Clone(o).(*Settings)
	*out = *clone
}

var (
	SettingsCrd = crd.NewCrd(
		"settings",
		SettingsGVK.Group,
		SettingsGVK.Version,
		SettingsGVK.Kind,
		"st",
		false,
		&Settings{})
)

var (
	SettingsGVK = schema.GroupVersionKind{
		Version: "v1",
		Group:   "gloo.solo.io",
		Kind:    "Settings",
	}
)
