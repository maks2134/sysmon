<script setup>
import {KillProcess} from "../../wailsjs/go/main/App.js";

defineProps(['processes']);

const killProc = (pid) => {
  KillProcess(pid).then((result) => {
    alert(result);
  });
}
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
          <th>Killing</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="proc in processes" :key="proc.pid">
          <td class="pid">{{ proc.pid }}</td>
          <td class="name">{{ proc.name }}</td>
          <td class="ram">{{ proc.ram.toFixed(1) }}%</td>
          <td>
            <button @click="killProc(proc.pid)" class="kill-btn">Kill</button>
          </td>
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
  max-height: 300px;
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
  position: sticky;
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

.kill-btn {
  background: #ef233c;
  border: none;
  color: white;
  padding: 4px 8px;
  cursor: pointer;
  border-radius: 4px;
  font-size: 0.8rem;
  text-align: center;
}
.kill-btn:hover {
  background: #d90429;
}
</style>