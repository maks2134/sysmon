<script setup>
import { reactive, onMounted } from 'vue';
import { EventsOn } from '../wailsjs/runtime/runtime';
import SysChart from './components/SysChart.vue';
import ProcessTable from './components/ProcessTable.vue';

const MAX_POINTS = 60;
const initData = new Array(MAX_POINTS).fill(0);
const initLabels = new Array(MAX_POINTS).fill("");

const history = reactive({
  cpu: [...initData],
  ram: [...initData],
  disk: [...initData],
  labels: [...initLabels],
  processes: []
});

onMounted(() => {
  EventsOn("system_stats", (data) => {
    const cpuVal = data.cpu ? parseFloat(data.cpu.toFixed(1)) : 0;
    const ramVal = data.ram ? parseFloat(data.ram.toFixed(1)) : 0;
    const diskVal = data.disk ? parseFloat(data.disk.toFixed(1)) : 0;
    const timeStr = new Date().toLocaleTimeString('ru-RU');

    history.cpu = [...history.cpu.slice(1), cpuVal];
    history.ram = [...history.ram.slice(1), ramVal];
    history.disk = [...history.disk.slice(1), diskVal];
    history.labels = [...history.labels.slice(1), timeStr];

    history.processes = data.processes || [];
  });
});
</script>

<template>
  <div class="container">
    <h1 class="title">System Monitor</h1>

    <div class="grid">
      <SysChart label="CPU Usage" :values="history.cpu" :labels="history.labels" color="#ef233c" />
      <SysChart label="RAM Usage" :values="history.ram" :labels="history.labels" color="#4cc9f0" />
      <SysChart label="Disk Usage (/)" :values="history.disk" :labels="history.labels" color="#fee440" />
    </div>

    <ProcessTable :processes="history.processes" />
  </div>
</template>

<style>
body {
  margin: 0;
  background-color: #2b2d42;
  color: white;
  font-family: 'Segoe UI', sans-serif;
}

.container {
  padding: 20px;
  max-width: 1000px;
  margin: 0 auto;
}

.title {
  text-align: center;
  margin-bottom: 30px;
  font-weight: 300;
  letter-spacing: 2px;
}

.grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 20px;
  margin-bottom: 20px;
}

@media (min-width: 600px) {
  .grid { grid-template-columns: 1fr 1fr; }
}
@media (min-width: 900px) {
  .grid { grid-template-columns: 1fr 1fr 1fr; }
}
</style>