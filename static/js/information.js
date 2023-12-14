const searchInput = document.getElementById("searchInput");
const clusterInfoField = document.getElementById("clusterInfoField");
const nodeInfoField = document.getElementById("nodeInfoField");
const podInfoTable = document.getElementById("podInfoFiels");


// 엔터 키 이벤트 처리
searchInput.addEventListener("keyup", function (event) {
    const searchText = searchInput.value;
    
    console.log("검색어:", searchText);

    filterCheckboxes();
});

function filterCheckboxes() {
    const searchText = searchInput.value.toLowerCase();
    const checkboxes = document.querySelectorAll('.form-check');

    

    checkboxes.forEach(function (checkbox) {
        const label = checkbox.innerText.toLowerCase();
        if (label.includes(searchText)) {
            checkbox.style.display = 'block';
        } else {
            checkbox.style.display = 'none';
        }
    });
    console.log(checkboxes)
}


$("#getClusterInfo").click(function() {
    // 선택된 라디오 버튼의 값을 가져오기
    const selectedCluster = $('input[name="form-type"]:checked').val();
    // AJAX 요청 보내기
    $.ajax({
        type: "GET",
        url: "/confirm-changes",
        data: { cluster: selectedCluster },
        dataType: "json",
        success: function(response) {
            clusterInfoField.innerHTML = response.innerHTML;
        },
        error: function() {
            // 오류 처리
            alert("에러 발생");
        }
    });
});
function nodeSelectBtnclick(nodeName) {
    $.ajax({
        type: "GET",
        url: "/nodeInfo",
        data: { node: nodeName },
        dataType: "json",
        success: function(response) {
            nodeInfoField.innerHTML = "";
            nodeInfoField.innerHTML = response.innerHTML; 
        },
        error: function() {
            // 오류 처리
            alert("에러 발생");
        }
    });
    $.ajax({
        type: "GET",
        url: "/podInfo",
        data: { node: nodeName },
        dataType: "json",
        success: function(response) {
            podInfoTable.innerHTML = "";
            podInfoTable.innerHTML = response.podHTML;
        },
        error: function() {
            // 오류 처리
            alert("에러 발생");
        }
    });
}
