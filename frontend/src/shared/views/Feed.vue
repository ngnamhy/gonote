<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/features/auth/stores/authStore'

const router = useRouter()
const authStore = useAuthStore()

// Sample posts data (Reddit-style feed)
const posts = ref([
  {
    id: 1,
    author: { name: 'John Doe', avatar: 'JD', role: 'Engineering Lead' },
    team: { name: 'Engineering', color: '#3B82F6' },
    timestamp: '2 hours ago',
    title: 'New deployment pipeline is live! 🚀',
    content: 'We\'ve successfully migrated to the new CI/CD pipeline. Deployment time reduced by 60%. Check out the docs for more details.',
    upvotes: 24,
    comments: 8,
    hasUpvoted: false,
  },
  {
    id: 2,
    author: { name: 'Sarah Smith', avatar: 'SS', role: 'Product Manager' },
    team: { name: 'General', color: 'var(--color-accent)' },
    timestamp: '4 hours ago',
    title: 'Team lunch this Friday!',
    content: 'Hey everyone! We\'re organizing a team lunch this Friday at 12 PM. Please confirm your attendance in the comments.',
    upvotes: 18,
    comments: 15,
    hasUpvoted: true,
  },
  {
    id: 3,
    author: { name: 'Mike Chen', avatar: 'MC', role: 'Designer' },
    team: { name: 'Design', color: '#8B5CF6' },
    timestamp: '6 hours ago',
    title: 'New brand guidelines available',
    content: 'Updated our brand guidelines with new color palette and typography. Please review before starting new projects.',
    upvotes: 12,
    comments: 5,
    hasUpvoted: false,
  },
  {
    id: 4,
    author: { name: 'Emily Johnson', avatar: 'EJ', role: 'Marketing Manager' },
    team: { name: 'Marketing', color: '#10B981' },
    timestamp: '8 hours ago',
    title: 'Q1 Campaign Results 📊',
    content: 'Amazing news team! Our Q1 campaign exceeded targets by 35%. Total reach: 2.4M users, conversion rate: 4.2%. Great work everyone!',
    upvotes: 42,
    comments: 12,
    hasUpvoted: false,
  },
  {
    id: 5,
    author: { name: 'David Park', avatar: 'DP', role: 'Senior Developer' },
    team: { name: 'Engineering', color: '#3B82F6' },
    timestamp: '10 hours ago',
    title: 'Code review best practices',
    content: 'Reminder: Please keep PRs under 400 lines of code. Smaller PRs = faster reviews = happier developers. Let\'s maintain our 24h review SLA!',
    upvotes: 28,
    comments: 9,
    hasUpvoted: true,
  },
  {
    id: 6,
    author: { name: 'Lisa Anderson', avatar: 'LA', role: 'UX Designer' },
    team: { name: 'Design', color: '#8B5CF6' },
    timestamp: '12 hours ago',
    title: 'User research findings - Mobile app',
    content: 'Completed 15 user interviews. Key insight: 78% of users want dark mode. Navigation needs simplification. Full report in Figma.',
    upvotes: 31,
    comments: 18,
    hasUpvoted: false,
  },
  {
    id: 7,
    author: { name: 'Alex Rivera', avatar: 'AR', role: 'DevOps Engineer' },
    team: { name: 'Engineering', color: '#3B82F6' },
    timestamp: '14 hours ago',
    title: 'Scheduled maintenance - Saturday 2AM',
    content: 'Database migration planned for Saturday 2-4 AM EST. Expected downtime: 30 minutes. All services will be back online by 4:30 AM.',
    upvotes: 15,
    comments: 6,
    hasUpvoted: false,
  },
  {
    id: 8,
    author: { name: 'Rachel Kim', avatar: 'RK', role: 'Content Strategist' },
    team: { name: 'Marketing', color: '#10B981' },
    timestamp: '1 day ago',
    title: 'Blog post ideas for February',
    content: 'Looking for input on February blog topics. Current ideas: AI trends, remote work tips, productivity hacks. What else should we cover?',
    upvotes: 22,
    comments: 14,
    hasUpvoted: true,
  },
  {
    id: 9,
    author: { name: 'Tom Wilson', avatar: 'TW', role: 'Project Manager' },
    team: { name: 'General', color: 'var(--color-accent)' },
    timestamp: '1 day ago',
    title: 'Sprint retrospective highlights',
    content: 'Great sprint team! Completed 32/35 story points. Velocity improving steadily. Let\'s keep this momentum going into next sprint!',
    upvotes: 36,
    comments: 11,
    hasUpvoted: false,
  },
  {
    id: 10,
    author: { name: 'Nina Patel', avatar: 'NP', role: 'QA Lead' },
    team: { name: 'Engineering', color: '#3B82F6' },
    timestamp: '1 day ago',
    title: 'Testing automation milestone 🎯',
    content: 'We\'ve reached 85% test coverage! Automated test suite now runs in 12 minutes (down from 45). Huge thanks to the QA team!',
    upvotes: 45,
    comments: 16,
    hasUpvoted: true,
  },
  {
    id: 11,
    author: { name: 'Chris Martinez', avatar: 'CM', role: 'UI Designer' },
    team: { name: 'Design', color: '#8B5CF6' },
    timestamp: '2 days ago',
    title: 'Design system v2.0 is here!',
    content: 'New component library is live! Includes 40+ new components, accessibility improvements, and better documentation. Check it out on Storybook.',
    upvotes: 38,
    comments: 22,
    hasUpvoted: false,
  },
  {
    id: 12,
    author: { name: 'Jessica Lee', avatar: 'JL', role: 'HR Manager' },
    team: { name: 'General', color: 'var(--color-accent)' },
    timestamp: '2 days ago',
    title: 'Welcome our new team members! 👋',
    content: 'Please join me in welcoming Mark (Engineering), Sophie (Design), and James (Marketing) to the team. Say hi in the comments!',
    upvotes: 52,
    comments: 28,
    hasUpvoted: true,
  },
  {
    id: 13,
    author: { name: 'Kevin Brown', avatar: 'KB', role: 'Backend Developer' },
    team: { name: 'Engineering', color: '#3B82F6' },
    timestamp: '2 days ago',
    title: 'API performance improvements',
    content: 'Optimized database queries. Average response time down from 320ms to 85ms. P95 latency improved by 65%. Details in the tech blog.',
    upvotes: 41,
    comments: 10,
    hasUpvoted: false,
  },
  {
    id: 14,
    author: { name: 'Amanda White', avatar: 'AW', role: 'Social Media Manager' },
    team: { name: 'Marketing', color: '#10B981' },
    timestamp: '3 days ago',
    title: 'Instagram followers milestone! 🎉',
    content: 'We just hit 100K followers on Instagram! Engagement rate up 25% this month. Thank you to everyone who contributed content ideas!',
    upvotes: 67,
    comments: 31,
    hasUpvoted: true,
  },
  {
    id: 15,
    author: { name: 'Ryan Cooper', avatar: 'RC', role: 'Frontend Developer' },
    team: { name: 'Engineering', color: '#3B82F6' },
    timestamp: '3 days ago',
    title: 'New dashboard features deployed',
    content: 'Real-time analytics, custom widgets, and dark mode are now live! Check your dashboard and let us know what you think.',
    upvotes: 34,
    comments: 19,
    hasUpvoted: false,
  },
])

// Trending topics
const trending = ref([
  { tag: 'product-launch', posts: 12 },
  { tag: 'team-building', posts: 8 },
  { tag: 'engineering', posts: 15 },
  { tag: 'design-system', posts: 6 },
])

const newPostContent = ref('')
const showCreatePost = ref(false)

function toggleUpvote(post) {
  if (post.hasUpvoted) {
    post.upvotes--
  } else {
    post.upvotes++
  }
  post.hasUpvoted = !post.hasUpvoted
}

function createPost() {
  showCreatePost.value = true
}

function submitPost() {
  // Handle post creation
  showCreatePost.value = false
  newPostContent.value = ''
}
</script>

<template>
  <div class="flex justify-between gap-6 p-6" style="background-color: var(--color-bg-primary); min-height: 100%;">
    <!-- Main Feed -->
    <main class="flex-1 max-w-4xl space-y-4">
      <!-- Create Post Card -->
      <div 
        class="rounded-xl p-4"
        style="background-color: var(--color-bg-secondary)"
      >
        <div class="flex items-center gap-3">
          <div 
            class="w-10 h-10 rounded-full flex items-center justify-center"
            style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))"
          >
            <i class="pi pi-user text-white"></i>
          </div>
          <button 
            @click="createPost"
            class="flex-1 px-4 py-3 rounded-lg text-left transition-all hover:bg-opacity-80"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-muted)"
          >
            Share something with your team...
          </button>
        </div>
        <div class="flex gap-2 mt-3 pt-3 border-t" style="border-color: var(--color-border-default)">
          <button 
            class="flex-1 py-2 px-3 rounded-lg text-sm font-semibold transition-all hover:bg-opacity-80 flex items-center justify-center gap-2"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)"
          >
            <i class="pi pi-image"></i>
            Image
          </button>
          <button 
            class="flex-1 py-2 px-3 rounded-lg text-sm font-semibold transition-all hover:bg-opacity-80 flex items-center justify-center gap-2"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)"
          >
            <i class="pi pi-link"></i>
            Link
          </button>
          <button 
            class="flex-1 py-2 px-3 rounded-lg text-sm font-semibold transition-all hover:bg-opacity-80 flex items-center justify-center gap-2"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)"
          >
            <i class="pi pi-calendar"></i>
            Event
          </button>
        </div>
      </div>

      <!-- Posts Feed -->
      <div 
        v-for="post in posts" 
        :key="post.id"
        class="rounded-xl overflow-hidden"
        style="background-color: var(--color-bg-secondary)"
      >
        <!-- Post Header -->
        <div class="p-4">
          <div class="flex items-start gap-3">
            <div 
              class="w-10 h-10 rounded-full flex items-center justify-center flex-shrink-0"
              style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))"
            >
              <span class="text-white text-sm font-bold">{{ post.author.avatar }}</span>
            </div>
            <div class="flex-1 min-w-0">
              <div class="flex items-center gap-2 flex-wrap">
                <span class="font-semibold text-white">{{ post.author.name }}</span>
                <span class="text-xs" style="color: var(--color-text-muted)">{{ post.author.role }}</span>
                <span class="text-xs" style="color: var(--color-text-muted)">•</span>
                <span 
                  class="text-xs px-2 py-0.5 rounded-full"
                  :style="`background-color: ${post.team.color}20; color: ${post.team.color}`"
                >
                  {{ post.team.name }}
                </span>
                <span class="text-xs" style="color: var(--color-text-muted)">•</span>
                <span class="text-xs" style="color: var(--color-text-muted)">{{ post.timestamp }}</span>
              </div>
            </div>
          </div>

          <!-- Post Content -->
          <div class="mt-3">
            <h3 class="font-bold text-lg text-white mb-2">{{ post.title }}</h3>
            <p style="color: var(--color-text-secondary)">{{ post.content }}</p>
          </div>
        </div>

        <!-- Post Actions -->
        <div 
          class="flex items-center gap-4 px-4 py-3 border-t"
          style="background-color: var(--color-bg-secondary); border-color: var(--color-border-default)"
        >
          <button 
            @click="toggleUpvote(post)"
            class="flex items-center gap-2 px-3 py-2 rounded-lg transition-all hover:bg-opacity-80"
            :style="post.hasUpvoted ? 'background-color: var(--color-accent-muted); color: var(--color-accent)' : 'color: var(--color-text-muted)'"
          >
            <i :class="post.hasUpvoted ? 'pi pi-arrow-up-right' : 'pi pi-arrow-up'"></i>
            <span class="font-semibold text-sm">{{ post.upvotes }}</span>
          </button>
          <button 
            class="flex items-center gap-2 px-3 py-2 rounded-lg transition-all hover:bg-opacity-80"
            style="color: var(--color-text-muted)"
          >
            <i class="pi pi-comment"></i>
            <span class="font-semibold text-sm">{{ post.comments }}</span>
          </button>
          <button 
            class="flex items-center gap-2 px-3 py-2 rounded-lg transition-all hover:bg-opacity-80"
            style="color: var(--color-text-muted)"
          >
            <i class="pi pi-share-alt"></i>
            <span class="font-semibold text-sm">Share</span>
          </button>
        </div>
      </div>
    </main>

    <!-- Right Sidebar - Trending & Info -->
    <aside class="hidden xl:block w-80">
      <div class="sticky top-6 space-y-4">
        <!-- Trending Topics -->
        <div 
          class="rounded-xl p-4"
          style="background-color: var(--color-bg-secondary)"
        >
        <h3 class="font-bold text-white mb-4 flex items-center gap-2">
          <i class="pi pi-bolt" style="color: var(--color-accent)"></i>
          Trending Topics
        </h3>
        <div class="space-y-3">
          <div 
            v-for="topic in trending" 
            :key="topic.tag"
            class="flex items-center justify-between p-3 rounded-lg cursor-pointer transition-all hover:bg-opacity-80"
            style="background-color: var(--color-bg-tertiary)"
          >
            <div>
              <p class="font-semibold text-sm" style="color: var(--color-accent)">#{{ topic.tag }}</p>
              <p class="text-xs" style="color: var(--color-text-muted)">{{ topic.posts }} posts</p>
            </div>
            <i class="pi pi-arrow-right" style="color: var(--color-text-muted)"></i>
          </div>
        </div>
      </div>

      <!-- Quick Actions -->
      <div 
        class="rounded-xl p-4"
        style="background-color: var(--color-bg-secondary)"
      >
        <h3 class="font-bold text-white mb-4">Quick Actions</h3>
        <div class="space-y-2">
          <button 
            class="w-full px-4 py-3 rounded-lg font-semibold text-sm transition-all hover:bg-opacity-80 flex items-center gap-3"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)"
          >
            <i class="pi pi-users" style="color: var(--color-accent)"></i>
            Create Team
          </button>
          <button 
            class="w-full px-4 py-3 rounded-lg font-semibold text-sm transition-all hover:bg-opacity-80 flex items-center gap-3"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)"
          >
            <i class="pi pi-calendar" style="color: var(--color-accent)"></i>
            Schedule Event
          </button>
          <button 
            class="w-full px-4 py-3 rounded-lg font-semibold text-sm transition-all hover:bg-opacity-80 flex items-center gap-3"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)"
          >
            <i class="pi pi-megaphone" style="color: var(--color-accent)"></i>
            Make Announcement
          </button>
        </div>
      </div>

      <!-- Team Stats -->
      <div 
        class="rounded-xl p-4"
        style="background-color: var(--color-bg-secondary)"
      >
        <h3 class="font-bold text-white mb-4">Your Activity</h3>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <span class="text-sm" style="color: var(--color-text-muted)">Posts</span>
            <span class="font-bold text-white">24</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-sm" style="color: var(--color-text-muted)">Comments</span>
            <span class="font-bold text-white">156</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-sm" style="color: var(--color-text-muted)">Upvotes Received</span>
            <span class="font-bold text-white">342</span>
          </div>
        </div>
      </div>
      </div>
    </aside>
  </div>
</template>
