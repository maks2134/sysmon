<script setup>
import { computed } from 'vue';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler
} from 'chart.js'
import { Line } from 'vue-chartjs'

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
    Filler
)

const props = defineProps(['values', 'labels', 'color', 'label']);

const chartData = computed(() => {
  return {
    labels: props.labels,
    datasets: [
      {
        label: props.label,
        backgroundColor: props.color + '33',
        borderColor: props.color,
        data: props.values,
        fill: true,
        tension: 0.4,
        pointRadius: 0
      }
    ]
  }
});

const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  scales: {
    y: {
      min: 0,
      max: 100,
      grid: { color: '#ffffff11' },
      ticks: { color: '#ccc' }
    },
    x: {
      display: true,
      grid: { display: false },
      ticks: {
        color: '#888',
        maxTicksLimit: 6,
        maxRotation: 0
      }
    }
  },
  plugins: {
    legend: { display: false }
  },
  animation: { duration: 0 }
}
</script>

<template>
  <div class="chart-container">
    <div class="header">
      <h3>{{ label }}</h3>
      <span class="current-val" :style="{ color: color }">
            {{ values[values.length - 1] }}%
        </span>
    </div>
    <div class="chart-wrapper">
      <Line :data="chartData" :options="chartOptions" />
    </div>
  </div>
</template>

<style scoped>
.chart-container {
  background: #3d405b;
  padding: 15px;
  border-radius: 12px;
  width: 100%;
  box-sizing: border-box;
}
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}
.current-val {
  font-weight: bold;
  font-size: 1.2rem;
}
.chart-wrapper {
  height: 150px;
  width: 100%;
}
h3 {
  margin: 0;
  font-size: 1rem;
  color: #ccc;
}
</style>