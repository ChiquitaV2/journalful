interface LayoutConfig {
  pageTitle: string
  backRoute: string
  showSearchButton: boolean
  showAddButton: boolean
  showActionButton: boolean
  actionIcon: string
  showProfile: boolean
}

interface HeaderButtonConfig {
  search?: boolean
  add?: boolean
  action?: boolean
  actionIcon?: string
  profile?: boolean
}

/**
 * Layout composable for managing layout state and configuration
 */
export const useLayout = () => {
  // Layout configuration state
  const layoutConfig = ref<LayoutConfig>({
    pageTitle: 'Academic Reader',
    backRoute: '/',
    showSearchButton: true,
    showAddButton: false,
    showActionButton: false,
    actionIcon: 'heroicons:ellipsis-horizontal',
    showProfile: true
  })

  // Update layout configuration
  const setLayoutConfig = (config: Partial<LayoutConfig>) => {
    layoutConfig.value = { ...layoutConfig.value, ...config }
  }

  // Set page title
  const setPageTitle = (title: string) => {
    layoutConfig.value.pageTitle = title
  }

  // Set back route
  const setBackRoute = (route: string) => {
    layoutConfig.value.backRoute = route
  }

  // Configure header buttons
  const configureHeaderButtons = (config: HeaderButtonConfig) => {
    if (config.search !== undefined) layoutConfig.value.showSearchButton = config.search
    if (config.add !== undefined) layoutConfig.value.showAddButton = config.add
    if (config.action !== undefined) layoutConfig.value.showActionButton = config.action
    if (config.actionIcon !== undefined) layoutConfig.value.actionIcon = config.actionIcon
    if (config.profile !== undefined) layoutConfig.value.showProfile = config.profile
  }

  return {
    layoutConfig: readonly(layoutConfig),
    setLayoutConfig,
    setPageTitle,
    setBackRoute,
    configureHeaderButtons
  }
}

interface ReaderConfig {
  showProgress: boolean
  progress: number
  progressText: string
  showBookmark: boolean
  isBookmarked: boolean
  showShare: boolean
  showSettings: boolean
}

interface ReaderButtonConfig {
  bookmark?: boolean
  share?: boolean
  settings?: boolean
}

/**
 * Reader layout composable for article reading pages
 */
export const useReaderLayout = () => {
  // Reader-specific state
  const readerConfig = ref<ReaderConfig>({
    showProgress: false,
    progress: 0,
    progressText: '',
    showBookmark: false,
    isBookmarked: false,
    showShare: false,
    showSettings: false
  })

  // Update reader configuration
  const setReaderConfig = (config: Partial<ReaderConfig>) => {
    readerConfig.value = { ...readerConfig.value, ...config }
  }

  // Update reading progress
  const setProgress = (progress: number, text = '') => {
    readerConfig.value.progress = Math.max(0, Math.min(100, progress))
    readerConfig.value.progressText = text
    readerConfig.value.showProgress = true
  }

  // Toggle bookmark state
  const toggleBookmark = () => {
    readerConfig.value.isBookmarked = !readerConfig.value.isBookmarked
    return readerConfig.value.isBookmarked
  }

  // Configure reader buttons
  const configureReaderButtons = (config: ReaderButtonConfig) => {
    if (config.bookmark !== undefined) readerConfig.value.showBookmark = config.bookmark
    if (config.share !== undefined) readerConfig.value.showShare = config.share
    if (config.settings !== undefined) readerConfig.value.showSettings = config.settings
  }

  return {
    readerConfig: readonly(readerConfig),
    setReaderConfig,
    setProgress,
    toggleBookmark,
    configureReaderButtons
  }
}

interface PageMetaConfig {
  title?: string
  description?: string
  keywords?: string
}

/**
 * Page meta composable for managing page metadata
 */
export const usePageMeta = () => {
  // Set page title and meta tags
  const setPageMeta = (meta: PageMetaConfig) => {
    const defaultMeta = {
      title: 'Academic Reading List Manager',
      description: 'Organize, track, and share your academic article reading progress.',
      keywords: 'academic, reading, research, articles, papers'
    }

    const finalMeta = { ...defaultMeta, ...meta }

    useHead({
      title: finalMeta.title,
      meta: [
        { name: 'description', content: finalMeta.description },
        { name: 'keywords', content: finalMeta.keywords },
        { property: 'og:title', content: finalMeta.title },
        { property: 'og:description', content: finalMeta.description },
        { property: 'twitter:title', content: finalMeta.title },
        { property: 'twitter:description', content: finalMeta.description }
      ]
    })
  }

  // Set page title only
  const setTitle = (title: string) => {
    useHead({
      title: `${title} - Academic Reading List Manager`
    })
  }

  return {
    setPageMeta,
    setTitle
  }
}

/**
 * Loading state composable for layout
 */
export const useLayoutLoading = () => {
  const isLoading = ref(false)
  const loadingText = ref('Loading...')

  const setLoading = (loading: boolean, text = 'Loading...') => {
    isLoading.value = loading
    loadingText.value = text
  }

  const withLoading = async (fn: () => Promise<void>, text = 'Loading...') => {
    setLoading(true, text)
    try {
      await fn()
    } finally {
      setLoading(false)
    }
  }

  return {
    isLoading: readonly(isLoading),
    loadingText: readonly(loadingText),
    setLoading,
    withLoading
  }
}
