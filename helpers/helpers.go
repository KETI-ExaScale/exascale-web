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
	podPrefix := clientset.CoreV1().Pods("gpu")
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

// func (nm *NodeManager) GetMetric(podIP string) (*pb.MultiMetric, error) {
// 	host := podIP + ":9322"
// 	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		fmt.Println("Did not connect", err)
// 	}
// 	defer conn.Close()
// 	metricClient := pb.NewMetricCollectorClient(conn)
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	r, err := metricClient.GetMultiMetric(ctx, &pb.Request{})

// 	return r, err
// }

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

	for _, node := range nm.Nodes {
		returnStr = returnStr + nm.GetNodeList(node)
	}

	return returnStr
}

func (nm *NodeManager) GetNodeList(node *pb.MultiMetric) string {
	// node IP parser 필요
	fmt.Println("getNodeInfo", node.NodeName)
	memoryTotal := 0
	memoryUsed := 0

	for _, gpuMetric := range node.GpuMetrics {
		memoryTotal += int(gpuMetric.MemoryTotal)
		memoryUsed += int(gpuMetric.MemoryUsed)
	}

	memoryUsed = int(float64(memoryUsed) * 0.000001)
	memoryTotal = int(float64(memoryTotal) * 0.000001)
	returnStr := ``
	returnStr = returnStr + fmt.Sprintf(`
		<div class="col-md-6 col-xl-3" >
			<a class="card card-link" onclick="nodeSelectBtnclick('%s')">
				<div class="card-body">
				<div class="row">
					<div class="col-auto">
					<span class="avatar rounded">EP</span>
					</div>
					<div class="col">
					<div class="font-weight-medium">%s</div>
					<div class="text-secondary">Used Memory (MB): %s</div>
					<div class="text-secondary">Total Memory (MB): %s </div>
					</div>
				</div>
				</div>
			</a>
		</div>
		`, node.NodeName, node.NodeName, strconv.Itoa(memoryUsed), strconv.Itoa(memoryTotal))
	return returnStr
}

func (nm *NodeManager) GetNodeGPUInfo(nodeName string) string {
	returnStr := ``
	index := 0
	totlaAllocated := 0
	totalUsed := 0
	i := 0
	for _, PodIP := range nm.IPMapper {
		res, err := api.GetMultiMetric(PodIP)
		if err != nil {
			fmt.Println("Error:", err)
		}
		nm.Nodes[i] = res
		i += 1
	}

	for j, node := range nm.Nodes {
		fmt.Println("Node Name :", node.NodeName)
		if node.NodeName == nodeName {
			index = j
			break
		}
	}

	for _, GPUMetric := range nm.Nodes[index].GpuMetrics {
		totlaAllocated += int(GPUMetric.MemoryUsed)
		totalUsed += int(GPUMetric.MemoryUsed)
	}

	returnStr = returnStr + `
	<div class="card">
		<div class="card-body">
			<div class="row row-deck row-cards">
				<div class="col-lg-6 mt-1">
					<div class="row row-deck row-cards">
						<div class="col-lg-12 mb-0">
							<h1>Node Information</h1>
						</div>
						<div class="col-lg-12 mt-0 mb-0" style="display:inline;">
							<h2 class="mb-0">GPU Info</h2>
						</div>
			`
	if len(nm.Nodes[index].GpuMetrics) == 0 {
		returnStr = returnStr + `
		<div class="col-12">
			<h2>This node does not have GPU</h2>
		</div>
	`
		return returnStr
	}
	// for i := 0; i < int(nodeRes.Data.VirtualGPU); i++ {
	//	if podName, ok := nodeRes.Data.GpuPodForPrint[i]; ok {
	//		returnStr = returnStr + generateUsedGPU(podName)
	//	} else {
	//		returnStr = returnStr + generateAllocateGPU()
	// 	}
	// }
	totalMemory := 0
	for _, gpuMetric := range nm.Nodes[index].GpuMetrics {
		totalMemory += int(gpuMetric.MemoryTotal)
	}
	usedCount := ((totalUsed * 100) / (totalMemory)) * 2

	fmt.Println("usedCount :", usedCount)

	totalUsedMB := float32(totalUsed) * 0.000001
	totalMemoryMB := float32(totalMemory) * 0.000001

	for i := 0; i < 20; i++ {
		if i < usedCount {
			returnStr += nm.generateUsedGPU()
		} else {
			returnStr = returnStr + nm.generateAllocateGPU()
		}
	}
	returnStr = returnStr + fmt.Sprintf(`
					</div>
				</div>
				<div class="col-lg-6 mt-1">
					<div class="row row-deck row-cards">
						<div class="card">
						<div class="ribbon bg-red custom-pd-1">%s</div>
							<div class="card-body custom-pd-3">
								<h3 class="card-title mb-1">Summary</h3>
								<p class="text-secondary mb-0">GPU(Used/total) : %s/%s (MB)</p>
							</div>
						</div>
	`, nodeName, strconv.Itoa(int(totalUsedMB)), strconv.Itoa(int(totalMemoryMB)))
	return returnStr
}

func (nm *NodeManager) generateUsedGPU() string {
	return `
		<div class="col-2 div-with-background">
			<img src="/static/img/gpuUsed.png">
		</div>
	`
}

func (nm *NodeManager) generateAllocateGPU() string {
	return `
		<div class="col-2 div-with-background">
			<img src="/static/img/gpuallocate.png">
		</div>
	`
}

func (nm *NodeManager) GetNodeMetricInfo(nodeName string, returnStr string) string {
	index := 0
	for i, node := range nm.Nodes {
		if node.NodeName == nodeName {
			//
			index = i
			break
		}
	}

	totlaAllocated := 0

	for _, GPUMetric := range nm.Nodes[index].GpuMetrics {
		totlaAllocated += int(GPUMetric.MemoryUsed)
	}

	for _, GPUMetric := range nm.Nodes[index].GpuMetrics {
		//
		MemoryUsedMB := float32(GPUMetric.MemoryUsed) * 0.000001
		CPUNodeUseMilli := float32(nm.Nodes[index].NodeMetric.MilliCpuUsage)
		MemoryNodeUseMB := float32(nm.Nodes[index].NodeMetric.MemoryUsage) * 0.000001
		StorageUsedMB := float32(nm.Nodes[index].NodeMetric.StorageUsage) / 1000000000
		returnStr = returnStr + fmt.Sprintf(`
							<div class="card mt-1">
								<div class="card-body">
								<h3 class="card-title mb-0">Metrics</h3>
								<table class="table table-sm table-borderless">
									<thead>
									<tr>
										<th>Usage</th>
										<th></th>
									</tr>
									</thead>
									<tbody>
									<tr>
										<td style="width:50%%">
										<div class="progressbg">
											<div class="progress progressbg-progress">
											<div class="progress-bar bg-primary-lt" style="width:100%%" role="progressbar" aria-valuenow="100" aria-valuemin="0" aria-valuemax="100">
											</div>
											</div>
											<div class="progressbg-text">GPUCore</div>
										</div>
										</td>
										<td class="w-1 fw-bold text-end">%s</td>
									</tr>
									<tr>
										<td style="width:50%%">
										<div class="progressbg">
											<div class="progress progressbg-progress">
											<div class="progress-bar bg-primary-lt" style="width:100%%" role="progressbar" aria-valuenow="100" aria-valuemin="0" aria-valuemax="100">
											</div>
											</div>
											<div class="progressbg-text">GPUMemory</div>
										</div>
										</td>
										<td class="w-1 fw-bold text-end">%s MB</td>
									</tr>
									<tr>
										<td style="width:50%%">
										<div class="progressbg">
											<div class="progress progressbg-progress">
											<div class="progress-bar bg-primary-lt" style="width:100%%" role="progressbar" aria-valuenow="100" aria-valuemin="0" aria-valuemax="100">
											</div>
											</div>
											<div class="progressbg-text">GPUPower</div>
										</div>
										</td>
										<td class="w-1 fw-bold text-end">%s </td>
									</tr>
									<tr>
										<td style="width:50%%">
										<div class="progressbg">
											<div class="progress progressbg-progress">
											<div class="progress-bar bg-primary-lt" style="width:100%%" role="progressbar" aria-valuenow="100" aria-valuemin="0" aria-valuemax="100">
											</div>
											</div>
											<div class="progressbg-text">CPU Milli</div>
										</div>
										</td>
										<td class="w-1 fw-bold text-end">%s</td>
									</tr>
									<tr>
										<td style="width:50%%">
										<div class="progressbg">
											<div class="progress progressbg-progress">
											<div class="progress-bar bg-primary-lt" style="width:100%%" role="progressbar" aria-valuenow="100" aria-valuemin="0" aria-valuemax="100">
											</div>
											</div>
											<div class="progressbg-text">Memory</div>
										</div>
										</td>
										<td class="w-1 fw-bold text-end">%s MB</td>
									</tr>
									<tr>
										<td style="width:50%%">
										<div class="progressbg">
											<div class="progress progressbg-progress">
											<div class="progress-bar bg-primary-lt" style="width:100%%" role="progressbar" aria-valuenow="100" aria-valuemin="0" aria-valuemax="100">
											</div>
											</div>
											<div class="progressbg-text">Storage</div>
										</div>
										</td>
										<td class="w-1 fw-bold text-end">%s GB</td>
									</tr>
									</tbody>
								</table>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
		`, strconv.Itoa(int(GPUMetric.Cudacore)),
			strconv.Itoa(int(MemoryUsedMB)),
			strconv.Itoa(int(GPUMetric.PowerUsed)),
			strconv.Itoa(int(CPUNodeUseMilli)),
			strconv.Itoa(int(MemoryNodeUseMB)),
			strconv.Itoa(int(StorageUsedMB)))
	}

	return returnStr
}

func (nm *NodeManager) GetPodInfo(nodeName string) string {
	index := 0
	i := 0
	for _, PodIP := range nm.IPMapper {
		res, err := api.GetMultiMetric(PodIP)
		if err != nil {
			fmt.Println("Error:", err)
		}
		nm.Nodes[i] = res
		i += 1
	}

	for j, node := range nm.Nodes {
		if node.NodeName == nodeName {
			index = j
			break
		}
	}

	for i, node := range nm.Nodes {
		if nodeName == node.NodeName {
			index = i
			break
		}
	}

	returnStr := `
	<div class="card">
		<div class="card-header">
			<h3 class="card-title">Pod Information</h3>
		</div>
		<div class="table-responsive">
			<table class="table card-table table-vcenter text-nowrap datatable" id="podInfoTable">
				<thead>
				<tr>
					<th>Pod Name</th>
					<th>CPU Used (Milli Core)</th>
					<th>Memory Used (MB)</th>
					<th>Storage Used (GB)</th>
					<th>GPU Memory Used (MB)</th>
				</tr>
				</thead>
				<tbody id="podInfoTableBody">`
	for podName, pod := range nm.Nodes[index].PodMetrics {
		GPUMemoryUsage := 0

		for _, podGPU := range pod.PodGpuMetrics {
			fmt.Println("Sub GPU Usage:", podGPU.GpuMemoryUsed)
			GPUMemoryUsage += int(podGPU.GpuMemoryUsed)
		}

		fmt.Println("Pod GPU Usage : ", GPUMemoryUsage)
		returnStr = returnStr + fmt.Sprintf(`
					<tr>
						<td>%s</td>
						<td>%s</td>
						<td>%s</td>
						<td>%s</td>
						<td>%s</td>
					</tr>
		`, podName, strconv.Itoa(int(pod.CpuUsage)), strconv.Itoa(int(pod.MemoryUsage/1000000)),
			strconv.Itoa(int(pod.StorageUsage/1000000000)), strconv.Itoa(int(GPUMemoryUsage/1000000)))
	}

	returnStr = returnStr + `
				</tbody>
			</table>
		</div>	
	</div>`
	return returnStr
}

// func podTr(pod corev1.Pod) (string, string, string) {
// 	nameTdVal := fmt.Sprintf(`<td class="u-table-cell u-text-custom-color-11">%s</td>`, pod.Name)
// 	namespaceTdVal := fmt.Sprintf(`<td class="u-table-cell u-text-custom-color-1">%s</td>`, pod.Namespace)
// 	age := time.Since(pod.CreationTimestamp.Time)
// 	totalSec := int(math.Round(age.Seconds()))
// 	day := totalSec / 86400
// 	hour := (totalSec % 86400) / 3600
// 	minute := ((totalSec % 86400) % 3600) / 60
// 	sec := ((totalSec % 86400) % 3600) % 60
// 	ageStr := ""
// 	if day > 0 {
// 		ageStr = fmt.Sprintf("%dd %dh", day, hour)
// 	} else if hour > 0 {
// 		ageStr = fmt.Sprintf("%dh %dm", hour, minute)
// 	} else if minute > 0 {
// 		ageStr = fmt.Sprintf("%dm %ds", minute, sec)
// 	} else {
// 		ageStr = fmt.Sprintf("%ds", sec)
// 	}
// 	ageTdVal := fmt.Sprintf(`<td class="u-table-cell u-text-custom-color-1">%s</td>`, ageStr)
// 	return nameTdVal, namespaceTdVal, ageTdVal
// }

// func containerTr(pod corev1.Pod) ([]string, []string) {
// 	containers := pod.Spec.Containers
// 	nameTdVals := make([]string, len(containers))
// 	imageTdVals := make([]string, len(containers))
// 	for i, container := range containers {
// 		nameTdVals[i] = fmt.Sprintf(`<td class="u-table-cell u-text-custom-color-1">%s</td>`, container.Name)
// 		imageTdVals[i] = fmt.Sprintf(`<td class="u-table-cell u-text-custom-color-1">%s</td>`, container.Image)
// 	}
// 	return nameTdVals, imageTdVals
// }

// func gpuTr(pod corev1.Pod) string {
// 	gpuTrVal := ""
// 	if pod.Namespace == "kube-system" || pod.Namespace == "keti-system" {
// 		gpuTrVal = `<td class="u-table-cell u-text-custom-color-12">False</td>`

// 	} else {
// 		gpuTrVal = `<td class="u-table-cell u-text-custom-color-13">True</td>`
// 	}
// 	// if pod.Labels["gpu"] == "true" {
// 	// 	gpuTrVal = `<td class="u-table-cell u-text-custom-color-13">True</td>`
// 	// } else {
// 	// 	gpuTrVal = `<td class="u-table-cell u-text-custom-color-12">False</td>`
// 	// }
// 	return gpuTrVal
// }
