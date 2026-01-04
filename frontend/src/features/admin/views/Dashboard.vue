<script setup>
import { ref } from 'vue'
import Chart from 'primevue/chart'
import Card from 'primevue/card'

const stats = ref([
  { title: 'Total Users', value: '2,543', change: '+12%', icon: 'pi-users', color: '#E8985E' },
  { title: 'Active Posts', value: '1,234', change: '+8%', icon: 'pi-file', color: '#A9714B' },
  { title: 'New Today', value: '89', change: '+23%', icon: 'pi-chart-line', color: '#54442B' },
  { title: 'Engagement', value: '94%', change: '+5%', icon: 'pi-heart', color: '#E8985E' }
])

const userChartData = ref({
  labels: ['Admin', 'Users', 'Moderators', 'Guests'],
  datasets: [{
    label: 'User Distribution',
    data: [45, 540, 125, 280],
    backgroundColor: ['#54442B', '#A9714B', '#E8985E', '#262A10'],
    borderColor: '#141204',
    borderWidth: 2
  }]
})

const activityChartData = ref({
  labels: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
  datasets: [{
    label: 'Activity',
    data: [65, 78, 90, 81, 95, 105, 88],
    fill: true,
    backgroundColor: 'rgba(232, 152, 94, 0.2)',
    borderColor: '#E8985E',
    borderWidth: 3,
    tension: 0.4
  }]
})

const chartOptions = ref({
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      labels: {
        color: '#E8985E',
        font: { size: 14 }
      }
    }
  },
  scales: {
    y: {
      ticks: { color: '#A9714B' },
      grid: { color: '#54442B' }
    },
    x: {
      ticks: { color: '#A9714B' },
      grid: { color: '#54442B' }
    }
  }
})

const recentActivities = ref([
  { user: 'John Doe', action: 'created a new post', time: '2 minutes ago' },
  { user: 'Jane Smith', action: 'updated profile', time: '15 minutes ago' },
  { user: 'Mike Johnson', action: 'commented on post', time: '1 hour ago' },
  { user: 'Sarah Wilson', action: 'joined the platform', time: '3 hours ago' }
])
</script>

<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-4xl font-bold text-white mb-2">Dashboard</h1>
      <p class="text-lg" style="color: var(--color-text-tertiary)">
        Welcome back! Here's what's happening today.
      </p>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <div 
        v-for="stat in stats" 
        :key="stat.title"
        class="p-6 rounded-2xl border transition-all hover:scale-105 cursor-pointer"
        style="background-color: var(--color-bg-secondary); border-color: var(--color-border-default)"
      >
        <div class="flex items-center justify-between mb-4">
          <i 
            :class="`pi ${stat.icon} text-3xl`" 
            :style="`color: ${stat.color}`"
          ></i>
          <span 
            class="text-sm font-semibold px-3 py-1 rounded-full"
            style="background-color: rgba(232, 152, 94, 0.1); color: var(--color-accent)"
          >
            {{ stat.change }}
          </span>
        </div>
        <h3 class="text-3xl font-bold text-white mb-1">{{ stat.value }}</h3>
        <p class="text-sm" style="color: var(--color-text-tertiary)">{{ stat.title }}</p>
      </div>
    </div>

    <!-- Charts Row -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
      <!-- Activity Chart -->
      <div 
        class="p-6 rounded-2xl border"
        style="background-color: var(--color-bg-secondary); border-color: var(--color-border-default)"
      >
        <h2 class="text-2xl font-bold text-white mb-6">Weekly Activity</h2>
        <div style="height: 300px">
          <Chart type="line" :data="activityChartData" :options="chartOptions" />
        </div>
      </div>

      <!-- User Distribution Chart -->
      <div 
        class="p-6 rounded-2xl border"
        style="background-color: var(--color-bg-secondary); border-color: var(--color-border-default)"
      >
        <h2 class="text-2xl font-bold text-white mb-6">User Distribution</h2>
        <div style="height: 300px">
          <Chart type="doughnut" :data="userChartData" :options="chartOptions" />
        </div>
      </div>
    </div>

    <!-- Recent Activity -->
    <div 
      class="p-6 rounded-2xl border"
      style="background-color: var(--color-bg-secondary); border-color: var(--color-border-default)"
    >
      <h2 class="text-2xl font-bold text-white mb-6">Recent Activity</h2>
      <div class="space-y-4">
        <div 
          v-for="(activity, index) in recentActivities" 
          :key="index"
          class="flex items-center justify-between p-4 rounded-xl transition-colors hover:bg-opacity-50"
          style="background-color: var(--color-bg-tertiary)"
        >
          <div class="flex items-center gap-4">
            <div 
              class="w-10 h-10 rounded-full flex items-center justify-center"
              style="background-color: var(--color-accent-muted)"
            >
              <i class="pi pi-user" style="color: var(--color-accent)"></i>
            </div>
            <div>
              <p class="text-white font-semibold">{{ activity.user }}</p>
              <p class="text-sm" style="color: var(--color-text-tertiary)">
                {{ activity.action }}
              </p>
            </div>
          </div>
          <span class="text-sm" style="color: var(--color-text-muted)">
            {{ activity.time }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* Chart.js canvas styling */
:deep(canvas) {
  background: transparent !important;
}
</style>