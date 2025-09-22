/**
 * Composable to track and handle new user onboarding
 */
export const useOnboarding = () => {
  const isNewUser = useState('isNewUser', () => false)
  const hasSeenWelcome = useState('hasSeenWelcome', () => false)
  
  const checkIfNewUser = async () => {
    try {
      const setupStatus = await $fetch('/api/setup/status')
      
      // If user just completed setup, they're new
      if (!setupStatus.needsSetup && setupStatus.hasProfile) {
        // Check if this is their first time on the app
        const welcomeShown = localStorage.getItem('journalful-welcome-shown')
        if (!welcomeShown) {
          isNewUser.value = true
        }
      }
      
      return setupStatus
    } catch (error) {
      console.warn('Failed to check user status:', error)
      return null
    }
  }
  
  const markWelcomeShown = () => {
    localStorage.setItem('journalful-welcome-shown', 'true')
    hasSeenWelcome.value = true
    isNewUser.value = false
  }
  
  const showWelcomeGuide = () => {
    // This could trigger a guided tour or welcome modal
    const { addNotification } = useNotifications()
    
    addNotification({
      type: 'success',
      message: '🎉 Welcome to Journalful! Your default library has been created.',
      duration: 8000
    })
    
    markWelcomeShown()
  }
  
  return {
    isNewUser: readonly(isNewUser),
    hasSeenWelcome: readonly(hasSeenWelcome),
    checkIfNewUser,
    markWelcomeShown,
    showWelcomeGuide
  }
}