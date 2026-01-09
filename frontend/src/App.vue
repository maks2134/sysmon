<script setup>
import { reactive, onMounted, onUnmounted } from 'vue';
import { EventsOn } from '../wailsjs/runtime/runtime';

const stats = reactive({
  cpu: 0,
  ram: 0
});

onMounted(() => {
  EventsOn("system_stats", (data) => {
    stats.cpu = data.cpu.toFixed(1);
    stats.ram = data.ram.toFixed(1);
  });
});
</script>

<template>
  <div class="container">
    <h1>System Monitor</h1>

    <div class="card">
      <div class="stat-row">
        <span>CPU Usage:</span>
        <span class="value">{{ stats.cpu }}%</span>
      </div>
      <div class="progress-bg">
        <div class="progress-fill cpu-fill" :style="{ width: stats.cpu + '%' }"></div>
      </div>
    </div>

    <div class="card">
      <div class="stat-row">
        <span>RAM Usage:</span>
        <span class="value">{{ stats.ram }}%</span>
      </div>
      <div class="progress-bg">
        <div class="progress-fill ram-fill" :style="{ width: stats.ram + '%' }"></div>
      </div>
    </div>
  </div>
</template>

<style>
body {
  margin: 0;
  background-color: #2b2d42;
  color: white;
  font-family: Arial, sans-serif;
}

.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 50px;
}

.card {
  background-color: #3d405b;
  width: 300px;
  padding: 20px;
  margin-bottom: 20px;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0,0,0,0.3);
}

.stat-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 10px;
  font-weight: bold;
}

.progress-bg {
  width: 100%;
  height: 20px;
  background-color: #1b1d2e;
  border-radius: 10px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  transition: width 0.5s ease-in-out;
}

.cpu-fill {
  background-color: #ef233c;
}

.ram-fill {
  background-color: #8d99ae;
}
</style>