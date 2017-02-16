/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cm

import (
	"fmt"
	"path"

	"k8s.io/kubernetes/pkg/api/v1"
)

type QOSContainerManager interface {
	Start(nodeInfo *v1.Node) error
	GetQOSContainersInfo() QOSContainersInfo
}

type qosContainerManagerImpl struct {
	NodeConfig
	// nodeInfo stores information about the node resource capacity
	nodeInfo *v1.Node
	// qosContainersInfo hold absolute paths of the top level qos containers
	qosContainersInfo QOSContainersInfo
	// Stores the mounted cgroup subsystems
	subsystems *CgroupSubsystems
	// cgroupManager is the cgroup Manager Object responsible for managing all
	// pod cgroups.
	cgroupManager CgroupManager
}

func NewQOSContainerManager(subsystems *CgroupSubsystems, nodeConfig NodeConfig) (QOSContainerManager, error) {
	if nodeConfig.CgroupsPerQOS {
		return &qosContainerManagerNoop{
			cgroupRoot: CgroupName(nodeConfig.CgroupRoot),
		}, nil
	}

	// this does default to / when enabled, but this tests against regressions.
	if nodeConfig.CgroupRoot == "" {
		return nil, fmt.Errorf("invalid configuration: cgroups-per-qos was specified and cgroup-root was not specified. To enable the QoS cgroup hierarchy you need to specify a valid cgroup-root")
	}

	// we need to check that the cgroup root actually exists for each subsystem
	// of note, we always use the cgroupfs driver when performing this check since
	// the input is provided in that format.
	// this is important because we do not want any name conversion to occur.
	cgroupManager := NewCgroupManager(subsystems, "cgroupfs")
	if !cgroupManager.Exists(CgroupName(nodeConfig.CgroupRoot)) {
		return nil, fmt.Errorf("invalid configuration: cgroup-root doesn't exist")
	}

	return &qosContainerManagerImpl{
		NodeConfig:    nodeConfig,
		subsystems:    subsystems,
		cgroupManager: NewCgroupManager(subsystems, nodeConfig.CgroupDriver),
	}, nil
}

func (m *qosContainerManagerImpl) GetQOSContainersInfo() QOSContainersInfo {
	return m.qosContainersInfo
}

func (m *qosContainerManagerImpl) Start(nodeInfo *v1.Node) error {
	// Top level for Qos containers are created only for Burstable
	// and Best Effort classes
	qosClasses := [2]v1.PodQOSClass{v1.PodQOSBurstable, v1.PodQOSBestEffort}

	// Create containers for both qos classes
	for _, qosClass := range qosClasses {
		// get the container's absolute name
		absoluteContainerName := CgroupName(path.Join(m.CgroupRoot, string(qosClass)))
		// containerConfig object stores the cgroup specifications
		containerConfig := &CgroupConfig{
			Name:               absoluteContainerName,
			ResourceParameters: &ResourceConfig{},
		}
		// check if it exists
		if !m.cgroupManager.Exists(absoluteContainerName) {
			if err := m.cgroupManager.Create(containerConfig); err != nil {
				return fmt.Errorf("failed to create top level %v QOS cgroup : %v", qosClass, err)
			}
		}
	}
	// Store the top level qos container names
	m.qosContainersInfo = QOSContainersInfo{
		Guaranteed: m.CgroupRoot,
		Burstable:  path.Join(m.CgroupRoot, string(v1.PodQOSBurstable)),
		BestEffort: path.Join(m.CgroupRoot, string(v1.PodQOSBestEffort)),
	}
	m.nodeInfo = nodeInfo
	return nil
}

type qosContainerManagerNoop struct {
	cgroupRoot CgroupName
}

var _ QOSContainerManager = &qosContainerManagerNoop{}

func (m *qosContainerManagerNoop) GetQOSContainersInfo() QOSContainersInfo {
	return QOSContainersInfo{}
}

func (m *qosContainerManagerNoop) Start(_ *v1.Node) error {
	return nil
}
