package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
)

const MaterClusterHost = "http://10.0.5.20:30850/"

type ClusterInfo struct {
	Nodes string
	GPUs  string
}
type ClusterListResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		ClusterName string   `json:"clusterName"`
		MasterNode  string   `json:"masterNode"`
		Nodes       []string `json:"nodes"`
		TotalGPU    int32    `json:"totalGPU"`
	} `json:"data"`
}

func GetClusterList() string {
	clusterRes := &ClusterListResponse{}

	response, err := http.Get(MaterClusterHost + "clusters")
	if err != nil {
		klog.Errorln(err)
	}
	defer response.Body.Close()

	// 응답 본문 읽기
	body, err := io.ReadAll(response.Body)
	if err != nil {
		klog.Errorln(err)
	}

	err = json.Unmarshal(body, clusterRes)
	if err != nil {
		klog.Errorln(err)
	}
	returnStr := ``
	for _, data := range clusterRes.Data {
		returnStr = returnStr + fmt.Sprintf(`
		<label class="form-check">
                        <input type="Radio" class="form-check-input" name="form-type" value="%s">
                        <span class="form-check-label">%s</span>
                      </label>
		`, data.MasterNode, data.MasterNode)
	}
	return returnStr
}

type ClusterInfoResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ClusterName string   `json:"clusterName"`
		MasterNode  string   `json:"masterNode"`
		Nodes       []string `json:"nodes"`
		TotalGPU    int32    `json:"totalGPU"`
	} `json:"data"`
}

func GetClusterInfo(clusterName string) string {
	clusterRes := &ClusterInfoResponse{}

	response, err := http.Get(MaterClusterHost + "cluster/" + clusterName)
	if err != nil {
		klog.Errorln(err)
	}
	defer response.Body.Close()

	// 응답 본문 읽기
	body, err := io.ReadAll(response.Body)
	if err != nil {
		klog.Errorln(err)
	}

	err = json.Unmarshal(body, clusterRes)
	if err != nil {
		klog.Errorln(err)
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
					VirtualGPUS
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
		`, clusterRes.Data.ClusterName, strconv.Itoa(len(clusterRes.Data.Nodes)), strconv.Itoa(int(clusterRes.Data.TotalGPU*20)))
	returnStr = returnStr + GetNodeList(clusterRes.Data.Nodes)
	return returnStr
}

type NodeInfoResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ClusterName    string         `json:"clusterName"`
		VirtualGPU     int32          `json:"virtualGPU"`
		Age            string         `json:"age"`
		GpuPods        map[string]int `json:"gpuPods"`
		GpuPodForPrint map[int]string `json:"gpuPodForPrint"`
	} `json:"data"`
}

func GetNodeList(nodeNames []string) string {
	fmt.Println("getNodeInfo", nodeNames)
	returnStr := ``
	for _, node := range nodeNames {
		nodeRes := &NodeInfoResponse{}

		response, err := http.Get(MaterClusterHost + "node/" + node)
		if err != nil {
			klog.Errorln(err)
		}
		defer response.Body.Close()

		// 응답 본문 읽기
		body, err := io.ReadAll(response.Body)
		if err != nil {
			klog.Errorln(err)
		}

		err = json.Unmarshal(body, nodeRes)
		if err != nil {
			klog.Errorln(err)
		}

		totalAllocated := int(nodeRes.Data.VirtualGPU)

		totalUsed := len(nodeRes.Data.GpuPodForPrint)

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
						<div class="text-secondary">%s Allocate GPUs</div>
						<div class="text-secondary">%s Used GPUs</div>
						</div>
					</div>
					</div>
				</a>
			</div>
			`, node, node, strconv.Itoa(totalAllocated), strconv.Itoa(totalUsed))
	}
	return returnStr
}

func GetNodeInfo(nodeName string) string {
	returnStr := ``
	nodeRes := &NodeInfoResponse{}

	response, err := http.Get(MaterClusterHost + "node/" + nodeName)
	if err != nil {
		klog.Errorln(err)
	}
	defer response.Body.Close()

	// 응답 본문 읽기
	body, err := io.ReadAll(response.Body)
	if err != nil {
		klog.Errorln(err)
	}

	err = json.Unmarshal(body, nodeRes)
	if err != nil {
		klog.Errorln(err)
	}

	totalAllocated := int(nodeRes.Data.VirtualGPU)

	totalUsed := len(nodeRes.Data.GpuPodForPrint)

	returnStr = returnStr + `
	<div class="card">
		<div class="card-body">
			<div class="row row-deck row-cards">
				<div class="col-lg-6 mt-1">
					<div class="row row-deck row-cards">
						<div class="col-lg-12 mb-0">
							<h1>Node Information</h1>
						</div>
						<div class="col-lg-12 mt-0 mb-0">
							<h2 class="mb-0">GPU Info</h2>
							<div class="div-with-background col-7"> <img src="/static/img/legend.png"/> </div>
							
						</div>
			`
	if nodeRes.Data.VirtualGPU == 0 {
		returnStr = returnStr + `
		<div class="col-12">
			<h2>This Node Has No GPU</h2>
		</div>
	`
	}
	for i := 0; i < int(nodeRes.Data.VirtualGPU); i++ {
		if podName, ok := nodeRes.Data.GpuPodForPrint[i]; ok {
			returnStr = returnStr + generateUsedGPU(podName)
		} else {
			returnStr = returnStr + generateAllocateGPU()
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
								<p class="text-secondary mb-0">GPU(Used/Allocate) : %s/%s</p>
								<p class="text-secondary mb-0">Age : %s</p>
							</div>
						</div>
	`, nodeName, strconv.Itoa(totalUsed), strconv.Itoa(totalAllocated), nodeRes.Data.Age)
	return returnStr
}

func generateUsedGPU(podName string) string {
	return `
		<div class="col-2 div-with-background">
			<img src="/static/img/gpuUsed.png">
		</div>
	`
}

func generateAllocateGPU() string {
	return `
		<div class="col-2 div-with-background">
			<img src="/static/img/gpuallocate.png">
		</div>
	`
}

type NodeMetricResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		GPUCore   string `json:"gpuCore,omitempty"`
		GPUMemory string `json:"gpuMemory,omitempty"`
		GPUPower  string `json:"gpuPower,omitempty"`
		CPUCore   string `json:"cpuCore,omitempty"`
		Memory    string `json:"memory,omitempty"`
		Storage   string `json:"storage,omitempty"`
		NetworkRX string `json:"networkRX,omitempty"`
		NetworkTX string `json:"networkTX,omitempty"`
	} `json:"data"`
}

func GetNodeMetricInfo(nodeName string, returnStr string) string {
	nodeRes := &NodeMetricResponse{}

	response, err := http.Get(MaterClusterHost + "node/" + nodeName + "/metrics")
	if err != nil {
		klog.Errorln(err)
	}
	defer response.Body.Close()

	// 응답 본문 읽기
	body, err := io.ReadAll(response.Body)
	if err != nil {
		klog.Errorln(err)
	}

	err = json.Unmarshal(body, nodeRes)
	if err != nil {
		klog.Errorln(err)
	}
	returnStr = returnStr + fmt.Sprintf(`
						<div class="card mt-1">
							<div class="card-body">
							<h3 class="card-title mb-0">Metrics</h3>
							<table class="table table-sm table-borderless">
								<thead>
								<tr>
									<th>Utilization</th>
									<th></th>
								</tr>
								</thead>
								<tbody>
								<tr>
									<td>
									<div class="progressbg">
										<div class="progress progressbg-progress">
										<div class="progress-bar bg-primary-lt" style="width:%s%%" role="progressbar" aria-valuenow="%s" aria-valuemin="0" aria-valuemax="100">
										</div>
										</div>
										<div class="progressbg-text">GPUCore</div>
									</div>
									</td>
									<td class="w-1 fw-bold text-end">%s%%</td>
								</tr>
								<tr>
									<td>
									<div class="progressbg">
										<div class="progress progressbg-progress">
										<div class="progress-bar bg-primary-lt" style="width:%s%%" role="progressbar" aria-valuenow="%s" aria-valuemin="0" aria-valuemax="100">
										</div>
										</div>
										<div class="progressbg-text">GPUMemory</div>
									</div>
									</td>
									<td class="w-1 fw-bold text-end">%s%%</td>
								</tr>
								<tr>
									<td>
									<div class="progressbg">
										<div class="progress progressbg-progress">
										<div class="progress-bar bg-primary-lt" style="width:0%%" role="progressbar" aria-valuenow="0" aria-valuemin="0" aria-valuemax="100">
										</div>
										</div>
										<div class="progressbg-text">GPUPower</div>
									</div>
									</td>
									<td class="w-1 fw-bold text-end">%sW</td>
								</tr>
								<tr>
									<td>
									<div class="progressbg">
										<div class="progress progressbg-progress">
										<div class="progress-bar bg-primary-lt" style="width:%s%%" role="progressbar" aria-valuenow="%s" aria-valuemin="0" aria-valuemax="100">
										</div>
										</div>
										<div class="progressbg-text">CPUCore</div>
									</div>
									</td>
									<td class="w-1 fw-bold text-end">%s%%</td>
								</tr>
								<tr>
									<td>
									<div class="progressbg">
										<div class="progress progressbg-progress">
										<div class="progress-bar bg-primary-lt" style="width:%s%%" role="progressbar" aria-valuenow="%s" aria-valuemin="0" aria-valuemax="100">
										</div>
										</div>
										<div class="progressbg-text">Memory</div>
									</div>
									</td>
									<td class="w-1 fw-bold text-end">%s%%</td>
								</tr>
								<tr>
									<td>
									<div class="progressbg">
										<div class="progress progressbg-progress">
										<div class="progress-bar bg-primary-lt" style="width:%s%%" role="progressbar" aria-valuenow="%s" aria-valuemin="0" aria-valuemax="100">
										</div>
										</div>
										<div class="progressbg-text">Storage</div>
									</div>
									</td>
									<td class="w-1 fw-bold text-end">%s%%</td>
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
	`, nodeRes.Data.GPUCore, nodeRes.Data.GPUCore, nodeRes.Data.GPUCore,
		nodeRes.Data.GPUMemory, nodeRes.Data.GPUMemory, nodeRes.Data.GPUMemory,
		nodeRes.Data.GPUPower,
		nodeRes.Data.CPUCore, nodeRes.Data.CPUCore, nodeRes.Data.CPUCore,
		nodeRes.Data.Memory, nodeRes.Data.Memory, nodeRes.Data.Memory,
		nodeRes.Data.Storage, nodeRes.Data.Storage, nodeRes.Data.Storage)
	return returnStr
}

type PodInfo struct {
	PodName            string `json:"podName"`
	Namespace          string `json:"namespace"`
	ContainerName      string `json:"containerName"`
	ContainerImageName string `json:"containerImageName"`
	IsGPU              string `json:"isGPU"`
	Age                string `json:"age"`
}

type PodResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    []PodInfo `json:"data"`
}

func GetPodInfo(nodeName string) string {
	podRes := &PodResponse{}

	response, err := http.Get(MaterClusterHost + "pod/" + nodeName)
	if err != nil {
		klog.Errorln(err)
	}
	defer response.Body.Close()

	// 응답 본문 읽기
	body, err := io.ReadAll(response.Body)
	if err != nil {
		klog.Errorln(err)
	}

	err = json.Unmarshal(body, podRes)
	if err != nil {
		klog.Errorln(err)
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
					<th>Name</th>
					<th>Namespace</th>
					<th>Container</th>
					<th>Container Image</th>
					<th>IsGPU</th>
					<th>Age</th>
				</tr>
				</thead>
				<tbody id="podInfoTableBody">`
	for _, pod := range podRes.Data {
		returnStr = returnStr + fmt.Sprintf(`
					<tr>
						<td>%s</td>
						<td>%s</td>
						<td>%s</td>
						<td>%s</td>
						<td>%s</td>
						<td>%s</td>
					</tr>
		`, pod.PodName, pod.Namespace, pod.ContainerName, pod.ContainerImageName, pod.IsGPU, pod.Age)
	}

	returnStr = returnStr + `
				</tbody>
			</table>
		</div>	
	</div>`
	return returnStr
}

func podTr(pod corev1.Pod) (string, string, string) {
	nameTdVal := fmt.Sprintf(`<td class="u-table-cell u-text-custom-color-11">%s</td>`, pod.Name)
	namespaceTdVal := fmt.Sprintf(`<td class="u-table-cell u-text-custom-color-1">%s</td>`, pod.Namespace)
	age := time.Since(pod.CreationTimestamp.Time)
	totalSec := int(math.Round(age.Seconds()))
	day := totalSec / 86400
	hour := (totalSec % 86400) / 3600
	minute := ((totalSec % 86400) % 3600) / 60
	sec := ((totalSec % 86400) % 3600) % 60
	ageStr := ""
	if day > 0 {
		ageStr = fmt.Sprintf("%dd %dh", day, hour)
	} else if hour > 0 {
		ageStr = fmt.Sprintf("%dh %dm", hour, minute)
	} else if minute > 0 {
		ageStr = fmt.Sprintf("%dm %ds", minute, sec)
	} else {
		ageStr = fmt.Sprintf("%ds", sec)
	}
	ageTdVal := fmt.Sprintf(`<td class="u-table-cell u-text-custom-color-1">%s</td>`, ageStr)
	return nameTdVal, namespaceTdVal, ageTdVal
}

func containerTr(pod corev1.Pod) ([]string, []string) {
	containers := pod.Spec.Containers
	nameTdVals := make([]string, len(containers))
	imageTdVals := make([]string, len(containers))
	for i, container := range containers {
		nameTdVals[i] = fmt.Sprintf(`<td class="u-table-cell u-text-custom-color-1">%s</td>`, container.Name)
		imageTdVals[i] = fmt.Sprintf(`<td class="u-table-cell u-text-custom-color-1">%s</td>`, container.Image)
	}
	return nameTdVals, imageTdVals
}

func gpuTr(pod corev1.Pod) string {
	gpuTrVal := ""
	if pod.Namespace == "kube-system" || pod.Namespace == "keti-system" {
		gpuTrVal = `<td class="u-table-cell u-text-custom-color-12">False</td>`

	} else {
		gpuTrVal = `<td class="u-table-cell u-text-custom-color-13">True</td>`
	}
	// if pod.Labels["gpu"] == "true" {
	// 	gpuTrVal = `<td class="u-table-cell u-text-custom-color-13">True</td>`
	// } else {
	// 	gpuTrVal = `<td class="u-table-cell u-text-custom-color-12">False</td>`
	// }
	return gpuTrVal
}
