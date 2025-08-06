<template>
  <div class="min-h-screen">
    <!-- Mobile Header -->
    <header class="glass-nav fixed top-0 left-0 right-0 z-40 h-16">
      <div class="flex items-center justify-between h-full mobile-padding">
        <div class="flex items-center space-x-3">
          <NuxtLink to="/profile" class="glass-button !p-2">
            <Icon name="heroicons:arrow-left" class="h-6 w-6" />
          </NuxtLink>
          <h1 class="text-lg font-bold text-white">Statistics</h1>
        </div>
        
        <div class="flex items-center space-x-2">
          <select v-model="selectedPeriod" class="glass-select text-sm">
            <option value="week">Week</option>
            <option value="month">Month</option>
            <option value="quarter">Quarter</option>
            <option value="year">Year</option>
            <option value="all">All Time</option>
          </select>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="pt-16">
      <div class="mobile-padding max-w-6xl mx-auto pb-20">
      <!-- Stats Header -->
      <div class="glass-card p-6 mb-6">
        <div class="flex items-center gap-4 mb-4">
          <NuxtLink to="/profile" class="glass-button">
            <Icon name="heroicons:arrow-left" class="h-4 w-4 mr-2" />
            Back
          </NuxtLink>
          <h1 class="text-2xl font-bold text-white flex-1">Reading Statistics</h1>
          
          <select v-model="selectedPeriod" class="glass-select">
            <option value="week">This Week</option>
            <option value="month">This Month</option>
            <option value="quarter">This Quarter</option>
            <option value="year">This Year</option>
            <option value="all">All Time</option>
          </select>
        </div>
        <p class="text-gray-300">Track your reading progress and discover insights about your academic journey.</p>
      </div>

      <!-- Key Metrics -->
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-6">
        <div class="glass-card p-6 text-center">
          <div class="text-3xl font-bold text-purple-400 mb-2">{{ stats.totalArticles }}</div>
          <div class="text-sm text-gray-400">Total Articles</div>
          <div class="text-xs text-green-400 mt-1">+{{ stats.newThisPeriod }} this {{ selectedPeriod }}</div>
        </div>
        
        <div class="glass-card p-6 text-center">
          <div class="text-3xl font-bold text-green-400 mb-2">{{ stats.completedArticles }}</div>
          <div class="text-sm text-gray-400">Completed</div>
          <div class="text-xs text-gray-500 mt-1">{{ Math.round((stats.completedArticles / stats.totalArticles) * 100) }}% of total</div>
        </div>
        
        <div class="glass-card p-6 text-center">
          <div class="text-3xl font-bold text-yellow-400 mb-2">{{ stats.inProgressArticles }}</div>
          <div class="text-sm text-gray-400">In Progress</div>
          <div class="text-xs text-blue-400 mt-1">{{ stats.avgProgress }}% avg progress</div>
        </div>
        
        <div class="glass-card p-6 text-center">
          <div class="text-3xl font-bold text-blue-400 mb-2">{{ stats.readingStreak }}</div>
          <div class="text-sm text-gray-400">Day Streak</div>
          <div class="text-xs text-purple-400 mt-1">{{ stats.longestStreak }} day record</div>
        </div>
      </div>

      <!-- Charts Section -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
        <!-- Reading Progress Chart -->
        <div class="glass-card p-6">
          <h3 class="text-lg font-semibold text-white mb-4 flex items-center gap-2">
            <Icon name="heroicons:chart-bar" class="h-5 w-5 text-purple-400" />
            Reading Progress Over Time
          </h3>
          
          <div class="h-64 flex items-end justify-between gap-2 px-4">
            <div 
              v-for="(day, index) in progressChartData"
              :key="index"
              class="flex flex-col items-center gap-2"
            >
              <div class="relative flex-1 w-8 bg-gray-700 rounded-t-lg overflow-hidden min-h-[100px]">
                <div 
                  class="absolute bottom-0 w-full bg-gradient-to-t from-purple-500 to-blue-400 transition-all duration-500 rounded-t-lg"
                  :style="{ height: `${Math.max(day.completed * 4, 8)}px` }"
                ></div>
                <div 
                  class="absolute bottom-0 w-full bg-gradient-to-t from-yellow-500 to-orange-400 opacity-60 transition-all duration-500"
                  :style="{ height: `${Math.max(day.started * 4, 4)}px` }"
                ></div>
              </div>
              <div class="text-xs text-gray-400 text-center">
                {{ day.label }}
              </div>
            </div>
          </div>
          
          <div class="flex items-center justify-center gap-6 mt-4 text-sm">
            <div class="flex items-center gap-2">
              <div class="w-3 h-3 bg-gradient-to-r from-purple-500 to-blue-400 rounded"></div>
              <span class="text-gray-300">Completed</span>
            </div>
            <div class="flex items-center gap-2">
              <div class="w-3 h-3 bg-gradient-to-r from-yellow-500 to-orange-400 rounded"></div>
              <span class="text-gray-300">Started</span>
            </div>
          </div>
        </div>

        <!-- Topic Distribution -->
        <div class="glass-card p-6">
          <h3 class="text-lg font-semibold text-white mb-4 flex items-center gap-2">
            <Icon name="heroicons:chart-pie" class="h-5 w-5 text-purple-400" />
            Topics Distribution
          </h3>
          
          <div class="space-y-4">
            <div 
              v-for="topic in topicsData"
              :key="topic.name"
              class="flex items-center gap-4"
            >
              <div class="flex-1">
                <div class="flex items-center justify-between mb-1">
                  <span class="text-gray-300 text-sm">{{ topic.name }}</span>
                  <span class="text-white font-medium">{{ topic.count }}</span>
                </div>
                <div class="w-full h-2 bg-gray-700 rounded-full overflow-hidden">
                  <div 
                    class="h-full rounded-full transition-all duration-500"
                    :style="{ 
                      width: `${(topic.count / Math.max(...topicsData.map(t => t.count))) * 100}%`,
                      background: topic.color 
                    }"
                  ></div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Detailed Statistics -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-6">
        <!-- Reading Habits -->
        <div class="glass-card p-6">
          <h3 class="text-lg font-semibold text-white mb-4 flex items-center gap-2">
            <Icon name="heroicons:clock" class="h-5 w-5 text-purple-400" />
            Reading Habits
          </h3>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-gray-300">Average per week</span>
              <span class="text-white font-medium">{{ habits.avgPerWeek }} articles</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-gray-300">Best reading day</span>
              <span class="text-white font-medium">{{ habits.bestDay }}</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-gray-300">Preferred time</span>
              <span class="text-white font-medium">{{ habits.preferredTime }}</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-gray-300">Avg. completion time</span>
              <span class="text-white font-medium">{{ habits.avgCompletionTime }} days</span>
            </div>
          </div>
        </div>

        <!-- Library Stats -->
        <div class="glass-card p-6">
          <h3 class="text-lg font-semibold text-white mb-4 flex items-center gap-2">
            <Icon name="heroicons:book-open" class="h-5 w-5 text-purple-400" />
            Library Stats
          </h3>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-gray-300">Total libraries</span>
              <span class="text-white font-medium">{{ libraryStats.total }}</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-gray-300">Largest library</span>
              <span class="text-white font-medium">{{ libraryStats.largest.name }}</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-gray-300">Most active</span>
              <span class="text-white font-medium">{{ libraryStats.mostActive.name }}</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="text-gray-300">Completion rate</span>
              <span class="text-white font-medium">{{ libraryStats.completionRate }}%</span>
            </div>
          </div>
        </div>

        <!-- Recent Achievements -->
        <div class="glass-card p-6">
          <h3 class="text-lg font-semibold text-white mb-4 flex items-center gap-2">
            <Icon name="heroicons:trophy" class="h-5 w-5 text-purple-400" />
            Achievements
          </h3>
          
          <div class="space-y-3">
            <div 
              v-for="achievement in achievements"
              :key="achievement.id"
              class="flex items-center gap-3 p-3 glass-elevated rounded-lg"
            >
              <div class="text-2xl">{{ achievement.icon }}</div>
              <div class="flex-1">
                <div class="text-white text-sm font-medium">{{ achievement.title }}</div>
                <div class="text-gray-400 text-xs">{{ achievement.description }}</div>
              </div>
              <div v-if="achievement.isNew" class="text-xs text-green-400 font-medium">New!</div>
            </div>
          </div>
        </div>
      </div>

      <!-- Monthly Calendar -->
      <div class="glass-card p-6 mb-6">
        <h3 class="text-lg font-semibold text-white mb-4 flex items-center gap-2">
          <Icon name="heroicons:calendar-days" class="h-5 w-5 text-purple-400" />
          Reading Activity Calendar
        </h3>
        
        <div class="grid grid-cols-7 gap-1 mb-4">
          <div v-for="day in ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']" :key="day" class="text-center text-xs text-gray-400 p-2">
            {{ day }}
          </div>
        </div>
        
        <div class="grid grid-cols-7 gap-1">
          <div 
            v-for="(day, index) in calendarData"
            :key="index"
            :class="[
              'aspect-square rounded p-1 text-xs flex items-center justify-center cursor-pointer transition-all',
              getActivityClass(day.activity),
              day.isToday ? 'ring-2 ring-purple-400' : ''
            ]"
            :title="`${day.date}: ${day.activity} articles read`"
          >
            {{ day.day }}
          </div>
        </div>
        
        <div class="flex items-center justify-between mt-4 text-xs text-gray-400">
          <span>Less</span>
          <div class="flex items-center gap-1">
            <div class="w-3 h-3 rounded bg-gray-700"></div>
            <div class="w-3 h-3 rounded bg-green-900"></div>
            <div class="w-3 h-3 rounded bg-green-700"></div>
            <div class="w-3 h-3 rounded bg-green-500"></div>
            <div class="w-3 h-3 rounded bg-green-400"></div>
          </div>
          <span>More</span>
        </div>
      </div>

      <!-- Export Options -->
      <div class="glass-card p-6">
        <h3 class="text-lg font-semibold text-white mb-4 flex items-center gap-2">
          <Icon name="heroicons:document-arrow-down" class="h-5 w-5 text-purple-400" />
          Export Statistics
        </h3>
        
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <button class="glass-button justify-center p-4">
            <Icon name="heroicons:document-text" class="h-5 w-5 mr-2" />
            Export as PDF
          </button>
          
          <button class="glass-button justify-center p-4">
            <Icon name="heroicons:table-cells" class="h-5 w-5 mr-2" />
            Export as CSV
          </button>
          
          <button class="glass-button justify-center p-4">
            <Icon name="heroicons:photo" class="h-5 w-5 mr-2" />
            Share Image
          </button>
        </div>
      </div>
    </div>
    </main>

    <!-- Mobile Bottom Navigation -->
    <MobileBottomNav />
  </div>
</template>
<script setup>
const selectedPeriod = ref('month')

// Mock data - replace with real API calls
const stats = ref({
  totalArticles: 156,
  completedArticles: 89,
  inProgressArticles: 23,
  newThisPeriod: 12,
  avgProgress: 67,
  readingStreak: 8,
  longestStreak: 23
})

const progressChartData = ref([
  { label: 'Mon', completed: 3, started: 5 },
  { label: 'Tue', completed: 1, started: 2 },
  { label: 'Wed', completed: 4, started: 3 },
  { label: 'Thu', completed: 2, started: 4 },
  { label: 'Fri', completed: 5, started: 2 },
  { label: 'Sat', completed: 3, started: 6 },
  { label: 'Sun', completed: 2, started: 1 }
])

const topicsData = ref([
  { name: 'Machine Learning', count: 45, color: 'linear-gradient(135deg, #8b5cf6, #3b82f6)' },
  { name: 'Computer Vision', count: 32, color: 'linear-gradient(135deg, #06b6d4, #0891b2)' },
  { name: 'NLP', count: 28, color: 'linear-gradient(135deg, #10b981, #059669)' },
  { name: 'Deep Learning', count: 24, color: 'linear-gradient(135deg, #f59e0b, #d97706)' },
  { name: 'AI Ethics', count: 15, color: 'linear-gradient(135deg, #ef4444, #dc2626)' },
  { name: 'Other', count: 12, color: 'linear-gradient(135deg, #6b7280, #4b5563)' }
])

const habits = ref({
  avgPerWeek: 5.2,
  bestDay: 'Wednesday',
  preferredTime: '9:00 AM',
  avgCompletionTime: 3.5
})

const libraryStats = ref({
  total: 8,
  largest: { name: 'ML Research', count: 45 },
  mostActive: { name: 'Current Reading', count: 23 },
  completionRate: 68
})

const achievements = ref([
  {
    id: 1,
    icon: '🔥',
    title: 'Reading Streak',
    description: '7 days in a row',
    isNew: true
  },
  {
    id: 2,
    icon: '📚',
    title: 'Century Reader',
    description: '100 articles read',
    isNew: false
  },
  {
    id: 3,
    icon: '⚡',
    title: 'Speed Reader',
    description: '5 articles in one day',
    isNew: false
  },
  {
    id: 4,
    icon: '🎯',
    title: 'Focused',
    description: 'Completed 90% of started articles',
    isNew: true
  }
])

// Generate calendar data for current month
const calendarData = computed(() => {
  const today = new Date()
  const year = today.getFullYear()
  const month = today.getMonth()

  const firstDay = new Date(year, month, 1)
  const lastDay = new Date(year, month + 1, 0)
  const startDate = new Date(firstDay)
  startDate.setDate(startDate.getDate() - firstDay.getDay() + 1)

  const days = []
  const currentDate = new Date(startDate)

  for (let i = 0; i < 42; i++) {
    const activity = Math.floor(Math.random() * 6) // Random activity level 0-5
    days.push({
      day: currentDate.getDate(),
      date: currentDate.toISOString().split('T')[0],
      activity,
      isToday: currentDate.toDateString() === today.toDateString(),
      isCurrentMonth: currentDate.getMonth() === month
    })
    currentDate.setDate(currentDate.getDate() + 1)
  }

  return days
})

// Methods
const getActivityClass = (activity) => {
  const classes = [
    'bg-gray-700 text-gray-500',
    'bg-green-900 text-green-200',
    'bg-green-800 text-green-100',
    'bg-green-600 text-white',
    'bg-green-500 text-white',
    'bg-green-400 text-white'
  ]
  return classes[activity] || classes[0]
}

// Meta
useHead({
  title: 'Reading Statistics - Academic Reader',
  meta: [
    { name: 'description', content: 'Track your reading progress and discover insights' }
  ]
})
</script>
