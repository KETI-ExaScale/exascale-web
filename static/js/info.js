function getNodeInfo(clusterInfo) {
    firstText1 = document.getElementById('firstText1');
    firstText1.style.display = "none";

    firstText2 = document.getElementById('firstText2');
    firstText2.style.display = "none";

    nonClick = document.querySelectorAll(".cluster-info");
    nonClick.forEach((e) => {
        e.classList.remove("cluster-click");
    });
    $(clusterInfo).addClass('cluster-click')
    clusterNum = $(clusterInfo).find("#clusterNum").text();
    clusterNodes = $(clusterInfo).find("#clusterNodes").text();
    clusterGPUs = $(clusterInfo).find("#clusterGPUs").text();

    nodeAndCnt = clusterNodes.split(' ');
    gpuAndCnt = clusterGPUs.split(' ');
    clusterData = { "clusterNum": clusterNum, "nodeCnt": nodeAndCnt[0], "gpuCnt": gpuAndCnt[0] };
    infoCards = document.getElementById('nodeInfoCards');
    infoCards.innerHTML = "";

    podInfoBody = document.getElementById('podInfoBody');
    podInfoBody.innerHTML = "";

    $.ajax({
        type: "GET",
        url: "/nodeInfo",
        data: clusterData,
        error: function () {
            alert("에러발생");
        },
        success: function (json) {
            infoCards.innerHTML = json.innerHTML;
        }
    });

    $.ajax({
        type: "GET",
        url: "/podInfo",
        data: clusterData,
        error: function () {
            alert("에러발생");
        },
        success: function (json) {
            podInfoBody.innerHTML = json.innerHTML;
        }
    });
}


function showNodeInfo(nodeInfo) {
    nodeNum = $(nodeInfo).find("#getID").text();
    window.open("/", "", "width=500, height=500");
}

