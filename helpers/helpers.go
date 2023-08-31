package helpers

import (
	"context"
	"fmt"
	"gin_session_auth/globals"
	"math"
	"math/rand"
	"strconv"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

type ClusterInfo struct {
	Nodes string
	GPUs  string
}

func GetClusterInfo() string {
	//TODO --> 클러스터 정보 보내주는 GRPC 모듈 생성(현재는 BlackBox)
	rand.Seed(time.Now().UnixNano())
	//totalCluster := rand.Intn(5) + 1
	totalCluster := 1
	infos := make([]ClusterInfo, totalCluster)

	// for i := 0; i < totalCluster; i++ {
	// 	nodes := rand.Intn(10) + 1
	// 	gpus := rand.Intn(nodes*35-nodes) + nodes

	// 	infos[i] = ClusterInfo{
	// 		Nodes: strconv.Itoa(nodes),
	// 		GPUs:  strconv.Itoa(gpus),
	// 	}

	// }
	infos[0] = ClusterInfo{
		Nodes: strconv.Itoa(2),
		GPUs:  strconv.Itoa(4),
	}
	returnStr := ``
	for i, info := range infos {
		returnStr = returnStr + fmt.Sprintf(`
		<div
        class="u-border-2 u-border-grey-60 u-container-style u-group u-hover-feature u-radius-17 u-shape-round u-group-1 cluster-info"
        data-animation-name="" data-animation-duration="0" data-animation-direction="" onclick='getNodeInfo(this)'>
        <div class="u-container-layout u-container-layout-1">
          <p class="u-text u-text-custom-color-1 u-text-default u-text-2">Cluster</p>
          <img class="u-image u-image-contain u-image-default u-preserve-proportions u-image-1"
            src="/static/img/network.png" alt="" data-image-width="512" data-image-height="512">
          <p id="clusterNum" class="u-text u-text-custom-color-1 u-text-default u-text-3">%s</p>
          <p id="clusterNodes" class="u-text u-text-custom-color-1 u-text-default u-text-4">%s Nodes</p>
          <p class="u-text u-text-custom-color-1 u-text-default u-text-5"><br>
          </p>
          <p id="clusterGPUs" class="u-text u-text-custom-color-1 u-text-default u-text-6">%s GPUS</p>
          <p class="u-text u-text-custom-color-1 u-text-default u-text-7"><br>
          </p>
        </div>
      </div>
		`, strconv.Itoa(i), info.Nodes, info.GPUs)
	}
	return returnStr
}

type NodeInfo struct {
	AllocGPU int
	UsedGPU  int
	BlockGPU int
	GPUUsage float32
	NodeName string
}

func GetNodeInfo(clusterID string, nodeNum string, gpuNum string) string {
	//TODO --> 클러스터 정보를 토대로 Node 정보 보내주는 GRPC 모듈 생성(현재는 BlackBox)
	totalNode, err := strconv.Atoi(nodeNum)
	if err != nil {
		klog.Errorln(err)
	}
	totalGPU, err := strconv.Atoi(gpuNum)
	if err != nil {
		klog.Errorln(err)
	}
	rand.Seed(time.Now().UnixNano())
	gpuMod := 0.0
	infos := make([]NodeInfo, totalNode)
	returnStr := ``
	for i := 0; i < totalNode; i++ {
		seedStr := `<div
		class="u-border-1 u-border-white u-container-style u-group u-hover-feature u-radius-8 u-shape-round u-group-1"
		data-animation-name="" data-animation-duration="0" data-animation-direction="">
		<div class="u-container-layout u-container-layout-1">
		  <div class="u-container-style u-custom-color-8 u-group u-radius-50 u-shape-round u-group-2">
			<div class="u-container-layout u-container-layout-2">
			  <p class="u-align-center u-text u-text-custom-color-9 u-text-2">7d 15h</p>
			</div>
		  </div>
		  <div class="u-table u-table-responsive u-table-1" style="position: relative; max-height: 125px; ">
			<div style="width: 100%; height: 100%; display: flex; justify-content: center;">
			  <img src="/static/img/graphics-card_negative.png"
				style="content: ''; display: block; position: absolute; background-size: cover; z-index: 1; height: 100%; margin : auto;">
			  </img>
			</div>
			<table class="u-table-entity" style="max-height: 100%;min-height: 0px;">
			  <colgroup>
				<col width="14.28%">
				<col width="14.28%">
				<col width="14.28%">
				<col width="14.28%">
				<col width="14.28%">
				<col width="14.28%">
				<col width="14.28%">
			  </colgroup>
			  <tbody class="u-table-body u-table-body-1">
				<tr style="height: 25px;">`
		gpuMod = gpuMod + (float64(totalGPU) / float64(totalNode))
		allocatedGPU := totalGPU / totalNode
		gpuMod = gpuMod - float64(allocatedGPU)
		if i+1 == totalNode {
			allocatedGPU = allocatedGPU + int(math.Ceil(gpuMod))
		}

		infos[i].UsedGPU = rand.Intn(allocatedGPU) + 1

		infos[i].BlockGPU = 35 - allocatedGPU

		infos[i].AllocGPU = allocatedGPU - infos[i].UsedGPU

		infos[i].GPUUsage = float32(infos[i].UsedGPU) / float32(allocatedGPU) * 100
		roundedNum := math.Round(float64(infos[i].GPUUsage)*10) / 10

		percent := fmt.Sprint(roundedNum, "%")
		useTd := `<td class="u-border-2 u-border-white u-custom-color-3 u-table-cell u-table-cell-1"></td>`
		allocateTd := `<td class="u-border-2 u-border-white u-custom-color-4 u-table-cell u-table-cell-1"></td>`
		blockTd := `<td class="u-border-2 u-border-white u-custom-color-8 u-table-cell u-table-cell-1"></td>`
		uCnt := infos[i].UsedGPU
		aCnt := infos[i].AllocGPU
		bCnt := infos[i].BlockGPU

		colCnt := 7

		for j := 0; j < 35; j++ {
			if colCnt == 0 {
				colCnt = 7
				seedStr = seedStr + `
				</tr>
				<tr style="height: 25px;">`

			}
			if uCnt > 0 {
				seedStr = seedStr + useTd
				uCnt--
				colCnt--
			} else if aCnt > 0 {
				seedStr = seedStr + allocateTd
				aCnt--
				colCnt--
			} else if bCnt > 0 {
				seedStr = seedStr + blockTd
				bCnt--
				colCnt--
			}
		}
		seedStr = seedStr + `
				</tr>
			  </tbody>
			</table>
		  </div>`
		seedStr = seedStr + fmt.Sprintf(`
		<p class="u-text u-text-custom-color-1 u-text-default u-text-3">Node</p>
		<p class="u-text u-text-custom-color-1 u-text-default u-text-4"> %s </p>
		<p class="u-text u-text-custom-color-1 u-text-default u-text-5">Current GPU Usage</p>
		<p class="u-text u-text-custom-color-1 u-text-default u-text-6">%s</p>
		<div class="u-container-style u-gradient u-group u-radius-6 u-shape-round u-group-3">
		  <div class="u-container-layout" onclick='showNodeInfo(this)' style="z-index:10000">
			<p class="u-text u-text-8">Show Detail</p>
			<div id="getID" style="display: none;"> %s </div>
		  </div>
		</div>
	  </div>
	</div>`, strconv.Itoa(i+1), percent, strconv.Itoa(i+1))
		returnStr = returnStr + seedStr
	}

	return returnStr
}

func GetPodInfo(clusterName string) string {
	podList := &corev1.PodList{}
	var err error
	if globals.Client[clusterName] != nil {
		podList, err = globals.Client[clusterName].CoreV1().Pods(corev1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			klog.Errorln(err)
		}
	} else {
		podList, err = globals.TestClient.CoreV1().Pods(corev1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			klog.Errorln(err)
		}
	}
	tableHTML := ``
	for _, pod := range podList.Items {
		name, namespace, age := podTr(pod)
		containerName, image := containerTr(pod)
		gpu := gpuTr(pod)
		for i := 0; i < len(containerName); i++ {
			tableHTML = tableHTML + `<tr style="height: 65px;">`
			tableHTML = tableHTML + name
			tableHTML = tableHTML + namespace
			tableHTML = tableHTML + containerName[i]
			tableHTML = tableHTML + image[i]
			tableHTML = tableHTML + gpu
			tableHTML = tableHTML + age
			tableHTML = tableHTML + `</tr>`
		}
	}

	return tableHTML
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
