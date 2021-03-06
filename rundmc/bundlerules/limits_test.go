package bundlerules_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opencontainers/specs"

	"github.com/cloudfoundry-incubator/garden"
	"github.com/cloudfoundry-incubator/goci"
	"github.com/cloudfoundry-incubator/guardian/gardener"
	"github.com/cloudfoundry-incubator/guardian/rundmc/bundlerules"
)

var _ = Describe("LimitsRule", func() {
	It("sets the correct memory limit in bundle resources", func() {
		newBndl := bundlerules.Limits{}.Apply(goci.Bundle(), gardener.DesiredContainerSpec{
			Limits: garden.Limits{
				Memory: garden.MemoryLimits{LimitInBytes: 4096},
			},
		})

		Expect(*(newBndl.Resources().Memory.Limit)).To(BeNumerically("==", 4096))
	})

	It("does not clobber other fields of the resources sections", func() {
		foo := "foo"
		bndl := goci.Bundle().WithResources(
			&specs.Resources{
				Devices: []specs.DeviceCgroup{{Access: &foo}},
			},
		)

		newBndl := bundlerules.Limits{}.Apply(bndl, gardener.DesiredContainerSpec{
			Limits: garden.Limits{
				Memory: garden.MemoryLimits{LimitInBytes: 4096},
			},
		})

		Expect(*(newBndl.Resources().Memory.Limit)).To(BeNumerically("==", 4096))
		Expect(newBndl.Resources().Devices).To(Equal(bndl.Resources().Devices))
	})
})
