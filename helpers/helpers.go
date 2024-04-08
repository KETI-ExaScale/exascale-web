package helpers

import (
	"context"
	"fmt"
	"strconv"

	"gin_session_auth/pkg/api/api"
	pb "gin_session_auth/pkg/api/api/metric"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const PORT = "9322"

func (nm *NodeManager) GetClusterList() string {
	returnStr := ``
	returnStr = returnStr + fmt.Sprintf(`
	<label class="form-check">
					<input type="Radio" class="form-check-input" name="form-type" value="%s">
					<span class="form-check-label">%s</span>
					</label>
	`, "Cluster1", "Cluster1")
	return returnStr
}

type NodeManager struct {
	Nodes    []*pb.MultiMetric
	IPMapper map[string]string
}

func NewNodeManager() *NodeManager {
	config, _ := rest.InClusterConfig()
	clientset, _ := kubernetes.NewForConfig(config)
	podPrefix := clientset.CoreV1().Pods("")
	labelMap := make(map[string]string)
	labelMap["name"] = "gpu-metric-collector"

	options := metav1.ListOptions{
		LabelSelector: labels.SelectorFromSet(labelMap).String(),
	}
	metricPods, err := podPrefix.List(context.Background(), options)

	if err != nil {
		fmt.Println(err)
	}

	podIPMap := make(map[string]string)

	for _, pod := range metricPods.Items {
		podIPMap[pod.Spec.NodeName] = pod.Status.PodIP
	}

	NodeInformation := make([]*pb.MultiMetric, len(podIPMap))

	return &NodeManager{
		Nodes:    NodeInformation,
		IPMapper: podIPMap,
	}
}

func (nm *NodeManager) GetClusterInfo(nodeName string) string {
	// gRPC 요청
	i := 0
	for _, PodIP := range nm.IPMapper {
		res, err := api.GetMultiMetric(PodIP)
		if err != nil {
			fmt.Println("Error:", err)
		}
		nm.Nodes[i] = res
		i += 1
	}

	total_gpu := 0
	for _, node := range nm.Nodes {
		total_gpu += int(node.GpuCount)
	}

	returnStr := ``
	returnStr = returnStr + fmt.Sprintf(`
		<div class="col-lg-4">
			<div class="card card-sm">
				<div class="card-body">
				<div class="row align-items-center">
					<div class="col-auto">
					<span class="bg-red text-white avatar"><!-- Download SVG icon from http://tabler-icons.io/i/heart -->
						<svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M19.5 12.572l-7.5 7.428l-7.5 -7.428a5 5 0 1 1 7.5 -6.566a5 5 0 1 1 7.5 6.572"></path></svg>
					</span>
					</div>
					<div class="col">
						<div class="text-secondary">
							ClusterName
						</div>
						<div class="font-weight-medium">
							%s
						</div>					
					</div>
					<div class="col-auto">
					</div>
				</div>
				</div>
			</div>
		</div>
		<div class="col-lg-4">
			<div class="card card-sm">
				<div class="card-body">
				<div class="row align-items-center">
					<div class="col-auto">
					<span class="bg-red text-white avatar"><!-- Download SVG icon from http://tabler-icons.io/i/heart -->
						<svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M19.5 12.572l-7.5 7.428l-7.5 -7.428a5 5 0 1 1 7.5 -6.566a5 5 0 1 1 7.5 6.572"></path></svg>
					</span>
					</div>
					<div class="col">
					<div class="text-secondary">
					Nodes
					</div>
					<div class="font-weight-medium">
					%s
					</div>					
					</div>
					<div class="col-auto">
					</div>
				</div>
				</div>
			</div>
		</div>
		<div class="col-lg-4">
			<div class="card card-sm">
				<div class="card-body">
				<div class="row align-items-center">
					<div class="col-auto">
					<span class="bg-red text-white avatar"><!-- Download SVG icon from http://tabler-icons.io/i/heart -->
						<svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M19.5 12.572l-7.5 7.428l-7.5 -7.428a5 5 0 1 1 7.5 -6.566a5 5 0 1 1 7.5 6.572"></path></svg>
					</span>
					</div>
					<div class="col">					
					<div class="text-secondary">
					GPUS
					</div>
					<div class="font-weight-medium">
					%s
					</div>
					</div>
					<div class="col-auto">
					</div>
				</div>
				</div>
			</div>
		</div>
		`, "Cluster1", strconv.Itoa(len(nm.Nodes)), strconv.Itoa(total_gpu))

	usedMemoryTotal := 0
	capacityMemoryTotal := 0

	for _, node := range nm.Nodes {
		returnStr = returnStr + nm.GetNodeList(node)

		for _, gpumetric := range node.GpuMetrics {
			usedMemoryTotal += int(gpumetric.MemoryUsed)
			capacityMemoryTotal += int(gpumetric.MemoryTotal)
		}

	}

	memoryUsagePercent := (usedMemoryTotal * 100 / capacityMemoryTotal)
	fmt.Println("Used Percent:", memoryUsagePercent)

	returnStr = returnStr + fmt.Sprintf(`
	<div class="card card-sm">
		<div class="col-md-6 col-xl-3">
			<div class="card-body">
				<div class="row">
					<div class="col-auto">
						<span class="avartar rounded">T</span>
					</div>
					<div class="col">
						<div class="font-weight-medium">Total GPU Memory</div>
						<div class="text-secondary">Used Memory (percent): %s</div>
					</div>
				</div>
			</div>
		</div>
	</div>
	`, strconv.Itoa(memoryUsagePercent))

	returnStr = returnStr + `
	<div class="card card-sm" style="padding:1rem;">
		<h1 class="card-title mb-1">Node GPU Information</h1>
		<div class="row">
	`

	for _, node := range nm.Nodes {
		returnStr += nm.GetNodeGPU(node)
	}

	returnStr = returnStr + `
			</div>
		</div>
	</div>
	`

	return returnStr
}

func (nm *NodeManager) GetNodeGPU(node *pb.MultiMetric) string {
	totalCapacity := 0
	totalUsed := 0
	returnStr := ``

	returnStr = returnStr + fmt.Sprintf(`
		<div class="col-lg-3 mt-1 custom-pd-3">
			<h3>%s</h3>
			<div class="row row-deck row-cards">
	`, node.NodeName)

	for _, GPUMetric := range node.GpuMetrics {
		totalCapacity += int(GPUMetric.MemoryTotal)
		totalUsed += int(GPUMetric.MemoryUsed)
	}

	if totalCapacity == 0 {
		return ``
	}

	cal_value := (totalUsed * 100) / totalCapacity

	fmt.Println("cal value:", cal_value)

	for i := 0; i < 10; i++ {
		if i < (cal_value / 10) {
			returnStr += nm.generateUsedGPU()
		} else {
			returnStr += nm.generateAllocateGPU()
		}
	}

	returnStr += `</div></div>`

	return returnStr
}

// show node information & pod information
func (nm *NodeManager) GetNodeList(node *pb.MultiMetric) string {
	memoryTotal := 0
	memoryUsed := 0

	for _, gpuMetric := range node.GpuMetrics {
		memoryTotal += int(gpuMetric.MemoryTotal)
		memoryUsed += int(gpuMetric.MemoryUsed)
	}

	memoryUsed = int(float64(memoryUsed) * 0.000001)
	memoryTotal = int(float64(memoryTotal) * 0.000001)

	// get graph
	returnStr := ``
	returnStr = returnStr + fmt.Sprintf(`
		<div class="col-md-6 col-xl-3" >
			<div class="card card-link">
				<div class="card-body">
				<div class="row">
					<div class="col-auto">
					<span class="avatar rounded">N</span>
					</div>
					<div class="col">
					<div class="font-weight-medium">Node : %s</div>
					<div class="text-secondary">Used Memory (MB): %s</div>
					<div class="text-secondary">Total Memory (MB): %s </div>
					</div>
				</div>
				</div>
			</div>
		</div>
		`, node.NodeName, strconv.Itoa(memoryUsed), strconv.Itoa(memoryTotal))

	return returnStr
}

func (nm *NodeManager) generateUsedGPU() string {
	return `
		<div class="col-2" style="max-width: 3rem; max-height: 3rem; padding: 0.2rem;">
			<img src="/static/img/useLegend.png">
		</div>
	`
}

func (nm *NodeManager) generateAllocateGPU() string {
	return `
		<div class="col-2" style="max-width: 3rem; max-height: 3rem; padding: 0.2rem;">
			<img src="/static/img/allocateLegend.png">
		</div>
	`
}
