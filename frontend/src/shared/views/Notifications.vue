<script setup>
import { ref } from 'vue'

const notifications = ref([
  {
    id: 1,
    type: 'upvote',
    icon: 'pi-arrow-up',
    color: 'var(--color-accent)',
    user: 'Sarah Smith',
    action: 'upvoted your post',
    content: 'New deployment pipeline is live! 🚀',
    timestamp: '5 minutes ago',
    read: false
  },
  {
    id: 2,
    type: 'comment',
    icon: 'pi-comment',
    color: '#3B82F6',
    user: 'Mike Chen',
    action: 'commented on your post',
    content: 'Great work on the design system!',
    timestamp: '1 hour ago',
    read: false
  },
  {
    id: 3,
    type: 'mention',
    icon: 'pi-at',
    color: '#10B981',
    user: 'John Doe',
    action: 'mentioned you in',
    content: 'Team Engineering',
    timestamp: '2 hours ago',
    read: false
  },
  {
    id: 4,
    type: 'team',
    icon: 'pi-users',
    color: '#8B5CF6',
    user: 'Lisa Anderson',
    action: 'invited you to join',
    content: 'UX Research Team',
    timestamp: '3 hours ago',
    read: true
  },
  {
    id: 5,
    type: 'post',
    icon: 'pi-megaphone',
    color: '#F59E0B',
    user: 'Emily Johnson',
    action: 'posted in',
    content: 'Marketing Team',
    timestamp: '5 hours ago',
    read: true
  },
  {
    id: 6,
    type: 'upvote',
    icon: 'pi-arrow-up',
    color: 'var(--color-accent)',
    user: 'David Park',
    action: 'upvoted your comment',
    content: 'I agree with this approach...',
    timestamp: '1 day ago',
    read: true
  },
  {
    id: 7,
    type: 'announcement',
    icon: 'pi-bell',
    color: '#EF4444',
    user: 'System',
    action: 'Company-wide announcement',
    content: 'Team lunch this Friday at 12 PM',
    timestamp: '1 day ago',
    read: true
  },
  {
    id: 8,
    type: 'follow',
    icon: 'pi-user-plus',
    color: '#EC4899',
    user: 'Rachel Kim',
    action: 'started following you',
    content: '',
    timestamp: '2 days ago',
    read: true
  },
])

function markAsRead(notification) {
  notification.read = true
}

function markAllAsRead() {
  notifications.value.forEach(n => n.read = true)
}

const unreadCount = ref(notifications.value.filter(n => !n.read).length)
</script>

<template>
  <div class="p-6 space-y-6" style="background-color: var(--color-bg-primary); min-height: 100%;">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-bold text-white mb-2">Notifications</h1>
        <p style="color: var(--color-text-muted)">
          <span v-if="unreadCount > 0">You have {{ unreadCount }} unread notification{{ unreadCount > 1 ? 's' : '' }}</span>
          <span v-else>You're all caught up!</span>
        </p>
      </div>
      <button 
        v-if="unreadCount > 0"
        @click="markAllAsRead"
        class="px-6 py-3 rounded-xl font-semibold transition-all hover:bg-opacity-80"
        style="background-color: var(--color-bg-tertiary); color: var(--color-accent)"
      >
        Mark all as read
      </button>
    </div>

    <!-- Notifications List -->
    <div class="space-y-3">
      <div 
        v-for="notification in notifications" 
        :key="notification.id"
        @click="markAsRead(notification)"
        class="rounded-xl p-5 transition-all hover:bg-opacity-80 cursor-pointer"
        :style="notification.read ? 'background-color: var(--color-bg-secondary)' : 'background-color: var(--color-bg-tertiary); border-left: 3px solid var(--color-accent)'"
      >
        <div class="flex items-start gap-4">
          <!-- Icon -->
          <div 
            class="w-12 h-12 rounded-full flex items-center justify-center flex-shrink-0"
            :style="`background-color: ${notification.color}20`"
          >
            <i :class="`pi ${notification.icon} text-lg`" :style="`color: ${notification.color}`"></i>
          </div>

          <!-- Content -->
          <div class="flex-1 min-w-0">
            <div class="flex items-start justify-between gap-3 mb-2">
              <div>
                <p class="text-white">
                  <span class="font-semibold">{{ notification.user }}</span>
                  <span style="color: var(--color-text-muted)"> {{ notification.action }}</span>
                </p>
                <p 
                  v-if="notification.content" 
                  class="text-sm mt-1"
                  :style="notification.read ? 'color: var(--color-text-muted)' : 'color: var(--color-text-secondary)'"
                >
                  {{ notification.content }}
                </p>
              </div>
              <div 
                v-if="!notification.read"
                class="w-2 h-2 rounded-full flex-shrink-0 mt-1"
                style="background-color: var(--color-accent)"
              ></div>
            </div>
            <p class="text-xs" style="color: var(--color-text-muted)">
              <i class="pi pi-clock mr-1"></i>
              {{ notification.timestamp }}
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div 
      v-if="notifications.length === 0"
      class="rounded-xl p-12 text-center"
      style="background-color: var(--color-bg-secondary)"
    >
      <div 
        class="w-20 h-20 rounded-full flex items-center justify-center mx-auto mb-4"
        style="background: linear-gradient(135deg, var(--color-accent-muted), transparent)"
      >
        <i class="pi pi-bell text-4xl" style="color: var(--color-accent)"></i>
      </div>
      <h3 class="text-xl font-semibold mb-2 text-white">No notifications yet</h3>
      <p style="color: var(--color-text-muted)">We'll notify you when something important happens</p>
    </div>
  </div>
</template>
