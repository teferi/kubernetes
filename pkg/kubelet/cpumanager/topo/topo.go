/*
Copyright 2017 The Kubernetes Authors.

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

package topo

import (
	"sort"
)

// TODO
type CPUTopology struct {
	NumCPUs           int
	Hyperthreading    bool
	NumNodes           int
	CPUtopoDetails    map[int]CPUInfo
}

type CPUInfo struct {
	NodeId int
	CoreId int
}


func (ct *CPUTopology) GetCPUsfromPhysicalCore(availableCPUs []int) (selectedCPUs []int) {


	return []int{0}
}


func (ct *CPUTopology) GetAnyCPU(availableCPUs []int) (selectedCPU int) {


	return 0
}



//Helpers for sorting
type mapEntry struct {
	key int
	val int
}
type mapAsSlice []mapEntry
func (n mapAsSlice) Len() int { return len(n) }
func (n mapAsSlice) Less(i, j int) bool { return n[i].val < n[j].val }
func (n mapAsSlice) Swap(i, j int) { n[i], n[j] = n[j], n[i] }



//return sorted slice with nodeIds with least utilization
func (ct CPUTopology) GetNodeUtilization(availableCPUs []int) (NodeId []int) {

	if ct.NumNodes == 1 {
		return []int{0}
	}
	freeCPUsPerNode := make(map[int]int,ct.NumNodes)
	for i := 0; i < ct.NumNodes; i++ {
		freeCPUsPerNode[i] = 0
	}

	for _, cpu := range availableCPUs {
		freeCPUsPerNode[ct.CPUtopoDetails[cpu].NodeId]++
	}

	freeCPUsPerNodeSorted := make(mapAsSlice,len(freeCPUsPerNode))

	for k, v := range freeCPUsPerNode {
		freeCPUsPerNodeSorted = append(freeCPUsPerNodeSorted, mapEntry{
			key: k,
			val: v,
		})
	}

	sort.Sort(freeCPUsPerNodeSorted)

	res := []int{}
	for _, item := range freeCPUsPerNodeSorted {
		res = append(res,item.key)
	}
	return res
}





