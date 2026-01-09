<script setup>
// Принимаем список процессов как проп
defineProps(['processes']);
</script>

<template>
  <div class="table-container">
    <h3>Top Processes (by RAM)</h3>
    <div class="table-scroll">
      <table>
        <thead>
        <tr>
          <th>PID</th>
          <th>Name</th>
          <th>RAM %</th>
        </tr>
        </thead>
        <tbody>
        <!-- Цикл по процессам -->
        <tr v-for="proc in processes" :key="proc.pid">
          <td class="pid">{{ proc.pid }}</td>
          <td class="name">{{ proc.name }}</td>
          <!-- Округляем до 1 знака -->
          <td class="ram">{{ proc.ram.toFixed(1) }}%</td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped>
.table-container {
  background: #3d405b;
  border-radius: 12px;
  padding: 15px;
  color: white;
  /* Занимаем всю ширину */
  width: 100%;
  box-sizing: border-box;
  margin-top: 20px;
}

h3 {
  margin: 0 0 15px 0;
  font-size: 1.2rem;
  color: #ccc;
}

.table-scroll {
  max-height: 300px; /* Если процессов много, появится скролл */
  overflow-y: auto;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th {
  text-align: left;
  padding: 8px;
  background-color: #2b2d42;
  position: sticky; /* Заголовок прилипает при скролле */
  top: 0;
}

td {
  padding: 8px;
  border-bottom: 1px solid #ffffff11;
  font-size: 0.9rem;
}

.pid {
  color: #8d99ae;
  font-family: monospace;
  width: 60px;
}

.name {
  font-weight: bold;
}

.ram {
  color: #4cc9f0;
  text-align: right;
  width: 80px;
}
</style>