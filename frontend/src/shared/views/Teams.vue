<script setup>
import { ref } from 'vue'

const allTeams = ref([
  { 
    id: 1, 
    name: 'Engineering', 
    icon: 'pi-code', 
    color: '#3B82F6', 
    members: 24,
    description: 'Software development and technical discussions',
    joined: true
  },
  { 
    id: 2, 
    name: 'Marketing', 
    icon: 'pi-megaphone', 
    color: '#10B981', 
    members: 12,
    description: 'Marketing campaigns and brand strategy',
    joined: true
  },
  { 
    id: 3, 
    name: 'Design', 
    icon: 'pi-palette', 
    color: '#8B5CF6', 
    members: 8,
    description: 'UX/UI design and creative work',
    joined: true
  },
  { 
    id: 4, 
    name: 'General', 
    icon: 'pi-comment', 
    color: 'var(--color-accent)', 
    members: 45,
    description: 'Company-wide announcements and discussions',
    joined: true
  },
  { 
    id: 5, 
    name: 'Sales', 
    icon: 'pi-chart-line', 
    color: '#F59E0B', 
    members: 15,
    description: 'Sales team updates and strategies',
    joined: false
  },
  { 
    id: 6, 
    name: 'Support', 
    icon: 'pi-question-circle', 
    color: '#EF4444', 
    members: 10,
    description: 'Customer support and service',
    joined: false
  },
  { 
    id: 7, 
    name: 'HR', 
    icon: 'pi-users', 
    color: '#EC4899', 
    members: 5,
    description: 'Human resources and team culture',
    joined: false
  },
  { 
    id: 8, 
    name: 'Product', 
    icon: 'pi-box', 
    color: '#06B6D4', 
    members: 12,
    description: 'Product management and roadmap',
    joined: false
  },
])

function toggleJoin(team) {
  team.joined = !team.joined
  if (team.joined) {
    team.members++
  } else {
    team.members--
  }
}
</script>

<template>
  <div class="p-6 space-y-6" style="background-color: var(--color-bg-primary); min-height: 100%;">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-bold text-white mb-2">Teams</h1>
        <p style="color: var(--color-text-muted)">Discover and join teams to connect with your colleagues</p>
      </div>
      <button 
        class="px-6 py-3 rounded-xl font-semibold transition-all hover:scale-105 flex items-center gap-2"
        style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon)); color: white"
      >
        <i class="pi pi-plus"></i>
        Create Team
      </button>
    </div>

    <!-- My Teams Section -->
    <div>
      <h2 class="text-xl font-bold text-white mb-4">My Teams</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        <div 
          v-for="team in allTeams.filter(t => t.joined)" 
          :key="team.id"
          class="rounded-xl p-6 transition-all hover:scale-105 cursor-pointer"
          style="background-color: var(--color-bg-secondary)"
        >
          <div class="flex items-start gap-4 mb-4">
            <div 
              class="w-14 h-14 rounded-xl flex items-center justify-center flex-shrink-0"
              :style="`background-color: ${team.color}20`"
            >
              <i :class="`pi ${team.icon} text-2xl`" :style="`color: ${team.color}`"></i>
            </div>
            <div class="flex-1">
              <h3 class="font-bold text-lg text-white mb-1">{{ team.name }}</h3>
              <p class="text-sm" style="color: var(--color-text-muted)">{{ team.members }} members</p>
            </div>
          </div>
          <p class="text-sm mb-4" style="color: var(--color-text-secondary)">{{ team.description }}</p>
          <button 
            @click.stop="toggleJoin(team)"
            class="w-full px-4 py-2 rounded-lg font-semibold text-sm transition-all"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)"
          >
            <i class="pi pi-check mr-2"></i>
            Joined
          </button>
        </div>
      </div>
    </div>

    <!-- Discover Teams Section -->
    <div>
      <h2 class="text-xl font-bold text-white mb-4">Discover Teams</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        <div 
          v-for="team in allTeams.filter(t => !t.joined)" 
          :key="team.id"
          class="rounded-xl p-6 transition-all hover:scale-105 cursor-pointer"
          style="background-color: var(--color-bg-secondary)"
        >
          <div class="flex items-start gap-4 mb-4">
            <div 
              class="w-14 h-14 rounded-xl flex items-center justify-center flex-shrink-0"
              :style="`background-color: ${team.color}20`"
            >
              <i :class="`pi ${team.icon} text-2xl`" :style="`color: ${team.color}`"></i>
            </div>
            <div class="flex-1">
              <h3 class="font-bold text-lg text-white mb-1">{{ team.name }}</h3>
              <p class="text-sm" style="color: var(--color-text-muted)">{{ team.members }} members</p>
            </div>
          </div>
          <p class="text-sm mb-4" style="color: var(--color-text-secondary)">{{ team.description }}</p>
          <button 
            @click.stop="toggleJoin(team)"
            class="w-full px-4 py-2 rounded-lg font-semibold text-sm transition-all hover:scale-105"
            style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon)); color: white"
          >
            <i class="pi pi-plus mr-2"></i>
            Join Team
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
