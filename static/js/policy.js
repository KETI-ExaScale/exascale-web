function metricTimeSet(metricInterval) {
    mValue = metricInterval.value;
    label = document.getElementById("collectingTime");
    label.innerHTML = mValue + " Seconds";
}

function confidentialSet(confidential) {
    cValue = confidential.value;
    label = document.getElementById("confidence");
    label.innerHTML = cValue + " Percent";

}

function allocateGPUSet(gpuNum) {
    gValue = gpuNum.value;
    label = document.getElementById("alocGPUNum");
    label.innerHTML = gValue + " Devices";

}